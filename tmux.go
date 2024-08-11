package main

import (
//  "fmt"
)

func newTmux() {
  if sessionActive() == true {
    killCurrentSession() 
    createSession()
  } else {
    createSession()
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

func createSession() {
  commands := []string {
    "tmux",
    "new-session",
    "-d",
    "-s",
    "'fundal process'",
    "/home/catalin/play/fundal/bin/main loop",
  }

  runCommand(commands)
}
