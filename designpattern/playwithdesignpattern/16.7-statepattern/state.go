package main

import "fmt"

type StateInterface interface {
	writeProgram(work *Work)
}

type ForenoonState struct {
}

func (f *ForenoonState) writeProgram(work *Work) {
	if work.getHour() < 12 {
		fmt.Printf("当前时间: %v 点 上午工作,精神百倍 \n", work.getHour())
	} else {
		work.SetState(&NoonState{})
		work.writeProgram()
	}
}

type NoonState struct {
}

func (f *NoonState) writeProgram(work *Work) {
	if work.getHour() < 13 {
		fmt.Printf("当前时间: %v 点 饿了, 午饭;犯困, 午休 \n", work.getHour())
	} else {
		work.SetState(&AfternoonState{})
		work.writeProgram()
	}
}

type AfternoonState struct {
}

func (f *AfternoonState) writeProgram(work *Work) {
	if work.getHour() < 17 {
		fmt.Printf("当前时间: %v 点 下午状态还不错,继续努力 \n", work.getHour())
	} else {
		work.SetState(&EveningState{})
		work.writeProgram()
	}
}

type EveningState struct {
}

func (f *EveningState) writeProgram(work *Work) {
	if work.getWorkFinished() {
		work.SetState(&RestState{})
		work.writeProgram()
	} else {
		if work.getHour() < 21 {
			fmt.Printf("当前时间: %v 点 加班哦, 疲累至极 \n", work.getHour())
		} else {
			work.SetState(&SleepingState{})
			work.writeProgram()
		}
	}
}

type RestState struct {
}

func (f *RestState) writeProgram(work *Work) {
	fmt.Printf("当前时间: %v 点 下班回家 \n", work.getHour())
}

type SleepingState struct {
}

func (f *SleepingState) writeProgram(work *Work) {
	fmt.Printf("当前时间: %v 点 不行了,睡觉了 \n", work.getHour())
}
