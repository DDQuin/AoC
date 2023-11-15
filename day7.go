package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type FileNode struct {
	name        string
	isDirectory bool
	size        int
	parent      *FileNode
	files       []*FileNode
}

func addFile(name string, isDirectory bool, size int, parent *FileNode) *FileNode {
	newFile := FileNode{
		name:        name,
		isDirectory: isDirectory,
		size:        size,
		parent:      parent,
		files:       make([]*FileNode, 0),
	}
	if parent != nil {
		parent.files = append(parent.files, &newFile)
	}
	return &newFile
}

func getOrAddFile(name string, isDirectory bool, size int, parent *FileNode) *FileNode {
	if parent == nil {

		return addFile(name, isDirectory, size, parent)
	}
	for _, child := range parent.files {
		if child.name == name {

			return child
		}
	}

	return addFile(name, isDirectory, size, parent)
}

func getFullPath(file *FileNode) string {
	var fullPath = file.name
	var curFile = file
	for curFile.parent != nil {
		curFile = curFile.parent
		if curFile.parent == nil {
			fullPath = curFile.name + fullPath
		} else {
			fullPath = curFile.name + "/" + fullPath
		}
	}
	return fullPath
}

func preOrderWalk(file *FileNode) {

	fmt.Println("Name: ", file.name, " Size: ", file.size, " isDirectory: ", file.isDirectory, "Full path: ", getFullPath(file))
	if len(file.files) == 0 {
		return
	}

	for _, child := range file.files {
		preOrderWalk(child)
	}
}

func directorySize(file *FileNode, dirs map[string]int) int {
	if len(file.files) == 0 {
		return file.size
	}

	var sum int = 0
	for _, child := range file.files {
		sum += directorySize(child, dirs)
	}
	if file.isDirectory {

		dirs[getFullPath(file)] = sum
		fmt.Println("Directorty ", file.name, " Size: ", sum)

	}
	return sum
}

func createTestTree() {
	root := FileNode{
		name:        "/",
		isDirectory: true,
		size:        -1,
		parent:      nil,
		files:       make([]*FileNode, 0),
	}
	a := addFile("a", true, -1, &root)
	addFile("b.txt", false, 14848514, &root)
	addFile("c.dat", false, 8504156, &root)
	d := addFile("d", true, -1, &root)

	addFile("j", false, 4060174, d)
	addFile("d.log", false, 8033020, d)
	addFile("d.ext", false, 5626152, d)
	addFile("k", false, 7214296, d)

	addFile("f", false, 29116, a)
	addFile("g", false, 2557, a)
	addFile("h.lst", false, 62596, a)
	e := addFile("e", true, -1, a)

	addFile("i", false, 584, e)

	preOrderWalk(&root)
}

func Day7() {
	lines, err := ReadLines("resources/day7input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var currentFile *FileNode = nil
	var root *FileNode = nil
	for _, line := range lines {
		tokens := strings.Split(line, " ")

		if tokens[0] == "$" {
			if tokens[1] == "cd" {
				if tokens[2] == ".." {
					currentFile = currentFile.parent
				} else {
					f := getOrAddFile(tokens[2], true, 0, currentFile)
					if f.name == "/" {
						root = f
					}
					currentFile = f
				}
			} else if tokens[1] == "ls" {

			}
		} else {
			if tokens[0] == "dir" {
				addFile(tokens[1], true, 0, currentFile)
			} else {
				size, err := strconv.Atoi(tokens[0])
				if err != nil {
					log.Fatalf("converting string to num: %s", err)
				}
				addFile(tokens[1], false, size, currentFile)
			}
		}
	}

	dirMap := make(map[string]int)
	var usedSpace int = directorySize(root, dirMap)

	var sum int = 0
	for _, val := range dirMap {
		if val <= 100_000 {
			sum += val
		}
	}
	fmt.Println(sum)

	var totalSpace int = 70000000
	var spaceNeeded int = 30000000

	var unusedSpace int = totalSpace - usedSpace

	fmt.Println(unusedSpace)

	var smallestDictSize int = totalSpace
	for _, val := range dirMap {
		if val+unusedSpace > spaceNeeded && val < smallestDictSize {
			smallestDictSize = val
		}
	}

	fmt.Println(smallestDictSize)

}
