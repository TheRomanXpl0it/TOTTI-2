package submitter

import (
	"sub/db"
	"sub/submitter/protocols"
	"sub/utils/config"
	"sub/utils/log"
	"sub/utils/ordered_set"
	"time"
)

func Loop(conf *config.Config) {
	var (
		start_time int64
		end_time   int64
		duration   int64
		sub_round  int64
		queueSize  int
	)

	interval_duration := int64((time.Duration(conf.SubInterval) * time.Second).Seconds())
	queue := ordered_set.NewOrderedSet()
	protocol := protocols.SelectProtocol(conf.SubUrl, conf.TeamToken, conf.SubInterval, conf.SubProtocol)

	log.Info("Starting submitter loop")
	for {
		start_time = time.Now().Unix()
		sub_round = (start_time-conf.FirstRound)/conf.RoundDuration + 1

		flags, err := db.FetchFlags(sub_round)
		if err != nil {
			log.Error("Error querying flags", "err", err)
			continue
		}
		log.Debug("Fetched flags", "round", sub_round, "flags", len(flags))
		for _, flag := range flags {
			queue.Add(flag.Flag)
		}
		queueSize = queue.Size()

		if queueSize == 0 {
			log.Info("No flags to submit")
		}

		counter := 0
		for i := 0; i < min(conf.SubLimit, queueSize); {
			flags := make([]string, 0, min(conf.SubMaxPayloadSize, queueSize))
			training := make([]string, 0, min(conf.SubMaxPayloadSize, queueSize)/10)
			for range min(conf.SubMaxPayloadSize, queueSize) {
				front := queue.Pop(true)
				if front == nil {
					log.Critical("Queue is empty")
					break
				}
				counter++
				if conf.Training && counter%10 == 3 {
					training = append(training, front.(string))
				} else {
					flags = append(flags, front.(string))
				}
			}
			log.Info("Submitting flags", "flags", len(flags))

			responses := protocol.Submit(flags)
			if responses != nil {
				err := db.UpdateFlags(responses)
				if err != nil {
					log.Error("Error updating flags", "err", err)
				}
				if conf.Training && len(training) > 0 {
					log.Info("Submitting training flags", "flags", len(training))
					trainingResponses := make([]db.Response, 0, len(training))
					for _, flag := range training {
						trainingResponses = append(trainingResponses, db.Response{
							Flag:   flag,
							Status: db.STATUS_SUCCESS,
						})
					}
					err := db.UpdateFlags(trainingResponses)
					if err != nil {
						log.Error("Error updating training flags", "err", err)
					}
				}
			}
			i += len(flags)
		}

		log.Debug("Updating expired flags", "round", sub_round)
		expired, err := db.UpdateExpiredFlags(sub_round)
		if err != nil {
			log.Error("Error updating expired flags", "err", err)
		} else {
			if expired > 0 {
				log.Warn("Updated expired flags", "expired", expired)
			} else {
				log.Debug("No expired flags to update")
			}
		}

		end_time = time.Now().Unix()
		duration = end_time - start_time
		if duration < interval_duration {
			time.Sleep(time.Duration(interval_duration-duration) * time.Second)
		}
	}
}
