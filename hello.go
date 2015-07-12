package main

import "fmt"
import "syscall"
import "io/ioutil"

func main() {
  fmt.Printf("starting\n")
  files, _ := ioutil.ReadDir("./")
  for _, f := range files {
    fmt.Println(f.Name())
  }

  syscall.Reboot()
}
