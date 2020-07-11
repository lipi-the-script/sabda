package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "strings"
//     "net/http"
)

func GetFiles(root string) []os.FileInfo {
    var (
        files []os.FileInfo
        filteredFiles []os.FileInfo
        err   error
    )

    ignore := []string{".git", ".idea", "src", "README", "pages", ".DS_Store"}

    files, err = ioutil.ReadDir(root)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        ignoreFile := false
        for _, strignore := range ignore {
            if strings.Contains(file.Name(), strignore) {
                ignoreFile = true
            }
        }

        if ignoreFile {
            continue
        }

        filteredFiles = append(filteredFiles, file)
    }
    return filteredFiles
}

func GenerateFiles(root string, outputRoot string) bool{
    var (
        files []os.FileInfo
    )

    files = GetFiles(root)
    GenerateDirectoryPage(root, files, outputRoot)

    for _, file := range files {
        if file.IsDir() {
               folderToCreate := outputRoot+"/"+file.Name()
             // create folder
             os.Mkdir(folderToCreate, os.ModePerm)
            // recurse
            GenerateFiles(root + "/" + file.Name(), folderToCreate)
        } else {
            // generate lyrics page
            outputFileName := outputRoot + "/" + file.Name() + ".html"
            fileToRead := root + "/" + file.Name()

            data, err := ioutil.ReadFile(fileToRead)
            if err != nil {
                fmt.Println("File reading error", err)
                return false
            }
            lines := strings.Split(string(data), "\n")
            GenerateDetailPage(root, file.Name(), lines, outputFileName)
        }
        fmt.Println(file.Name())
    }
    return true
}

func main() {
    GenerateFiles(".", "./pages")

//     http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
//         fmt.Fprint(w, GenerateDirectoryPage("/"))
//     })
//     http.ListenAndServe(":8009", nil)
}