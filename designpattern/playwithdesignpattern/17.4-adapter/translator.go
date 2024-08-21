package main

type translator struct {
	f *foreignCenter
}

func NewTranslator(name string) translator {
	return translator{
		f: NewForeignCenter(name),
	}
}

func (t *translator) attach() {
	t.f.foreignCenterAttack()
}

func (t *translator) defense() {
	t.f.foreignCenterDefense()
}
