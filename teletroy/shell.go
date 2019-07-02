package main

import (
	"./teletroy"
	"fmt"
	"github.com/mattn/go-shellwords"
	"os/exec"
)

func runCmdStr(cmdstr string) (string, error) {
	c, err := shellwords.Parse(cmdstr)
	if err != nil {
		return "", err
	}
	switch len(c) {
	case 0:
		return "", nil
	case 1:
		out, err := exec.Command(c[0]).Output()

		return string(out), err
	default:
		out, err := exec.Command(c[0], c[1:]...).Output()
		return string(out), err
	}
	if err != nil {
		return "", err
	}
	return "", nil
}
func main() {
	out, err := runCmdStr("ls /tmp | grep py ")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(out)
	}
	fmt.Println(telebot.Me())
	fmt.Println(telebot.Me())
	fmt.Println(telebot.Me())
}
