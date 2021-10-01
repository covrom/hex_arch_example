package manager

// This is the core domain logic that contains all the core entities
// on the inside, but not on the outside.
// All other packages must import and use this package, but not vice versa

// this package does not import any packages from outer layers

type Mail struct{}

type MailAPI interface {
	Send(m Mail)
}

type AbstractDB interface {
	Select() Mail
}
