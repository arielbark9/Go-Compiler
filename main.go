package main

import (
	"bufio"
	"fmt"
	ist "github.com/arielbark9/Go-Compiler/instructions"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter absolute path of directory: ")
	path, _ := reader.ReadString('\n')
	path = path[:len(path)-2] // remove \r\n from path

	// get all files in directory
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("error opening directory path")
		return
	}

	vmFiles := extractFormatFiles(files, "vm")

	// Foreach file in the list
	for _, vmFile := range vmFiles {
		// open the file using its name
		currentFile, err := os.Open(filepath.Join(path, vmFile.Name()))
		if err != nil {
			panic("Panicking. Could not open file: " + vmFile.Name())
		} // not deferring close here because we're inside a loop

		wd, _ := os.Getwd()
		outputFileName := wd + strings.TrimSuffix(vmFile.Name(), ".vm") + ".asm"
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			panic("Panicking. Could not open output file: " + outputFile.Name())
		} // not deferring close here because we're inside a loop

		// scanning the file line after line
		scanner := bufio.NewScanner(currentFile)
		for scanner.Scan() {
			var asmCommands []ist.Instruction
			asmCommands = append(asmCommands, handleVmLine(scanner.Text())...)
		}
		outputFile.Close()
		currentFile.Close()
	}
}

func handleVmLine(text string) []ist.Instruction {
	var res []ist.Instruction
	if strings.HasPrefix(text, "//") {
		return []ist.Instruction{}
	}
	return res
}

// extractFormatFiles extract some format of files from list of files
func extractFormatFiles(files []fs.FileInfo, fmt string) []fs.FileInfo {
	var vmFiles []fs.FileInfo
	for _, file := range files {
		if strings.HasSuffix(file.Name(), "."+fmt) {
			vmFiles = append(vmFiles, file)
		}
	}
	return vmFiles
}
