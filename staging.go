package main

import (
	"errors"
	"fmt"
	"os"
)

const TEMP_FILE_PERM string = "033333"

func checkStagingParams() error {
	if len(os.Args) < 3 {
		return errors.New("Args Error")
	}

	return nil
}

/*
How to Recursively read files in a file tree
Generate their blob objects and get their hashes
write those hashes in our index
That is the main problem that this function solves
*/

func createIndex() error {
	result_str := ""

	for i := 0; i < len(os.Args)-2; i++ {
		err := createBlobObject(os.Args[2+i])
		hash, err := HashFile(os.Args[2+i])

		if err != nil {
			return err
		}

		result_str += fmt.Sprintf("%v%v%v#", TEMP_FILE_PERM, hash, os.Args[2+i])
	}

	err := os.WriteFile(".mygit/index.txt", []byte(result_str), 0777)

	if err != nil {
		return err
	}

	return nil
}

func createBlobObject(filename string) error {
	err := HashFileHandler("-w", filename)

	if err != nil {
		return err
	}

	return nil

}
