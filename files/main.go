package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "log"
)

func main() {
    file, err := os.Open("dat.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()


  b, err := ioutil.ReadAll(file)
  fmt.Print(b)
}