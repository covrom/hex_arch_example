package dbimpl1

import "github.com/covrom/hex_arch_example/bad_example/source_bad/model"

type ConcreteDB struct{}

// Mistake #3: returning a model from a package without behavior 
func (*ConcreteDB) Select() model.Mail
