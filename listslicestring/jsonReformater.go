package main

import (
	"encoding/json"
	"fmt"
)

// Response represents body of JSON.
type Response struct {
	One   []int
	Three []bool
	Two   []string
}

// Reformated represents reformated type.
type Reformated struct {
	One   int
	Three bool
	Two   string
}

// JSONReform takes JSON body and returns slice of Reformated type.
func JSONReform(res Response) []Reformated {
	ref := make([]Reformated, len(res.One))
	for i := range res.One {
		ref[i] = Reformated{One: res.One[i], Three: res.Three[i], Two: res.Two[i]}
	}
	return ref
}

func main() {
	b := []byte(`{"one":[1,2,3,4,5],"three":[true,false,true,false,true],"two":["1","2","3","4","5"]}`)
	res := Response{}
	json.Unmarshal(b, &res)
	fmt.Println(JSONReform(res))
}
