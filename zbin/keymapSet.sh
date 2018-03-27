# !/bin/bash

# caps <-> esc
# setxkbmap -option "caps:swapescape"

# caps <-> ctrl
setxkbmap -option "ctrl:swapcaps"

# alt <-> win
# setxkbmap -option altwin:swap_alt_win

# esc <-> grave
xmodmap -e "keycode 9 = grave asciitilde"
xmodmap -e "keycode 49 = Escape"

