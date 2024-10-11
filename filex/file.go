package filex

import (
	"os"
	"path/filepath"
)

// LoadFilesWithFilter load files with filter function func(filePath string) bool
func LoadFilesWithFilter(filePath string, filter func(filePath string) bool) ([]string, error) {
	files := make([]string, 0)
	err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filter(path) {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
