package pkg

import (
	"os"
	"path/filepath"
)

func ListFiles(dir string) ([]string, error) {
	var result []string

	d, err := os.Open(dir)

	if err != nil {
		return nil, err
	}
	defer d.Close()

	files, err := d.Readdir(-1)

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			newDir := filepath.Join(dir, file.Name())
			files, err := ListFiles(newDir)
			if err != nil {
				continue
			}
			result = append(result, files...)
			continue
		}
		result = append(result, filepath.Join(dir, file.Name()))
	}

	return result, nil
}
