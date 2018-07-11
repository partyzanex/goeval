package goeval

import (
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {
	s := NewScope()
	s.Set("print", fmt.Println)
	t.Log(s.Eval(`count := 0`))
	t.Log(s.Eval(`for i:=0;i<100;i=i+1 { 
			count=count+i
		}`))
	t.Log(s.Eval(`print(count)`))
}

func TestScope_Eval(t *testing.T) {
	type data struct {
		Number      int
		Str         string
		Real        float64
		IntPointer  *int
		StrPointer  *string
		RealPointer *string
	}

	i := 100
	s := "pointer"

	ev := NewScope()
	ev.Set("data", &data{
		Number:      1000,
		Str:         "test",
		Real:        99.98,
		IntPointer:  &i,
		StrPointer:  &s,
		RealPointer: nil,
	})

	expr := `data.Number > 0 && data.Number <= 1000 && data.Str != "" && data.Str == "test"`

	res, err := ev.Eval(expr)
	if err != nil {
		t.Fatal(err)
	}

	if !res.(bool) {
		t.Fatal("expected true, assert false")
	}
}
