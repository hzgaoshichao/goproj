package main

func main() {
	emergencyProjects := NewWork()
	emergencyProjects.setHour(9)
	emergencyProjects.writeProgram()

	emergencyProjects.setHour(10)
	emergencyProjects.writeProgram()

	emergencyProjects.setHour(12)
	emergencyProjects.writeProgram()

	emergencyProjects.setHour(13)
	emergencyProjects.writeProgram()

	emergencyProjects.setHour(14)
	emergencyProjects.writeProgram()

	emergencyProjects.setHour(17)
	emergencyProjects.setWorkFinished(false)
	emergencyProjects.writeProgram()
	emergencyProjects.setHour(19)
	emergencyProjects.writeProgram()
	emergencyProjects.setHour(22)
	emergencyProjects.writeProgram()
}

// 参考文档: https://learnku.com/docs/go-patterns/1.0.0/zhuang-tai-mo-shi-state-pattern/14755
