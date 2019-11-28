package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {

	gobin, err := exec.LookPath("go")

	if err != nil {
		panic(err)
	}

    fi, _ := os.Stdin.Stat()

    if (fi.Mode() & os.ModeCharDevice) != 0 {
    	fmt.Println("'modify go build' (if needed) and exec this command")
        fmt.Println("go build 2>&1 >/dev/null | GoMiss")
        os.Exit(0)
    }

	reader := bufio.NewReader(os.Stdin)

	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		line := string(input)
		if strings.Contains(line, "cannot find package") {
			slice := strings.Split(line, `"`)

			if len(slice) == 3 {
				// get package url
				packageUrl := slice[1]

				// build go get -u command
				get := exec.Command(gobin, "get", "-u", packageUrl)

				message := "Downloading " + packageUrl

				if err := get.Run(); err != nil {
					message = message + " : " + err.Error()
				}

				fmt.Println(message)
			} else {
				fmt.Println("Invalid message : " + line)
			}
		}
	}
}
