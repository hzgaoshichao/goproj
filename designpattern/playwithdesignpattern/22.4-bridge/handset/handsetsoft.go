package handset

import "fmt"

type HandsetSoftInterface interface {
	Run()
}

type HandsetGame struct {
}

func (h *HandsetGame) Run() {
	fmt.Printf("手机游戏 \n")
}

type HandsetAddressList struct {
}

func (h *HandsetAddressList) Run() {
	fmt.Printf("手机通讯录 \n")
}

type HandsetMusicPlay struct {
}

func (h *HandsetMusicPlay) Run() {
	fmt.Printf("音乐播放器 \n")
}
