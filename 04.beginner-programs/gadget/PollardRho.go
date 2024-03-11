package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// 将要分解的数字赋值给变量 n
	n := new(big.Int)
	//     n.SetString("50296446669475900103173373917676050797315238472655018625792237958753851302131", 10)
	n.SetString("50296446669475902694584875652761273678224636096879696755629890223059048071168", 10)

	// 分解质因数
	factors := factor(n)

	// 输出质因数
	for _, f := range factors {
		fmt.Println(f)
	}
}

// 分解质因数
func factor(n *big.Int) []*big.Int {
	factors := []*big.Int{}

	// 尝试用 2 整除
	two := big.NewInt(2)
	for n.Rem(n, two).Sign() == 0 {
		factors = append(factors, two)
		n.Div(n, two)
	}

	// 使用 Pollard's rho 算法分解质因数
	for n.Cmp(big.NewInt(1)) > 0 {
		d := pollardRho(n)
		factors = append(factors, d)
		n.Div(n, d)
	}

	return factors
}

// Pollard's rho 算法
func pollardRho(n *big.Int) *big.Int {
	a := random(n) // Directly assign the result of random function to a
	x := big.NewInt(2)
	y := big.NewInt(2)

	for {
		x.Mul(x, x).Add(x, a).Mod(x, n)
		y.Mul(y, y).Mul(y, y).Mod(y, n)

		d := gcd(new(big.Int).Sub(x, y), n) // Make sure to pass a new big.Int to gcd as the subtraction modifies its first operand
		if d.Cmp(big.NewInt(1)) > 0 {
			return d
		}
	}
}

// 随机数
func random(n *big.Int) *big.Int {
	r, err := rand.Int(rand.Reader, n)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate random number: %v", err))
	}
	return r
}

// 最大公约数
func gcd(a, b *big.Int) *big.Int {
	for b.Sign() > 0 {
		t := a.Mod(a, b)
		a.Set(b)
		b.Set(t)
	}
	return a
}
