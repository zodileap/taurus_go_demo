package main

import (
	"fmt"
	"testing"

	"github.com/yohobala/taurus_go/asset"
)

func TestAdd(t *testing.T) {
	var assets asset.Assets
	assets.Add("file.txt", []byte("Hellow, World!"))
	assets.Add("../file_2.go", []byte(`
	package main
	import "fmt"
	func Main() {fmt.Println("Hello, World!")}
	`))
	err := assets.Write()
	if err != nil {
		fmt.Print(err)
	}
}

func TestAddDir(t *testing.T) {
	var assets asset.Assets
	assets.AddDir("dir")
	assets.AddDir("./dir2")
	err := assets.Write()
	if err != nil {
		fmt.Print(err)
	}
}

func TestWrite(t *testing.T) {
	var assets asset.Assets
	assets.Add("file.txt", []byte("Hellow, World!"))
	assets.AddDir("dir")
	err := assets.Write()
	if err != nil {
		fmt.Print(err)
	}
}

func TestFormat(t *testing.T) {
	var assets asset.Assets
	assets.Add("test.go", []byte(`
	package main
	import "fmt"
	func Main() {fmt.Println("Hello, World!")}
	`))
	err := assets.Write()
	if err != nil {
		fmt.Print(err)
	}
	err = assets.Format()
	if err != nil {
		fmt.Print(err)
	}
}

func TestCopy(t *testing.T) {
	err := asset.CopyFile("file.go", "file2.go")
	if err != nil {
		fmt.Print(err)
	}
}

func TestCopyDir(t *testing.T) {
	err := asset.CopyDir("dir", "dir2")
	if err != nil {
		fmt.Print(err)
	}
}
