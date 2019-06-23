package parsec

// TODO
type Bank interface {
}

type BankKeeper struct {
	Bank
}

func NewBankKeeper() BankKeeper {
	return BankKeeper{}
}
