package postfile

// import "fmt"

// func InsertCode() {
// 	p := "green"
// 	index := 2
// 	q := p[:index] + "HI" + p[index:]
// 	fmt.Println(p, q)
// }

import (
	"fmt"
	"io"
	"os"
)

var path = "Dockerfile"

func CreateFile() {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create("fileWrite/" + path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully", path)
}

func WriteFile() {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile("fileWrite/"+path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()
	var insert string
	fmt.Scan(&insert)
	// Write some text line-by-line to file.
	_, err = file.WriteString(insert)
	if isError(err) {
		return
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("File Updated Successfully.")
}

func ReadFile() {
	// Open file for reading.
	var file, err = os.OpenFile("fileWrite/"+path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	fmt.Println("Reading from file.")
	fmt.Println(string(text))
}

func DeleteFile() {
	// delete file
	var err = os.Remove("fileWrite/" + path)
	if isError(err) {
		return
	}

	fmt.Println("File Deleted")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
