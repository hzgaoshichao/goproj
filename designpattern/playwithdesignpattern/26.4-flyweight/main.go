package main

import (
	w "flyweight/website"
	"fmt"
)

func main() {
	f := w.NewWebsiteFactory()
	fx := f.GetWebsiteCategory("产品展示")
	fx.Use(w.NewUser("小菜"))
	fy := f.GetWebsiteCategory("产品展示")
	fy.Use(w.NewUser("大鸟"))
	fz := f.GetWebsiteCategory("产品展示")
	fz.Use(w.NewUser("娇娇"))
	fl := f.GetWebsiteCategory("博客")
	fl.Use(w.NewUser("老顽童"))
	fm := f.GetWebsiteCategory("博客")
	fm.Use(w.NewUser("桃谷六仙"))
	fn := f.GetWebsiteCategory("博客")
	fn.Use(w.NewUser("南海鳄神"))

	fmt.Printf("网站分类总数为:%v", f.GetWebsiteCount())
}
