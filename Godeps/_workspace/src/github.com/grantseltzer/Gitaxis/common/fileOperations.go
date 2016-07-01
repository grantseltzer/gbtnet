package common

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// AppendString appends a string as a new line in a file
func AppendString(newLine string, fileName string) {
	file, openError := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	ErrorCheck(openError)
	bytesWritten, writeError := file.Write([]byte(newLine + "\n"))
	fmt.Println(bytesWritten)
	ErrorCheck(writeError)
}

// RemoveString searches a file for a string and removes the line
func RemoveString(lineToRemove string, fileName string) {
	inputFromFile, inputError := ioutil.ReadFile(fileName)
	ErrorCheck(inputError)

	allLines := strings.Split(string(inputFromFile), "\n")

	for i, line := range allLines {
		if strings.Contains(line, lineToRemove) {
			allLines = append(allLines[:i], allLines[i+1:]...)
		}
	}

	outputToFile := strings.Join(allLines, "\n")
	writeError := ioutil.WriteFile(fileName, []byte(outputToFile), 0600)
	ErrorCheck(writeError)
}

// ReadFirstLine reads in the first line of a file
func ReadFirstLine(fileName string) string {
	infile, inError := os.Open(fileName)
	ErrorCheck(inError)
	defer infile.Close()

	scanner := bufio.NewScanner(infile)
	scanner.Scan()
	text := scanner.Text()
	return text
}

// OverwriteFile overwrites a file with content
func OverwriteFile(filename string, content []byte) {
	fmt.Println("Test 1")
	newFile, createError := os.Create(filename)
	fmt.Println("Test 2")
	ErrorCheck(createError)
	defer newFile.Close()

	bytesWritten, writeError := newFile.Write(content)
	ErrorCheck(writeError)
	fmt.Println(bytesWritten)
}

// DeleteFile ...
func DeleteFile(filePath string) {
	removeError := os.Remove(filePath)
	ErrorCheck(removeError)
}
