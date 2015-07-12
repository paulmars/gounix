package main

import "fmt"
import "syscall"
// import "io/ioutil"
// import "os/exec"

func main() {
  fmt.Printf("starting\n")
  // files, _ := ioutil.ReadDir("./")
  // for _, f := range files {
  //   fmt.Println(f.Name())
  // }

  // RB_AUTOBOOT
  syscall.Syscall(55,55,55,55)
}
