package utils

import (
	"fmt"
	"math/big"
)

func IntToHex(n *big.Int) string {
	return fmt.Sprintf("0x%x", n)
}
