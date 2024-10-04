package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var ErrStruct error = fmt.Errorf("directory structure is not expected")
var ErrSyntax error = fmt.Errorf("syntax error")

func main() {
	// check command-line args
	if len(os.Args) != 2 {
		log.Fatal("wrong argments")
	}

	dir, err := readDirStructDef(readJson())
	if err != nil {
		fmt.Println("false")
		return
	}
	err = checkAllDirStruct(dir)
	if err != nil {
		fmt.Println("false")
		return
	}
	fmt.Println("true")
}

func readJson() []byte {
	jsonData, _ := io.ReadAll(os.Stdin)
	return jsonData
}

func readDirStructDef(jsonData []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func checkDirContents(dir Directory, que *Queue) error {
	dirEntries, err := os.ReadDir(dir.path)
	if err != nil {
		return err
	}

	if len(dirEntries) != len(dir.dir) {
		return ErrStruct
	}
	for _, dirEntry := range dirEntries {
		name := dirEntry.Name()
		value, ok := dir.dir[name]
		if !ok {
			return ErrStruct
		}

		childDir, ok := value.(map[string]interface{})
		if ok {
			// if directory
			if !dirEntry.IsDir() {
				return ErrStruct
			}
			que.push(Directory{dir: childDir, path: filepath.Join(dir.path, name)})
		} else if value == nil {
			// if file
			if dirEntry.IsDir() {
				return ErrStruct
			}
		} else {
			return ErrSyntax
		}
	}
	return nil
}

func checkAllDirStruct(totalDirStruct map[string]interface{}) error {
	var que Queue
	que.push(Directory{dir: totalDirStruct, path: os.Args[1]})
	for !que.isEmpty() {
		err := checkDirContents(que.pop(), &que)
		if err != nil {
			return err
		}
	}
	return nil
}
