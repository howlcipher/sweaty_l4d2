package main

import (
	"os"
	"path/filepath"
	"testing"
)

// TestInstallCfg verifies that the configuration file is correctly written
// to the expected directory structure (left4dead2/cfg/autoexec.cfg)
func TestInstallCfg(t *testing.T) {
	// Create an isolated temporary directory for testing
	tempDir := t.TempDir()

	testPayload := []byte("rate 100000\ncl_cmdrate 100")

	// Run the function
	err := InstallCfg(tempDir, testPayload)
	if err != nil {
		t.Fatalf("InstallCfg failed unexpectedly: %v", err)
	}

	// Verify the file was created in the correct location
	expectedPath := filepath.Join(tempDir, appFolder, cfgFolder, cfgFileName)
	content, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("Failed to read expected cfg file at %s: %v", expectedPath, err)
	}

	// Verify the content matches
	if string(content) != string(testPayload) {
		t.Errorf("Expected content %q, got %q", string(testPayload), string(content))
	}
}

// TestInstallVPKs verifies that VPK files are properly extracted from the embed.FS
// Note: This uses the actual embedded filesystem in main.go
func TestInstallVPKs(t *testing.T) {
	tempDir := t.TempDir()

	// Run extraction
	count, err := InstallVPKs(tempDir, addonsFS)
	if err != nil {
		t.Fatalf("InstallVPKs failed unexpectedly: %v", err)
	}

	// Ensure at least some addons were extracted (since we know the real repo has them)
	// If the real repo is empty, count could be 0, so we check if it matches the directory contents.
	expectedAddonsDir := filepath.Join(tempDir, appFolder, addonsFolder)
	
	entries, err := os.ReadDir(expectedAddonsDir)
	if err != nil {
		t.Fatalf("Failed to read expected addons directory at %s: %v", expectedAddonsDir, err)
	}

	// Verify the number of files matches the returned count
	actualCount := 0
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == vpkExtension {
			actualCount++
		}
	}

	if actualCount != count {
		t.Errorf("InstallVPKs returned count %d, but found %d files on disk", count, actualCount)
	}
}

// TestDetectDefaultPath ensures the path detection doesn't panic
// and returns a sensible string.
func TestDetectDefaultPath(t *testing.T) {
	path := detectDefaultPath()
	
	// Since CI environments or arbitrary developer machines might not have L4D2 installed,
	// path can validly be an empty string. We just want to ensure it executes cleanly.
	if path != "" && !filepath.IsAbs(path) {
		t.Errorf("Expected an absolute path or empty string, got: %s", path)
	}
}
