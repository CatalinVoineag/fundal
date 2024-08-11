package main

import (
  "fmt"
  "os"
  "os/exec"
)

func main() {
  fmt.Println("hello")
  if os.Args[1] == "change" {
    commands := []string {"feh", "--bg-scale", "home/catalin/Pictures/wallpapers/firewatch.jpg"}
    runCommand(commands)
  }
}

func runCommand(commands []string) {
  cmd := exec.Command(commands[0], commands[1:]...)
  cmd.Dir = "/"
  err := cmd.Run()

  if err != nil {
    fmt.Println(err.Error())
    return
  }
}
