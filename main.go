package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/charmbracelet/huh"
)

//go:embed data/addons/*.vpk
var addonsFS embed.FS

//go:embed data/autoexec.cfg
var autoexecCfg []byte

func main() {
	fmt.Println("===========================================")
	fmt.Println(" Left 4 Dead 2 Competitive Mod Installer")
	fmt.Println("===========================================")
	fmt.Println()

	defaultPath := detectDefaultPath()

	var targetDir string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Left 4 Dead 2 Installation Directory").
				Description("Enter the path to your Left 4 Dead 2 folder").
				Value(&targetDir).
				Placeholder(defaultPath).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("path cannot be empty")
					}
					info, err := os.Stat(filepath.Join(s, "left4dead2"))
					if err != nil || !info.IsDir() {
						return fmt.Errorf("could not find 'left4dead2' inside this directory. Are you sure this is the right path?")
					}
					return nil
				}),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Installation cancelled.")
		os.Exit(1)
	}

	if targetDir == "" {
		targetDir = defaultPath
	}

	fmt.Printf("\nInstalling to: %s\n\n", targetDir)

	// 1. Install CFG
	cfgPath := filepath.Join(targetDir, "left4dead2", "cfg", "autoexec.cfg")
	fmt.Printf("Writing %s...\n", cfgPath)
	err = os.WriteFile(cfgPath, autoexecCfg, 0644)
	if err != nil {
		fmt.Printf("Error writing autoexec.cfg: %v\n", err)
	} else {
		fmt.Println(" -> Success!")
	}

	// 2. Install VPKs
	addonsDir := filepath.Join(targetDir, "left4dead2", "addons")
	fmt.Println("\nExtracting Addons (VPKs)...")
	
	entries, err := fs.ReadDir(addonsFS, "data/addons")
	if err != nil {
		fmt.Printf("Error reading embedded addons: %v\n", err)
		os.Exit(1)
	}

	count := 0
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".vpk") {
			continue
		}
		
		targetPath := filepath.Join(addonsDir, entry.Name())
		fmt.Printf("Extracting %s...\n", entry.Name())
		
		data, err := addonsFS.ReadFile("data/addons/" + entry.Name())
		if err != nil {
			fmt.Printf(" -> Error reading %s: %v\n", entry.Name(), err)
			continue
		}
		
		err = os.WriteFile(targetPath, data, 0644)
		if err != nil {
			fmt.Printf(" -> Error writing %s: %v\n", entry.Name(), err)
		} else {
			count++
		}
	}

	fmt.Printf("\n===========================================\n")
	fmt.Printf(" Installation Complete! Installed %d addons.\n", count)
	fmt.Printf(" You can now safely unsubscribe from these mods on the Steam Workshop.\n")
	fmt.Printf("===========================================\n")
}

func detectDefaultPath() string {
	if runtime.GOOS == "windows" {
		return "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Left 4 Dead 2"
	}
	
	// Check common Linux paths
	home, _ := os.UserHomeDir()
	paths := []string{
		filepath.Join(home, ".local", "share", "Steam", "steamapps", "common", "Left 4 Dead 2"),
		filepath.Join(home, ".steam", "steam", "steamapps", "common", "Left 4 Dead 2"),
		filepath.Join(home, ".var", "app", "com.valvesoftware.Steam", ".local", "share", "Steam", "steamapps", "common", "Left 4 Dead 2"), // Flatpak
	}

	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	
	return ""
}
