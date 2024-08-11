package main

import (
  "fmt"
  "os"
  "os/exec"
  "bytes"
  "slices"
  "strings"
	"regexp"
  "math/rand"
  //"reflect"
)

const WALLPAPERS_PATH = "home/catalin/Pictures/wallpapers/"

func main() {
  fmt.Println("hello")
  if os.Args[1] == "change" {
    changeBackground()
  }
}

func changeBackground() {
  wallPaperPath := WALLPAPERS_PATH + backgrounds()[randomBackgroundIndex()]
  commands := []string {"feh", "--bg-scale", wallPaperPath}
  runCommand(commands)
  fmt.Println("New background:", currentBackground())
}

func backgrounds() []string {
  commands := []string { "ls", "home/catalin/Pictures/wallpapers/" }
  out := runCommand(commands)
  if out != "" {
    return strings.Split(out, "\n")
  }

  return []string {}
} 

func currentBackground() string {
  commands := []string { "cat", "/home/catalin/.fehbg" }
  runCommand(commands)
  re := regexp.MustCompile(`'([^"]*)'`)
  fehContent := strings.Trim(re.FindString(runCommand(commands)), "'")
  splitFeh := strings.Split(fehContent, "/")

  return splitFeh[len(splitFeh) -1]
}

func randomBackgroundIndex() int {
  currentIndex := slices.Index(backgrounds(), currentBackground())
  index := rand.Intn(len(backgrounds()))

  for currentIndex == index {
    index = rand.Intn(len(backgrounds()))
  }

  return index 
}

func runCommand(commands []string) string {
  cmd := exec.Command(commands[0], commands[1:]...)
  cmd.Dir = "/"
  var out bytes.Buffer
  cmd.Stdout = &out

  err := cmd.Run()

  if err != nil {
    fmt.Println(err.Error())
    return ""
  }

  return out.String()
}
