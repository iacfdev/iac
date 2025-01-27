package funcs

import (
	"database/sql"
	"log"
)

type TransCodeInterface interface {
	Execute(string, map[string]interface{}, *sql.Tx) (map[string]interface{}, error)
}

type SubTranCodeFuncs struct {
	TranFlowstr TransCodeInterface
}

func New(tci TransCodeInterface) *SubTranCodeFuncs {
	return &SubTranCodeFuncs{
		TranFlowstr: tci,
	}

}

func (cf *SubTranCodeFuncs) Execute(f *Funcs) {
	tcode := f.Fobj.Content
	_, _, mappedinputs := f.SetInputs()

	outputs, err := cf.TranFlowstr.Execute(tcode, mappedinputs, f.DBTx)
	if err != nil {
		log.Println(err)
		return
	}
	f.SetOutputs(outputs)
}

func (cf *SubTranCodeFuncs) Validate(f *Funcs) (bool, error) {

	return true, nil
}
