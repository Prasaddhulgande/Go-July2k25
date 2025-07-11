package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	// Simple way to read file most easiest way.If file size small then only we are using this way. Other wise
	// at a time load full file size like in GB's
	/* data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data)) */

	// Read folders
	/*
		dir, err := os.Open("../") // ../ ---> return all dir
		if err != nil {
			panic(err)
		}
		defer dir.Close()

		fileInfo, err := dir.ReadDir(-1) //If here int <0 then return all files

		for _, fi := range fileInfo {
			fmt.Println(fi.Name(), fi.IsDir())
		}
	*/

	/*f, err := os.Open("example.txt") // Open the file
	if err != nil {
		//log the error
		panic(err)
	}

	defer f.Close() // Close the file

	buf := make([]byte, 20)
	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data in File: ", d, buf)
	// Read data in string fmt
	for i := 0; i < len(buf); i++ {
		fmt.Println("Data in File: ", d, string(buf[i]))
	} */
	/*if err != nil {
		//log the error
		panic(err)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("File Name is: ", fileInfo.Name())
	fmt.Println("File Name is: ", fileInfo.Mode())
	fmt.Println("File Name is: ", fileInfo.ModTime()) */

	//CreateFile()
	// OneFileToAnotherFileText()
	// deleteFile()
	CopyPasteFile()
}

func CreateFile() {

	f, err := os.Create("example3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// f.WriteString("Hello Mr. Golang, How are you?")

	bytes := []byte("Hello Golang")
	f.Write(bytes)

}

// READ & WRITE TO ANOTHER FILE (STRIMING FASHION)

func OneFileToAnotherFileText() {
	SourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer SourceFile.Close()

	// create new file and paste data file example.txt

	destFile, err := os.Create("example3.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(SourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break

		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}

		writer.Flush()

		fmt.Println("Written to new file successfully: ", e)
	}

}

func CopyPasteFile() {
	srcFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	destFile, err := os.Create("exam2.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	newFile, err := io.Copy(destFile, srcFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("NewFile written contetn:", newFile)
}

func deleteFile() {

	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Removed file successfully")

}
