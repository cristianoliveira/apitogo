package api

import (
  "encoding/json"
	"io/ioutil"
  "path/filepath"
)

type Collection struct {
  path string
  data map[string]interface{}
}

func (c *Collection) Name() string {
  extension := filepath.Ext(c.path)
  base := filepath.Base(c.path)
  return base[0:len(base)-len(extension)]
}

func (c *Collection) AsJson() (map[string]interface{}) {
  return c.data
}

func (c *Collection) AsBytes() ([]byte, error) {
  return json.Marshal(c.data)
}

func (c *Collection) Get(key string) (interface{}) {
  data := c.data[key]
  return data
}

func (c *Collection) GetAsList(key string) ([]interface{}) {
  return c.data[key].([]interface{})
}

func (c *Collection) GetById(id float64) *Collection {
	for i := range c.GetAsList("data") {
    item := items[i].(map[string]interface{})
		if item["id"] == id {
      return &Collection {
        path: c.path,
        data: items[i].(map[string]interface{}),
      }
		}
	}

  return &Collection { path: c.path, data: nil }
}

func CollectionLoad(path string) (*Collection, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)

  return &Collection { path: path, data: jsonData }, err
}
