package main

import (
    "io/ioutil"
    "bytes"
    "github.com/yuin/goldmark"
    "log"
    "os"
    "fmt"
)


func main() {
    // #1 check arguments
    argLength := len(os.Args[1:])
    if argLength != 1 {
	fmt.Printf("Please Specify the mdfile name you wanna transfer !")
	os.Exit(-1)
    }

    // #2 get file name
    var mdfile string
    mdfile = os.Args[1]


    // #3 create new html file
    mdnamebytes := []byte(mdfile)
    var namebytes []byte

    for _, b := range mdnamebytes {
	if b == '.' {
	    break;
	} else {
	    namebytes = append(namebytes, b)
	}
    }
    htmlfile := string(namebytes[:])
    if htmlfile == mdfile {
	fmt.Printf("a4md only support converting markdown format files !")
	os.Exit(-1)
    }

    htmlfile += ".html"
    fmt.Printf(htmlfile)

    // #4 md -> html (using goldmark)
    var buf bytes.Buffer
    content, err := ioutil.ReadFile(mdfile)
    if err != nil {
	log.Fatal(err)
    }

    if err := goldmark.Convert(content, &buf); err != nil {
	panic(err)
    }

    usrdir, uerr := os.UserHomeDir()
    if uerr != nil {
	log.Fatal(uerr)
    }

    // #5 Final step, combine them into the target html file
    htmlf, err := os.OpenFile(htmlfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

    article1, err1 := os.ReadFile(usrdir + "/Library/a4md/article1.html")
    // article1, err1 := os.ReadFile("article1.html")
    if err1 != nil {
	log.Fatal(err1)
    }

    article2, err2 := os.ReadFile(usrdir + "/Library/a4md/article2.html")
    // article2, err2 := os.ReadFile("article2.html")
    if err2 != nil {
	log.Fatal(err2)
    }

    if _, errw1 := htmlf.Write(article1); err != nil {
        log.Fatal(errw1)
    }

    if _, errwmd := htmlf.Write(buf.Bytes()); err != nil {
	log.Fatal(errwmd)
    }

    if _, errw2 := htmlf.Write(article2); err != nil {
        log.Fatal(errw2)
    }

    if htmlerr := htmlf.Close(); err != nil {
        log.Fatal(htmlerr)
    }
}
