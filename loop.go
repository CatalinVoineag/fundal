package main

import (
  "time"
  "fmt"
  "strconv"
  "errors"
)

func newLoop(args []string) {
  minutes, err := validDurationArgs(args)

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  for true {
    timer1 := time.NewTimer(time.Duration(minutes) * time.Minute)

    <-timer1.C
    changeBackground() 
  }
}

func validDurationArgs(args []string) (int, error) {
  if len(args) < 3 {
    return 0, errors.New("You need to pass the minutes")
  }

  minutes, err := strconv.Atoi(args[2])

  if err != nil {
    return 0, errors.New("You need to pass the minutes as an int")
  }

  return minutes, nil
}
