package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type myJson struct {
	Intvalue        int       `json:"intvalue"`
	Stringvalue     string    `json:"stringvalue"`
	BoolValue       bool      `json:"boolvalue"`
	DateValue       time.Time `json:"time"`
	ObjectValue     *myObj    `json:"objectvalue"`
	NullIntValue    *int      `json:"nullIntvalue,omitempty"`
	NullStringValue *string   `json:"nullstringvalue"`
}

type myObj struct {
	ArrayValue []int
}

var otherint = 123

func main() {
	data := &myJson{
		Intvalue:    1,
		Stringvalue: "Harsh",
		BoolValue:   true,
		DateValue:   time.Date(2022, 3, 2, 8, 10, 0, 0, time.UTC),
		ObjectValue: &myObj{
			ArrayValue: []int{1, 2, 3, 5},
		},
		NullIntValue:    nil,
		NullStringValue: nil,
	}
	marshalData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error processing json data: %v", err)
	}
	fmt.Printf("Json data: %s\n", marshalData)
}
