// Showing source of page by URL

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		src, err := ioutil.ReadAll(resp.Body)
		dst := os.Stdout
		b, err := io.Copy(dst, src)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: odczytywanie %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
