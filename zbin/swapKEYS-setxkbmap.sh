# !/bin/bash

#
# *** search 
# $ grep -E "(ctrl|caps):" /usr/share/X11/xkb/rules/base.lst

# caps <-> ctrl
setxkbmap -option "ctrl:swapcaps"


# caps <-> esc
# setxkbmap -option "caps:swapescape"


# alt <-> win
# setxkbmap -option altwin:swap_alt_win


# hangul
#setxkbmap -option "korean:rctrl_ralt" # Right Ctrl as Hangul, right Alt as Hanja
#setxkbmap -option "korean:ralt_rctrl"   
