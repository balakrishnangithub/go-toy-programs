package main

import (
	"log"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	rmDupeGenFile(usr.HomeDir+"/.bash_history", ".bash_history__unique")
	printArgsFile()
}
