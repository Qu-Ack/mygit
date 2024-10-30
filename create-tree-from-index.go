package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

const SEP_INDEX string = "#"

type entry struct {
	hash    string
	dir_arr []string
}

func GenerateRootTree(filename string) (string, error) {
	fs, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer fs.Close()

	var content bytes.Buffer
	_, err = io.Copy(&content, fs)
	if err != nil {
		return "", err
	}

	split_arr := strings.Split(string(content.Bytes()), SEP_INDEX)
	tree_map := make(map[int][]Node)
	tree_map[0] = make([]Node, 0)
	tree_map[0] = append(tree_map[0], &DirNode{
		Name:     ".",
		Children: make([]Node, 0),
	})

	for i := range split_arr {
		if len(split_arr[i]) < 38 {
			continue
		}

		entry := entry{
			hash:    split_arr[i][6:70],
			dir_arr: strings.Split(split_arr[i][70:], "/"),
		}
		for x := 0; x < len(entry.dir_arr); x++ {

			if tree_map[x] == nil {
				tree_map[x] = make([]Node, 0)
			}

			nodeExists := 0

			for val := range tree_map[x] {
				if tree_map[x][val].GetName() == entry.dir_arr[x] {
					nodeExists = 1
					break
				}
			}

			if nodeExists != 1 {
				if x == len(entry.dir_arr)-1 {
					file_node := &FileNode{
						Name: entry.dir_arr[x],
						Hash: entry.hash,
					}
					tree_map[x] = append(tree_map[x], file_node)

					// add to parent
					if x-1 > -1 {
						for item := range tree_map[x-1] {
							if tree_map[x-1][item].GetName() == entry.dir_arr[x-1] {
								switch n := tree_map[x-1][item].(type) {
								case *DirNode:
									n.Children = append(n.Children, file_node)
								}
							}
						}
					}

				} else {
					dir_node := &DirNode{
						Name:     entry.dir_arr[x],
						Children: make([]Node, 0),
					}
					tree_map[x] = append(tree_map[x], dir_node)

					// add to parent
					if x-1 > -1 {
						for item := range tree_map[x-1] {
							if tree_map[x-1][item].GetName() == entry.dir_arr[x-1] {
								switch n := tree_map[x-1][item].(type) {
								case *DirNode:
									n.Children = append(n.Children, dir_node)
								}
							}
						}
					}
				}
			}

		}
	}

	hash, _ := commitTreeHash(tree_map[0][0])

	return hash, nil
}
