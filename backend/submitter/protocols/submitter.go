package protocols

import (
	"sub/db"
	"sub/utils/log"
)

type Submitter interface {
	Submit(flags []string) []db.Response
}

func SelectProtocol(subUrl string, teamToken string, subInterval int, protocol string) Submitter {
	logger := log.WithPrefix("protocol")

	switch protocol {
	case "ccit":
		return NewCCITSubmitter(subUrl, teamToken, subInterval, logger)
	case "ccit-new":
		return NewCCITNewSubmitter(subUrl, teamToken, subInterval, logger)
	default:
		log.Fatal("Unknown protocol", "protocol", protocol)
		return nil
	}
}
