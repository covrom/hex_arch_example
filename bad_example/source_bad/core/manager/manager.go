package manager

// Mistake # 1: This is the core logic, but it is importing some packages from the outer layers.
import (
	"github.com/covrom/hex_arch_example/bad_example/source_bad/api/mailapi"
	"github.com/covrom/hex_arch_example/bad_example/source_bad/db/abstractdb"
)

type ManagerAll struct {
	Api mailapi.MailAPI
	Db  abstractdb.AbstractDB
}

func (m *ManagerAll) Process() {
	eml := m.Db.Select()
	m.Api.Send(eml)
}
