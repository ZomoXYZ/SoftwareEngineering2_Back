#!/bin/bash
for path in ./avatars_originals/*.png; do
    filename=${path##*/}
    name=${filename%.*}
    convert \( -size 256x256 xc:#ffffff -colorspace sRGB \) \( ./avatars_originals/$filename -resize 256x256 -colorspace sRGB \) -gravity center -compose over -composite -type TrueColor ./avatars/$name.png
done