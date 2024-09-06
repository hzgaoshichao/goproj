package unitednation

import "fmt"

type Country interface {
	Declare(message string)
}

type CountryCommon struct {
	unitedNations UnitedNations
}

func NewCountryCommon(un UnitedNations) *CountryCommon {
	return &CountryCommon{
		unitedNations: un,
	}
}

type USA struct {
	CountryCommon
}

func NewUSA(un UnitedNations) *USA {
	return &USA{
		CountryCommon: *NewCountryCommon(un),
	}
}

func (u *USA) Declare(message string) {
	u.unitedNations.Declare(message, u)
}

func (u *USA) getMessage(message string) {
	fmt.Printf("美国获得对方信息：%v", message)
}

type Irap struct {
	CountryCommon
}

func NewIrap(un UnitedNations) *Irap {
	return &Irap{
		CountryCommon: *NewCountryCommon(un),
	}
}

func (i *Irap) Declare(message string) {
	i.unitedNations.Declare(message, i)
}

func (u *Irap) getMessage(message string) {
	fmt.Printf("伊拉克获得对方信息：%v", message)
}
