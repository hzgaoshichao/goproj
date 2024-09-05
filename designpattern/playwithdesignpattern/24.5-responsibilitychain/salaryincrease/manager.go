package salaryincrease

import "fmt"

type ManagerInterface interface {
	SetSuperior(superior ManagerInterface)
	RequestApplication(request Request)
}

type Manager struct {
	name     string
	superior ManagerInterface
}

func NewManager(name string) *Manager {
	return &Manager{
		name: name,
	}
}

func (m *Manager) SetSuperior(superior ManagerInterface) {
	m.superior = superior
}

func (m *Manager) RequestApplication(request Request) {
	fmt.Println("Not Implemented")
}

type CommonManager struct {
	Manager
}

func NewCommonManager(name string) *CommonManager {
	return &CommonManager{
		Manager: *NewManager(name),
	}
}

func (cm *CommonManager) RequestApplication(request Request) {
	if request.RequestType() == "请假" && request.Number() <= 2 {
		fmt.Printf("%v : %v 数量: %v 天,被批准 \n", cm.name, request.RequestContent(), request.Number())
	} else {
		if cm.Manager.superior != nil {
			cm.Manager.superior.RequestApplication(request)
		}
	}
}

type Director struct {
	Manager
}

func NewDirector(name string) *Director {
	return &Director{
		Manager: *NewManager(name),
	}
}

func (d *Director) RequestApplication(request Request) {
	if request.RequestType() == "请假" && request.Number() <= 5 {
		fmt.Printf("%v : %v 数量: %v 天,被批准 \n", d.name, request.RequestContent(), request.Number())
	} else {
		if d.Manager.superior != nil {
			d.Manager.superior.RequestApplication(request)
		}
	}
}

type GeneralManager struct {
	Manager
}

func NewGeneralManager(name string) *GeneralManager {
	return &GeneralManager{
		Manager: *NewManager(name),
	}
}

func (gm *GeneralManager) RequestApplication(request Request) {
	if request.RequestType() == "请假" && request.Number() <= 5 {
		fmt.Printf("%v : %v 数量: %v 天,被批准 \n", gm.name, request.RequestContent(), request.Number())
	} else if request.RequestType() == "加薪" && request.Number() <= 5000 {
		fmt.Printf("%v : %v 数量: %v 元,被批准 \n", gm.name, request.RequestContent(), request.Number())
	} else if request.RequestType() == "加薪" && request.Number() > 5000 {
		fmt.Printf("%v : %v 数量: %v 元, 再说吧 \n", gm.name, request.RequestContent(), request.Number())
	}
}
