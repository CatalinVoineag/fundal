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
  if os.Args[1] == "change" {
    changeBackground()
  } else if os.Args[1] == "next" {
    nextBackground()
  } else if os.Args[1] == "prev" {
    prevBackground()
  } else if os.Args[1] == "loop" {
    newLoop(os.Args)
  } else if os.Args[1] == "start" {
    newTmux()
  } else if os.Args[1] == "stop" {
    exitTmux()
  }
}

func changeBackground() {
  wallPaperPath := WALLPAPERS_PATH + backgrounds()[randomBackgroundIndex()]
  commands := []string {"feh", "--bg-scale", wallPaperPath}
  runCommand(commands)
  fmt.Println("New background:", currentBackground())
}

func nextBackground() {
  currentIndex := slices.Index(backgrounds(), currentBackground())
  wallPaperPath := WALLPAPERS_PATH + backgrounds()[currentIndex + 1]
  commands := []string {"feh", "--bg-scale", wallPaperPath}
  runCommand(commands)
  fmt.Println("New background:", currentBackground())
}

func prevBackground() {
  currentIndex := slices.Index(backgrounds(), currentBackground())
  wallPaperPath := WALLPAPERS_PATH + backgrounds()[currentIndex + -1]
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
  re := regexp.MustCompile(`'([^"]*)'`)
  commands := []string { "cat", "/home/catalin/.fehbg" }
  catCommand:= runCommand(commands)

  fehContent := strings.Trim(re.FindString(catCommand), "'")
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

  fmt.Println("Commands", commands)
  cmd.Run()

  returnString := ""

  if out.Len() > 0 {
    returnString = out.String()
  }

  return returnString
}
