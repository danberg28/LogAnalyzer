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

	for {
		read_from_std_in
		c := Read(buf)
		byteCount += c
		// counts the bytes between the line seperations
		count += bytes.Count(buf[:c], lineSep)
	}

	// filescanner := io.Reader(os.Stdin)

	fmt.Println(os.Stdin)
	// for liness := range linesProcessedChannel {
	// 	fmt.Println(liness)
	wg.Done()
	// }

}

func writeReport(wg *sync.WaitGroup, startTime time.Time, bytesChannel, linesProcessedChannel chan int) {
	// since sharing the channel prints the bytes and lines at the end.
	defer wg.Done()
	fmt.Println("dibne")
	for bytess := range bytesChannel {
		fmt.Println(bytess)
	}

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

	linesProcessedChannel <- count
	bytesChannel <- 5
	linesProcessedChannel <- 3
	bytesChannel <- 4

	wg.Wait()
	fmt.Println("done")
}
