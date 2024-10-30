package main

import (
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

func ChecklsTreeArgs() error {
	if len(os.Args) == 3 {
		return nil
	} else {
		if len(os.Args) == 4 {
			if os.Args[2] == "--name-only" {
				return nil
			} else {
				return errors.New("Args Error")
			}
		} else {
			return errors.New("Args Error")
		}
	}

}

func ReadTree(hash string) error {
	path := fmt.Sprintf(".git/objects/%v/%v", hash[:2], hash[2:])
	source_file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer source_file.Close()

	zlib_reader, err := zlib.NewReader(source_file)

	if err != nil {
		return err
	}

	defer zlib_reader.Close()

	var content bytes.Buffer

	_, err = io.Copy(&content, zlib_reader)

	if err != nil {
		return err
	}

	ParseTreeFile(content.Bytes())

	return nil
}

func ParseTreeFile(buffer []byte) error {

	line_count := 0
	for i := 0; i < len(buffer); i++ {
		if (line_count) == 0 {
			if buffer[i] == '\x00' {
				fmt.Println("")
				line_count = line_count + 1
				continue
			}
			fmt.Print(string(buffer[i]))

		} else {

			if buffer[i] == '\x00' {
				fmt.Printf(" %v", hex.EncodeToString(buffer[i+1:i+21]))
				fmt.Println("")
				i = i + 20
				continue
			}

			fmt.Print(string(buffer[i]))
		}
	}
	return nil
}
