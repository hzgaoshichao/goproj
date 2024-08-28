package handset

import "fmt"

type HandsetBrandInterface interface {
	SetHandsetSoft(soft HandsetSoftInterface)
	Run()
}

// 手机品牌 M
type HandsetBrandM struct {
	soft HandsetSoftInterface
}

func (hm *HandsetBrandM) SetHandsetSoft(soft HandsetSoftInterface) {
	hm.soft = soft
}

func (hm *HandsetBrandM) Run() {
	fmt.Printf("品牌 M \n")
	hm.soft.Run()
}

// 手机品牌 N
type HandsetBrandN struct {
	soft HandsetSoftInterface
}

func (hm *HandsetBrandN) SetHandsetSoft(soft HandsetSoftInterface) {
	hm.soft = soft
}

func (hm *HandsetBrandN) Run() {
	fmt.Printf("品牌 N \n")
	hm.soft.Run()
}

// 手机品牌 S
type HandsetBrandS struct {
	soft HandsetSoftInterface
}

func (hm *HandsetBrandS) SetHandsetSoft(soft HandsetSoftInterface) {
	hm.soft = soft
}

func (hm *HandsetBrandS) Run() {
	fmt.Printf("品牌 S \n")
	hm.soft.Run()
}
