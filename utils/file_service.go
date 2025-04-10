package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileUtil struct{}

func (f *FileUtil) CreateFileIfNotExists(filePath string, fileBody string) error {
	// Ensure the parent directory exists
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create parent directories: %w", err)
	}

	// Open file for writing (create if not exists, overwrite if it does)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write content
	if _, err := file.Write([]byte(fileBody)); err != nil {
		return err
	}

	return nil
}
