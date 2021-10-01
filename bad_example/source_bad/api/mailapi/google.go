package mailapi

import "github.com/covrom/hex_arch_example/bad_example/source_bad/model"

// Mistake #2: too many responsibilities in the package - MailAPI, GoogleAPI and YandexAPI

type GoogleAPI struct{}

// Mistake #3: using a model from a package without behavior 
func (*GoogleAPI) Send(m model.Mail)
