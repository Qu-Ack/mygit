package main

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

var catFileCliParams = make([]string, 0)

func catfileCheckArgs() error {
	catFileCliParams = append(catFileCliParams, "-p")
	catFileCliParams = append(catFileCliParams, "-t")
	catFileCliParams = append(catFileCliParams, "-s")

	if len(os.Args) != 4 {
		return errors.New("Args Error")
	}

	if !slices.Contains(catFileCliParams, os.Args[2]) {
		return errors.New("Args Error")
	}

	return nil
}

// cat-file only runs from root dir of project
// should be able to run from anywhere inside the project

func catfileReadBlob() ([]byte, error) {

	zlib_data := make([]byte, 2048)
	string_data := make([]byte, 2048)

	directory_name := os.Args[3][:2]

	file_name := os.Args[3][2:]

	path := fmt.Sprintf("./.mygit/objects/%v/%v", directory_name, file_name)

	fs, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer fs.Close()

	n, err := fs.Read(zlib_data)
	zlib_data = zlib_data[:n]

	payload := bytes.NewReader(zlib_data)

	if err != nil {
		return nil, err
	}

	uncompressed_data, err := zlib.NewReader(payload)

	defer uncompressed_data.Close()

	if err != nil {
		return nil, err
	}

	n, err = uncompressed_data.Read(string_data)

	string_data = string_data[:n]

	return string_data, nil
}

func catFileOutputResult(param string) error {
	data, err := catfileReadBlob()
	if err != nil {
		return err
	}

	if param == "-t" {
		result := ""
		for _, val := range data {
			if val == 0 {
				break
			}
			result += string(val)
		}
		newResult := strings.Split(result, " ")
		fmt.Println(newResult[0])

	} else if param == "-s" {
		result := ""
		for _, val := range data {
			if val == 0 {
				break
			}
			result += string(val)
		}
		newResult := strings.Split(result, " ")
		fmt.Println(newResult[1])
	} else if param == "-p" {
		afterNull := false
		result := ""
		for _, val := range data {
			if afterNull {
				result += string(val)
			}
			if val == 0 {
				afterNull = true
			}
		}
		fmt.Print(result)

	}

	return nil
}
