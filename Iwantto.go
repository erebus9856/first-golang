package main

import (
  "fmt"
  "bufio"
  "os"
  "time"
  // "math"
  "strconv"
  // "reflect"
  "strings"
)

func main() {
  fmt.Printf("Welcome to this. Whatever this is. Time to Go.\n")

  // Declare variables
  var thisText string
  var birthMonth int
  var birthDay int
  var birthYear int

  thisText = "Enter the month you were born (1-12): "
  birthMonth = GetIntInput(thisText)
  thisText = "Enter the day you were born (1-31): "
  birthDay = GetIntInput(thisText)
  thisText = "Enter the year you were born (YYYY): "
  birthYear = GetIntInput(thisText)
  fmt.Println(birthMonth,"-",birthDay,"-",birthYear)

  // reader := bufio.NewReader(os.Stdin)
  // // fmt.Print("What is your name? ")
  // // name, _ := reader.ReadString('\n')
  // // clean_name := strings.TrimSuffix(name, "\n")
  // // fmt.Println(clean_name)
  // //
  // // fmt.Print("Enter the month you were born (1-12): ")
  // // bmonth, _ := reader.ReadString('\n')
  // // clean_month := strings.TrimSuffix(bmonth, "\n")
  // // fmt.Println(clean_month)
  // //
  // // fmt.Print("Enter the day you were born (1-31): ")
  // // bday, _ := reader.ReadString('\n')
  // // clean_day := strings.TrimSuffix(bday, "\n")
  // // fmt.Println(clean_day)
  // //
  // // fmt.Print("Enter the year you were born (YYYY): ")
  // // byear, _ := reader.ReadString('\n')
  // // clean_year := strings.TrimSuffix(byear, "\n")
  // // fmt.Println(clean_year)
  //
  // fmt.Println("\n######################################\n")
  // year, month, day := time.Now().Date()
  // fmt.Println("Year   :", year)
  // fmt.Println("Month  :", int(month))
  // fmt.Println("Day    :", day)
  //
  // ibyear, _ := strconv.Atoi(clean_year)
  // ibmonth, _ := strconv.Atoi(clean_month)
  // ibday, _ := strconv.Atoi(clean_day)
  //
  // bdate := Date(ibyear, ibmonth, ibday)
  // tdate := Date(year, int(month), day)
  // days := tdate.Sub(bdate).Hours() / 24
  // fmt.Println("You are", days, "days old")

// @// TODO: Fix math
  // years := days / 365.25
  // floor_year := math.Floor(years)
  // remainder_year := years - floor_year
  // months := remainder_year * 12
  // floor_months := math.Floor(months)
  // remainder_months := months - floor_months
  // days = remainder_months * 24

  // fmt.Println(clean_name, ", you are", floor_year, "years", floor_months, "months", math.Floor(days), "days old")


}

// Function to get input for an int
func GetIntInput(displayText string) int {
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
    validInt = CheckInput(clean_input)
  }

  myInt, _ = strconv.Atoi(clean_input)
  return myInt
}

// Function to check if an int is input or not
func CheckInput(clean_input string) bool {
  // Convert string to int
  _, err := strconv.Atoi(clean_input)
  // if it converts correctly, return true
  if err == nil {
    return true
  } else { // else return false
    return false
  }
}

func Date(year, month, day int) time.Time {
  return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
