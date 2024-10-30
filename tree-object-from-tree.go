package main

import (
	"compress/zlib"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

type Node interface {
	GetName() string
}

type FileNode struct {
	Hash string
	Name string
}

type DirNode struct {
	Name     string
	Children []Node
}

func (f *FileNode) GetName() string {
	return f.Name
}
func (f *DirNode) GetName() string {
	return f.Name
}

func commitTreeHash(node Node) (string, string) {
	switch n := node.(type) {
	case *FileNode:
		return n.Hash, n.Name
	case *DirNode:

		temp_store := make(map[string]string)
		for val := range n.Children {
			hash, key := commitTreeHash(n.Children[val])
			temp_store[key] = hash
		}
		final_hash, err := generateTreeHash(temp_store)

		if err != nil {
			fmt.Println(err.Error())
			panic("error in generateTree Hash")
		}

		return final_hash, n.Name
	default:
		return "", ""
	}
}

func generateTreeHash(hashstore map[string]string) (string, error) {
	result_str := ""

	for name, hash := range hashstore {
		fmt.Println(name)
		result_str += fmt.Sprintf("%v\x00%v", name, hash)
	}

	hashBytes := sha256.Sum256([]byte(result_str))
	result_hash := fmt.Sprintf("%x", hashBytes)

	err := CreateTreeObject(result_hash, []byte(result_str))
	if err != nil {
		return "", err
	}

	return result_hash, nil

}

func CreateTreeObject(hash string, str []byte) error {
	dir_path := fmt.Sprintf("./.mygit/objects/%v/%v", hash[:2], hash[2:])
	err := os.MkdirAll(fmt.Sprintf("./.mygit/objects/%v", hash[:2]), 0777)
	if err != nil {
		return err
	}
	fs, err := os.Create(dir_path)
	if err != nil {
		return err
	}
	zlib_writer := zlib.NewWriter(fs)
	_, err = zlib_writer.Write(str)
	if err != nil {
		zlib_writer.Close()
		fs.Close()
		return err
	}
	err = zlib_writer.Close()
	if err != nil {
		fs.Close()
		return err
	}
	err = fs.Close()

	if err != nil {
		return err
	}
	return nil
}

func DecompressZlib(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a zlib reader
	zlibReader, err := zlib.NewReader(file)
	if err != nil {
		return "", err
	}
	defer zlibReader.Close()

	decompressedData, err := ioutil.ReadAll(zlibReader)
	if err != nil {
		return "", err
	}

	fmt.Println(string(decompressedData))

	return string(decompressedData), nil
}
