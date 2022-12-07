package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type fileInfo struct {
	size int
}

type directoryNode struct {
	parent         *directoryNode
	subdirectories map[string]*directoryNode
	files          map[string]fileInfo
}

func newDirectoryNode() directoryNode {
	return directoryNode{
		nil,
		map[string]*directoryNode{},
		map[string]fileInfo{},
	}
}

func main() {
	fileContent, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(fileContent), "\n")

	rootDir := newDirectoryNode()
	currentDir := &rootDir

	// Skip $ cd /
	for i := 1; i < len(lines); i++ {
		currentLine := lines[i]
		cmdParts := strings.Split(currentLine, " ")
		if len(cmdParts) <= 1 {
			break
		}
		command := cmdParts[1]
		if command == "ls" {
			for j := i + 1; j < len(lines); j++ {
				lsLineParts := strings.Split(lines[j], " ")
				// Start of next command; stop processing
				if lsLineParts[0] == "$" || len(lsLineParts) == 1 {
					i = j - 1
					break
				} else if lsLineParts[0] == "dir" {
					dirName := lsLineParts[1]
					newDir := newDirectoryNode()
					newDir.parent = currentDir
					currentDir.subdirectories[dirName] = &newDir
				} else {
					fileName := lsLineParts[1]
					fileSize, _ := strconv.Atoi(lsLineParts[0])
					currentDir.files[fileName] = fileInfo{fileSize}
				}
			}
		} else if command == "cd" {
			targetDir := cmdParts[2]
			if targetDir == ".." {
				currentDir = currentDir.parent
			} else {
				currentDir, _ = currentDir.subdirectories[targetDir]
			}
		}
	}

	p1 := func() {
		dirSizes := listDirSizes(&rootDir)
		total := 0
		for _, dirSize := range dirSizes {
			if dirSize <= 100_000 {
				total += dirSize
			}
		}
		fmt.Printf("%d\n", total)
	}

	p2 := func() {
		dirSizes := listDirSizes(&rootDir)
		totalSize := dirSizes[len(dirSizes)-1]
		minSize := math.MaxInt
		sizeToDelete := totalSize - 40_000_000
		for _, dirSize := range dirSizes {
			if dirSize >= sizeToDelete && dirSize < minSize {
				minSize = dirSize
			}
		}
		fmt.Printf("%d\n", minSize)
	}

	_ = p1
	p2()
}

func listDirSizes(dir *directoryNode) []int {
	_, dirSizes := listDirSizesImpl(dir, []int{})
	return dirSizes
}

func listDirSizesImpl(dir *directoryNode, dirSizes []int) (int, []int) {
	// Returns directory total, list of all directory sizes
	totalSize := 0
	for _, file := range dir.files {
		totalSize += file.size
	}
	for _, dir := range dir.subdirectories {
		subTotalSize, newDirSizes := listDirSizesImpl(dir, dirSizes)
		dirSizes = newDirSizes
		totalSize += subTotalSize
	}
	dirSizes = append(dirSizes, totalSize)
	return totalSize, dirSizes
}
