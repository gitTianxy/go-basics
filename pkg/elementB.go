package pkg

type ElementB struct {
	name string
}

func (a ElementB) GetName() string  {
	return a.name
}

func (a *ElementB) SetName(name string)  {
	a.name = name
}