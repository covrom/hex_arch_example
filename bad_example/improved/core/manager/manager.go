package manager

// interfaces must be defined inside the package that uses them
// this package does not import any packages from outer layers

type ManagerAll struct {
	// dependency inversion by interfaces
	Api MailAPI
	Db  AbstractDB
}

func (m *ManagerAll) Process() {
	// here we use our interfaces
	eml := m.Db.Select()
	m.Api.Send(eml)
}
