// File : wordCount.go
// File name is the first command line argument
// Search string is the second command line argument
//
// Compile this program using :-
// 	go build wordCount.go
//
// Run this program using :-
// 	./wordCount <file_name> <search_string>

package main

import (
	"os"
	"fmt"
	"log"
	"runtime"
	"sync"
	"bufio"
)

func main() {

	args := os.Args

	file := args[1]
	text := args[2]


// runtime primitive NumCPU returns the number of cores that can be used by the program
	num_threads := runtime.NumCPU()

// channel d to see if the thread has finished working
// channel c to send the lines one by one to the threads
	d = make(chan string)
	c = make(chan string)

// count variable to keep track of lines with occurrence of string
	var count = 0

// Mutex to allow safe access of commonly used count varibale : "count"
	var mutex = &sync.Mutex{}

// open the inpute file specified by the argument "file" and read the file line by line
	inputFile, err := os.Open(file)
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}

// close the file when we leave the scope of this function
	defer inputFile.Close()

	bf := bufio.NewReader(inputFile)

	for {
		line, isPrefix, err := bf.ReadLine()
		if err == io.EOF {
			close(c)
			break
		}
		if err != nil {
			log.Fatal("Error reading line:", err)
		}
		if isPrefix {
			log.Fatal("Error: Unexpected long line reading", inputFile.Name())
		}

		c <- line
	}

	for i := 0; i < num_threads; i++ {
		go func() {
			for {
// read from the channel that feeds in lines from the main thread 
				mutex.Lock()
				line, ok <- c
				mutex.Unlock()

// check if the channel is empty, if so ... output "Done"
				if !ok {
					d <- "Done"
				} else {
// process the string to check if it contains a valid string
					if Contains(line, text) {
						mutex.Lock()
						count++
						mutex.Unlock()
					}
				}
			}

		} ()

	}

// check if every process has finished
	for i := 0; i < num_threads; i++ {
		msg <- d
		if msg != "Done" {
			log.Fatal("Error: Incorrect message read", msg)
		}
	}

	fmt.Printf("%d\n", count)

}
