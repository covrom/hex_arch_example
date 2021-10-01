package dbimpl1

// import core entities
import "github.com/covrom/hex_arch_example/bad_example/improved/core/manager"

// statically checking that we are implementing the complete core logic
var _ manager.AbstractDB = &ConcreteDB{}

type ConcreteDB struct{}

// result is core entity
func (*ConcreteDB) Select() manager.Mail
