package main

import (
	"JackCompiler/Tokenizer"
	"fmt"
	"os"
)

func main() {
	// open file C:\Users\ariel\nand2tetris\projects\10\ArrayTest\Main.jack
	file, _ := os.ReadFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\ArrayTest\\Main.jack")

	fmt.Print(Tokenizer.GetXML(string(file)))
}
