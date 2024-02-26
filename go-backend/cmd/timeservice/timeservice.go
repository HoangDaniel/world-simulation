package timeservice

import (
	"sync"
	"time"
)

type TimeObserver interface {
	NotifyTimeChange(time.Time)
}

type TimeService struct {
	sync.Mutex
	currentTime time.Time
	speed       float64
	isPaused    bool
	observers   []TimeObserver
}

func NewTimeService(startTime time.Time) *TimeService {
	ts := &TimeService{
		currentTime: startTime,
		speed:       1,
		isPaused:    false,
	}

	go ts.startClock()
	return ts
}

func (ts *TimeService) SetSpeed(s float64) {
	ts.Lock()
	defer ts.Unlock()

	ts.speed = s
}

func (ts *TimeService) Pause() {
	ts.Lock()
	defer ts.Unlock()

	ts.isPaused = true
}

func (ts *TimeService) Resume() {
	ts.Lock()
	defer ts.Unlock()

	ts.isPaused = false
}

func (ts *TimeService) GetCurrentTime() time.Time {
	ts.Lock()
	defer ts.Unlock()

	return ts.currentTime
}

func (ts *TimeService) startClock() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		ts.advanceTime()
	}
}

func (ts *TimeService) advanceTime() {
	ts.Lock()
	defer ts.Unlock()

	if !ts.isPaused {
		ts.currentTime = ts.currentTime.Add(time.Second * time.Duration(ts.speed))
	}
}
