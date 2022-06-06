package main

import (
	"bufio"
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	p, _ := build.Default.Import("github.com/twreporter/go-api", "", build.FindOnly)

	fname := filepath.Join(p.Dir, "CHANGELOG.md")
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println("Cannot open CHANGELOG.md")
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile(`#{1,}\s+(?:\d+\.)(?:\d+\.)(?:\d+)`)
	for scanner.Scan() {
		ver := re.FindString(scanner.Text())
		if ver != "" {
			// remove #
			ver = strings.Replace(ver, "#", "", -1)
			// remove whitespace
			ver = strings.Replace(ver, " ", "", -1)
			fmt.Print(ver)
			break
		}
	}
}
