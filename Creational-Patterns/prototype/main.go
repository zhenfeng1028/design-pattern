package main

import "fmt"

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Printf("Printing hierarchy for Folder2\n")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Printf("\nPrinting hierarchy for clone Folder\n")
	cloneFolder.print("  ")
}
