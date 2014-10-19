
package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"container/list"
)

func main() {
	//First example of opening a file
	file, err := os.Open("/home/luca/Java/workspace/Learning_go/src/chapter13/test.txt")
	if err != nil {
		fmt.Println("Error opening test.txt")
		return
	}
	defer file.Close()
	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	 
	if err != nil { 
		return
	}
	
	str := string(bs)
	fmt.Println(str)
	
	
	//Second way of opening a file
	bs, err = ioutil.ReadFile("/home/luca/Java/workspace/Learning_go/src/chapter13/test2.txt")
	if err != nil {
		fmt.Println("Error opening test2.txt")
		return
	}
	str = string(bs)
	fmt.Println(str)	
	
	
	//This is an example of writing a file
	file, err = os.Create("/home/luca/Java/workspace/Learning_go/src/chapter13/test3.txt")
	if err != nil {
		fmt.Println("Error creating test3.txt")
		return
	}
	
	defer file.Close()
	file.WriteString("file test 3")
	
	//Read the file(s) in a folder
	dir, err := os.Open("/home/luca/Java/workspace/Learning_go/src/chapter13/")
	if err != nil {
		fmt.Println("Error opening the folder")
		return
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Errorr executing Readdir")
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
	
	//Example of using lists
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	fmt.Println("Printing a list")
	for e := x.Front(); e != nil; e=e.Next() {
		fmt.Println(e.Value)
	}
}