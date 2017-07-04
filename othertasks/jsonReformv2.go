package jsonreformv2

import (
	"encoding/json"
	"fmt"
)

// Builder makes correct form of output and returns it.
func Builder(i int, r map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for k, vi := range r {
		switch v := vi.(type) {
		case []interface{}:
			for range v {
				res[k] = v[i]
			}
		case map[string]interface{}:
			res[k] = Builder(i, v)
		default:
			res[k] = v
		}
	}
	return res
}

// LenFinder finds slice type and returns his len.
func LenFinder(r map[string]interface{}, n *int) {
	for _, vi := range r {
		switch v := vi.(type) {
		case []interface{}:
			*n = len(v)
		case map[string]interface{}:
			LenFinder(v, n)
		default:
			continue
		}
	}
}

func main() {
	input := []byte(`{"id":"1234","forecast":{"temp":[14,15,26,25,16],"day":{"moon":[2,2,3,3,2],"wind":[23,24,17,17,25]}}}`)
	response := make(map[string]interface{})
	json.Unmarshal(input, &response)
	n := 0
	LenFinder(response, &n)
	var rezult []map[string]interface{}
	for i := 0; i < n; i++ {
		rez := Builder(i, response)
		rezult = append(rezult, rez)
	}
	fmt.Println(rezult)
}
