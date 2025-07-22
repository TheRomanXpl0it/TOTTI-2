package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sub/db"
	"sub/utils/config"
	"sub/utils/log"
	"time"
)

var conf *config.Config

func currentRound() int64 {
	return (time.Now().Unix()-conf.FirstRound)/conf.RoundDuration + 1
}

func flagsHandler(w http.ResponseWriter, r *http.Request) {
	var flags []db.FlagReceived
	err := json.NewDecoder(r.Body).Decode(&flags)
	if err != nil {
		log.Errorf("json decoder error", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(flags) != 0 {
		log.Infof("Received %d flags form %s", len(flags), flags[0].Exploit)
	} else {
		log.Warnf("Received empty flags from %s", flags[0].Exploit)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, f := range flags {
		var team int
		_, err := fmt.Sscanf(f.TeamIP, conf.TeamFormat, &team)
		if err != nil {
			log.Error("Failed to parse team from IP:", "ip", f.TeamIP, "err", err)
			team = -1
		}
		flags[i].TeamID = team
		flags[i].Round = currentRound()
	}

	err = db.InsertFlags(flags)
	if err != nil {
		log.Error("DB insert error", "err", err)
	}

	w.WriteHeader(http.StatusOK)
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	teams := make(map[string]string)
	for i, ip := range conf.Teams {
		teams[fmt.Sprintf("team%d", i+1)] = ip
	}

	response := map[string]interface{}{
		"flag_format":    conf.FlagFormat,
		"round_duration": conf.RoundDuration,
		"teams":          teams,
		"flag_id_url":    conf.FlagIdUrl,
		"flag_lifetime":  conf.FlagAlive,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Errorf("json marshal error", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func manualHandler(w http.ResponseWriter, r *http.Request) {
	var flags []db.FlagReceived
	err := json.NewDecoder(r.Body).Decode(&flags)
	if err != nil {
		log.Errorf("json decoder error", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(flags) != 0 {
		log.Infof("Received a flags form manual")
	} else {
		log.Warnf("Received empty flags from manual")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	flags[0].Exploit = "manual"
	flags[0].TeamID = 0
	flags[0].Round = currentRound()
	flags[0].Flag = conf.FlagRegex.FindString(flags[0].Flag)
	if flags[0].Flag == "" {
		log.Warnf("Received invalid flag from manual")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = db.InsertFlags(flags)
	if err != nil {
		log.Errorf("DB insert error", "err", err)
	}

	w.WriteHeader(http.StatusOK)
}

func ServeAPI(c *config.Config) {
	conf = c

	http.HandleFunc("POST /api/flags", flagsHandler)
	http.HandleFunc("GET /api/config", configHandler)

	http.HandleFunc("POST /api/manual", manualHandler)
	dataHandler("GET /api/exploit", exploitHandler)
	dataHandler("GET /api/timeline", timelineHandler)
	dataHandler("GET /api/teams", teamsHandler)
	http.Handle("GET /", http.FileServer(http.Dir("./static")))

	if err := http.ListenAndServe("0.0.0.0:5000", nil); err != nil {
		log.Fatal("error starting the api server", "err", err)
	}
}
