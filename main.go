package main

import (
	"log"
	"os"
	"os/user"
	"path"
)

func main() {
	home, err := getHome()
	if err != nil {
		log.Fatal(err)
	}

	sshConfigPath := path.Join(home, ".ssh", "config")

	sshConfigReader, err := os.Open(sshConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	defer sshConfigReader.Close()

	results, err := check(sshConfigReader, tcpTryConnect)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(os.Stdout)
	for r := range results {
		log.Print(simpleStringFormatter(r))
	}
}

func getHome() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return user.HomeDir, nil
}
