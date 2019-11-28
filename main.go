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

	reader := bufio.NewReader(os.Stdin)

	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		line := string(input)
		if strings.Contains(line, "cannot find package") {
			slice := strings.Split(line, `"`)
			// get package url
			packageUrl := slice[1]

			// build go get -u command
			get := exec.Command(gobin, "get", "-u", packageUrl)

			message := "Downloading " + packageUrl

			if err := get.Run(); err != nil {
				message = message + " : " + err.Error()
			}

			fmt.Println(message)
		}
	}
}
