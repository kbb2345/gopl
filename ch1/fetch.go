package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url,"http://") {
			url = "http://" + url
		}
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		out, err := os.Create("./gopl.html")
		wt := bufio.NewWriter(out)
		defer out.Close()
		n, err := io.Copy(wt, resp.Body)
		// b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println(n)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:reading%s:%v\n", url, err)
			os.Exit(1)
		}
		wt.Flush()
		// fmt.Printf("%s", b)
	}
}
