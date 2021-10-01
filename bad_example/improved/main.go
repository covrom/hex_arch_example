package bad_example

import (
	"github.com/covrom/hex_arch_example/bad_example/improved/api/abstractapi/mailapi/yandex"
	"github.com/covrom/hex_arch_example/bad_example/improved/core/manager"
	"github.com/covrom/hex_arch_example/bad_example/improved/db/abstractdb/dbimpl1"
)

// improved example

func Serve() {
	var mapi manager.MailAPI = &yandex.YandexAPI{} // easy switch to &google.GoogleAPI{}
	var db manager.AbstractDB = &dbimpl1.ConcreteDB{}

	// dependency injection
	logic := manager.ManagerAll{
		Db:  db,
		Api: mapi,
	}

	logic.Process()
}
