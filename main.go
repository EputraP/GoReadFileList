package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "Uploading File\n")
	// 1.parse input, type multipart/form data
	r.ParseMultipartForm(10 << 20)
	fmt.Println(r)
	// 2. retrieve file from posted form data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving file from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Upload File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("NIME Header: %+v\n", handler.Header)
	// 3. write temporary file on our server
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("temp-images/"+handler.Filename, fileBytes, 0666)
	if err != nil {
		fmt.Println("Error writing to temporary file:", err)
		return
	}
	fmt.Printf("type: %T\n", tempFile.Name())
	fmt.Println("name: ", tempFile.Name())

	// err = os.Remove("temp-images/" + tempFile.Name())
	substrings := strings.Split(tempFile.Name(), `\`)
	substring := substrings[1]
	fmt.Println("substring: ", substring)
	// time.Sleep(30 * time.Second)
	tempFile.Close()
	err = os.Remove("temp-images/" + substring)
	if err != nil {
		fmt.Println("Error removing old file:", err)
		return
	}
	// tempFile.Write(fileBytes)
	// err = os.Rename(tempFile.Name(), handler.Filename)
	// if err != nil {
	// 	fmt.Println("Error renaming temporary file:", err)
	// 	return
	// }
	// 4. return wether or not this has been successful
	fmt.Fprintf(w, "successfully Uploaded File")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Go File Upload Tutorial")
	setupRoutes()
}

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// )

// func main() {
// 	files, err := ioutil.ReadDir("temp-images/")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	for _, file := range files {
// 		fmt.Println(file.Name(), file.Size())
// 	}
// }
