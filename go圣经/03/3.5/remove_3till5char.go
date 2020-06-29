package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("/tmp/test")
	outputFile, _ := os.OpenFile("/tmp/goprogramT", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)

	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF???")
			f , _ := os.Open("/tmp/goprogramT")
			defer f.Close()

			scanner := bufio.NewScanner(f)
			for scanner.Scan(){
				fmt.Println(scanner.Text())
			}
			return
		}
		outputString := string(inputString[2:5]) + "\r\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
		outputWriter.Flush()
	}
	fmt.Println("Conversion done")


}