package website

import "fmt"

type Website interface {
	Use(user *User)
}

type ConcreteWebsite struct {
	name string
}

func NewConcreteWebsite(name string) *ConcreteWebsite {
	return &ConcreteWebsite{name: name}
}

func (c *ConcreteWebsite) Use(user *User) {
	fmt.Printf("网站分类: %v , 用户: %v\n", c.name, user.Name)
}

type WebsiteFactory struct {
	flyweights map[string]*ConcreteWebsite
}

func NewWebsiteFactory() *WebsiteFactory {
	f := &WebsiteFactory{}
	f.flyweights = make(map[string]*ConcreteWebsite)
	return f
}

func (wf *WebsiteFactory) GetWebsiteCategory(key string) Website {
	if wf.flyweights[key] == nil {
		wf.flyweights[key] = NewConcreteWebsite(key)
	}
	return wf.flyweights[key]
}

func (wf *WebsiteFactory) GetWebsiteCount() int {
	return len(wf.flyweights)
}
