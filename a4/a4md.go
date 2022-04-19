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
  fmt.Println(argLength);
  if argLength < 1 {
      fmt.Printf("Please Specify the mdfile name you wanna transfer !")
      os.Exit(-1)
  }

  // #2 get file name
  var mdfile string
  var path string
  var comment bool
  comment = false;
  mdfile = os.Args[1]


  if mdfile == "-d" {
    // enable commends (disqus)
    // a4md -d path filename
    argpath := os.Args[2];
    mdfile = os.Args[3];
    comment = true;
    // deal with the path (do not check)
    bpath := []byte(argpath)
    path = string(bpath[:])
  }


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

  if comment {
    // concatenation
    urlpre := "this.page.url = ";
    urlpre += "\"https://angold4.org/";
    urlpre += path;
    urlpre += "/";
    urlpre += htmlfile;
    urlpre += "\"\n";
    fmt.Println(urlpre);
    // "https://angold4.org/path/htmlfile

    idpre := "this.page.identifier = ";
    idpre += "\"";
    idpre += path;
    idpre += "/";
    idpre += htmlfile;
    idpre += "\"\n";
    fmt.Println(idpre);

    // combine them a1 + a2 + md + urlpre + idpre + a3
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

    article3, err3 := os.ReadFile(usrdir + "/Library/a4md/article3.html")
    // article3, err3 := os.ReadFile("article3.html")
    if err3 != nil {
	log.Fatal(err3)
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

    htmlf.WriteString(urlpre);
    htmlf.WriteString(idpre);

    if _, errw3 := htmlf.Write(article3); err != nil {
	log.Fatal(errw3)
    }

    if htmlerr := htmlf.Close(); err != nil {
	log.Fatal(htmlerr)
    }

  } else {
    // no commends
    article1, err1 := os.ReadFile(usrdir + "/Library/a4md/article1.html")
    // article1, err1 := os.ReadFile("article1.html")
    if err1 != nil {
	log.Fatal(err1)
    }

    article4, err4 := os.ReadFile(usrdir + "/Library/a4md/article4.html")
    // article4, err4 := os.ReadFile("article4.html")

    if err4 != nil {
	log.Fatal(err4)
    }

    if _, errw1 := htmlf.Write(article1); err != nil {
	log.Fatal(errw1)
    }

    if _, errwmd := htmlf.Write(buf.Bytes()); err != nil {
	log.Fatal(errwmd)
    }

    if _, errw2 := htmlf.Write(article4); err != nil {
	log.Fatal(errw2)
    }

    if htmlerr := htmlf.Close(); err != nil {
	log.Fatal(htmlerr)
    }
  }
}
