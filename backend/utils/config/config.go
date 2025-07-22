package config

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	LogLevel string `yaml:"log_level"`
	LogFile  string `yaml:"log_file"`

	Team          int    `yaml:"team"`
	NumberOfTeams int    `yaml:"number_of_teams"`
	TeamToken     string `yaml:"team_token"`

	Training bool `yaml:"training"`

	TeamFormat string   `yaml:"team_format"`
	TeamIp     string   `yaml:"team_ip"`
	NopTeam    string   `yaml:"nop_team"`
	Teams      []string `yaml:"teams"`

	RoundDuration  int64  `yaml:"round_duration"`
	RoundFlagAlive int64  `yaml:"flag_alive"`
	FlagFormat     string `yaml:"flag_format"`
	FlagAlive      int64
	FlagRegex      *regexp.Regexp

	FlagIdUrl string `yaml:"flagid_url"`

	SubProtocol       string `yaml:"sub_protocol"`
	SubLimit          int    `yaml:"sub_limit"`
	SubInterval       int    `yaml:"sub_interval"`
	SubMaxPayloadSize int    `yaml:"sub_max_payload_size"`
	SubUrl            string `yaml:"sub_url"`
	StartRound        string `yaml:"start_round"`
	Rounds            int64  `yaml:"rounds"`
	FirstRound        int64
	LastRound         int64

	Database string `yaml:"database"`
}

func LoadConfig(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	dec := yaml.NewDecoder(file)
	if err = dec.Decode(&config); err != nil {
		return nil, err
	}

	config.TeamIp = fmt.Sprintf(config.TeamFormat, config.Team)
	config.NopTeam = fmt.Sprintf(config.TeamFormat, 0)
	config.FlagAlive = config.RoundFlagAlive * int64(config.RoundDuration)
	config.FlagRegex = regexp.MustCompile(config.FlagFormat)

	config.Teams = make([]string, 0, config.NumberOfTeams-2)
	for i := range config.NumberOfTeams + 1 {
		if i == 0 || i == config.Team {
			continue
		}
		config.Teams = append(config.Teams, fmt.Sprintf(config.TeamFormat, i))
	}

	startRound, err := time.Parse("2006-01-02T15:04-07", config.StartRound)
	if err != nil {
		return nil, err
	}
	config.FirstRound = startRound.Unix()
	config.LastRound = config.FirstRound + config.RoundDuration*config.Rounds

	return &config, nil
}
