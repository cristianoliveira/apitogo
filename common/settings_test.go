package common

import "testing"

import "path/filepath"

func TestItComposeTheFilePath(t *testing.T) {
  settings := Settings()
  settings.Dir = "path/to/file"

	expected, _ := filepath.Abs("./path/to/file/name.json")
  result := settings.PathFile("name")

  if expected != result {
    t.Errorf("Expected %d got %d", expected, result)
  }
}
