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

	t.Run("Maps", func(t *testing.T) {
		aMap := map[string]string{
			"foo": "bar",
			"baz": "boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "bar")
		assertContains(t, got, "bar")
	})
}

func assertContains(t *testing.T, got []string, want string) {
	t.Helper()
	contains := false
	for _, x := range got {
		if x == want {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", want, got)
	}
}