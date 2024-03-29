package main

import (
    "io/ioutil"
    "text/template"
    "encoding/csv"
    "bytes"
    "flag"
    "os"
    "fmt"
    "strings"
    "time"
    "github.com/fatih/color"
)

func main() {
    // Starting execution timer
    start := time.Now()

    // Flags
    csvFileDir := flag.String("csv", "data.csv",
        "A CSV file containing data to be inserted into the template.")
    tmplDir := flag.String("tmpl", "template.tmpl",
        "A .tmpl file to insert data into.")
    outputDir := flag.String("output", "../index.html",
        "The .html output file.")
    flag.Parse()

    // Correcting any missing extensions for the given flags
    if strings.LastIndex(*csvFileDir, ".") < 0 {
        *csvFileDir += ".csv"
    }
    if strings.LastIndex(*tmplDir, ".") < 0 {
        *tmplDir += ".tmpl"
    }
    if strings.LastIndex(*outputDir, ".") < 0 {
        *outputDir += ".html"
    }

    // Do all the work & get filesize
    var fileSize float64
    data := getData(*csvFileDir)
    fileSize += createAndSaveFile(*tmplDir, *outputDir, data)

    // Pretty font styling
    boldGreen := color.New(color.FgGreen, color.Bold).SprintFunc()
    bold := color.New(color.Bold).SprintFunc()

    // Get execution time in seconds
    exeTime := time.Since(start).Seconds()

    // Success message :)
    fmt.Printf("%s Generated %s (%.1fkB total) in %.3f seconds.\n",
    boldGreen("[Success!]"), bold(outputDir), fileSize, exeTime)
}

// Struct of everything to be inserted into template
type portfolioData struct {
    Projects []projectData
    Articles []articleData
    Footer []footerData
}

type projectData struct {
    Title string
    Desc string
    Tools string    // languages, tools, & technologies used
    Repo string     // gh repo name
    Url string      // url to live project
}

type articleData struct {
    Title string
    Tag string
    Desc string
    Url string
}

type footerData struct {
    Icon string
    Desc string
    Url string
}


func readCsvFile(csvFile string) [][]string {
    // Open the csv file, throw err if something goes wrong
    f, err := os.Open(csvFile)
    checkErr(err, fmt.Sprintf("while opening %s!\n", csvFile))
    defer f.Close()

    // Read the csv file, throw err if something goes wrong
    csvReader := csv.NewReader(f)
    csvData, err := csvReader.ReadAll()
    checkErr(err, fmt.Sprintf("while reading %s!\n", csvFile))

    // Return 2d slice of the data from the given csv file
    return csvData
}

func getData(csvFileDir string) portfolioData {
    // Get raw data from csv
    data := readCsvFile(csvFileDir)

    // Setup vars for later
    var projects []projectData
    var articles []articleData
    var footer []footerData
    // var item itemData

    // Sort thru the data from the csv
    for _, line := range data {
        // Skip iteration if empty line or missing fields
        if line[0] == "" || len(line) < 6 {
            continue
        }

        // Append to appropriate slice
        switch line[0] {
            case "P": // Project
                projects = append(projects, projectData {
                  Title: line[1],
                  Tools: line[2],
                  Desc: line[3],
                  Repo: line[4],
                  Url: line[5],
                })
            case "A": // Article
                articles = append(articles, articleData {
                  Title: line[1],
                  Tag: line[2],
                  Desc: line[3],
                  Url: line[4],
                })
            case "F": // Footer
                footer = append(footer, footerData {
                  Icon: line[1],
                  Desc: line[2],
                  Url: line[3],
                })
        }
    }
    // Return struct of all the data packaged together for the template
    return portfolioData{Projects: projects, Articles: articles, Footer: footer}
}

func createAndSaveFile(tmplDir string, outputDir string, data portfolioData) float64 {
    // Parse template
    tmpl := template.Must(template.New(tmplDir).ParseFiles(tmplDir))

    //  Insert data into template & save bytes to buffer, throw err if something goes wrong
    buffer := new(bytes.Buffer)
    err := tmpl.Execute(buffer, data)
    checkErr(err, fmt.Sprintf("while parsing template %s!\n", tmplDir))

    // Convert write bytes into output file, throw err if something goes wrong
    bytesToWrite := []byte(buffer.String())
    err = ioutil.WriteFile(outputDir, bytesToWrite, 0644)
    checkErr(err, fmt.Sprintf("while creating the output file %s!\n", outputDir))

    // Return filesize (kB) of the new html file
    return float64(len(bytesToWrite)) / float64(1000)
}

func checkErr(err error, msg string) {
    // Check for error
    if err != nil {
        // String formatting
        boldRed := color.New(color.FgRed, color.Bold).SprintFunc()
        // Print pretty error message
        fmt.Println(boldRed("[Error]"), msg)
        // Print the raw err msg below the pretty one for more info
        panic(err)
    }
}
