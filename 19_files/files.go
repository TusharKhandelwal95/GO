/*
File operations
*/

package main

import (
	// "bufio"
	"fmt"
	"os"
)

// BASIC FILE OPERATIONS

// func main(){
// 	f, err := os.Open("example.txt")
// 	if err != nil{
// 		//log the error or panic
// 		panic(err)
// 	}

// 	fileInfo, err := f.Stat()
// 	if err != nil{
// 		//log the error or panic
// 		panic(err)
// 	}
// 	fmt.Println("Filename: ", fileInfo.Name()) // example.txt
// 	fmt.Println("File or folder: ", fileInfo.IsDir())  // false
// 	fmt.Println("File size: ", fileInfo.Size())        // 13
// 	fmt.Println("File permissions: ", fileInfo.Mode()) // -rw-rw-rw-
// 	fmt.Println("File modified at: ", fileInfo.ModTime()) // 2025-05-14 20:12:18.8213987 +0530 IST

// }
//----------------------------------------------------------------------

// READING FILE
// func main(){
// 	f, err := os.Open("example.txt")
// 	if err != nil{
// 		panic(err)
// 	}

// 	defer f.Close()

// 	// we read the file ans store in buffer(array of bites)
// 	buf := make([]byte, 20)

// 	d, err := f.Read(buf)
// 	if err != nil{
// 		panic(err)
// 	}

// 	for i:= 0; i<len(buf); i++{
// 		fmt.Print(string(buf[i]))
// 	}
// 	fmt.Println()
// 	fmt.Println(d, buf) // 10 [72 101 108 108 111 32 71 111 108 97]
// }

// -------------------------------------------------------------

//SIMPLE WAY TO READ FILE
// func main(){
// 	// ReadFile loads all content in memory
// 	data, err := os.ReadFile("example.txt")  // If the files are too big than it could not get enough resources
// 	if err != nil{
// 		panic(err)
// 	}

// 	fmt.Println(string(data)) // Hello Golang!

// }

//-------------------------------------------------------------------

//READING FOLDERS
// func main(){
// 	dir, err := os.Open("../")
// 	if err != nil{
// 		panic(err)
// 	}

// 	defer dir.Close()

// 	fileInfo, err := dir.ReadDir(0) // if >=0 return all, if <0 return than n no. of files

// 	for _, fi := range(fileInfo){
// 		fmt.Println(fi.Name())
// 	}
// }

// CREATING A FILE AND WRITING
// func main(){
// 	f, err := os.Create("example2.txt")
// 	if err != nil{
// 		panic(err)
// 	}

// 	defer f.Close()

// 	fmt.Println("Writing in", f.Name())
// 	f.WriteString("hii GO ")
// 	f.WriteString("How do you do")  //appending

// 	// another way using stream
// 	bytes := []byte("Wassup go")
// 	f.Write(bytes)  //append

// }

//---------------------------------------------

//READ AND WRITE TO ANOTHER FILE IN STREAM FASHION
// WE CAN ALSO DIRECTLY COPY WITHOUT DOING ALL THIS USING COPY function (read about it)
// func main(){
// 	sourceFile, err := os.Open("example.txt")
// 	if err != nil{
// 		panic(err)
// 	}	

// 	defer sourceFile.Close()

// 	destFile, err := os.Create("example2.txt")
// 	if err != nil{
// 		panic(err)
// 	}	

// 	defer destFile.Close()

// 	reader := bufio.NewReader(sourceFile)
// 	writer := bufio.NewWriter(destFile)

// 	for {
// 		b, err := reader.ReadByte()
// 		if err != nil{
// 			if err.Error() != "EOF"{
// 				panic(err)
// 			}
// 			break
// 		}
// 		e := writer.WriteByte(b)
// 		if e != nil{
// 			panic(err)
// 		}

// 	}
// 	writer.Flush()

// 	fmt.Println("File written succesfully")

// }


//----------------------------------------------------------

// DELETING A FILE
func main(){
	err := os.Remove("example2.txt")
	if err != nil{
		panic(err)
	}

	fmt.Println("File deleted successfully")	
}