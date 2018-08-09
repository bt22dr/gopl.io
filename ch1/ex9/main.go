// ex7
package main

import (
	"fmt"
        "io"
	"net/http"
        "strings"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
                if !strings.HasPrefix(url, "http://") {
                    url = "http://" + url
                }
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
                _, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR!! (%s)\n", err)
			os.Exit(1)
		}

                fmt.Println(resp.Status)
	}
}

//!-
