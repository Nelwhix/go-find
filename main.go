package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "go-find by Nelson Isioma")
		fmt.Fprintln(flag.CommandLine.Output(), "Copyright " + strconv.Itoa(time.Now().Local().Year()) + "\n")
	}

	flag.Parse()
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("I can only find one item at a time")
		os.Exit(1)

	}

	cmd := exec.Command("pwd")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	working_dir := out.String()

	if err := run(working_dir, args[0]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(working_dir, fileName string) error {
	return filepath.WalkDir(".", func (path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == fileName {
			if d.IsDir() {
				fmt.Println("found: " + filepath.Join(working_dir, d.Name()))
			} else {
				fmt.Println("found: " + filepath.Join(working_dir, d.Name()))
			}
		}
		
		return nil
	})
}