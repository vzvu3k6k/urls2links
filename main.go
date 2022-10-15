package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var b strings.Builder

	fmt.Fprint(&b, "<ul>")

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := s.Text()
		fmt.Fprintf(&b, "<li><a href=\"%s\">%s</a>", t, t)
	}

	fmt.Fprint(&b, "</ul>")

	fmt.Printf("data:text/html,%s", b.String())
}
