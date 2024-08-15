package main

import "fmt"

type testPaperA struct {
	testPaper
}

func (t *testPaperA) answer1() string {
	fmt.Printf("答案是: A \n")
	return ""
}

func (t *testPaperA) answer2() string {
	fmt.Printf("答案是: B \n")
	return ""
}

func (t *testPaperA) answer3() string {
	fmt.Printf("答案是: C \n")
	return ""
}
