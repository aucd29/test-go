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

    switch os.Args[1] {
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
      argv = append(argv, os.Args[2])

    case "s":
      argv = append(argv, "status")

    case "co":
      argv = append(argv, "checkout")
      argv = append(argv, os.Args[2])

    case "cob":
      argv = append(argv, "checkout")
      argv = append(argv, "-b")
      argv = append(argv, os.Args[2])

    case "cob2":
      prefix := "dev_aucd29_"
      prefix = append(prefix, os.Args[2])

      argv = append(argv, "log")
      argv = append(argv, "-b")
      argv = append(argv, prefix)

    case "delb":
      argv = append(argv, "branch")
      argv = append(argv, "-d")
      argv = append(argv, os.Args[2])

    // case "m":
    // 	argv = append(argv, "log")
    //
    // case "mp":
    // 	argv = append(argv, "log")
    //
    // case "mpp":
    // 	argv = append(argv, "log")
    //
    case "ci":
      argv = append(argv, "commit")

    case "cip":
      argv = append(argv, "commit")
      gitCommand(argv)

      var s_argv []string
      s_argv = append(s_argv, "push")
      s_argv = append(s_argv, "origin")
      s_argv = append(s_argv, "$")
      gitCommand(s_argv)

    case "df":
      argv = append(argv, "diff")
      argv = append(argv, ">")

      name := "$"
      name = append(name, ".patch")

      argv = append(argv, name)

    case "p":
      argv = append(argv, "push")
      argv = append(argv, "origin")
      argv = append(argv, "$")

    case "j":
      argv = append(argv, "log")

    case "cp":
      argv = append(argv, "cherry-pick")
      argv = append(argv, os.Args[2])
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
