package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func packageZip(fileList []string, zipName string) error {
	zipFile, err := os.Create(zipName)
	if err != nil {
		return fmt.Errorf("create zip file: %w", err)
	}
	defer zipFile.Close()

	writer := zip.NewWriter(zipFile)
	defer writer.Close()

	for _, fileName := range fileList {
		if err := writeToZip(writer, fileName); err != nil {
			return fmt.Errorf("add %q to zip: %w", fileName, err)
		}
	}
	return nil
}

func writeToZip(zw *zip.Writer, srcPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("open source file: %w", err)
	}
	defer srcFile.Close()

	// Use filepath.Base to get just the filename
	zipEntry, err := zw.Create(filepath.Base(srcPath))
	if err != nil {
		return fmt.Errorf("create zip entry: %w", err)
	}

	_, err = io.Copy(zipEntry, srcFile)
	if err != nil {
		return fmt.Errorf("copy file content: %w", err)
	}
	return nil
}

func unpackZip(zipPath, targetDir string) error {
	zr, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("open zip file: %w", err)
	}
	defer zr.Close()

	targetDir = filepath.Clean(targetDir)
	base := targetDir + string(os.PathSeparator)

	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return fmt.Errorf("create target directory: %w", err)
	}

	for _, f := range zr.File {
		// Validate/sanitize entry name
		name := filepath.Clean(f.Name)
		if !filepath.IsLocal(name) {
			return fmt.Errorf("invalid entry path: %q", f.Name)
		}

		// Skip empty names
		if name == "" || name == "." {
			continue
		}

		outPath := filepath.Join(targetDir, name)
		outPath = filepath.Clean(outPath)

		// Prevent zip-slip attacks
		if !strings.HasPrefix(outPath, base) {
			return fmt.Errorf("path escapes target dir: %q", f.Name)
		}

		// Create directories
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(outPath, 0o755); err != nil {
				return fmt.Errorf("create directory %q: %w", outPath, err)
			}
			continue
		}

		// Ensure parent directory exists
		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			return fmt.Errorf("create parent directory for %q: %w", outPath, err)
		}

		rc, err := f.Open()
		if err != nil {
			return fmt.Errorf("open zip entry %q: %w", f.Name, err)
		}

		// Use os.O_EXCL to prevent overwriting existing files unintentionally
		outFile, err := os.OpenFile(outPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, f.Mode())
		if err != nil {
			rc.Close()
			return fmt.Errorf("create output file %q: %w", outPath, err)
		}

		_, copyErr := io.Copy(outFile, rc)

		// Close resources
		closeErr1 := outFile.Close()
		closeErr2 := rc.Close()

		// Prioritize copy error, then closing errors
		if copyErr != nil {
			return fmt.Errorf("copy file content for %q: %w", outPath, copyErr)
		}
		if closeErr1 != nil {
			return fmt.Errorf("close output file %q: %w", outPath, closeErr1)
		}
		if closeErr2 != nil {
			return fmt.Errorf("close zip entry %q: %w", f.Name, closeErr2)
		}
	}

	return nil
}

func main() {
	// Generate test content
	documents := []string{"report.txt", "summary.txt"}
	for i, doc := range documents {
		content := fmt.Sprintf("Report section %d\nAnalysis results here", i+1)
		if err := os.WriteFile(doc, []byte(content), 0644); err != nil {
			log.Fatal(err)
		}
		defer os.Remove(doc) // Clean up test files
	}

	// Package and unpack
	zipName := "reports.zip"
	extractDir := "unpacked"

	if err := packageZip(documents, zipName); err != nil {
		log.Fatalf("Failed to create zip: %v", err)
	}
	defer os.Remove(zipName) // Clean up zip file

	if err := unpackZip(zipName, extractDir); err != nil {
		log.Fatalf("Failed to extract zip: %v", err)
	}
	defer os.RemoveAll(extractDir) // Clean up extracted directory

	fmt.Println("ZIP packaging completed successfully")
}