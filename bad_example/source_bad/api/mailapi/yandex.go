package mailapi

import "github.com/covrom/hex_arch_example/bad_example/source_bad/model"

// Mistake #2: too many responsibilities in the package - MailAPI, GoogleAPI and YandexAPI

type YandexAPI struct{}

// Mistake #3: using a model from a package without behavior 
func (*YandexAPI) Send(m model.Mail)
