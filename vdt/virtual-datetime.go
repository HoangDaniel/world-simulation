package vdt

import "fmt"

type VirtualDateTime struct {
	Day    int
	Month  int
	Year   int
	Second int
	Minute int
	Hour   int
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

func (t *VirtualDateTime) increaseYear() {
	t.Year++
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

func (t *VirtualDateTime) ForwardSecond() {
	if t.Second >= 59 {
		t.Second = 0
		t.increaseMinute()
	}
	t.Second++
}

func (t *VirtualDateTime) ForwardMinute() {
	for i := 0; i < 60; i++ {
		t.ForwardSecond()
	}
}

func (t *VirtualDateTime) ForwardHour() {
	for i := 0; i < 60; i++ {
		t.ForwardMinute()
	}
}

func (t *VirtualDateTime) ForwardDay() {
	for i := 0; i < 24; i++ {
		t.ForwardHour()
	}
}

func (t *VirtualDateTime) ForwardMonth() {
	for i := 0; i < 30; i++ {
		t.ForwardDay()
	}
}

func (t *VirtualDateTime) ForwardYear() {
	for i := 0; i < 12; i++ {
		t.ForwardMonth()
	}
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
