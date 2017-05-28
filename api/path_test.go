package api

import "testing"

import "path/filepath"

func TestItComposeTheFilePath(t *testing.T) {
	expected, _ := filepath.Abs("./path/to/file/name.json")
  result := PathFile("path/to/file", "name")

  if expected != result {
    t.Errorf("Expected %d got %d", expected, result)
  }
}
