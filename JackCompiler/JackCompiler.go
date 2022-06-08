package main

import (
	"JackCompiler/Parser"
	"os"
)

func main() {
	//// region Ex4
	//file, _ := os.ReadFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\ExpressionLessSquare\\Main.jack")
	//s := Parser.ParseToXML(string(file))
	//// save to file
	//ioutil.WriteFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\ExpressionLessSquare\\MainTm.xml", []byte(s), 0644)
	//
	//file, _ = os.ReadFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\ExpressionLessSquare\\Square.jack")
	//s = Parser.ParseToXML(string(file))
	//// save to file
	//ioutil.WriteFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\ExpressionLessSquare\\SquareTm.xml", []byte(s), 0644)
	//
	//file, _ = os.ReadFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\ExpressionLessSquare\\SquareGame.jack")
	//s = Parser.ParseToXML(string(file))
	//// save to file
	//ioutil.WriteFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\ExpressionLessSquare\\SquareGameTm.xml", []byte(s), 0644)
	//
	//file, _ = os.ReadFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\Square\\Main.jack")
	//s = Parser.ParseToXML(string(file))
	//// save to file
	//ioutil.WriteFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\Square\\MainTm.xml", []byte(s), 0644)
	//
	//file, _ = os.ReadFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\Square\\Square.jack")
	//s = Parser.ParseToXML(string(file))
	//// save to file
	//ioutil.WriteFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\Square\\SquareTm.xml", []byte(s), 0644)
	//
	//file, _ = os.ReadFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\Square\\SquareGame.jack")
	//s = Parser.ParseToXML(string(file))
	//// save to file
	//ioutil.WriteFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\Square\\SquareGameTm.xml", []byte(s), 0644)
	//
	//file, _ = os.ReadFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\ArrayTest\\Main.jack")
	//s = Parser.ParseToXML(string(file))
	//// save to file
	//ioutil.WriteFile("C:\\Users\\ariel\\nand2tetris\\projects\\10\\ArrayTest\\MainTm.xml", []byte(s), 0644)
	//// endregion

	// region Ex5
	file, _ := os.Open("C:\\Users\\ariel\\nand2tetris\\projects\\11\\Average\\Main.jack")
	Parser.ParseToVM(file)
	file.Close()

	file, _ = os.Open("C:\\Users\\ariel\\nand2tetris\\projects\\11\\Seven\\Main.jack")
	Parser.ParseToVM(file)
	file.Close()

	file, _ = os.Open("C:\\Users\\ariel\\nand2tetris\\projects\\11\\ConvertToBin\\Main.jack")
	Parser.ParseToVM(file)
	file.Close()

	file, _ = os.Open("C:\\Users\\ariel\\nand2tetris\\projects\\11\\ComplexArrays\\Main.jack")
	Parser.ParseToVM(file)
	file.Close()

	file, _ = os.Open("C:\\Users\\ariel\\nand2tetris\\projects\\11\\Square\\Main.jack")
	Parser.ParseToVM(file)
	file.Close()

	file, _ = os.Open("C:\\Users\\ariel\\nand2tetris\\projects\\11\\Square\\Square.jack")
	Parser.ParseToVM(file)
	file.Close()

	file, _ = os.Open("C:\\Users\\ariel\\nand2tetris\\projects\\11\\Square\\SquareGame.jack")
	Parser.ParseToVM(file)
	file.Close()

	file, _ = os.Open("C:\\Users\\ariel\\nand2tetris\\projects\\11\\Pong\\Main.jack")
	Parser.ParseToVM(file)
	file.Close()

	file, _ = os.Open("C:\\Users\\ariel\\nand2tetris\\projects\\11\\Pong\\Ball.jack")
	Parser.ParseToVM(file)
	file.Close()

	file, _ = os.Open("C:\\Users\\ariel\\nand2tetris\\projects\\11\\Pong\\Bat.jack")
	Parser.ParseToVM(file)
	file.Close()

	file, _ = os.Open("C:\\Users\\ariel\\nand2tetris\\projects\\11\\Pong\\PongGame.jack")
	Parser.ParseToVM(file)
	file.Close()
	// endregion
}
