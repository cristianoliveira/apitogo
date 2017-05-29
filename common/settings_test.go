package common

import "testing"

import "path/filepath"

func TestItComposeTheFilePath(t *testing.T) {
  args := map[string]interface{}{
    "-d": "path/to/file",
  }

  settings := Settings()
  settings.UpdateByArgs(args)

	expected, _ := filepath.Abs("./path/to/file/name.json")
  result := settings.PathFile("name")

  if expected != result {
    t.Errorf("Expected %d got %d", expected, result)
  }
}

func TestUpdateByArgs(t *testing.T) {
  port := "5000"
  dir := "somedir"

  args := map[string]interface{}{ "-p": port, "-d": dir }
  settings := Settings()
  settings.UpdateByArgs(args)

  if settings.Dir != dir {
    t.Error("Expected directory %s got %s", dir, settings.Dir)
  }

  if settings.Port != port {
    t.Error("Expected port %s got %s", dir, settings.Port)
  }
}
