package timeutils

import (
	"testing"
	"time"
)

var format = "2006-01-02 15:04:05"

func TestTomorrow(t *testing.T) {

	n := time.Date(2015, 2, 8, 11, 41, 49, 0, time.Local)

	now := New(n)

	if now.Tomorrow().Format(format) != "2015-02-09 11:41:49" {
		t.Errorf("Tomorrow")
	}

	n = time.Date(2015, 2, 28, 11, 41, 49, 0, time.Local)

	now = New(n)

	if now.Tomorrow().Format(format) != "2015-03-01 11:41:49" {
		t.Errorf("Tomorrow")
	}

	n = time.Date(2016, 2, 28, 11, 41, 49, 0, time.Local)

	now = New(n)

	if now.Tomorrow().Format(format) != "2016-02-29 11:41:49" {
		t.Errorf("Tomorrow")
	}

}

func TestTomorrowAtTime(t *testing.T) {

	n := time.Date(2015, 2, 8, 11, 41, 49, 0, time.Local)

	o := time.Date(2010, 1, 28, 8, 10, 10, 0, time.Local)

	now := New(n)

	if now.TomorrowAtTime(o).Format(format) != "2015-02-09 08:10:10" {
		t.Errorf("TomorrowAtTime")
	}

	n = time.Date(2015, 2, 28, 11, 41, 49, 0, time.Local)

	now = New(n)

	if now.TomorrowAtTime(o).Format(format) != "2015-03-01 08:10:10" {
		t.Errorf("TomorrowAtTime")
	}

	n = time.Date(2016, 2, 28, 11, 41, 49, 0, time.Local)

	now = New(n)

	if now.TomorrowAtTime(o).Format(format) != "2016-02-29 08:10:10" {
		t.Errorf("TomorrowAtTime")
	}

}

func TestIsLeap(t *testing.T) {

	n := time.Date(2015, 2, 8, 11, 41, 49, 0, time.Local)

	now := New(n)

	if now.IsLeap() != false {
		t.Errorf("IsLeap")
	}

	n = time.Date(2016, 2, 28, 11, 41, 49, 0, time.Local)

	now = New(n)

	if now.IsLeap() != true {
		t.Errorf("IsLeap")
	}

	n = time.Date(2015, 2, 28, 11, 41, 49, 0, time.Local)

	now = New(n)

	if now.IsLeap() != false {
		t.Errorf("IsLeap FALSE")
	}

}

func TestDaysInMonth(t *testing.T) {

	n := time.Date(2015, 2, 8, 11, 41, 49, 0, time.Local)

	now := New(n)

	if now.DaysInMonth() != 28 {
		t.Errorf("DaysInMonth")
	}

	n = time.Date(2016, 2, 28, 11, 41, 49, 0, time.Local)

	now = New(n)

	if now.DaysInMonth() != 29 {
		t.Errorf("DaysInMonth")
	}
}

func TestYesterday(t *testing.T) {

	n := time.Date(2015, 2, 8, 11, 41, 49, 0, time.Local)

	now := New(n)

	if now.Yesterday().Format(format) != "2015-02-07 11:41:49" {
		t.Errorf("Yesterday")
	}

	n = time.Date(2016, 2, 29, 11, 41, 49, 0, time.Local)
	now = New(n)

	if now.Yesterday().Format(format) != "2016-02-28 11:41:49" {
		t.Errorf("Yesterday")
	}

	o := time.Date(2010, 1, 28, 8, 10, 10, 0, time.Local)

	if now.YesterdayAtTime(o).Format(format) != "2016-02-28 08:10:10" {
		t.Errorf("YesterdayAtTime")
	}

	if now.YesterdayAtTime(o).Format(format) != "2016-02-28 08:10:10" {
		t.Errorf("YesterdayAtTime")
	}

}

func TestMisc(t *testing.T) {

	n := time.Date(2015, 2, 8, 11, 41, 49, 0, time.Local)

	now := New(n)

	if now.Ago(86400).Format(format) != "2015-02-07 11:41:49" {
		t.Errorf("Ago")
	}

	if now.Since(86400).Format(format) != "2015-02-09 11:41:49" {
		t.Errorf("Since")
	}

	if now.NextMonth() != time.March {
		t.Errorf("NextMonth")
	}

	n = time.Date(2015, 2, 8, 13, 00, 00, 0, time.Local)
	//Interval start time
	is := time.Date(2006, 2, 1, 14, 00, 00, 0, time.Local)
	//Interval end time
	ie := time.Date(2006, 2, 1, 16, 00, 00, 0, time.Local)

	now = New(n)

	if now.PreviousMonth() != time.January {
		t.Errorf("PreviousMonth")
	}

	if now.NewTimeAfterInterval(is, ie, 7200).Format(format) != "2015-02-08 17:00:00" {
		t.Errorf("NewTimeAfterInterval")
	}

	if now.TomorrowDayOfWeek() != time.Monday {
		t.Errorf("TomorrowDayOfWeek")
	}

	if now.DayOfWeekAfterNDays(8) != time.Monday {
		t.Errorf("DayOfWeekAfterNDays")
	}

	if now.YesterdayDayOfWeek() != time.Saturday {
		t.Errorf("YesterdayDayOfWeek")
	}

	if now.AddDays(2).Format(format) != "2015-02-10 13:00:00" {
		t.Errorf("AddDays")
	}

	if now.SecondsToEndOfDay() != 39599 {
		t.Errorf("SecondsToEndOfDay")
	}

	if now.EndOfDay().Format(format) != "2015-02-08 23:59:59" {
		t.Errorf("EndOfDay")
	}

	if now.BeginningOfDay().Format(format) != "2015-02-08 00:00:00" {
		t.Errorf("BeginningOfDay")
	}

}
