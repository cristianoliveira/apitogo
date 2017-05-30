package common

import "path/filepath"
import "sync"

type Setting struct {
	Port string
	Dir string
}

var instance *Setting
var once sync.Once

func (s *Setting) PathFile(filename string) string {
	path, err := filepath.Abs(s.Dir)
	if err != nil {
		return ""
	}

  return path + "/" + filename + ".json"
}

func initSettings() {
  instance = &Setting {
    Port: "8080",
    Dir:  "./",
  }
}

func Settings() *Setting {
  once.Do(initSettings)
  return instance
}
