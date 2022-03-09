package main

import (
	"bufio"
	"errors"
	"fmt"
	. "github.com/arielbark9/Go-Compiler/arithmetic"
	. "github.com/arielbark9/Go-Compiler/instructions"
	. "github.com/arielbark9/Go-Compiler/logical"
	. "github.com/arielbark9/Go-Compiler/memory"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		panic("panicking. Exactly one variable should be passed.\n" +
			"use like this:\n" +
			"VMtranslator /path/to/dir/of/vm-files")
	}

	path := os.Args[1]

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

		var asmCommands []Instruction
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
func handleVmLine(text string) ([]Instruction, error) {
	if strings.HasPrefix(text, "//") {
		return []Instruction{}, nil
	}
	res := []Instruction{Comment{Text: text}}
	var splitInstruction = strings.Split(text, " ")

	//if splitInstruction[0] == "push" && splitInstruction[1] == "constant" {
	//	parameter, _ := strconv.Atoi(splitInstruction[2])
	//	res = append(res, PushConstant(parameter)...)
	//} else if splitInstruction[0] == "add" {
	//	res = append(res, Add()...)
	//} else {
	//	return nil, errors.New("no matching instruction found")
	//}
	switch splitInstruction[0] {
	case "push":
		switch splitInstruction[1] {
		case "constant":
			parameter, _ := strconv.Atoi(splitInstruction[2])
			res = append(res, PushConstant(parameter)...)
		}
	case "add":
		res = append(res, AddSet...)
	case "and":
		res = append(res, AndSet...)
	case "neg":
		res = append(res, NegSet...)
	case "sub":
		res = append(res, SubSet...)
	case "or":
		res = append(res, OrSet...)
	case "not":
		res = append(res, NotSet...)
	case "eq":
		res = append(res, Eq()...)
	case "lt":
		res = append(res, Lt()...)
	case "gt":
		res = append(res, Gt()...)
	default:
		return nil, errors.New("no matching instruction found")
	}

	return res, nil
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
