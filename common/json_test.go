package main

import (
	"testing"
	"encoding/json"
	"bytes"
	"fmt"
	"crypto/rand"
	"errors"
)

func TestJsonEncoder(t *testing.T) {
	writer := new(bytes.Buffer)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode([]int64{12131, 131312})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(writer.Bytes()))
}

type People struct {
	Name string
	Age  string
	Id   []byte
}

func (p *People) MarshalJson() ([]byte, error) {
	result := make(map[string]interface{})
	result["Age"] = p.Age
	result["Name"] = p.Name
	rand.Read(result["Id"].([]byte))
	return json.Marshal(result)
}

func (p *People) UnmarshalJson(input []byte) error {
	if len(input) == 0 {
		return errors.New("input byte array can't be empty")
	}
	mapper := make(map[string]interface{})
	json.Unmarshal(input, mapper)
	p.Name = mapper["Name"].(string)
	p.Age = mapper["Age"].(string)
	p.Id = mapper["Id"].([]byte)
	return nil
}
func TestCustomMarshal(t *testing.T) {
	temp := make([]byte, 32)
	rand.Read(temp)
	p := &People{
		Name: "刘飞",
		Age:  "25",
		Id:   temp,
	}
	jsonStr, _ := json.Marshal(p)
	fmt.Println(string(jsonStr))
	p2 := &People{}
	json.Unmarshal([]byte(jsonStr), p2)
	jsonStr, _ = json.Marshal(p2)
	fmt.Println(string(jsonStr))
}
