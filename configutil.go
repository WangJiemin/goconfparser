package goconfparser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ConfigReadConfig(filepath string) map[string]string {
	res := map[string]string{}
	file, err := os.Open(filepath)
	if err != nil {
		return res
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				return res
			}
			if len(line) == 0 {
				break
			}
		}

		if len(line) == 0 || line == "\r\n" {
			//break
			continue
		}

		if line[0] == '/' {
			continue
		}

		fmt.Println(line)
		i := strings.IndexAny(line, "=")
		value := strings.TrimSpace(line[i+1 : len(line)])
		res[strings.TrimSpace(line[0:i])] = value
	}
	return res
}
