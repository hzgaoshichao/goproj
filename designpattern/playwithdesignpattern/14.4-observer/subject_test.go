package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Value struct {
	Name   string
	Gender *string
}

func TestDetach(t *testing.T) {
	v1 := Value{Name: "煎鱼", Gender: new(string)}
	v2 := Value{Name: "煎鱼", Gender: new(string)}
	//  if v1 == v2 {
	if reflect.DeepEqual(v1, v2) {
		fmt.Println("脑子进煎鱼了")
		return
	}

	fmt.Println("脑子没进煎鱼")
}
