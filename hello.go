package main

import "fmt"
import "syscall"
// import "io/ioutil"
// import "os/exec"

// import "github.com/mitchellh/go-ps"

func main() {
  fmt.Printf("starting\n")

  // RB_AUTOBOOT
  // syscall.Syscall(55,55,55,55)

  // procs, _ := ps.Processes()
  // for _, p := range procs {
  //   fmt.Println(p.Pid())
  //   // fmt.Println(p.PPid())
  //   fmt.Println(p.Executable())
  // }

  syscall.Kill(341, syscall.SIGINT)
}
