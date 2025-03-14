package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("Hello.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString("Hello 📚world!!!\n")
	//file.Write([]byte("123"))
	data := make([]byte, 64)

	for{
		n, err := file.Read(data)
		if err == io.EOF{   // если конец файла
			break           // выходим из цикла
		}
		fmt.Print(string(data[:n]))
	}
	fmt.Println(file.Name())
}
