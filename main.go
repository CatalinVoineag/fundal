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
)

var WALLPAPERS_PATH = fmt.Sprintf("%s/Pictures/wallpapers/", os.Getenv("HOME"))

func main() {
  if len(os.Args) < 2 {
    changeBackground()
  } else {
    switch os.Args[1] {
    case "next":
      nextBackground()
    case "prev":
      prevBackground()
    case "loop":
      newLoop(os.Args)
    case "start":
      newProcess(os.Args)
    case "stop":
      //exitTmux()
      // Stop process
    case "--help":
      fmt.Println(helpText())
    } 
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
  commands := []string { "ls", fmt.Sprintf("%s/Pictures/wallpapers/", os.Getenv("HOME")) }
  out := runCommand(commands)

  if out != "" {
    return strings.Split(out, "\n")
  }

  return []string {}
} 

func currentBackground() string {
  re := regexp.MustCompile(`'([^"]*)'`)
  commands := []string { "cat", fmt.Sprintf("%s/.fehbg", os.Getenv("HOME")) }
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

  cmd.Run()

  returnString := ""

  if out.Len() > 0 {
    returnString = out.String()
  }

  return returnString
}

func helpText() string {
  result := `fundal is a simple tool to change or loop through wallpapers.
The tools 'feh' and 'tmux' need to be installed for it to work

It also expects the wallpapers to be in 'Pictures/wallpapers'

Usage: fundal [options...]
next            Switches to the next wallpaper
prev            Switches to the previous wallpaper
loop minutes    Loops over wallpapers with a specified interval in minutes
start minutes   Starts new tmux session and runs the loop command with an interval in minutes
stop minutes    Cancels any tmux session and loops

Running only the fundal command will change the background to a random wallpaper 
`
  return result
}
