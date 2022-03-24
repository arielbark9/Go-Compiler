// Achikam Smila 213796071
// Ariel Bar Kalifa 214181604
// Group no. 5782.41
package main

import (
	"bufio"
	"errors"
	. "github.com/arielbark9/Go-Compiler/arithmetic"
	. "github.com/arielbark9/Go-Compiler/instructions"
	. "github.com/arielbark9/Go-Compiler/logical"
	. "github.com/arielbark9/Go-Compiler/memory"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Exactly one variable should be passed.\n" +
			"use like this:\n" +
			"VMtranslator /path/to/dir/of/vm-files")
	}

	path := os.Args[1]

	// get all files in directory
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal("error opening directory path")
	}

	vmFiles := extractFormatFiles(files, "vm")

	outputFileName := filepath.Join(path, filepath.Base(path)+".asm")
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("Fatal Error. Could not open output file: %s", outputFile.Name())
	}
	defer outputFile.Close()

	// foreach file in the list
	for _, vmFile := range vmFiles {
		// open the file using its name
		currentFile, err := os.Open(filepath.Join(path, vmFile.Name()))
		if err != nil {
			log.Fatalf("fatal error. Could not open file: %s", vmFile.Name())
		} // not deferring close here because we're inside a loop

		var asmCommands []Instruction
		// scanning the file line after line
		line := 1
		scanner := bufio.NewScanner(currentFile)
		for scanner.Scan() {
			if scanner.Text() != "" { // not an empty line
				currentLineInstructions, err := handleVmLine(scanner.Text(), strings.TrimSuffix(vmFile.Name(), ".vm"))
				if err != nil {
					log.Fatalf("Compilation error in file %s in line %d\n"+
						"%s\n "+
						"%s", currentFile.Name(), line, err.Error(), scanner.Text())
				}
				asmCommands = append(asmCommands, currentLineInstructions...)
				line++
			}
		}

		// write commands to output asm file
		for _, command := range asmCommands {
			outputFile.WriteString(command.Translate() + "\n")
		}

		currentFile.Close()
	}
}

// handleVmLine handle a line of VM code and return asm instructions
func handleVmLine(text string, fileName string) ([]Instruction, error) {
	if strings.HasPrefix(text, "//") {
		return []Instruction{}, nil
	}
	res := []Instruction{Comment{Text: text}}

	var splitInstruction = strings.Split(text, " ")
	switch splitInstruction[0] {
	case "push":
		parameter, _ := strconv.Atoi(splitInstruction[2])
		switch splitInstruction[1] {
		case "constant":
			res = append(res, PushConstant(parameter)...)
		case "local":
			res = append(res, PushLocal(parameter)...)
		case "argument":
			res = append(res, PushArgument(parameter)...)
		case "this":
			res = append(res, PushThis(parameter)...)
		case "that":
			res = append(res, PushThat(parameter)...)
		case "temp":
			res = append(res, PushTemp(parameter)...)
		case "static":
			res = append(res, PushStatic(parameter, fileName)...)
		case "pointer":
			switch splitInstruction[2] {
			case "0":
				res = append(res, PushPointer0Set...)
			case "1":
				res = append(res, PushPointer1Set...)
			default:
				return nil, errors.New("no matching instruction found")
			}
		default:
			return nil, errors.New("no matching instruction found")
		}
	case "pop":
		parameter, _ := strconv.Atoi(splitInstruction[2])
		switch splitInstruction[1] {
		case "local":
			res = append(res, PopLocal(parameter)...)
		case "argument":
			res = append(res, PopArgument(parameter)...)
		case "this":
			res = append(res, PopThis(parameter)...)
		case "that":
			res = append(res, PopThat(parameter)...)
		case "temp":
			res = append(res, PopTemp(parameter)...)
		case "static":
			res = append(res, PopStatic(parameter, fileName)...)
		case "pointer":
			switch splitInstruction[2] {
			case "0":
				res = append(res, PopPointer0Set...)
			case "1":
				res = append(res, PopPointer1Set...)
			default:
				return nil, errors.New("no matching instruction found")
			}
		default:
			return nil, errors.New("no matching instruction found")
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
