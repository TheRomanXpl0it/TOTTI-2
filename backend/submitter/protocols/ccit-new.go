package protocols

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sub/db"
	"sub/utils/log"
	"time"
)

type CCITNewResponse struct {
	Msg    string `json:"msg"`
	Flag   string `json:"flag"`
	Status string `json:"status"`
}

type CCITNewResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CCITNewGroup struct {
	Group string
	Msg   string
	Code  int
	Idx   int
}

type CCITNewSubmitter struct {
	Submitter
	resp_list   []CCITNewGroup
	response    map[string]CCITNewGroup
	logger      *log.Logger
	regex       *regexp.Regexp
	subUrl      string
	teamToken   string
	subInterval time.Duration
}

func NewCCITNewSubmitter(subUrl string, teamToken string, subInterval int, logger *log.Logger) *CCITNewSubmitter {
	resp_list := []CCITNewGroup{
		{Group: "accepted", Msg: "flag claimed", Code: db.STATUS_SUCCESS},
		{Group: "invalid", Msg: "invalid", Code: db.STATUS_ERROR},
		{Group: "old", Msg: "too old", Code: db.STATUS_EXPIRED},
		{Group: "yours", Msg: "your own", Code: db.STATUS_ERROR},
		{Group: "duplicate", Msg: "already claimed", Code: db.STATUS_ERROR},
		{Group: "nop", Msg: "from NOP team", Code: db.STATUS_ERROR},
		{Group: "unavailable", Msg: "not available", Code: db.STATUS_ERROR},
		{Group: "dispatch_err", Msg: "the check which dispatched this flag didn't terminate successfully", Code: db.STATUS_ERROR},
		{Group: "young", Msg: "the flag is not active yet", Code: db.STATUS_QUEUED},
		{Group: "critical", Msg: "notify the organizers", Code: db.STATUS_ERROR},
	}

	responses := make(map[string]CCITNewGroup)
	messages := make([]string, 0, len(responses))
	for i := range len(resp_list) {
		resp_list[i].Idx = i
		resp_list[i].Msg = strings.ToLower(resp_list[i].Msg)
		responses[resp_list[i].Msg] = resp_list[i]
		messages = append(messages, resp_list[i].Msg)
	}
	regex := regexp.MustCompile("(" + strings.Join(messages, "|") + ")")

	return &CCITNewSubmitter{
		resp_list:   resp_list,
		response:    responses,
		logger:      logger,
		regex:       regex,
		subUrl:      subUrl,
		teamToken:   teamToken,
		subInterval: time.Duration(subInterval) * time.Second,
	}
}

func (s *CCITNewSubmitter) Submit(flags []string) []db.Response {
	responses, err := s.submitFlags(flags)
	if err != nil {
		log.Error("failed to submit flags", "err", err)
		return nil
	}

	groups := make([]int, len(s.resp_list))
	resps := make([]db.Response, len(responses))
	for i, r := range responses {
		resps[i].Flag = r.Flag
		find := s.regex.FindString(strings.ToLower(r.Msg))
		if find == "" {
			s.logger.Errorf("unknown response: %s", r.Msg)
			continue
		}
		if status, ok := s.response[find]; ok {
			resps[i].Status = status.Code
			groups[status.Idx]++
			if status.Idx == len(s.resp_list)-1 {
				log.Critical("Contact The Organizers", "flag", r.Flag, "msg", r.Msg)
			}
		} else {
			s.logger.Errorf("unknown response: %s", r.Msg)
			continue
		}
	}

	if len(resps) != len(flags) {
		s.logger.Criticalf("response length mismatch: %d != %d", len(resps), len(flags))
	}

	log_groups := make([]interface{}, 0, len(groups)*2)
	for _, g := range s.resp_list {
		count := groups[g.Idx]
		if count == 0 {
			continue
		}
		log_groups = append(log_groups, g.Group, count)
	}
	s.logger.Info(fmt.Sprintf("Submitted %d flags:", len(flags)), log_groups...)
	return resps
}

func (s *CCITNewSubmitter) submitFlags(flags []string) ([]CCITNewResponse, error) {
	flagsJSON, err := json.Marshal(flags)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", s.subUrl, bytes.NewBuffer(flagsJSON))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Team-Token", s.teamToken)

	client := http.Client{Timeout: s.subInterval}
	resp, err := client.Do(req)
	if err != nil {
		urlErr, ok := err.(*url.Error)
		if ok && urlErr.Timeout() {
			return nil, fmt.Errorf("timeout: %v", err)
		}
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v: Response: %+v", err, resp)
	}

	var res []CCITNewResponse
	err = json.Unmarshal(body, &res)
	if err == nil {
		return res, nil
	}

	if resp.StatusCode == 500 {
		var res CCITNewResponseError
		err = json.Unmarshal(body, &res)
		if err == nil {
			if res.Code == "RATE_LIMIT" {
				return nil, fmt.Errorf("rate limit exceeded: %s", res.Message)
			}
			return nil, fmt.Errorf("server error: %s: %s", res.Code, res.Message)
		}
	}

	return nil, fmt.Errorf(
		"failed to unmarshal response: %v: Response: %+v Body: %v",
		err, resp, string(body))
}
