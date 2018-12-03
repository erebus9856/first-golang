package main

import (
  "testing"
  "time"
  // "fmt"
)

func TestDateCalculation(t *testing.T) {
  nowTime = func() time.Time {
      return time.Date(1985, 11, 9, 0, 0, 0, 0, time.UTC)
  }
  t.Log("Testing to see if this actually calculates the difference between a birthday and a fixed date correctly.")
  birthData := [3]int { 1982, 12, 2 }
  var ageData [3]float64
  ageData = CalcAge(birthData)

  t.Log("Year: ",ageData[0])
  t.Log("Month: ",ageData[1])
  t.Log("Day: ",ageData[2])

  if (ageData[0] != 2) {
    t.Error("Year Calculation WRONG!!")
  }
  if (ageData[1] != 11) {
    t.Error("Month Calculation WRONG!!!!")
  }
  if (ageData[2] != 7) {
    t.Error("Day Calculation WRONG!!!")
  }

  resetClock()
}

// Thanks for the help
// http://labs.yulrizka.com/en/2014/10/stubbing-time-dot-now-in-golang.html#.XAWrLvxOmEI
