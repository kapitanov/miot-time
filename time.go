package main

import (
	"log"
	"os"
	"sync"
	"time"
)

var (
	timezone  *time.Location
	timeValue *TimeInfo
	timeLock  sync.Mutex
)

type TimeInfo struct {
	H int `json:"h"`
	M int `json:"m"`
}

func timeInit() {
	tz := os.Getenv("TIMEZONE")
	if tz == "" {
		tz = "UTC"
	}

	timezone = time.FixedZone(tz, 0)
	log.Printf("time: using timezone %s\n", timezone)

	timeUpdate()

	go func() {
		for {
			time.Sleep(time.Second * 10)
			if timeUpdate() {
				mqttPublish()
			}
		}
	}()
}

func timeGet() *TimeInfo {
	timeLock.Lock()
	defer timeLock.Unlock()
	return timeValue
}

func timeUpdate() bool {
	t := time.Now().In(timezone)

	timeLock.Lock()
	defer timeLock.Unlock()
	val := &TimeInfo{H: t.Hour(), M: t.Minute()}
	if timeValue == nil || timeValue.H != val.H || timeValue.M != val.M {
		timeValue = val
		return true
	}

	return false
}
