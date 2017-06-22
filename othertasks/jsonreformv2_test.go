package jsonreformv2

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	input := []byte(`{"id":"1234","forecast":{"temp":[14,15,26,25,16],"day":{"moon":[2,2,3,3,2],"wind":[23,24,17,17,25]}}}`)
	response := make(map[string]interface{})
	json.Unmarshal(input, &response)
	res := Builder(4, response)
	result, _ := json.Marshal(res)
	exp := []byte(`{"forecast":{"day":{"moon":2,"wind":25},"temp":16},"id":"1234"}`)
	assert.Equal(t, string(exp), string(result), "The two words should be the same.")
}

func TestLenFinder(t *testing.T) {
	input := []byte(`{"id":"1234","forecast":{"temp":[14,15,26,25,16],"day":{"moon":[2,2,3,3,2],"wind":[23,24,17,17,25]}}}`)
	response := make(map[string]interface{})
	json.Unmarshal(input, &response)
	n := 0
	LenFinder(response, &n)
	if n != 5 {
		t.Errorf("Something going wrong!")
	}
}
