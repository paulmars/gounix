package main

import "fmt"
// import "syscall"
// import "io/ioutil"
// import "os/exec"

// import "github.com/mitchellh/go-ps"

import "github.com/pusher/pusher-http-go"

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

  // syscall.Kill(341, syscall.SIGINT)

  client := pusher.Client{
    AppId: "130573",
    Key: "bb9289e6521feadf6999",
    Secret: "1e97d762ff900c953e06",
  }

  data := map[string]string{"message": "hello world"}

  client.Trigger("test_channel", "my_event", data)

}
