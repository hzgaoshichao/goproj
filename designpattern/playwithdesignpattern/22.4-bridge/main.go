package main

import (
	hs "bridge/handset"
)

func main() {
	var ab hs.HandsetBrandInterface
	ab = &hs.HandsetBrandM{}
	ab.SetHandsetSoft(&hs.HandsetGame{})
	ab.Run()
	ab.SetHandsetSoft(&hs.HandsetAddressList{})
	ab.Run()

	ab2 := hs.HandsetBrandN{}
	ab2.SetHandsetSoft(&hs.HandsetGame{})
	ab2.Run()
	ab2.SetHandsetSoft(&hs.HandsetAddressList{})
	ab2.Run()
	ab2.SetHandsetSoft(&hs.HandsetMusicPlay{})
	ab2.Run()

	ab3 := hs.HandsetBrandS{}
	ab3.SetHandsetSoft(&hs.HandsetMusicPlay{})
	ab3.Run()
}
