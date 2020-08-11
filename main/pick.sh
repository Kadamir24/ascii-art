#! /bin/bash
width=$(tput cols)
if [ "$1" = "center" ]; then
    while IFS= read -r line
    do 
    printf "%$((width/2 - ${#line}/2))s%s\n" "" "$line"
    
    done < $2
fi
if [ "$1" = "right" ]; then
    while IFS= read -r line
    do 
    printf "%*s\n" "$width" "$line"
    
    done < $2
fi
if [ "$1" = "left" ]; then
while IFS= read -r line
    do 
    printf  "$line\n"
    
    done < $2
fi