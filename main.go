package main

import (
	"fmt"
	"time"
)
type TimeTest struct {
	Id int `json:"id"`
	User_Name string `json:"user_name"`
	Creation_Date time.Time `json:"creation_date"`
}

func main() {
	var srcTime time.Time
	//Assign A Date To
	fmt.Println("Difference of 2 date is:",DaysBeforeNow(srcTime))

	//fmt.Println(crTime)
}
func DaysBeforeNow(_date time.Time) int{
	var _days int
	_now:=time.Now()
	_days=Round((_now.Sub(_date).Hours())/24)
	return _days
}
func Round(val float64) int {
	if val < 0 {
		return int(val-0.5)
	}
	return int(val+0.5)
}
