package main

import s "responsibilitychain/salaryincrease"

func main() {
	manager := s.NewCommonManager("金利")
	director := s.NewDirector("宗剑")
	generalManager := s.NewGeneralManager("钟精励")
	manager.SetSuperior(director)
	director.SetSuperior(generalManager)

	request := s.Request{}
	request.SetRequestType("请假")
	request.SetRequestContent("小菜请假")
	request.SetNumber(1)
	manager.RequestApplication(request)

	request2 := s.Request{}
	request2.SetRequestType("请假")
	request2.SetRequestContent("小菜请假")
	request2.SetNumber(4)
	manager.RequestApplication(request2)

	request3 := s.Request{}
	request3.SetRequestType("加薪")
	request3.SetRequestContent("小菜请求加薪")
	request3.SetNumber(5000)
	manager.RequestApplication(request3)

	request4 := s.Request{}
	request4.SetRequestType("加薪")
	request4.SetRequestContent("小菜请求加薪")
	request4.SetNumber(10000)
	manager.RequestApplication(request4)

}
