package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// http://stackoverflow.com/questions/3356011/whats-gos-equivalent-of-argv0
	if len(os.Args) < 2 {
		panic("ERROR: ")
	}

	var argv []string
	for i := range os.Args {
		fmt.Printf("cmd[%d] : %s\n", i, os.Args[i])

		switch os.Args[i] {
		case "log":
			argv = append(argv, "log")
			argv = append(argv, "--abbrev-commit")
			argv = append(argv, "--name-status")

		case "logt":
			argv = append(argv, "log")
			argv = append(argv, "--graph")
			argv = append(argv, "--abbrev-commit")
			argv = append(argv, "--name-status")

		case "b":
			argv = append(argv, "branch")

		case "s":
			argv = append(argv, "status")

		case "co":
			argv = append(argv, "log")

		case "cob":
			argv = append(argv, "log")

		case "cob2":
			argv = append(argv, "log")

		case "delb":
			argv = append(argv, "log")

		case "m":
			argv = append(argv, "log")

		case "mp":
			argv = append(argv, "log")

		case "mpp":
			argv = append(argv, "log")

		case "ci":
			argv = append(argv, "log")

		case "cip":
			argv = append(argv, "log")

		case "df":
			argv = append(argv, "log")

		case "p":
			argv = append(argv, "log")

		case "j":
			argv = append(argv, "log")

		case "cp":
			argv = append(argv, "log")

		}
	}

	fmt.Printf("git cmd : %s\n", gitCommand(argv))

	// fmt.Println("ok argv %s", argv[0])
}

func gitCommand(gitcmd []string) string {
	cmd := exec.Command("git", gitcmd...)

	var outb bytes.Buffer
	cmd.Stdout = &outb
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()

	fmt.Println(outb.String())

	return outb.String()
}

func printLog() {
	// var currentUrl string
	// var masterBranch string
	// var currentBranch string
	// var mergeBranch string
	//
	// print("\n\n")
	//
	// print("*[Change info]*")
	// print("{noformat}")
	//
	// print("{noformat}")
	// print("")
	// print("*[Modified Files]*")
	// print("{noformat}")
	// //
	// print("{noformat}")
	// print("*[Branch]*")
	// print("{noformat}")
	// //
	// print("{noformat}")
	//
	// print("URL           : %s", currentUrl)
	// print("Master Branch : %s", currentUrl)
	// print("Merge Branch  : %s", currentUrl)
	// print("Dev Branch    : %s", currentUrl)
	//
	// print("{noformat}")
	// print("\n\n")
}
