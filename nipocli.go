package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/NipoDB/nipolib"
	"strings"
)

func checkCmd(cmd string, config *nipolib.Config) string {
	cmdFields := strings.Fields(cmd)
    result := ""
	if len(cmdFields) == 1 && cmdFields[0] == "ping" {
		result,_ = nipolib.Ping(config)
	}
	if len(cmdFields) == 1 && cmdFields[0] == "status" {
		result,_ = nipolib.Status(config)
	}
    if len(cmdFields) >= 2 {
        switch cmdFields[0] {
        case "set":
			value := ""
			for count:=2 ; count < len(cmdFields); count++ {
				value += cmdFields[count]+" "
			}
            result,_ = nipolib.Set(config, cmdFields[1], value)
            break
        case "get":
			keys := ""
			for count:=1 ; count < len(cmdFields); count++ {
				keys += cmdFields[count]+" "
			}
            result,_ = nipolib.Get(config, keys)
            break
		case "sum":
            result,_ = nipolib.Sum(config, cmdFields[1])
			break
        case "select":
            result,_ = nipolib.Select(config, cmdFields[1])
            break
        case "avg":
            result,_ = nipolib.Avg(config, cmdFields[1])
			break
		case "count":
            result,_ = nipolib.Count(config, cmdFields[1])
			break
        }
    } 
	return result
}

func Start(config *nipolib.Config) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Woclome to nipo")
	for {
		fmt.Print("nipo > ")
		var char byte
		cmd := ""
		var err error
		for char != byte('\n'){
			char, err = reader.ReadByte()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			cmd += string(char)
		} 
		result := checkCmd(cmd, config)
		fmt.Print(result)
	}
}