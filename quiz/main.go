package main

import (
  "os"
  "encoding/csv"
  "fmt"
  "flag"
)

func main() {
  fileName := flag.String("file", "problems.csv", "file name")
  flag.Parse()

  file, err := os.Open(*fileName)

  if err != nil {
    println("Oops alguma coisa deu errada.")
    println(err.Error())
  } else {
    println("File is opened!")
    csvReader := csv.NewReader(file)

    for {
      line, err := csvReader.Read()
      if err != nil && err.Error() == "EOF" {
        println("End of file bye!")
        break
      }
      fmt.Printf("Question: %s, Answer: %s\n", line[0], line[1])
    }
  }
}

func filename() {
}
