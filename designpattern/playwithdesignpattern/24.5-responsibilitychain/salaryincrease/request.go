package salaryincrease

type Request struct {
	requestType    string
	requestContent string
	number         int
}

func (r *Request) RequestType() string {
	return r.requestType
}

func (r *Request) RequestContent() string {
	return r.requestContent
}

func (r *Request) Number() int {
	return r.number
}

func (r *Request) SetRequestType(requestType string) {
	r.requestType = requestType
}

func (r *Request) SetRequestContent(requestContent string) {
	r.requestContent = requestContent
}

func (r *Request) SetNumber(number int) {
	r.number = number
}
