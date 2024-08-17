package main

func main() {
	person := NewPersonBuilder().
		WithName("LeoGao").
		WithAge(20).
		WithGender("Male").
		WithAddress("中国南京").
		Build()

	person.Speak("精忠报国")
	person.Sleep()
}

// 参考文档：
// 这个列子更加的容易理解： https://learnku.com/docs/go-patterns/1.0.0/jian-zao-zhe-mo-shi-builder-pattern/14747
// 这个列子更加的容易理解, 更加的优雅： https://github.com/ruanrunxue/Practice-Design-Pattern--Go-Implementation/blob/main/docs/go_practice_design_pattern__builder.md
