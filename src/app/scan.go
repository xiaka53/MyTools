package app

import (
	"MyTools/src/utils"
	"fmt"
	"strings"
)

const (
	InputCmdOpt int = iota
	InputCmdServer
	InputCmdGroupPrefix
)

var defaultServer = ""

func scanInput() {
	cmd, inputCmd, extInfo := checkInput()
	switch inputCmd {
	case InputCmdOpt:
		operation := operations[cmd]
		if operation.Process != nil {

		}
	}
}

func checkInput() (cmd string, inputCmd int, extInfo interface{}) {
	for {
		ipt := ""
		skipOut := false
		if defaultServer == "" {
			utils.Scanln(&ipt)
		} else {
			ipt = defaultServer
			defaultServer = ""
			skipOut = true
		}

		ipts := strings.Split(ipt, " ")
		cmd = ipts[0]

		if !skipOut {
			if _, exists := operations[cmd]; exists {
				inputCmd = InputCmdOpt
				extInfo = ipts[1:]
				break
			}
		}
		fmt.Println(cmd)
	}
	return
}
