package main

import (
	"bytes"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpictureshaiye"
	pathkey := CASPathTransformFunc(key)
	expectedOriginalKey := "97c1910a9d07e491e3a8c5e8bac04d5c41f6fd7d"
	expectedPathName := "97c19/10a9d/07e49/1e3a8/c5e8b/ac04d/5c41f/6fd7d"
	if pathkey.Pathname != expectedPathName {
		t.Errorf("Expected %s, got %s", expectedPathName, pathkey.Pathname)
	}
	if pathkey.Filename != expectedOriginalKey {
		t.Errorf("Expected %s, got %s", expectedOriginalKey, pathkey.Filename)
	}
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)
	key := "specialpicture"
	data := []byte("some data")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)
	key := "specialpicture"
	data := []byte("some data")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if ok := s.Has(key); !ok {
		t.Errorf("Expected key %s to exist", key)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(r)
	if string(b) != string(data) {
		t.Errorf("Expected %s, got %s", data, b)
	}

	s.Delete(key)
}