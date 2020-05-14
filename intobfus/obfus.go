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

type obfus struct {
	base   *big.Int // max+1
	encKey *big.Int
	decKey *big.Int
	key    uint64
	a      uint64 // enc(key), but used as enc(1), enc(key) is changed to b
	b      uint64 // enc(a), but used as enc(key), and enc(a) is changed to 1
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

func (o *obfus) enc(n uint64) (ret *big.Int) {
	ret = newUint(n)
	ret.Mul(ret, o.encKey)
	ret.Mod(ret, o.base)
	return ret
}

func (o *obfus) dec(n uint64) (ret *big.Int) {
	ret = newUint(n)
	ret.Mul(ret, o.decKey)
	ret.Mod(ret, o.base)
	return ret
}

func (o *obfus) Obfuscate(serial uint64) (ret uint64) {
	switch serial {
	case o.key:
		return o.b
	case o.a:
		return o.key
	case 1:
		return o.a
	case 0:
		return 0
	}

	return o.enc(serial).Uint64()
}

func (o *obfus) Explain(code uint64) (ret uint64, err error) {
	switch code {
	case o.b:
		return o.key, nil
	case o.key:
		return o.a, nil
	case o.a:
		return 1, nil
	case 0:
		return 0, nil
	}

	x := newUint(code)
	if x.Cmp(o.base) >= 0 {
		err = ErrCode(code)
		return
	}
	ret = o.dec(code).Uint64()
	return
}

// Restore creates an Obfuscator instance using previously generated key
//
// You can call GenKey() to obtain the key. Providing invalid parameters results
// ErrRestore in err.
func Restore(max, key uint64) (ret Obfuscator, err error) {
	base := (&big.Int{}).Add(newUint(max), big.NewInt(1))
	d := &obfus{
		base:   base,
		encKey: newUint(key),
		key:    key,
	}

	x := (&big.Int{}).ModInverse(d.encKey, d.base)
	if x == nil {
		err = ErrKey(key)
		return
	}

	d.decKey = x
	d.a = d.enc(key).Uint64()
	d.b = d.enc(d.a).Uint64()
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
