package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestStore(t *testing.T) {
	options := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(options)
	data := []byte("some jpg bytes")
	key := "mynewfile"

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if ok := s.Has(key); !ok {
		t.Errorf("expected to have key %s", key)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(r)

	fmt.Println(string(b))

	if string(b) != string(data) {
		t.Errorf("want %s have %s", data, b)
	}

	s.Delete(key)
}

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpictures"
	pathKey := CASPathTransformFunc(key)
	expectedPathName := ""
	expectedOriginalKey := ""

	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathName)
	}

	if pathKey.FileName != expectedOriginalKey {
		t.Errorf("have %s want %s", pathKey.FileName, expectedOriginalKey)
	}

	fmt.Println(pathKey)
}

func TestStoreDeleteKey(t *testing.T) {
	options := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(options)
	data := []byte("some jpg bytes")
	key := "mynewfile"

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}
