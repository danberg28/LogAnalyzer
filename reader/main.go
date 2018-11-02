package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

// readStream reads the slices of bytes from lineCounter and
// returns lines and how many bytes are in it.
func readStream() (int, int, error) {

	fileScanner := io.Reader(os.Stdin)
	lines, numBytes, err := lineCounter(fileScanner)
	// if there is an error will show how far it read through before error and prints error
	if err != nil {
		fmt.Println(err)
		return lines, numBytes, err
	}
	return lines, numBytes, nil

}

// lineCounter takes in the streaming data and counts the lines
func lineCounter(r io.Reader) (int, int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	// will seperate when there is a newline character
	lineSep := []byte{'\n'}
	byteCount := 0
	for {
		c, err := r.Read(buf)
		byteCount += c
		// counts the bytes between the line seperations
		count += bytes.Count(buf[:c], lineSep)

		if err == io.EOF {
			return count, byteCount, nil
		} else if err != nil {
			return count, byteCount, err
		}
	}
}

func main() {
	start := time.Now()
	lines, numBytes, err := readStream()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Lines processed::", lines)
	fmt.Println("Total bytes:", numBytes)
	d := time.Duration(time.Since(start)) * time.Millisecond
	ms := int64(d / time.Millisecond)
	// converts from int to int64 so we can get throughput
	// time is in int64 and needed bytes to match
	var numBytes64 int64
	numBytes64 = int64(numBytes)
	fmt.Println("Time to process:", time.Since(start))
	fmt.Println("Throughput:", (numBytes64 / (ms / 1000)), "Bytes/s")

}
