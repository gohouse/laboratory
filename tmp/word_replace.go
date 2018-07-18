package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func Replace(oldFilename, newFilename, oldWord, newWord string) (err error) {
	zipReader, err := zip.OpenReader(oldFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	newFile, err := os.Create(newFilename)
	if err != nil {
		return
	}
	defer newFile.Close()

	zipWriter := zip.NewWriter(newFile)
	for _, file := range zipReader.File {

		writer, err := zipWriter.Create(file.Name)
		if err != nil {
			return err
		}

		readCloser, err := file.Open()
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		buf.ReadFrom(readCloser)

		if file.Name == "word/document.xml" {
			newContent := strings.Replace(string(buf.Bytes()), oldWord, newWord, -1)
			writer.Write([]byte(newContent))
		} else {
			writer.Write(buf.Bytes())
		}
	}
	zipWriter.Close()
	return nil
}

func main() {
	if len(os.Args) != 5 {
		fmt.Println("\n Usage: go run Replace.go <source.docx> <target.docx> <originalWord> <newWord>" +
			"\n e.g. go run replace.go ./sample.docx ./newfile.docx xyz abc\n")
		return
	}

	Replace(os.Args[1], os.Args[2], os.Args[3], os.Args[4])
}
