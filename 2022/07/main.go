package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Folder struct {
	Name   string
	Parent *Folder
	SubDir []*Folder
	Files  []*File
}

type File struct {
	Name string
	Size int
}

func (folder *Folder) FolderSizeMap() map[string]int {
	sizeMap := make(map[string]int)
	for _, subFolder := range folder.SubDir {
		for name, size := range subFolder.FolderSizeMap() {
			sizeMap[name] = size
		}
	}
	sizeMap[folder.Name] = folder.Size()
	return sizeMap
}

func (folder *Folder) FolderLessThanSize(size int) []*Folder {
	var folders []*Folder
	for _, sub := range folder.SubDir {
		folders = append(folders, sub.FolderLessThanSize(size)...)
		if sub.Size() <= size {
			folders = append(folders, sub)
		}
	}
	return folders
}

func (folder *Folder) Size() int {
	// Sum size of files in folder (and subfolders)
	var size int
	for _, file := range folder.Files {
		size += file.Size
	}
	for _, folder := range folder.SubDir {
		size += folder.Size()
	}
	return size
}

func (folder *Folder) FindFolder(path string) *Folder {
	// Find folders in directory
	if path == "/" {
		return folder
	}
	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	if strings.Contains(path, "/") {
		tmp := strings.Split(path, "/")
		for _, sub := range folder.SubDir {
			if sub.Name == tmp[0] {
				return sub.FindFolder(strings.Join(tmp[1:], "/"))
			}
		}
	} else {
		for _, sub := range folder.SubDir {
			if sub.Name == path {
				return sub
			}
		}
	}
	return nil
}

func main() {
	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}
	fileLines := strings.Split(strings.TrimSpace(string(data)), "\n")

	rootDir := &Folder{}
	currentDir := rootDir
	for _, line := range fileLines {
		if strings.HasPrefix(line, "$ cd ") {
			path := strings.TrimPrefix(line, "$ cd ")
			if path == "/" {
				currentDir = rootDir
			} else if path == ".." {
				currentDir = currentDir.Parent
			} else {
				currentDir = currentDir.FindFolder(path)
			}
		} else if strings.HasPrefix(line, "$ ls") {
			//Do nothing
		} else {
			if strings.HasPrefix(line, "dir ") {
				currentDir.SubDir = append(currentDir.SubDir, &Folder{
					Name:   strings.TrimPrefix(line, "dir "),
					Parent: currentDir,
				})
			} else {
				tmp := strings.Split(line, " ")
				size, _ := strconv.Atoi(tmp[0])
				currentDir.Files = append(currentDir.Files, &File{
					Name: tmp[1],
					Size: size,
				})
			}
		}
	}

	folders := rootDir.FolderLessThanSize(100000)
	var size int
	for _, folder := range folders {
		size += folder.Size()
	}

	fmt.Println()
	fmt.Println("Day 7: No Space Left On Device")
	fmt.Println("What is the sum of the total sizes of those directories?")
	fmt.Println(size)

	unusedSpace := 70000000 - rootDir.Size()
	spaceNeedToFree := 30000000 - unusedSpace
	sizeMap := rootDir.FolderSizeMap()

	//sort sizeMap
	var sizes []int
	for _, size := range sizeMap {
		sizes = append(sizes, size)
	}
	for i := 0; i < len(sizes); i++ {
		for j := i + 1; j < len(sizes); j++ {
			if sizes[i] > sizes[j] {
				sizes[i], sizes[j] = sizes[j], sizes[i]
			}
		}
	}

	//find the smallest folder that can free up enough space
	for _, size := range sizes {
		if size >= spaceNeedToFree {
			// Found it!
			break
		}
	}

	fmt.Println()
	fmt.Println("Part 2:")
	fmt.Println("What is the total size of that directory?")
	fmt.Println(size)

}
