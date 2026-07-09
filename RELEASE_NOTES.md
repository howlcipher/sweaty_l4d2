# v1.0.0 - The Sweaty Debut

Welcome to the ultimate deployment tool for Left 4 Dead 2 competitive players. This release entirely automates the excruciating process of manually managing VPK files and writing complex `autoexec.cfg` scripts.

## 🩸 Features in v1.0.0

### The Installer
- **TUI (Terminal User Interface)**: Built in Go using the `huh` library, featuring a sleek, interactive terminal menu.
- **Smart Path Detection**: Automatically detects standard Steam installation paths for both Windows and Linux (including Flatpak and Bazzite/SteamOS defaults).
- **Embedded Payloads**: No external downloads required. All `.vpk` mods and the `autoexec.cfg` are embedded directly into the executable using Go's `//go:embed` for lightning-fast deployments.
- **Workshop Bypass**: Installs mods directly into the main `/addons` directory, completely bypassing the agonizing "Workshop Verification" screen when launching the game.

### The Payload (Included Mods & Config)
This installer currently embeds the definitive competitive meta:
- **Network Overhaul**: `autoexec.cfg` optimized for 100-tick servers (`rate 100000`, `cl_cmdrate 100`, `cl_interp 0.0`).
- **Sweaty Keybinds**: Pre-configured Lerp Toggle (F8), Quick-Melee (V), and Null-Canceling Movement.
- **Visual Clarity**:
  - Glow Items (Easier scavenging)
  - Glows Special Infected (Better visibility)
  - Hunter & Charger Fire Trail (Track their path)
  - avHUD 3 (Clean, centered interface)
  - Dot White Crosshair (Solid, RNG-Free aiming)
- **Audio Clarity**:
  - Quiet weapons COMPLETE pack
  - Mute & Quiet Pack (Hear spawns over gunfire)

## 🛠 Usage
Download the binary for your platform below.
* **Windows**: Double-click `l4d2-installer-windows.exe`.
* **Linux/Steam Deck**: Run `./l4d2-installer-linux` from your terminal.

*Note: Since the `.vpk` files are baked into the binary, the executable file sizes are large (~400MB). This is normal and ensures a 1-click offline installation.*
