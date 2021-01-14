package reflection

import (
	"testing"
	"reflect"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	} {
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Ilya"},
			[]string{"Ilya"},
		},
		{
			"Struct with Two string field",
			struct {
				Name string
				Age int
			}{"Ilya", 29},
			[]string{"Ilya"},
		},
		{
			"Struct with not only string field",
			struct {
				Name string
				City string
			}{"Ilya", "Sev"},
			[]string{"Ilya", "Sev"},
		},
		{
			"Nested fields",
			Person{
				"Ilya", 
				Profile{29, "Sev"}},
			[]string{"Ilya", "Sev"},
		},
		{
			"Pointers to things",
			&Person{
				"Ilya", 
				Profile{29, "Sev"}},
			[]string{"Ilya", "Sev"},
		},
		{
			"Slices",
			[]Profile{
				{29, "Sev"},
				{30, "Simf"},
			},
			[]string{"Sev", "Simf"},
		},
		{
			"Arrays",
			[2]Profile{
				{29, "Sev"},
				{30, "Simf"},
			},
			[]string{"Sev", "Simf"},
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