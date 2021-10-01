package abstractdb

import (
	"github.com/covrom/hex_arch_example/bad_example/source_bad/model"
)

type AbstractDB interface {
	Select() model.Mail
}
