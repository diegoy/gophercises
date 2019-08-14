package main

import (
  "os"
  "bufio"
  "enconding/csv"
)

func main() {
  file, err := os.Open("problems.csv")
  if err != nil {
    println("Oops alguma coisa deu errada.")
  } else {
    println("File is opened!")
    reader := bufio.NewReader(file)
    line, _ := reader.ReadString('\n')
    print(line)
    println("oi mundo")
  }
}
