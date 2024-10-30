package main

import (
	"compress/zlib"
	"crypto/sha256"
	"fmt"
	"os"
	"time"
)

/*
commit <size>\0commitmessage\0
*/

func WriteCommitObject(hash string) (string, error) {

	t := time.Now()
	result_str := fmt.Sprintf("commit \x00%vDakshSangal\x00%v", hash, t.Format("20060102150405"))

	result_hash_bytes := sha256.Sum256([]byte(result_str))

	result_hash := fmt.Sprintf("%x", result_hash_bytes)

	file_path := fmt.Sprintf("./.mygit/objects/%v/%v", result_hash[:2], result_hash[2:])
	dir_path := fmt.Sprintf("./.mygit/objects/%v", result_hash[:2])
	err := os.Mkdir(dir_path, 0777)

	if err != nil {
		return "", err
	}

	fs, err := os.Create(file_path)
	defer fs.Close()

	if err != nil {
		return "", err
	}

	zlib_writer := zlib.NewWriter(fs)

	defer zlib_writer.Close()

	zlib_writer.Write([]byte(fmt.Sprintf("commit \x00")))

	zlib_writer.Write([]byte(hash))

	zlib_writer.Write([]byte("DakshSangal\x00"))

	zlib_writer.Write([]byte(t.Format("20060102150405")))

	return result_hash, nil

}
