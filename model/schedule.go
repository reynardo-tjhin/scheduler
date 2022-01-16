package model

import (
	"strconv"
	"strings"
)

type Schedule struct {
	Name     string
	Days     []string // range of days: Monday to Sunday
	Interval int      // by minutes
	Range    string   // 00:00-23:59
	Contents []Content
}

type Content struct {
	Title       string
	Description string
	Duration    int64  // in hours
	Day         string // "Monday", "Tuesday", ..., "Sunday"
	StartTime   string // HH:MM
	EndTime     string // HH:MM
}

func CreateNewSchedule(Name string, Days []string, Interval int, Range string) Schedule {

	return Schedule{
		Name:     Name,
		Days:     Days,
		Interval: Interval,
		Range:    Range,
	}
}

func (s Schedule) AddContent(Title string, Description string, Day string, StartTime string, EndTime string) Schedule {

	var content Content
	content.Title = Title
	content.Description = Description
	content.Day = Day

	if CheckTimeValidity(StartTime) {
		content.StartTime = StartTime
	}
	if CheckTimeValidity(EndTime) {
		content.EndTime = EndTime
	}
	content.Duration = FindTimeDuration(StartTime, EndTime)

	s.Contents = append(s.Contents, content)
	return s
}

func (s Schedule) EditContentTitle(Title string, NewTitle string) bool {

	var isExist bool = false
	for _, content := range s.Contents {
		if content.Title == Title {
			content.Title = NewTitle
			isExist = true
		}
	}
	return isExist
}

func (s Schedule) EditContentDescription(Title string, NewDescription string) bool {

	var isExist bool = false
	for _, content := range s.Contents {
		if content.Title == Title {
			content.Description = NewDescription
			isExist = true
		}
	}
	return isExist
}

func (s Schedule) EditContentStartTime(Title string, NewStarTime string) bool {

	if !CheckTimeValidity(NewStarTime) {
		return false
	}
	var isExist bool = false
	for _, content := range s.Contents {
		if content.Title == Title {
			content.StartTime = NewStarTime
			isExist = true
		}
	}
	return isExist
}

func (s Schedule) EditContentEndTime(Title string, NewEndTime string) bool {

	if !CheckTimeValidity(NewEndTime) {
		return false
	}
	var isExist bool = false
	for _, content := range s.Contents {
		if content.Title == Title {
			content.StartTime = NewEndTime
			isExist = true
		}
	}
	return isExist
}

func CheckTimeValidity(Time string) bool {

	TimeList := strings.Split(Time, ":")
	if len(TimeList) > 2 {
		return false
	}

	Hour, err := strconv.ParseInt(TimeList[0], 10, 8)
	if err != nil {
		return false
	}
	Minute, err := strconv.ParseInt(TimeList[1], 10, 8)
	if err != nil {
		return false
	}

	if Hour > 23 || Hour < 0 {
		return false
	}
	if Minute > 60 || Minute < 0 {
		return false
	}
	return true
}

func FindTimeDuration(Time1 string, Time2 string) int64 {

	Time1List := strings.Split(Time1, ":")
	if len(Time1List) > 2 {
		return 0
	}
	Time2List := strings.Split(Time2, ":")
	if len(Time2List) > 2 {
		return 0
	}

	Hour1, err := strconv.ParseInt(Time1List[0], 10, 8)
	if err != nil {
		return 0
	}
	Minute1, err := strconv.ParseInt(Time1List[1], 10, 8)
	if err != nil {
		return 0
	}

	Hour2, err := strconv.ParseInt(Time2List[0], 10, 8)
	if err != nil {
		return 0
	}
	Minute2, err := strconv.ParseInt(Time2List[1], 10, 8)
	if err != nil {
		return 0
	}

	number1 := Hour1*60 + Minute1
	number2 := Hour2*60 + Minute2
	if number1 == 0 {
		number1 = 1440
	}
	if number2 == 0 {
		number2 = 1440
	}

	if number1 > number2 {
		return 0
	}
	return number2 - number1
}
