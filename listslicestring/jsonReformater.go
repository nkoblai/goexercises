package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{"one":[1,2,3,4,5],"three":[true,false,true,false,true],"two":["1","2","3","4","5"]}`)
	response := make(map[string]interface{})
	var rez []map[string]interface{}
	json.Unmarshal(b, &response)
	flag := false
	for k, v := range response {
		vt := v.([]interface{})
		if !flag {
			flag = true
			for range vt {
				rez = append(rez, make(map[string]interface{}))
			}
		}
		for i := range vt {
			rez[i][k] = vt[i]
		}
	}
	fmt.Println(rez)
}
