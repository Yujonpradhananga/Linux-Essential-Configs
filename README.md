# Linux Essential Configs

My personal dotfiles for Arch Linux.
---

##                                                    Hyprland
<img width="1917" height="1080" alt="Image" src="https://github.com/user-attachments/assets/21570495-d974-4722-8383-b684577bbb5c" />
<img width="1918" height="1081" alt="Image" src="https://github.com/user-attachments/assets/3ecf9205-7d0a-49bc-af17-350d3334dd47" />
<img width="1922" height="1082" alt="Image" src="https://github.com/user-attachments/assets/57966bde-0aa7-42ea-8fb1-fde80949ce0e" />
<img width="1918" height="1081" alt="Image" src="https://github.com/user-attachments/assets/aafc7ea1-36c1-46ba-ba72-65771bb142c4" />
<img width="1915" height="1079" alt="Image" src="https://github.com/user-attachments/assets/748693e0-e17f-4d53-82e4-19c30c851b77" />
<img width="1917" height="1079" alt="Image" src="https://github.com/user-attachments/assets/9dfae166-73da-481b-b003-5c707df3832a" />
---

## DWL
<img width="1921" height="1084" alt="Image" src="https://github.com/user-attachments/assets/b6cbfe09-a1ec-4ee3-998a-63ae456695da" />
<img width="1921" height="1086" alt="Image" src="https://github.com/user-attachments/assets/5666195e-3e0e-4888-8565-20f24e18c080" />
<img width="1919" height="1080" alt="Image" src="https://github.com/user-attachments/assets/1ebe95f4-b74e-4a08-930b-4e755f02a255" />
<img width="1921" height="1089" alt="Image" src="https://github.com/user-attachments/assets/52a95699-f339-4312-8e2a-e98957573baa" />
---

## What's included

- **foot** - Terminal emulator config
- **kitty** - Terminal emulator config
- **dwl** - Window manager config
- **hypr** - Window manager config
- **qutebrowser** - vim motion browser config
- **rofi** - application runner
- **Pictures** - wallpaper collection
- **nvim** - nvim config
- **wal** - pywal config with a custom starship script
- **waybar** - top bar, note i have multiple configs here, put any one in the ~/.config/waybar/ directory. Note: i am using the 'shuriken' config here.
  
## Usage

Copy the configs to your `~/.config/` directory:


For dwl, you'll need to compile from source:
```bash
cd dwl/dwl
sudo make clean install
cd dwl/dwl/slstatus
sudo make clean install
```

---

*Personal configurations - use at your own risk.*
