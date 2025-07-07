package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	type Profile struct {
		Age   int
		Place string
	}
	type Person struct {
		Name    string
		Profile Profile
	}
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"case with 1 string values",
			struct {
				name string
			}{"Harsh"},
			[]string{"Harsh"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with one integer fields",
			struct {
				Name string
				age  int
			}{"Chris", 24},
			[]string{"Chris"},
		},
		{
			"nested structs",
			Person{
				"Harsh",
				Profile{
					22,
					"Mumbai",
				},
			},
			[]string{"Harsh", "Mumbai"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
