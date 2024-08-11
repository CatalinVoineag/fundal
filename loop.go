package main

import (
  "time"
//  "fmt"
)

func newLoop(args []string) {
  for true {
    timer1 := time.NewTimer(10 * time.Second)

    <-timer1.C
    changeBackground() 
  }
}
