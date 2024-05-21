# Prototype in Go

## Introduction 

Prototype is a creational design pattern that allows cloning objects, even complex ones, without coupling to their specific classes.

All prototype classes should have a common interface that makes it possible to copy objects even if their concrete classes are unknown. Prototype objects can produce full copies since objects of the same class can access each other’s private fields.

## Conceptual Example

Prototype pattern using an example based on the operating system’s file system. The OS file system is recursive: the folders contain files and folders, which may also include files and folders, and so on.

Each file and folder can be represented by an `inode` interface. inode interface also has the `clone` function.

Both `file` and `folder` structs implement the `print` and `clone` functions since they are of the `inode` type. Also, notice the `clone` function in both `file` and `folder`. The `clone` function in both of them returns a copy of the respective file or folder. During the cloning, we append the suffix “_clone” to the name field.

###  inode.go: Prototype interface

```
package main

type Inode interface {
    print(string)
    clone() Inode
}
```

### file.go: Concrete prototype

```
package main

import "fmt"

type File struct {
    name string
}

func (f *File) print(indentation string) {
    fmt.Println(indentation + f.name)
}

func (f *File) clone() Inode {
    return &File{name: f.name + "_clone"}
}
```

### folder.go: Concrete prototype

```
package main

import "fmt"

type Folder struct {
    children []Inode
    name     string
}

func (f *Folder) print(indentation string) {
    fmt.Println(indentation + f.name)
    for _, i := range f.children {
        i.print(indentation + indentation)
    }
}

func (f *Folder) clone() Inode {
    cloneFolder := &Folder{name: f.name + "_clone"}
    var tempChildren []Inode
    for _, i := range f.children {
        copy := i.clone()
        tempChildren = append(tempChildren, copy)
    }
    cloneFolder.children = tempChildren
    return cloneFolder
}
```

### main.go: Client code

```
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
    fmt.Println("\nPrinting hierarchy for Folder2")
    folder2.print("  ")

    cloneFolder := folder2.clone()
    fmt.Println("\nPrinting hierarchy for clone Folder")
    cloneFolder.print("  ")
}
```

### output.txt: Execution result

```
Printing hierarchy for Folder2
  Folder2
    Folder1
        File1
    File2
    File3

Printing hierarchy for clone Folder
  Folder2_clone
    Folder1_clone
        File1_clone
    File2_clone
    File3_clone
```
