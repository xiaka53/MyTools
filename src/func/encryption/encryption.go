package encryption

//执行加密
func Run(types, str string) {
	switch types {
	case "AESECB":
		EncAesEcb(str)
	case "":

	}
}

//执行加密
func (rncryprion rncryprion) DecRun() {
	switch rncryprion.Type {
	case "AESECB":
		rncryprion.DecAesEcb()
	case "":

	}
}
