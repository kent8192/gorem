package body

import "fmt"

type GoremBody struct {
	isLegitimatelyInitialized bool
}

func Mold() *GoremBody {
	return &GoremBody{
		isLegitimatelyInitialized: true,
	}
}

func (g *GoremBody) CheckIntegirity() error {
	if g.isLegitimatelyInitialized {
		return nil
	} else {
		err := fmt.Errorf("this GolemBody is illegitimately molded using a forbidden method")
		return err
	}
}
