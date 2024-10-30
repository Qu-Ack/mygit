package main

import (
	"compress/zlib"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
)

var hashFileParams = make([]string, 0)

func hashFileCheckArgs() error {
	hashFileParams = append(hashFileParams, "-w")

	if len(os.Args) == 3 {
		return nil

	} else {

		if len(os.Args) != 4 {
			return errors.New("Args Error")
		}

		if !slices.Contains(hashFileParams, os.Args[2]) {
			return errors.New("Args Error")
		}

	}
	return nil
}

/*
git uses SHA-1 but is migrating to SHA-256
SHA-1 has a security flaw that lets you generate
2 files with the same signature
thus different files can be considered same
https://shattered.io
that's why we will  use SHA-256 in our implementation
*/
func HashFile(filename string) (string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return "", errors.New("Couldn't open the file")

	}

	defer file.Close()

	hasher := sha256.New()

	// get the file size
	fileinfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	// calculating the header hash
	header := fmt.Sprintf("blob %v\x00", fileinfo.Size())
	hasher.Write([]byte(header))

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)

		if err != nil && err != io.EOF {
			return "", err
		}

		if n > 0 {
			hasher.Write(buffer[:n])
		}

		if err == io.EOF {
			break
		}
	}

	result := hex.EncodeToString(hasher.Sum(nil))

	return result, nil
}

/*
handler to do things based on different params
*/
func HashFileHandler(param string, filename string) error {
	hash, err := HashFile(filename)

	if err != nil {
		return err
	}

	if param == "-w" {
		dir_path := fmt.Sprintf("./.mygit/objects/%v/", hash[:2])
		err := os.Mkdir(dir_path, 0777)

		if err != nil {
			return err
		}

		source_file, err := os.Open(filename)

		if err != nil {
			return err
		}

		defer source_file.Close()

		source_file_stat, err := source_file.Stat()

		if err != nil {
			return err
		}

		result_file, err := os.Create(fmt.Sprintf("%v%v", dir_path, hash[2:]))

		if err != nil {
			return err
		}

		defer result_file.Close()

		zlib_writer := zlib.NewWriter(result_file)

		defer zlib_writer.Close()

		zlib_writer.Write([]byte(fmt.Sprintf("blob %v\x00", source_file_stat.Size())))

		_, err = io.Copy(zlib_writer, source_file)
		if err != nil {
			return err
		}

	} else if param == "" {
		fmt.Println(hash)
	}

	return nil

}
