package main

import "fmt"

type testPaperB struct {
	testPaper
}

func (t *testPaperB) answer1() string {
	fmt.Printf("答案是: C \n")
	return ""
}

func (t *testPaperB) answer2() string {
	fmt.Printf("答案是: D \n")
	return ""
}

func (t *testPaperB) answer3() string {
	fmt.Printf("答案是: A \n")
	return ""
}
