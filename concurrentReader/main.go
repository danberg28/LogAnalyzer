package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

func readFromInput(wg *sync.WaitGroup, bytesChannel, linesProcessedChannel chan int) {
	// take in reading data
	// pipe in the lines to linesprocessedchannel
	// pipe the bytes in each line into byteschannel
	buf := make([]byte, 32*1024)
	count := 0
	// will seperate when there is a newline character
	lineSep := []byte{'\n'}
	byteCount := 0
	// work in progress
	//get streaming data into channels
	for line := range scan(os.Stdin) {
		c, err := Read(line)
		if err != nil {
			fmt.Println(err)
			wg.Done()
			return
		}
		bytesChannel <- c
		// counts the bytes between the line seperations
		linesProcessedChannel <- bytes.Count(buf[:c], lineSep)
	}

}

func writeReport(wg *sync.WaitGroup, startTime time.Time, bytesChannel, linesProcessedChannel chan int) {
	// since sharing the channel prints the bytes and lines at the end.
	numBytesCount := 0
	linesCount := 0
	var numBytes64 int64

	defer wg.Done()
	for numBytes := range bytesChannel {
		numBytesCount += numBytes
	}
	for lines := range linesProcessedChannel {
		linesCount += lines
	}
	numBytes64 = int64(numBytesCount)
	d := time.Duration(time.Since(startTime)) * time.Millisecond
	ms := int64(d / time.Millisecond)
	fmt.Println("Total bytes:", numBytesCount)
	fmt.Println("Time to process:", time.Since(startTime))
	fmt.Println("Throughput:", (numBytes64 / (ms / 1000)), "Bytes/s")
	fmt.Println("Lines processed::", linesCount)

}

func main() {

	var wg sync.WaitGroup
	bytesChannel := make(chan int)
	linesProcessedChannel := make(chan int)
	start := time.Now()
	wg.Add(1)
	go readFromInput(&wg, bytesChannel, linesProcessedChannel)
	wg.Add(1)
	go writeReport(&wg, start, bytesChannel, linesProcessedChannel)
	wg.Wait()
	fmt.Println("done")
}
