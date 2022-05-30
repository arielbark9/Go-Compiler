package main

import (
	"JackCompiler/Parser"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := os.ReadFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\Square\\Main.jack")
	s := Parser.ParseToXML(string(file))
	// save to file
	ioutil.WriteFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\Square\\MainTm.xml", []byte(s), 0644)

	file, _ = os.ReadFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\Square\\Square.jack")
	s = Parser.ParseToXML(string(file))
	// save to file
	ioutil.WriteFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\Square\\SquareTm.xml", []byte(s), 0644)

	file, _ = os.ReadFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\Square\\SquareGame.jack")
	s = Parser.ParseToXML(string(file))
	// save to file
	ioutil.WriteFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\Square\\SquareGameTm.xml", []byte(s), 0644)

}
