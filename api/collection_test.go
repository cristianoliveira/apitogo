package api

import "testing"

func TestCreatingWithWrongPath(t *testing.T) {
  _, err := CollectionLoad("./randon")

  if err == nil {
    t.Error("When has no file must return error")
  }
}

func TestWhenHasFile(t *testing.T) {
  collection, err := CollectionLoad("../examples.json")
  if err != nil {
    t.Error("it must not contain error")
  }


  t.Run("Its name is the same as filename", func(tt *testing.T) {
    expected := "examples"
    result := collection.Name()

    if expected != result {
      tt.Errorf("Expected %s got %s", expected, result)
    }
  })

  t.Run("Its contains collection with the same name", func(tt *testing.T) {
    data := collection.Get("examples")

    if data == nil {
      t.Error("It must contain data")
    }
  })

  t.Run("Its contains more than one items", func(tt *testing.T) {
    items := collection.GetAsList("examples")

    expected := 4
    result := len(items)

    if expected != result {
      tt.Errorf("Expected %d got %d", expected, result)
    }
  })

  t.Run("Its contains more than one items", func(tt *testing.T) {
    items := collection.GetAsList("examples")

    expected := 4
    result := len(items)

    if expected != result {
      tt.Errorf("Expected %d got %d", expected, result)
    }
  })

  t.Run("Its expects that collection 1 has name foo", func(tt *testing.T) {
    item := collection.GetById(1)

    expected := "foo"
    result := item.Get("name")

    if expected != result {
      tt.Errorf("Expected %d got %d", expected, result)
    }
  })

  t.Run("Its expects that collection 2 has name foo bar", func(tt *testing.T) {
    item := collection.GetById(2)

    expected := "foo bar"
    result := item.Get("name")

    if expected != result {
      tt.Errorf("Expected %s got %s", expected, result)
    }
  })
}
