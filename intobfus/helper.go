/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package intobfus

// I wraps Obfuscator for int
//
// You have to ensure 0 <= max <= IntMax yourself
type I struct {
	Obfuscator
}

func (i I) Obfuscate(serial int) (ret int) {
	return int(i.Obfuscator.Obfuscate(uint64(serial)))
}

func (i I) Explain(code int) (ret int, err error) {
	x, err := i.Obfuscator.Explain(uint64(code))
	ret = int(x)

	return
}

// I8 wraps Obfuscator for int8
//
// You have to ensure 0 <= max <= Int8Max yourself
type I8 struct {
	Obfuscator
}

func (i I8) Obfuscate(serial int8) (ret int8) {
	return int8(i.Obfuscator.Obfuscate(uint64(serial)))
}

func (i I8) Explain(code int8) (ret int8, err error) {
	x, err := i.Obfuscator.Explain(uint64(code))
	ret = int8(x)

	return
}

// I16 wraps Obfuscator for int16
//
// You have to ensure 0 <= max <= Int16Max yourself
type I16 struct {
	Obfuscator
}

func (i I16) Obfuscate(serial int16) (ret int16) {
	return int16(i.Obfuscator.Obfuscate(uint64(serial)))
}

func (i I16) Explain(code int16) (ret int16, err error) {
	x, err := i.Obfuscator.Explain(uint64(code))
	ret = int16(x)

	return
}

// I32 wraps Obfuscator for int32
//
// You have to ensure 0 <= max <= Int32Max yourself
type I32 struct {
	Obfuscator
}

func (i I32) Obfuscate(serial int32) (ret int32) {
	return int32(i.Obfuscator.Obfuscate(uint64(serial)))
}

func (i I32) Explain(code int32) (ret int32, err error) {
	x, err := i.Obfuscator.Explain(uint64(code))
	ret = int32(x)

	return
}

// I64 wraps Obfuscator for int64
//
// You have to ensure 0 <= max <= Int64Max yourself
type I64 struct {
	Obfuscator
}

func (i I64) Obfuscate(serial int64) (ret int64) {
	return int64(i.Obfuscator.Obfuscate(uint64(serial)))
}

func (i I64) Explain(code int64) (ret int64, err error) {
	x, err := i.Obfuscator.Explain(uint64(code))
	ret = int64(x)

	return
}

// U wraps Obfuscator for uint
//
// You have to ensure max <= UIntMax yourself
type U struct {
	Obfuscator
}

func (i U) Obfuscate(serial uint) (ret uint) {
	return uint(i.Obfuscator.Obfuscate(uint64(serial)))
}

func (i U) Explain(code uint) (ret uint, err error) {
	x, err := i.Obfuscator.Explain(uint64(code))
	ret = uint(x)

	return
}

// U8 wraps Obfuscator for uint8
//
// You have to ensure max <= UInt8Max yourself
type U8 struct {
	Obfuscator
}

func (i U8) Obfuscate(serial uint8) (ret uint8) {
	return uint8(i.Obfuscator.Obfuscate(uint64(serial)))
}

func (i U8) Explain(code uint8) (ret uint8, err error) {
	x, err := i.Obfuscator.Explain(uint64(code))
	ret = uint8(x)

	return
}

// U16 wraps Obfuscator for uint16
//
// You have to ensure max <= UInt16Max yourself
type U16 struct {
	Obfuscator
}

func (i U16) Obfuscate(serial uint16) (ret uint16) {
	return uint16(i.Obfuscator.Obfuscate(uint64(serial)))
}

func (i U16) Explain(code uint16) (ret uint16, err error) {
	x, err := i.Obfuscator.Explain(uint64(code))
	ret = uint16(x)

	return
}

// U32 wraps Obfuscator for uint32
//
// You have to ensure max <= UInt32Max yourself
type U32 struct {
	Obfuscator
}

func (i U32) Obfuscate(serial uint32) (ret uint32) {
	return uint32(i.Obfuscator.Obfuscate(uint64(serial)))
}

func (i U32) Explain(code uint32) (ret uint32, err error) {
	x, err := i.Obfuscator.Explain(uint64(code))
	ret = uint32(x)

	return
}
