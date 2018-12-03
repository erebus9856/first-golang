package main

import (
  "fmt"
  "bufio"
  "os"
  "time"
  "math"
  "strconv"
  // "reflect"
  "strings"
)

func main() {
  fmt.Printf("Welcome to this. Whatever this is. Time to Go.\n")

  var birthData [3]int
  var theName string
  birthData, theName = GetInput()

  fmt.Println("\n######################################")

  var ageData [3]float64
  ageData = CalcAge(birthData)

  fmt.Println(theName, ", you are", ageData[0], "years", ageData[1], "months", ageData[2], "days old")

}

func CalcAge(birthData [3]int) [3]float64 {
  // Get today's date
  year, month, day := nowTime().Date()

  // fmt.Println("Year   :", year)
  // fmt.Println("Month  :", int(month))
  // fmt.Println("Day    :", day)

  birthConv := Date(birthData[0], birthData[1], birthData[2])
  todayConv := Date(year, int(month), day)

  seconds := todayConv.Sub(birthConv).Seconds()
  years := math.Floor(seconds / ( 86400 * 365.25 ))

  months := math.Floor(( years - math.Floor(years) ) * 12)

  days := math.Ceil(( seconds - ( math.Floor(years) * 86400 * 365.25) ) / 86400)

  CalcAge := [3]float64 { years, months, days }

  return CalcAge

}


// Function to get all reqired data
func GetInput() ([3]int, string) {
  // Declare variables
  var thisText [3]string
  var birthMonth int
  var birthDay int
  var birthYear int

  reader := bufio.NewReader(os.Stdin)
  fmt.Print("What is your name? ")
  name, _ := reader.ReadString('\n')
  clean_name := strings.TrimSuffix(name, "\n")

  thisText[0] = "Enter the month you were born (1-12): "
  thisText[1] = "Enter the day you were born (1-31): "
  thisText[2] = "Enter the year you were born (YYYY): "
  birthMonth = GetIntInput(thisText[0], 1, 12)
  birthDay = GetIntInput(thisText[1], 1, 31)
  birthYear = GetIntInput(thisText[2], 0, 9999)
  fmt.Println(birthMonth,"-",birthDay,"-",birthYear)

  birthData := [3]int { birthYear, birthMonth, birthDay }

  return birthData, clean_name
}

// Function to get input for an individual int
// @// TODO: need to make sure that the int is within the correct range.
func GetIntInput(displayText string, min int, max int) int {
  // Declare the variables
  var validInt bool
  var clean_input string
  var myInt int

  // Set a bool to false so this for (which is a while) will keep going
  validInt = false
  // do this until the user inputs a valid int
  for !validInt {
    // setup a reader
    reader := bufio.NewReader(os.Stdin)
    // Display text asking user to enter something
    fmt.Println(displayText)
    input, _ := reader.ReadString('\n')
    clean_input = strings.TrimSuffix(input, "\n")
    validInt = CheckInput(clean_input, min, max)
  }

  myInt, _ = strconv.Atoi(clean_input)
  return myInt
}

// Function to check if an int is input or not
func CheckInput(clean_input string, min int, max int) bool {
  // Convert string to int
  thisInt, err := strconv.Atoi(clean_input)

  // if it converts correctly, return true
  if (thisInt >= min && thisInt <= max && err == nil) {
    return true
  } else { // else return false
    return false
  }
}

func Date(year, month, day int) time.Time {
  return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}


// Functions to Stub time and set it for testing.
type TypeNowTime func() time.Time

var nowTime TypeNowTime

func init() {
    resetClock()
}

func resetClock() {
    nowTime = func() time.Time {
        return time.Now()
    }
}
