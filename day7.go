package main

import (
	"fmt"
)

type FileNode struct {
	name string
	isDirectory bool
	size int
	parent *FileNode
	files []*FileNode

}

func addFile(name string, isDirectory bool, size int, parent *FileNode) *FileNode {
	newFile := FileNode{
		name: name,
		isDirectory: isDirectory,
		size: size,
		parent: parent,
		files: make([]*FileNode, 0),
	}
	parent.files = append(parent.files, &newFile)
	return &newFile
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

func createTestTree() {
	root := FileNode{
		name: "/",
		isDirectory: true,
		size: -1,
		parent: nil,
		files: make([]*FileNode, 0),
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
	// lines, err := ReadLines("resources/day7test.txt")
	// if err != nil {
	// 	log.Fatalf("readLines: %s", err)
	// }
	
	createTestTree()
	// for _, line := range lines {
	// 	fmt.Println(line)
	// }

}
