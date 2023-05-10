package command

type Hello struct {
	hello string
}

func (h Hello) HelloHandle() string {
	return h.hello
}
