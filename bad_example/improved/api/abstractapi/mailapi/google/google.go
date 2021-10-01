package google

// import only core entities

import "github.com/covrom/hex_arch_example/bad_example/improved/core/manager"

// statically checking that we are implementing the complete core logic
var _ manager.MailAPI = &GoogleAPI{}

type GoogleAPI struct{}

// argument is core entity
func (*GoogleAPI) Send(m manager.Mail)
