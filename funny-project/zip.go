package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func packageZip(fileList []string, zipName string) error {
	zipFile, err := os.Create(zipName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	writer := zip.NewWriter(zipFile)
	defer writer.Close()

	for _, fileName := range fileList {
		if err := writeToZip(writer, fileName); err != nil {
			return err
		}
	}
	return nil
}

func writeToZip(zw *zip.Writer, srcPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create compressed entry.
	zipEntry, err := zw.Create(filepath.Base(srcPath))
	if err != nil {
		return err
	}

	_, err = io.Copy(zipEntry, srcFile)
	return err
}

func unpackZip(zipPath, targetDir string) error {
	zr, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer zr.Close()

	// Clean once; used for prefix check.
	targetDir = filepath.Clean(targetDir)
	base := targetDir + string(os.PathSeparator)

	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return err
	}

	for _, f := range zr.File {
		// Validate/sanitize entry name.
		name := filepath.Clean(f.Name)
		if !filepath.IsLocal(name) {
			return fmt.Errorf("invalid entry path: %q", f.Name)
		}

		outPath := filepath.Join(targetDir, name)
		outPath = filepath.Clean(outPath)

		// Ensure the final path is inside targetDir (prevents zip-slip).
		if outPath != targetDir && !strings.HasPrefix(outPath, base) {
			return fmt.Errorf("path escapes target dir: %q", f.Name)
		}

		// Directories.
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(outPath, 0o755); err != nil {
				return err
			}
			continue
		}

		// Ensure parent dirs exist.
		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		outFile, err := os.OpenFile(outPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			rc.Close()
			return err
		}

		_, copyErr := io.Copy(outFile, rc)

		// Close in the right order; preserve copy error if any.
		closeErr1 := outFile.Close()
		closeErr2 := rc.Close()

		if copyErr != nil {
			return copyErr
		}
		if closeErr1 != nil {
			return closeErr1
		}
		if closeErr2 != nil {
			return closeErr2
		}
	}

	return nil
}

func main() {
	// Generate test content.
	documents := []string{"report.txt", "summary.txt"}
	for i, doc := range documents {
		content := fmt.Sprintf("Report section %d\nAnalysis results here", i+1)
		if err := os.WriteFile(doc, []byte(content), 0644); err != nil {
			log.Fatal(err)
		}
	}

	// Package and unpack.
	packageZip(documents, "reports.zip")
	unpackZip("reports.zip", "unpacked/")
	fmt.Println("ZIP packaging completed")
}
