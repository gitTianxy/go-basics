package pkg

type ElementA struct {
	name string
}

func (a ElementA) GetName() string  {
	return a.name
}

func (a *ElementA) SetName(name string)  {
	a.name = name
}