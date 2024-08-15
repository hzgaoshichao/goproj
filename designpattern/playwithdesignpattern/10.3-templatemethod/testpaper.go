package main

import "fmt"

type testPaperInterface interface {
	testQuestion1()
	answer1() string
	testQuestion2()
	answer2() string
	testQuestion3()
	answer3() string
}

type testPaper struct {
}

func (t *testPaper) testQuestion1() {
	fmt.Printf("Question 1:xxxxxxxxxxxxx 答案是 [] : a.x b.x c.x d.x \n")
	t.answer1()
}

func (t *testPaper) answer1() string {
	fmt.Printf("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx============================================== \n")
	return ""
}

func (t *testPaper) testQuestion2() {
	fmt.Printf("Question 2:xxxxxxxxxxxxx 答案是 [] : a.x b.x c.x d.x \n")

}

func (t *testPaper) answer2() string {
	return ""
}

func (t *testPaper) testQuestion3() {
	fmt.Printf("Question 3:xxxxxxxxxxxxx 答案是 [] : a.x b.x c.x d.x \n")

}

func (t *testPaper) answer3() string {
	return ""
}

func templateMethod(t testPaperInterface) {
	t.testQuestion1()
	t.answer1()
	t.testQuestion2()
	t.answer2()
	t.testQuestion3()
	t.answer3()
}
