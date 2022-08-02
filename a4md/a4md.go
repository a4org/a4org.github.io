package main

import (
    "io/ioutil"
    // "bytes"
    // "github.com/yuin/goldmark"
    "log"
    "os/exec"
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
  bufile := "tmp.html"
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

  usrdir, uerr := os.UserHomeDir()
  if uerr != nil {
      log.Fatal(uerr)
  }


  // #4 md -> html (using pandoc)
  // pandoc --toc --standalone --mathjax -f markdown -t html test.md -o test.html 
  app := "pandoc"
  arg0 := "--toc"
  arg1 := "--standalone"
  arg2 := "--mathjax"
  arg3 := "-f"
  arg4 := "markdown"
  arg5 := "-t"
  arg6 := "html"
  arg7 := mdfile
  arg8 := "-o"
  arg9 := bufile

  cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
  stdout, err := cmd.Output()
  
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  // Print the output
  fmt.Println(string(stdout))

  app = usrdir + "/Library/a4md/a4filter"
  cmd = exec.Command(app)

  stdout, err = cmd.Output()
  
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  // Print the output
  fmt.Println(string(stdout))


  content, err := ioutil.ReadFile("out.html")
  if err != nil {
      log.Fatal(err)
  }


  // #5 Final step, combine them into the target html file
  htmlf, err := os.OpenFile(htmlfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
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

    if _, errwmd := htmlf.Write(content); err != nil {
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

    if _, errwmd := htmlf.Write(content); err != nil {
	log.Fatal(errwmd)
    }

    if _, errw2 := htmlf.Write(article4); err != nil {
	log.Fatal(errw2)
    }

    if htmlerr := htmlf.Close(); err != nil {
	log.Fatal(htmlerr)
    }
  }

  arg0 = "rm"
  arg1 = "-rf"
  arg2 = "tmp.html"
  arg3 = "out.html"

  cmd = exec.Command(arg0, arg1, arg2, arg3)

  stdout, err = cmd.Output()
  
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  // Print the output
  fmt.Println(string(stdout))
}
