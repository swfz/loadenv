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

	if len(flag.Args()) < 2 {
		fmt.Println("not enought argument!")
		fmt.Println("USAGE: loadenv envfilename command [option...]")
		os.Exit(1)
	}

	// fmt.Println("args:", flag.Args())
	envfile := flag.Args()[0]
	fp, err := os.Open(envfile)

	if err != nil {
		fmt.Println("load file error!")
		os.Exit(1)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		// fmt.Println("line:", scanner.Text())
		env := strings.Split(scanner.Text(), "=")

		if len(env) < 2 {
			fmt.Println("load env error!!", scanner.Text())
			continue
		}

		os.Setenv(env[0], env[1])
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("scan line error!")
		os.Exit(1)
	}

	//for _, e := range os.Environ() {
	//    pair := strings.Split(e, "=")
	//    fmt.Println(pair[0],pair[1])
	//}
	// command := exec.Command(strings.Join(flag.Args()[1:], ","))

	command := exec.Command(flag.Args()[1], flag.Args()[2:]...)
	// fmt.Println("%#v", command)
	command.Stdin  = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
}
