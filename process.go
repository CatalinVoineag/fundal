package main

import (
  "fmt"
  "os/exec"
  "strconv"
)

func newProcess(args []string) {
  minutes, err := validDurationArgs(args)
  if err != nil {
    fmt.Println(err.Error())
    return 
    // maybe return error from this method?
  }

  command_args := []string{"loop", strconv.Itoa(minutes)}
  error := exec.Command("fundal", command_args...).Start()
  // Kill process if already exists
  if error != nil {
    panic(error)
  }
}
