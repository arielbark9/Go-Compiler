package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/arielbark9/Go-Compiler/arithmetic"
	ist "github.com/arielbark9/Go-Compiler/instructions"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
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

		outputFileName := filepath.Join(path, strings.TrimSuffix(vmFile.Name(), ".vm")+".asm")
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			panic("Panicking. Could not open output file: " + outputFile.Name())
		} // not deferring close here because we're inside a loop

		var asmCommands []ist.Instruction
		// scanning the file line after line
		scanner := bufio.NewScanner(currentFile)
		for scanner.Scan() {
			if scanner.Text() != "" {
				currentLineInstructions, err := handleVmLine(scanner.Text())
				if err != nil {
					fmt.Println(err)
					panic("Compilation error in file in line")
				}
				asmCommands = append(asmCommands, currentLineInstructions...)
			}
		}

		// write commands to output asm file
		for _, command := range asmCommands {
			outputFile.WriteString(command.Translate() + "\n")
		}

		// close files
		outputFile.Close()
		currentFile.Close()
	}
}

// handleVmLine handle a line of VM code and return asm instructions
func handleVmLine(text string) ([]ist.Instruction, error) {
	if strings.HasPrefix(text, "//") {
		return []ist.Instruction{}, nil
	}

	var splitInstruction = strings.Split(text, " ")

	if splitInstruction[0] == "push" && splitInstruction[1] == "constant" {
		parameter, _ := strconv.Atoi(splitInstruction[2])
		return arithmetic.PushConstant(parameter), nil
	} else if splitInstruction[0] == "add" {
		return arithmetic.Add(), nil
	} else {
		return nil, errors.New("no matching instruction found")
	}
	// TODO: lt, gt, or, not Achikam
	// TODO: eq, and, neg, sub Ariel
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
