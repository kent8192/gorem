package body

import "fmt"

type GoremBody struct {
	isIntegrityChecked bool
}

func Mold() *GoremBody {
	return &GoremBody{
		isIntegrityChecked: true,
	}
}

func (g *GoremBody) CheckIntegirity() error {
	if g.isIntegrityChecked {
		return nil
	} else {
		err := fmt.Errorf("this GolemBody is illegally molded using a forbidden method")
		return err
	}
}
