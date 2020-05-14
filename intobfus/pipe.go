/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package intobfus

type pipe struct {
	parts []Obfuscator
}

func (p *pipe) Obfuscate(serial uint64) (ret uint64) {
	ret = serial
	for _, o := range p.parts {
		ret = o.Obfuscate(ret)
	}

	return
}

func (p *pipe) Explain(code uint64) (ret uint64, err error) {
	ret = code
	for x := len(p.parts) - 1; x >= 0; x-- {
		ret, err = p.parts[x].Explain(ret)
		if err != nil {
			return
		}
	}

	return
}

// Pipe returns an Obfuscator that interates through provided parts
//
// See example for how and what it does
func Pipe(parts ...Obfuscator) (ret Obfuscator) {
	return &pipe{parts: parts}
}

// PipeByKey is same as Pipe(), just accepts keys instead of Obfuscator instances
func PipeByKey(max uint64, keys ...uint64) (ret Obfuscator, err error) {
	parts := make([]Obfuscator, len(keys))

	for idx, k := range keys {
		var o Obfuscator
		o, err = Restore(max, k)
		if err != nil {
			return
		}
		parts[idx] = o
	}

	ret = &pipe{parts: parts}
	return
}

// MustPipeByKey is same as PipeByKey(), just wraps error with panic
func MustPipeByKey(max uint64, keys ...uint64) (ret Obfuscator) {
	ret, err := PipeByKey(max, keys...)
	if err != nil {
		panic(err)
	}

	return
}
