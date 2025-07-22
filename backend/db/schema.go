package db

const (
	STATUS_QUEUED = iota
	STATUS_SUCCESS
	STATUS_ERROR
	STATUS_EXPIRED
)

type Flag struct {
	Flag    string
	Round   int
	TeamID  int
	Exploit string
	Status  int
}

type Response struct {
	Flag   string
	Status int
}

type FlagReceived struct {
	Flag    string `json:"flag"`
	Exploit string `json:"exploit_name"`
	TeamIP  string `json:"team_ip"`
	TeamID  int
	Round   int64
}

type ExploitData struct {
	IsWorking bool    `json:"is_working"`
	Status    []int64 `json:"status"`
}

type TimelineData struct {
	Status []int64 `json:"status"`
}

type TeamsData struct {
	Status []int64 `json:"status"`
}
