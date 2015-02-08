package timeutils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

	assert.NotEqual(t, true, now.IsLeap(), "Error on IsLeap false")

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

	assert.Equal(t, "2016-02-28 08:10:10", now.YesterdayAtTime(o).Format(format), "YesterdayAtTime Failed")

}

func TestMisc(t *testing.T) {

	n := time.Date(2015, 2, 8, 11, 41, 49, 0, time.Local)

	now := New(n)

	assert.Equal(t, "2015-02-07 11:41:49", now.Ago(86400).Format(format), "Ago(86400) Error")

	assert.Equal(t, "2015-02-09 11:41:49", now.Since(86400).Format(format), "Since(86400) Error")

	assert.Equal(t, time.March, now.NextMonth(), "Error on NextMonth()")

	n = time.Date(2015, 2, 8, 13, 00, 00, 0, time.Local)

	//Interval start time
	is := time.Date(2006, 2, 1, 14, 00, 00, 0, time.Local)
	//Interval end time
	ie := time.Date(2006, 2, 1, 16, 00, 00, 0, time.Local)

	now = New(n)

	assert.Equal(t, time.March, now.NextMonth(), "Error on NextMonth()")

	assert.Equal(t, time.January, now.PreviousMonth(), "Error on PreviousMonth()")

	assert.Equal(t, "2015-02-08 17:00:00", now.NewTimeAfterInterval(is, ie, 7200).Format(format), "Error on NewTimeAfterInterval")

	assert.Equal(t, time.Monday, now.TomorrowDayOfWeek(), "Error on TomorrowDayOfWeek")

	assert.Equal(t, time.Monday, now.DayOfWeekAfterNDays(8), "Error on DayOfWeekAfterNDays")

	assert.Equal(t, time.Saturday, now.YesterdayDayOfWeek(), "Error on YesterdayDayOfWeek")

	assert.Equal(t, "2015-02-10 13:00:00", now.AddDays(2).Format(format), "Error on AddDays")

	assert.Equal(t, 39599, now.SecondsToEndOfDay(), "Error on SecondsToEndOfDay")

	assert.Equal(t, "2015-02-08 23:59:59", now.EndOfDay().Format(format), "Error on EndOfDay")

	assert.Equal(t, "2015-02-08 00:00:00", now.BeginningOfDay().Format(format), "Error on BeginningOfDay")

}
