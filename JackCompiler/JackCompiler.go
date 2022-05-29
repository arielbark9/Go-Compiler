package main

import (
	"JackCompiler/Parser"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := os.ReadFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\Square\\Main.jack")
	s := Parser.GetXML(string(file))
	// save to file
	ioutil.WriteFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\Square\\MainTm.xml", []byte(s), 0644)

}
