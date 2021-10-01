package bad_example

import (
	"github.com/covrom/hex_arch_example/bad_example/source_bad/api/mailapi"
	"github.com/covrom/hex_arch_example/bad_example/source_bad/core/manager"
	"github.com/covrom/hex_arch_example/bad_example/source_bad/db/abstractdb"
	"github.com/covrom/hex_arch_example/bad_example/source_bad/db/abstractdb/dbimpl1"
)

// bad example

func Serve() {
	var mapi mailapi.MailAPI = &mailapi.YandexAPI{} //&mailapi.GoogleAPI{}
	var db abstractdb.AbstractDB = &dbimpl1.ConcreteDB{}

	logic := manager.ManagerAll{
		Db:  db,
		Api: mapi,
	}

	logic.Process()
}
