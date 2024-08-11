package main

import (
  "fmt"
  "time"
  "github.com/robfig/cron"
)

type Loop struct {
  wallpapers []string
  currentWallpaper string
}

func newLoop(wallpapers []string, currentWallpaper string) {
  loop := Loop { wallpapers: wallpapers, currentWallpaper: currentWallpaper }
  call(loop)
  fmt.Println("Loop started")
}

func call(loop Loop) {
    cronJob := cron.New()

    cronJob.AddFunc("* * * * *", func() {
        fmt.Println("Hello world!")
    })

    // Start the Cron job scheduler
    cronJob.Start()

    // Wait for the Cron job to run
    time.Sleep(5 * time.Minute)

    // Stop the Cron job scheduler
    cronJob.Stop()
}

//func start(loop Loop) {
//  randomWallpaperIndex := randomWallpaperIndex(
//    loop.wallpapers,
//    loop.currentWallpaper,
//  )
//
//  feh := []string { "feh", "--bg-scale", loop.wallpapers[randomWallpaperIndex] }
//  executeCommand(feh)
//  fmt.Println("New Background:", loop.wallpapers[randomWallpaperIndex])
//}
