package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)


func ZipFile(sourceDir string, zipFileName string) {

	// Create a new zip file
	zipFile, err := os.Create(zipFileName + ".zip")
	if err != nil {
		fmt.Println("Error creating zip file:", err)
		return
	}
	defer zipFile.Close()

	// Create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk through the source directory and add files to the zip archive
	err = filepath.Walk(sourceDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Open the file
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// Create a zip entry for the file
		zipEntry, err := zipWriter.Create(filePath[len(sourceDir):])
		if err != nil {
			return err
		}

		// Copy the file content to the zip entry
		_, err = io.Copy(zipEntry, file)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking through directory:", err)
		return
	}

	fmt.Println("Zip file created successfully:", zipFileName)
}