# !/bin/bash

# caps <-> esc
# setxkbmap -option "caps:swapescape"

# caps <-> ctrl
setxkbmap -option "ctrl:swapcaps"

# esc <-> grave
xmodmap -e "keycode 9 = grave asciitilde"
xmodmap -e "keycode 49 = Escape"
