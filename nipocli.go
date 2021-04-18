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
	ok := false
	if len(cmdFields) == 1 && cmdFields[0] == "ping" {
		result,ok = nipolib.Ping(config)
		if !ok {
			fmt.Println("Error on connecting to server ");
		}
	}
	if len(cmdFields) == 1 && cmdFields[0] == "status" {
		result,ok = nipolib.Status(config)
		if !ok {
			fmt.Println("Error on connecting to server ");
		}
	}
    if len(cmdFields) >= 2 {
        switch cmdFields[0] {
        case "set":
			value := ""
			for count:=2 ; count < len(cmdFields); count++ {
				value += cmdFields[count]+" "
			}
            result,ok = nipolib.Set(config, cmdFields[1], value)
			if !ok {
				fmt.Println("Error on connecting to server ");
			}
            break
        case "get":
			keys := ""
			for count:=1 ; count < len(cmdFields); count++ {
				keys += cmdFields[count]+" "
			}
            result,ok = nipolib.Get(config, keys)
			if !ok {
				fmt.Println("Error on connecting to server ");
			}
            break
		case "sum":
            result,ok = nipolib.Sum(config, cmdFields[1])
			if !ok {
				fmt.Println("Error on connecting to server ");
			}
			break
        case "select":
            result,ok = nipolib.Select(config, cmdFields[1])
			if !ok {
				fmt.Println("Error on connecting to server ");
			}
            break
        case "avg":
            result,ok = nipolib.Avg(config, cmdFields[1])
			if !ok {
				fmt.Println("Error on connecting to server ");
			}
			break
		case "count":
            result,ok = nipolib.Count(config, cmdFields[1])
			if !ok {
				fmt.Println("Error on connecting to server ");
			}
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