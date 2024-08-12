`fundal` is a simple tool to change or loop through wallpapers. The tools 'feh' and 'tmux' need to be installed for it to work

It also expects the wallpapers to be in 'Pictures/wallpapers'

```
Usage: fundal [options...]
next            Switches to the next wallpaper
prev            Switches to the previous wallpaper
loop minutes    Loops over wallpapers with a specified interval in minutes
start minutes   Starts new tmux session and runs the loop command with an interval in minutes
stop minutes    Cancels any tmux session and loops
```

Running only the `fundal` command will change the background to a random wallpaper 
