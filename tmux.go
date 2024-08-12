package main

import (
  "strconv"
  "fmt"
)

func newTmux(args []string) {
  if sessionActive() == true {
    killCurrentSession() 
    createSession(args)
  } else {
    createSession(args)
  }
}

func exitTmux() {
  killCurrentSession() 
}

func sessionActive() bool {
  commands := []string {
    "tmux",
    "has-session",
    "-t",
    "'fundal process'",
  }

  return runCommand(commands) == ""
}

func killCurrentSession() {
  commands := []string {
    "tmux",
    "kill-session",
    "-t",
    "'fundal process'",
  }

  runCommand(commands)
}

func createSession(args []string) {
  minutes, err := validDurationArgs(args)

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  commands := []string {
    "tmux",
    "new-session",
    "-d",
    "-s",
    "'fundal process'",
    "/home/catalin/play/fundal/bin/main loop",
    strconv.Itoa(minutes),
  }

  runCommand(commands)
}
