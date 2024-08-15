package main

func main() {
	studentA := &testPaperA{}
	templateMethod(studentA)

	studentB := &testPaperB{}
	templateMethod(studentB)
}

// 参考:
// https://learnku.com/articles/33704
// https://learnku.com/docs/the-way-to-go/104-tag-structure/3642
