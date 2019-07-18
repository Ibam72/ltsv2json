package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main(){
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		keyValues := strings.Split(stdin.Text(), "\t")
		kvMap := map[string]interface{}{}
		for _,keyValue := range keyValues {
			tmp := strings.Split(keyValue, ":")
			if len(tmp)!= 2 {
				fmt.Fprintf(os.Stderr, "{\"err\": \"wrong format\"}\n")
				continue
			}
			// TODO 型アサーション
			kvMap[tmp[0]]=tmp[1];
		}
		jsonBytes, err := json.Marshal(kvMap)
		if err != nil {
			fmt.Fprintf(os.Stderr, "{\"err\": \"%v\"}\n", err)
			continue
		}
		fmt.Println(string(jsonBytes))
	}
}
