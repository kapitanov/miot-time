package main

import (
	"fmt"
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

func timeInit() error {
	tz := os.Getenv("TIMEZONE")
	if tz == "" {
		tz = "UTC"
	}

	var err error
	timezone, err = time.LoadLocation(tz)
	if err != nil {
		return err
	}
	fmt.Printf("time: using timezone %s\n", timezone)

	timeUpdate()

	go func() {
		for {
			time.Sleep(time.Second * 10)
			if timeUpdate() {
				mqttPublish()
			}
		}
	}()

	return nil
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

		fmt.Printf("time: now %2d:%02d\n", val.H, val.M)
		return true
	}

	return false
}
