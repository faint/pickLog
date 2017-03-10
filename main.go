package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	read = flag.String("read", "moeMobile.log", "file to read")
	key  = flag.String("search", "key", "key to search")
)

func main() {
	flag.Parse()
	fmt.Println("pick start!")
	b, e := ioutil.ReadFile(*read)
	if e != nil {
		log.Fatal(e.Error())
	}
	r := bytes.NewReader(b)
	scanner := bufio.NewScanner(r)
	var s string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, *key) {
			s = s + line + "\n"
		}
	}

	file, e := os.Create(*key + ".log")
	if e != nil {
		log.Fatal(e.Error())
	}
	file.WriteString(s)
	fmt.Println("pick ended!")
}

/*
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

*/
