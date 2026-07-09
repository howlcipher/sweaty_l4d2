# Sweaty L4D2 Competitive Installer

The ultimate, no-nonsense, one-click installer for the sweatiest Left 4 Dead 2 competitive players. 

This installer automatically deploys the most optimal, VAC-safe visual/audio addons and config tweaks so you can completely dominate in VS mode without having to manually drag-and-drop a single file.

## What's Included?
1. **Sweaty Autoexec.cfg**: Built-in 100-tick network rates, perfectly tuned interpolation, lerp toggles, null-canceling movement binds, and raw mouse input.
2. **Competitive Visual Addons**: 
   - Neon Special Infected (See them in the dark)
   - Neon Survivors (For when you're the Infected)
   - Transparent Viewmodels & No Weapon Smoke/Flash
   - avHUD 2 (Centered, clean HUD)
   - Solid RNG-Free Crosshair
3. **Audio Tweaks**: 
   - Quiet Weapons Pack (Hear hunter spawns over the gunfire)

## How to use
Just run the executable for your platform. It automatically detects your L4D2 installation, unpacks the embedded VPKs straight into your `addons` folder (skipping the slow Steam Workshop verification at launch), and drops the sweaty `autoexec.cfg` right where it belongs.

### Windows
Double click `l4d2-installer-windows.exe`. 

### Linux / Bazzite / Steam Deck
Run the `l4d2-installer-linux` binary from your terminal:
```bash
chmod +x l4d2-installer-linux
./l4d2-installer-linux
```

## Build from Source (Go)
If you want to build this yourself or modify the embedded addons:
```bash
go build -o l4d2-installer-linux
GOOS=windows GOARCH=amd64 go build -o l4d2-installer-windows.exe
```

*Stay sweaty.*

## 🚀 Recommended Launch Options
To guarantee the absolute highest frame rates and skip the intro videos, right-click Left 4 Dead 2 in your Steam Library -> Properties -> General -> **Launch Options** and paste this exactly:

```text
-novid -console -vulkan
```
> **Note on `-vulkan`**: This forces the game to use the modern DXVK translation layer which massively boosts FPS and reduces stutter on modern hardware, especially on Linux/Steam Deck. *Warning: The first time you launch with this, the game may heavily stutter or even freeze while compiling shaders. Let it finish or play a single player game first.*

## ⚙️ The `autoexec.cfg` Contents
For full transparency, here is the exact script injected into your `cfg` folder. It uses 100% VAC-safe built-in console variables:

```cfg
// --- LEFT 4 DEAD 2 COMPETITIVE AUTOEXEC ---
// These settings are 100% VAC-safe and utilize built-in Source engine console commands.

// 1. Network Settings
rate "100000"
cl_cmdrate "100"
cl_updaterate "100"
cl_interp "0.0"       
cl_interp_ratio "-1"  

// 2. Mouse Settings
m_rawinput "1"        
m_customaccel "0"     
m_mouseaccel1 "0"
m_mouseaccel2 "0"

// 3. Visual & Performance Tweaks
mat_queue_mode "-1"             
cl_viewmodelfovsurvivor "90"    
cl_crosshair_dynamic "0"        
cl_colorblind "1"               

// 4. Disable Useless Eye/Face Animations
r_eyemove "0"
r_eyegloss "0"
r_eyesize "0"
blink_duration "0"

// 5. Audio Settings
snd_pitchquality "1"
dsp_enhance_stereo "0"          

// 6. Quality of Life
cl_show_splashes "0" 

// 7. Advanced Competitive Binds
net_graphproportionalfont "0"
net_graphpos "1"
alias "+scoregraph" "+showscores; net_graph 1"
alias "-scoregraph" "-showscores; net_graph 0"
bind "TAB" "+scoregraph"

bind "F8" "toggle_lerp"
alias "toggle_lerp" "lerp1"
alias "lerp1" "cl_interp 0.000; echo [ Lerp: 0.000 - Best for <40 ping on 100Tick ]; alias toggle_lerp lerp2"
alias "lerp2" "cl_interp 0.016; echo [ Lerp: 40-70 ping on 60Tick ]; alias toggle_lerp lerp3"
alias "lerp3" "cl_interp 0.033; echo [ Lerp: 0.033 - Vanilla Default/Laggy Servers ]; alias toggle_lerp lerp1"

alias "+quickmelee" "slot2; +attack"
alias "-quickmelee" "-attack; slot1"
bind "v" "+quickmelee"

bind w +mfwd
bind s +mback
bind a +mleft
bind d +mright
alias +mfwd "-back;+forward;alias checkfwd +forward"
alias +mback "-forward;-back;+back;alias checkback +back"
alias +mleft "-moveright;-moveleft;+moveleft;alias checkleft +moveleft"
alias +mright "-moveleft;-moveright;+moveright;alias checkright +moveright"
alias -mfwd "-forward;checkback;alias checkfwd none"
alias -mback "-back;checkfwd;alias checkback none"
alias -mleft "-moveleft;checkright;alias checkleft none"
alias -mright "-moveright;checkleft;alias checkright none"
alias checkfwd none
alias checkback none
alias checkleft none
alias checkright none
alias none ""

echo ">>> L4D2 Competitive Autoexec Loaded Successfully! <<<"
```
