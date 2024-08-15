package main

type workExperience struct {
	timeArea string
	company  string
}

func (w *workExperience) setTimeArea(value string) {
	w.timeArea = value
}
func (w *workExperience) setCompany(value string) {
	w.company = value
}

func (w *workExperience) clone() cloneableInterface {
	newWorkExperience := *w
	return &newWorkExperience
}
