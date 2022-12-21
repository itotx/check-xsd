package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/antchfx/xmlquery"
)

func main() {
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(path) == ".xsd" {
				checkFile(path)
			}
			return nil
		})
	if err != nil {
		panic(err)
	}
}
func checkFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	parser, err := xmlquery.CreateStreamParser(f, "//xs:complexType")
	if err != nil {
		panic(err)
	}
	for {
		n, err := parser.Read()
		if err == io.EOF {
			break
		}
		if n.Attr[0].Value[0:1] != strings.ToUpper(n.Attr[0].Value[0:1]) {
			fmt.Println("Ошибка у аттрибута name типа:", n.Attr[0].Value, "в файле", fileName)
		}
		if err != nil {
			panic(err)
		}
	}
	defer f.Close()
}
