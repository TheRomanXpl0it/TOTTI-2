package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sub/db"
	"sub/utils/log"
	"time"
)

func dataHandler(pattern string, handler func(int, string) (interface{}, error)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var round int
		roundCookie, err := r.Cookie("round")
		if err == nil {
			round, err = strconv.Atoi(roundCookie.Value)
			if err != nil {
				log.Errorf("error parsing round cookie", "err", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			current_round := (time.Now().Unix()-conf.FirstRound)/conf.RoundDuration + 1
			round = int(current_round) - round
		} else {
			round = -1
		}

		var exploit string
		exploitCookie, err := r.Cookie("exploit")
		if err == nil {
			exploit = exploitCookie.Value
		}

		data, err := handler(round, exploit)
		if err != nil {
			log.Error("error getting data", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Errorf("json encoder error", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func exploitHandler(round int, exploit string) (interface{}, error) {
	var data interface{}
	var err error
	if exploit == "" {
		data, err = db.FetchAllExploitData(round)
	} else {
		data, err = db.FetchExploitData(round, exploit)
	}
	return data, err
}

func timelineHandler(round int, exploit string) (interface{}, error) {
	var data interface{}
	var err error
	if exploit == "" {
		data, err = db.FetchAllTimelineData(round)
	} else {
		data, err = db.FetchTimelineData(round, exploit)
	}
	return data, err
}

func teamsHandler(round int, exploit string) (interface{}, error) {
	var data interface{}
	var err error
	if exploit == "" {
		data, err = db.FetchAllTeamsData(round)
	} else {
		data, err = db.FetchTeamsData(round, exploit)
	}
	return data, err
}
