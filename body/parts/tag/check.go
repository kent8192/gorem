package tag

import (
	"fmt"
	"go/ast"
	"go/parser"
)

type Check struct {
	isLegitimatelyInitialized bool
	condition                 string
}

func (c *Check) InitTag(condition string) *Check {
	c.isLegitimatelyInitialized = true
	c.condition = condition
	if err := c.ValidateTag(); err != nil {
		panic(err)
	}
	return c
}

func (c *Check) RenderTag() string {
	if !c.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("check:%s", c.condition)
}

func (c *Check) IsValid() (bool, string) {
	if !c.isLegitimatelyInitialized {
		return false, "check is not initialized legitimately"
	}
	if c.condition == "" {
		return false, "check condition must not be empty"
	}
	// Goのパーサーを使用して条件式を検証
	expr, err := parser.ParseExpr(c.condition)
	if err != nil {
		return false, "check condition is not a valid expression"
	}
	// 条件式が比較演算子を含むかを確認
	valid := false
	ast.Inspect(expr, func(n ast.Node) bool {
		if _, ok := n.(*ast.BinaryExpr); ok {
			valid = true
			return false
		}
		return true
	})
	if !valid {
		return false, "check condition must contain a comparison operator"
	}
	return true, ""
}

func (c *Check) ValidateTag() error {
	if valid, errMessage := c.IsValid(); !valid {
		return fmt.Errorf("invalid check: %s (condition=%s)", errMessage, c.condition)
	}
	return nil
}

func (c *Check) GetIsLegitimatelyInitialized() bool {
	return c.isLegitimatelyInitialized
}
