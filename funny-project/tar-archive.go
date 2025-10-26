package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func createTar(sourceFiles []string, outputPath string) error {
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	tarWriter := tar.NewWriter(outFile)
	defer tarWriter.Close()

	for _, srcFile := range sourceFiles {
		if err := addToTar(tarWriter, srcFile); err != nil {
			return err
		}
	}
	return nil
}

func addToTar(tw *tar.Writer, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	hdr := &tar.Header{
		Name: filepath.Base(filePath),
		Mode: 0644,
		Size: stat.Size(),
	}

	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}
	_, err = io.Copy(tw, file)
	return err
}

func extractTar(tarPath, destDir string) error {
	file, err := os.Open(tarPath)
	if err != nil {
		return err
	}
	defer file.Close()

	tarReader := tar.NewReader(file)
	for {
		hdr, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Sanitize archive path to prevent directory traversal (zip-slip).
		cleanName := filepath.Clean(hdr.Name)

		if !filepath.IsLocal(cleanName) {
			return fmt.Errorf("invalid file path: %s", hdr.Name)
		}

		outPath := filepath.Join(destDir, cleanName)
		if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
			return err
		}

		outFile, err := os.Create(outPath)
		if err != nil {
			return err
		}
		if _, err := io.Copy(outFile, tarReader); err != nil {
			return err
		}
		if err := outFile.Close(); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// Create sample files.
	testFiles := []string{"notes.txt", "config.txt"}
	for i, name := range testFiles {
		if err := os.WriteFile(name, []byte(fmt.Sprintf("Sample data %d", i+1)), 0644); err != nil {
			log.Fatal(err)
		}
	}

	// Create and extract TAR.
	createTar(testFiles, "backup.tar")
	extractTar("backup.tar", "restored/")
	fmt.Println("TAR operations completed")
}
