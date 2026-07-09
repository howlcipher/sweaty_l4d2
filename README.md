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
