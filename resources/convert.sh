#!/bin/bash
for path in ./avatars_originals/*.png; do
    filename=$(basename "${path}")
    convert \( -size 256x256 xc:#ffffff -colorspace sRGB \) \( ./avatars_originals/$filename -resize 256x256 -colorspace sRGB \) -gravity center -compose over -composite -type TrueColor ./avatars/$filename
done