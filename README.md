# timeutils
Go collection of small time related utility functions (Alpha)
# timeutils
--
    import "gmaisto/timeutils"


## Usage

#### type Now

```go
type Now struct {
	time.Time
}
```

Now a simple wrapper of time.Time

#### func  New

```go
func New(n time.Time) *Now
```
New return a new instance of Now

#### func (*Now) AddDays

```go
func (t *Now) AddDays(n int) time.Time
```
AddDays add a number of days at Now

#### func (*Now) Ago

```go
func (t *Now) Ago(seconds int64) time.Time
```
Ago return a time instance representing the time a given number of seconds ago

#### func (*Now) BeginningOfDay

```go
func (t *Now) BeginningOfDay() time.Time
```
BeginningOfDay return a time instance for the beginning of the day

#### func (*Now) DayOfWeekAfterNDays

```go
func (t *Now) DayOfWeekAfterNDays(n int) time.Weekday
```
DayOfWeekAfterNDays return the day of week after a specified number of days from
the given time instance

#### func (*Now) DaysInMonth

```go
func (t *Now) DaysInMonth() int
```
DaysInMonth return the number of days in the month which t is in

#### func (*Now) EndOfDay

```go
func (t *Now) EndOfDay() time.Time
```
EndOfDay return a time instance for the end of the day

#### func (*Now) IsLeap

```go
func (t *Now) IsLeap() bool
```
IsLeap return true if the Year in t is a leap year

#### func (*Now) NewTimeAfterInterval

```go
func (t *Now) NewTimeAfterInterval(istart time.Time, iend time.Time, length int64) time.Time
```
NewTimeAfterInterval return a new time instance obtained adding "length" seconds
to the original time instance excluding an interval of time

#### func (*Now) NextMonth

```go
func (t *Now) NextMonth() time.Month
```
NextMonth return the next month (time.Month)

#### func (*Now) PreviousMonth

```go
func (t *Now) PreviousMonth() time.Month
```
PreviousMonth return the next month (time.Month)

#### func (*Now) SecondsToEndOfDay

```go
func (t *Now) SecondsToEndOfDay() int32
```
SecondsToEndOfDay return the number of seconds to the end of the day

#### func (*Now) Since

```go
func (t *Now) Since(seconds int64) time.Time
```
Since return a time instance representing the time a given seconds since the
instance time

#### func (*Now) Tomorrow

```go
func (t *Now) Tomorrow() time.Time
```
Tomorrow return a time instance for tomorrows date at the same time

#### func (*Now) TomorrowAtTime

```go
func (t *Now) TomorrowAtTime(w time.Time) time.Time
```
TomorrowAtTime return a time instance for tomorrow's date at a given time of day

#### func (*Now) TomorrowDayOfWeek

```go
func (t *Now) TomorrowDayOfWeek() time.Weekday
```
TomorrowDayOfWeek return the day of the week of tomorrow

#### func (*Now) Yesterday

```go
func (t *Now) Yesterday() time.Time
```
Yesterday return a time instance for yesterday's date at the same time

#### func (*Now) YesterdayAtTime

```go
func (t *Now) YesterdayAtTime(w time.Time) time.Time
```
YesterdayAtTime return a time instance for yesterday's date at a given time of
day

#### func (*Now) YesterdayDayOfWeek

```go
func (t *Now) YesterdayDayOfWeek() time.Weekday
```
YesterdayDayOfWeek return the day of week of yesterday
