package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	flag.Parse()
	fmt.Println("args:", flag.Args())

	envfile := flag.Args()[0]
	fp, err := os.Open(envfile)

	if err != nil {
		fmt.Println("load file error!")
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		fmt.Println("line:", scanner.Text())

		env := strings.Split(scanner.Text(), "=")
		os.Setenv(env[0], env[1])
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("scan line error!")
	}

	//for _, e := range os.Environ() {
	//    pair := strings.Split(e, "=")
	//    fmt.Println(pair[0],pair[1])
	//}
	// command := exec.Command(strings.Join(flag.Args()[1:], ","))
	command := exec.Command(flag.Args()[1], flag.Args()[2:]...)
	fmt.Println("%#v", command)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
}