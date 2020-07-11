package main

import (
    "html/template"
    "os"
    "log"
    "fmt"
)

type File struct {
    Name string
    Directory bool
}

type DirectoryPage struct {
    PageTitle string
    BaseDir string
    FileNames []File
}

func GenerateDirectoryPage(baseDir string, files []os.FileInfo, outputDir string) bool {
    tmpl := template.Must(template.ParseFiles("src/templates/DirectoryPage.gohtml", "src/templates/NavigationBar.gohtml"))

    FileNames := []File {}
    for _, file := range files {
        FileNames = append(FileNames, File{
            Name: file.Name(),
            Directory: file.IsDir(),
        })
    }
    data := DirectoryPage{
        BaseDir: baseDir,
        FileNames: FileNames,
    }

    f, err := os.Create(outputDir + "/index.html")
    if err != nil {
        log.Println("create file: ", err)
        return false
    }
    tmpl.ExecuteTemplate(f, "directorypage", data)
    return true
}


type DetailPage struct {
    BaseDir string
    FileName string
    Lines []string
}

func GenerateDetailPage(root string, title string, lines []string, outputFileName string) bool {
    data := DetailPage{
        BaseDir: root,
        FileName: title,
        Lines: lines,
    }
    tmpl := template.Must(template.ParseFiles("src/templates/DetailsPage.gohtml", "src/templates/NavigationBar.gohtml"))
    fmt.Println("helo", outputFileName)
    f, err := os.Create(outputFileName)
    if err != nil {
        log.Println("create file: ", err)
        return false
    }
    tmpl.ExecuteTemplate(f, "detailpage", data)
    return true
}

