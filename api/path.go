package api

import "path/filepath"

func PathFile(directory string, filename string) string {
	path, err := filepath.Abs(directory)
	if err != nil {
		return ""
	}

  return path + "/" + filename + ".json"
}
