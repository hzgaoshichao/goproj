package main

type Boss struct {
	Subject // 这里使用嵌套结构体的方式, 这样就可以提取出公共的方法和字段到嵌套结构体中了
}

func NewBoss(name string) *Boss {
	return &Boss{
		Subject{name: name},
	}
}

func (b *Boss) getSubject() *Subject {
	return &b.Subject
}
