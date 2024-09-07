package main

import w "flyweight/website"

func main() {
	f := w.NewWebsiteFactory()
	fx := f.GetWebsiteCategory("产品展示")
	fx.Use()
	fy := f.GetWebsiteCategory("产品展示")
	fy.Use()
	fz := f.GetWebsiteCategory("产品展示")
	fz.Use()
	fl := f.GetWebsiteCategory("博客")
	fl.Use()
	fm := f.GetWebsiteCategory("博客")
	fm.Use()
	fn := f.GetWebsiteCategory("博客")
	fn.Use()
}
