package main

func main() {
	resume1 := NewResume("大鸟")
	resume1.setPersonalInfo("男", "29")
	resume1.setWorkExperience("1998-2000", "xx公司")

	resume2 := resume1.clone().(*resume) // 这里使用了GoLang类型断言, 详细可参考: https://juejin.cn/post/7136359125738815524
	resume2.setWorkExperience("2000-2003", "yy集团")

	resume3 := resume1.clone().(*resume)
	resume3.setWorkExperience("2003-2006", "zz公司")

	// 从下面的这个赋值和显示结果来看, go 语言的结构体赋值就相当于是全部的值拷贝, 也就不会出现像Java那样的引用类型的深度拷贝的问题
	// 原型模式（Prototype Pattern）是一种创建型设计模式，其核心思想是通过复制现有对象来创建新的对象，而不是使用传统的构造函数创建。这种方式可以避免大量的构造函数调用，从而提高代码的性能和可读性。

	resume4 := resume1
	resume4.setWorkExperience("2007-2008", "躺平")

	resume1.display()
	resume2.display()
	resume3.display()
	resume4.display()

}
