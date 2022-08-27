package util

import (
	"os"
	logger "shopfeed/pkg/log"
	"time"
)

const HOUR_TO_MILLISECOND int64 = 3.6e+6

var log = logger.Logger{}

func AddTimestamp(timestamp int64, additionalHours int64) int64 {

	timestamp = timestamp + additionalHours*HOUR_TO_MILLISECOND

	return timestamp
}

func ConvertEpochTimeToTime(timeFloat int64) time.Time {
	tz := os.Getenv("TZ")
	if tz == "" {
		tz = "Asia/Ho_Chi_Minh"
	}

	location, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatal("Fatal error " + err.Error())
	}
	timeMsg := time.Unix(timeFloat/1000, 0).In(location)
	return timeMsg
}
