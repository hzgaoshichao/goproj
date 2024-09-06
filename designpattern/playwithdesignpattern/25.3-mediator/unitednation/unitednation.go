package unitednation

type UnitedNations interface {
	Declare(message string, country Country)
}

type UnitedNationsSecurityCouncil struct {
	countryUSA  *USA
	countryIraq *Irap
}

func (unsc *UnitedNationsSecurityCouncil) SetUSA(u *USA) {
	unsc.countryUSA = u
}

func (unsc *UnitedNationsSecurityCouncil) SetIrap(value *Irap) {
	unsc.countryIraq = value
}

func (unsc *UnitedNationsSecurityCouncil) Declare(message string, country Country) {
	if unsc.countryUSA == country {
		unsc.countryIraq.getMessage(message)

	} else if unsc.countryIraq == country {
		unsc.countryUSA.getMessage(message)

	}
}
