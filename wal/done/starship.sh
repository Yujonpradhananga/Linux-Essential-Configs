#!/bin/bash
# Source pywal colors
source ~/.cache/wal/colors.sh

# Replace color placeholders in template and save to starship config
sed "s/COLOR1/$color1/g; s/COLOR2/$color2/g; s/COLOR3/$color3/g; s/COLOR4/$color4/g; s/COLOR5/$color5/g; s/COLOR6/$color6/g" /home/yujon/.config/wal/templates/starship.toml >~/.config/starship.toml
