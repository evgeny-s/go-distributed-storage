package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)

	key := "momsspecials"

	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestPathTransformFunc(t *testing.T) {
	key := "momsbetspicture"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "b7f2c4114b9c073dc2c281fb216f66a8595b11a6"
	expectedPathName := "b7f2c/4114b/9c073/dc2c2/81fb2/16f66/a8595/b11a6"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathName)
	}
	if pathKey.Filename != expectedOriginalKey {
		t.Errorf("have %s want %s", pathKey.Filename, expectedOriginalKey)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)

	key := "momsspecials"

	data := []byte("some jpg bytes")

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

	b, _ := ioutil.ReadAll(r)
	if string(b) != string(data) {
		t.Errorf("want %s have %s", data, b)
	}

	s.Delete(key)
}
