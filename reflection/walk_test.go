package main

import "testing"

func TestWalk(t *testing.T){
	expected := "Chris"
	var got []string

	x:= struct{
		Name string
	}{expected}

	walk(x, func(input string)){
		got = append(got,input)est
	}
}	