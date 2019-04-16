package meander

// facade represents obj that provide a public view of themselves
type Facade interface {
	Public() interface{}
}

func Public(o interface{}) interface{} {
	if p, ok := o.(Facade); ok {
		return p.Public()
	}
	return o
}
