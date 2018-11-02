package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// from https://stackoverflow.com/questions/46365221/fill-os-stdin-for-function-that-reads-from-it
func TestReadStream(t *testing.T) {
	// creates a temp fil which will be filled with content
	content := []byte("header\ncat")

	tmpfile, err := ioutil.TempFile("", "example")
	// if an error occurs will log the error and stop the test
	if err != nil {
		log.Fatal(err)
	}
	// clean up
	defer os.Remove(tmpfile.Name())
	// if there is an error writing the temp fil will log error
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	// will seek the start of the temp file or log error if cant be found
	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}
	// if there is an Stdin already in progress will defer until temp is complete
	// then restore it.
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	// Creates new Stdin with the tempfile
	os.Stdin = tmpfile
	lines, numBytes, err := readStream()
	if err != nil {
		t.Errorf("readStream failed: %v", err)
	}
	if lines != 1 {
		t.Error("There was an error, expected 1 line but had", lines)
	}
	if numBytes != 10 {
		t.Error("There was an error, expected 10 bytes but had", numBytes)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

}
func TestLineCounter(t *testing.T) {
	// creates a byte array with some strings on new lines
	testByteStream := []byte("header\nThis\nis\na\ntest")
	// creates new reader with the testBytseStream data
	testRead := bytes.NewReader(testByteStream)
	// tests the linecounter to make sure they are correct for the test data
	count, byteCount, err := lineCounter(testRead)
	if count != 4 {
		t.Error("There was an error, expected 4 lines but received", count)
	}
	if byteCount != 21 {
		t.Error("There was an error, expected 21 bytes but received", byteCount)
	}
	if err != nil {
		t.Error("There was an error", err)
	}
}
