package main

import (
	"Blockchain/pkg"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("your filename->")
	var fileName string
	_, err := fmt.Scan(&fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	data := make([]byte, 64)
	for {
		_, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
	}
	text := string(data)
	list := strings.Split(text, "\n")
	root := pkg.New()
	c := root.Root
	for i := range list {
		b := pkg.MakeBlock(list[i], c)
		c = root.Add(*b)
	}

	text = ""
	c = root.Root.Next

	for c != root.Root {
		text += c.Value
		text += " -> " + c.TransactionHash + " \n "
		c = c.Next
	}
	blockchain, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	data = []byte(text)
	_, err = blockchain.Write(data)
	if err != nil {
		log.Fatal(err.Error())
	}
}
