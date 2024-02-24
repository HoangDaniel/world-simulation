package vdt

import "fmt"

func (t *VirtualDateTime) IncreaseSecond() {
	if t.Second >= 59 {
		t.Second = 0
		t.increaseMinute()
	}
	t.Second++
}

func (t *VirtualDateTime) increaseMinute() {
	if t.Minute >= 59 {
		t.Minute = 0
		t.increaseHour()
		return
	}
	t.Minute++
}

func (t *VirtualDateTime) increaseHour() {
	if t.Hour >= 23 {
		t.Hour = 0
		t.increaseDay()
		return
	}
	t.Hour++
}

func (t *VirtualDateTime) increaseDay() {
	if t.Day >= 30 {
		t.Day = 1
		t.increaseMonth()
		return
	}
	t.Day++
}

func (t *VirtualDateTime) increaseMonth() {
	if t.Month >= 12 {
		t.Month = 1
		t.increaseYear()
		return
	}
	t.Month++
}

func (t *VirtualDateTime) increaseYear() {
	t.Year++
}

type VirtualDateTime struct {
	Day    int
	Month  int
	Year   int
	Second int
	Minute int
	Hour   int
}

func (t *VirtualDateTime) Date() string {
	return fmt.Sprintf("%v-%v-%v", t.Day, t.Month, t.Year)
}

func (t *VirtualDateTime) Time() string {
	return fmt.Sprintf("%v:%v:%v", t.Hour, t.Minute, t.Second)
}

func (t *VirtualDateTime) DateTime() string {
	return fmt.Sprintf("%v %v", t.Date(), t.Time())
}

func NewVirtualDateTime() *VirtualDateTime {
	return &VirtualDateTime{
		Day:    1,
		Month:  1,
		Year:   1,
		Second: 0,
		Minute: 0,
		Hour:   0,
	}
}
