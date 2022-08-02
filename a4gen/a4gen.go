package main

import (
    // "bytes"
    // "log"
    "os/exec"
    // "os"
    "fmt"
    "strings"
    "regexp"
)


func main() {
  pwd := "pwd"
  cmd := exec.Command(pwd)
  stdout, err := cmd.Output();

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  i := strings.Index(string(stdout), "a4org.github.io")
  comment := string(stdout)[i+16:]

  fmt.Println(comment)

  ls := "ls"
  cmd = exec.Command(ls)
  stdout, err = cmd.Output()

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  var tmp string;
  files := []string{}
  for _, ch := range(string(stdout)) {
    if (string(ch) == "\n") {
      files = append(files, tmp)
      tmp = ""
    } else {
      tmp += string(ch)
    }
  }

  mdfiles := []string{}
  rem := regexp.MustCompile(`.md`)
  for _, f := range(files) {
    if (rem.Match([]byte(f))) {
      mdfiles = append(mdfiles, f)
    }
  }

  a4md := "a4md"
  arg0 := "-d"
  arg1 := comment
  arg2 :=  ""
  for _, mf := range(mdfiles) {
    fmt.Println(mf)
    arg2 = mf
    cmd = exec.Command(a4md, arg0, arg1, arg2)
    cmd.Output()
  }
}
