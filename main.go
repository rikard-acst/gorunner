package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 0 {
		log.Fatal("usage: gorunner <progname> [args]...")
	}

	os.Args = os.Args[1:]

	cmd := exec.Command(os.Args[0], getArgs(os.Args)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func getArgs(progAndArgs []string) []string {
	if len(progAndArgs) == 1 {
		return nil
	}

	return progAndArgs[1:]
}
