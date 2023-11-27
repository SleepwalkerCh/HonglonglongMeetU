package dao

const BlindMatchTableName = "blind_match"

type BlindMatchModelInterface interface {
}

type BlindMatchModelInterfaceImp struct{}

var IBlindMatchInterface = &BlindMatchModelInterfaceImp{}
