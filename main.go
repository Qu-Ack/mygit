package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
		os.Exit(1)
	}

	switch command := os.Args[1]; command {
	case "init":

		for _, dir := range []string{".mygit", ".mygit/objects", ".mygit/refs"} {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
			}
		}

		headFileContents := []byte("ref: refs/heads/main\n")
		if err := os.WriteFile(".mygit/HEAD", headFileContents, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
		}

		fmt.Println("Initialized git directory")
	case "cat-file":
		err := catfileCheckArgs()
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("usage: mygit <command> [<args>...]")
			os.Exit(1)
		}

		err = catFileOutputResult(os.Args[2])

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("usage: mygit <command> [<args>...]")
			os.Exit(1)
		}
	case "hash-file":
		err := hashFileCheckArgs()

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("usage: mygit <command> [<args>...]")
			os.Exit(1)
		}

		if len(os.Args) == 3 {
			err = HashFileHandler("", os.Args[2])
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("usage: mygit <command> [<args>...]")
				os.Exit(1)
			}

		} else {

			err = HashFileHandler(os.Args[2], os.Args[3])
			if err != nil {

				fmt.Println(err.Error())
				fmt.Println("usage: mygit <command> [<args>...]")
				os.Exit(1)
			}
		}
	case "ls-tree":
		err := ChecklsTreeArgs()

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("usage: mygit <command> [<args>...]")
			os.Exit(1)
		}

		err = ReadTree(os.Args[2])

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("usage: mygit <command> [<args>...]")
			os.Exit(1)

		}
	case "commit":
		hash, err := GenerateRootTree("./.mygit/index.txt")

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("usage: mygit <command> [<args>...]")
			os.Exit(1)
		}

		commit_object_hash, err := WriteCommitObject(hash)

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("usage: mygit <command> [<args>...]")
			os.Exit(1)
		}

		fmt.Println(commit_object_hash)

	case "decomp-zlib":
		DecompressZlib(os.Args[2])
	case "staging":
		err := checkStagingParams()
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("usage: mygit <command> [<args>...]")
			os.Exit(1)
		}
		err = createIndex()
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("usage: mygit <command> [<args>...]")
			os.Exit(1)
		}

	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}
}
