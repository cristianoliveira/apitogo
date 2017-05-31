package api

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/cristianoliveira/apitogo/api/json"
	"github.com/cristianoliveira/apitogo/common"
)

func TestWhenHasCollectionFile(t *testing.T) {
	err := copyFile("../posts.json", "./posts.json")
	if err != nil {
		t.Error(err)
	}

	router := Router()
	ts := httptest.NewServer(router)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/posts")
	if err != nil {
		t.Error(err)
	}

	t.Run("Responds with OK status", func(t *testing.T) {
		expected := http.StatusOK
		result := resp.StatusCode

		if expected != result {
			t.Errorf("Expected %s got %s", expected, result)
		}
	})

	t.Run("Provide collection data", func(t *testing.T) {

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		file, err := ioutil.ReadFile("./posts.json")
		if err != nil {
			t.Error(err)
		}

		result := strings.TrimSpace(string(body))
		expected := strings.TrimSpace(string(file))

		if expected == result {
			t.Errorf("Expected %s got %s", expected, result)
		}
	})

	t.Run("Provide json data", func(t *testing.T) {

		resp, err := http.Get(ts.URL + "/posts")
		if err != nil {
			t.Error(err)
		}

		result := resp.Header.Get("Content-Type")
		expected := "application/json; charset=UTF-8"

		if expected != result {
			t.Errorf("Expected %s got %s", expected, result)
		}
	})

	os.Remove("./posts.json")
}

func TestWhenHasNoCollectionFile(t *testing.T) {
	router := Router()
	ts := httptest.NewServer(router)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/postss")
	if err != nil {
		t.Error(err)
	}

	t.Run("Responds with Not Found", func(t *testing.T) {
		expected := http.StatusNotFound
		result := resp.StatusCode

		if expected != result {
			t.Errorf("Expected %d got %d", expected, result)
		}
	})

	t.Run("Responds with json", func(t *testing.T) {

		result := resp.Header.Get("Content-Type")
		expected := "application/json; charset=UTF-8"

		if expected != result {
			t.Errorf("Expected %s got %s", expected, result)
		}
	})

	t.Run("Provide error cause", func(t *testing.T) {
		body, err := ioutil.ReadAll(resp.Body)
		_, err = ioutil.ReadFile(common.Settings().PathFile("postss"))

		result := string(body)
		expected := string(json.NewError(404, err).AsBytes())

		if expected != result {
			t.Errorf("Expected %d got %d", expected, result)
		}
	})
}

func TestUsingWrongParameters(t *testing.T) {
	router := Router()
	ts := httptest.NewServer(router)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/posts/foo")
	if err != nil {
		t.Error(err)
	}

	t.Run("Responds with Bad Request", func(t *testing.T) {
		expected := http.StatusBadRequest
		result := resp.StatusCode

		if expected != result {
			t.Errorf("Expected %d got %d", expected, result)
		}
	})
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}
	return cerr
}
