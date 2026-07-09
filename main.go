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

// Embedded File System for Addons
//go:embed data/addons/*.vpk
var addonsFS embed.FS

// Embedded Configuration File
//go:embed data/autoexec.cfg
var autoexecCfg []byte

// Dynamic configuration constants to avoid hardcoding strings everywhere
const (
	appFolder      = "left4dead2"
	cfgFolder      = "cfg"
	addonsFolder   = "addons"
	cfgFileName    = "autoexec.cfg"
	dataAddonsDir  = "data/addons"
	vpkExtension   = ".vpk"
	filePerms      = 0644
	windowsOS      = "windows"
	defaultWinPath = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Left 4 Dead 2"
)

func main() {
	printBanner()

	defaultPath := detectDefaultPath()
	var targetDir string

	// Initialize and configure the TUI form
	form := createTUIForm(&targetDir, defaultPath)

	// Run the interactive prompt
	err := form.Run()
	if err != nil {
		fmt.Println("Installation cancelled.")
		os.Exit(1)
	}

	// Use default path if user left it blank
	if targetDir == "" {
		targetDir = defaultPath
	}

	fmt.Printf("\nInstalling to: %s\n\n", targetDir)

	// Step 1: Install Configuration File
	if err := InstallCfg(targetDir, autoexecCfg); err != nil {
		fmt.Printf("Error writing %s: %v\n", cfgFileName, err)
	}

	// Step 2: Install Addons (VPKs)
	count, err := InstallVPKs(targetDir, addonsFS)
	if err != nil {
		fmt.Printf("Critical error installing addons: %v\n", err)
		os.Exit(1)
	}

	printCompletion(count)
}

// printBanner displays the startup ASCII banner
func printBanner() {
	fmt.Println("===========================================")
	fmt.Println(" Left 4 Dead 2 Competitive Mod Installer")
	fmt.Println("===========================================")
	fmt.Println()
}

// createTUIForm configures the interactive terminal prompt
func createTUIForm(targetDir *string, defaultPath string) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Left 4 Dead 2 Installation Directory").
				Description("Enter the path to your Left 4 Dead 2 folder").
				Value(targetDir).
				Placeholder(defaultPath).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("path cannot be empty")
					}
					// Ensure the directory contains the expected app folder
					info, err := os.Stat(filepath.Join(s, appFolder))
					if err != nil || !info.IsDir() {
						return fmt.Errorf("could not find '%s' inside this directory. Are you sure this is the right path?", appFolder)
					}
					return nil
				}),
		),
	)
}

// InstallCfg writes the byte payload into the game's cfg directory
func InstallCfg(baseDir string, cfgData []byte) error {
	cfgPath := filepath.Join(baseDir, appFolder, cfgFolder, cfgFileName)
	fmt.Printf("Writing %s...\n", cfgPath)

	// Ensure destination directory exists before writing
	if err := os.MkdirAll(filepath.Dir(cfgPath), 0755); err != nil {
		return fmt.Errorf("failed to create cfg directory: %w", err)
	}

	err := os.WriteFile(cfgPath, cfgData, filePerms)
	if err == nil {
		fmt.Println(" -> Success!")
	}
	return err
}

// InstallVPKs iterates through the embedded FS and writes all valid VPKs to the game's addons directory
func InstallVPKs(baseDir string, embeddedFS embed.FS) (int, error) {
	addonsDir := filepath.Join(baseDir, appFolder, addonsFolder)
	fmt.Println("\nExtracting Addons (VPKs)...")

	// Ensure addons destination directory exists
	if err := os.MkdirAll(addonsDir, 0755); err != nil {
		return 0, fmt.Errorf("failed to create addons directory: %w", err)
	}

	entries, err := fs.ReadDir(embeddedFS, dataAddonsDir)
	if err != nil {
		return 0, fmt.Errorf("error reading embedded addons: %w", err)
	}

	count := 0
	for _, entry := range entries {
		// Skip directories and non-VPK files
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), vpkExtension) {
			continue
		}

		targetPath := filepath.Join(addonsDir, entry.Name())
		fmt.Printf("Extracting %s...\n", entry.Name())

		// Read from embedded binary
		data, err := embeddedFS.ReadFile(filepath.Join(dataAddonsDir, entry.Name()))
		if err != nil {
			fmt.Printf(" -> Error reading %s: %v\n", entry.Name(), err)
			continue
		}

		// Write to disk
		err = os.WriteFile(targetPath, data, filePerms)
		if err != nil {
			fmt.Printf(" -> Error writing %s: %v\n", entry.Name(), err)
		} else {
			count++
		}
	}

	return count, nil
}

// printCompletion displays the final status
func printCompletion(count int) {
	fmt.Printf("\n===========================================\n")
	fmt.Printf(" Installation Complete! Installed %d addons.\n", count)
	fmt.Printf(" You can now safely unsubscribe from these mods on the Steam Workshop.\n")
	fmt.Printf("===========================================\n")
}

// detectDefaultPath attempts to locate the most common installation directory based on OS
func detectDefaultPath() string {
	if runtime.GOOS == windowsOS {
		return defaultWinPath
	}

	// Check common Linux paths (Native, SteamCMD, Flatpak)
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	paths := []string{
		filepath.Join(home, ".local", "share", "Steam", "steamapps", "common", "Left 4 Dead 2"),
		filepath.Join(home, ".steam", "steam", "steamapps", "common", "Left 4 Dead 2"),
		filepath.Join(home, ".var", "app", "com.valvesoftware.Steam", ".local", "share", "Steam", "steamapps", "common", "Left 4 Dead 2"),
	}

	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}

	return ""
}
