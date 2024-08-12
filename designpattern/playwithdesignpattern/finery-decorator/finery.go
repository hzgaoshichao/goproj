package main

type Finery struct {
	characterInterface CharacterInterface
}

func (f *Finery) decorate(characterInterface CharacterInterface) {
	f.characterInterface = characterInterface
}

func (f *Finery) show() {
	f.characterInterface.show()
}
