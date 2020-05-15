/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package intobfus

import (
	"crypto/rand"
	"math"
	"math/big"
)

func newUint(i uint64) (ret *big.Int) {
	ret = &big.Int{}
	ret.SetUint64(i)
	return
}

func genRandPrime(max, lower *big.Int) (ret *big.Int, err error) {
	b := math.Log2(float64(max.Uint64()))
	lb := math.Log2(float64(lower.Uint64()))
	bits := int(b)
	lbits := int(lb)
	if b > float64(bits) {
		bits++
	}
	if lb > float64(lbits) {
		lbits++
	}

	for {
		n, err := rand.Prime(rand.Reader, bits)
		if err != nil {
			return nil, err
		}
		if n.Cmp(max) < 0 && lower.Cmp(n) <= 0 {
			return n, nil
		}

		n, err = rand.Prime(rand.Reader, lbits)
		if err != nil {
			return nil, err
		}
		if n.Cmp(max) < 0 && lower.Cmp(n) <= 0 {
			return n, nil
		}
	}
}

// GenKey creates a random prime number as key
//
// In most case, you will not use this in production.
func GenKey(max uint64) (prime uint64, err error) {
	m := newUint(max)
	lower := (&big.Int{}).Mul(m, newUint(618))
	lower.Div(lower, newUint(1000))
	for prime == 0 {
		var x *big.Int
		x, err = genRandPrime(m, lower)
		if err == nil && x.ProbablyPrime(0) {
			prime = x.Uint64()
			break
		}
	}

	return
}

// ModMul computes a * b mod (m+1)
//
// It implements Schrage's algorithm, which is 2x faster than math/big.Int for
// uint64 and below.
func ModMul(a, b, m uint64) (ret uint64) {
	if m <= math.MaxUint32 {
		return (a * b) % (m + 1)
	}

	return modmul(a, b, m)
}

func modmul(a, b, m uint64) (ret uint64) {
	if a == 0 || b == 0 {
		return 0
	}
	if a == 1 {
		return b
	}
	if b == 1 {
		return a
	}

	q := m / a
	r := m - a*q + 1
	if r == a {
		r = 0
		q++
	}

	x := a * (b % q)
	var y uint64
	if r > q {
		y = modmul(r, b/q, m)
	} else {
		y = b / q
		y *= r
	}

	if x >= y {
		return x - y
	}

	return m - (y - x) + 1
}

type obfus struct {
	max    uint64
	encKey uint64
	decKey uint64
	a      uint64
	b      uint64
}

func (o *obfus) enc(n uint64) (ret uint64) {
	return ModMul(o.encKey, n, o.max)
}

func (o *obfus) dec(n uint64) (ret uint64) {
	return ModMul(o.decKey, n, o.max)
}

func (o *obfus) Obfuscate(serial uint64) (ret uint64) {
	switch serial {
	case o.encKey:
		return o.b
	case o.a:
		return o.encKey
	case 1:
		return o.a
	case 0:
		return 0
	}

	return o.enc(serial)
}

func (o *obfus) Explain(code uint64) (ret uint64, err error) {
	switch code {
	case o.b:
		return o.encKey, nil
	case o.encKey:
		return o.a, nil
	case o.a:
		return 1, nil
	case 0:
		return 0, nil
	}

	if code > o.max {
		err = ErrCode(code)
		return
	}
	ret = o.dec(code)
	return
}

// Restore creates an Obfuscator instance using previously generated key
//
// You can call GenKey() to obtain the key. Providing invalid parameters results
// ErrRestore in err.
func Restore(max, key uint64) (ret Obfuscator, err error) {
	if key >= max {
		err = ErrKey(key)
		return
	}

	d := &obfus{
		max:    max,
		encKey: key,
	}

	x := (&big.Int{}).ModInverse(
		newUint(d.encKey),
		(&big.Int{}).Add(
			newUint(d.max),
			big.NewInt(1),
		),
	)
	if x == nil {
		err = ErrKey(key)
		return
	}

	d.decKey = x.Uint64()
	d.a = d.enc(key)
	d.b = d.enc(d.a)
	ret = d
	return
}

// MustRestore wraps Restore() with panic
func MustRestore(max, key uint64) (ret Obfuscator) {
	ret, err := Restore(max, key)
	if err != nil {
		panic(err)
	}

	return
}
