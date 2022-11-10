package utils

const (
	DATEFORMAT      = "2006-01-02"
	ACTIVE          = "active"
	INACTIVE        = "inactive"
	WITHDRAW        = "withdraw"
	DEPOSIT         = "deposit"
	CHECK           = "check"
	TRANSFER        = "transfer"
	DIGITSONLYREGEX = `(?m)^[0-9]+$`
	DATEFORMATREGEX = `(?m)^(\d{4})(-(0[1-9]|1[0-2])(-([12]\d|0[1-9]|3[01]))([T\s]((([01]\d|2[0-3])((:)[0-5]\d))([\:]\d+)?)?(:[0-5]\d([\.]\d+)?)?([zZ]|([\+-])([01]\d|2[0-3]):?([0-5]\d)?)?)?)$`
)
