/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/raohwork/intobfus/intobfus"
)

func bybit(bit uint) (max uint64) {
	if bit > 64 || bit < 8 {
		log.Fatal("intobfus supports only 8~64 bits serial numbers")
	}

	if bit == 64 {
		return math.MaxUint64
	}

	max = uint64(1)<<bit - 1
	return
}

func genkey(max uint64) (key uint64) {
	key, err := intobfus.GenKey(max)
	if err != nil {
		log.Fatal("cannot generate key: ", err)
	}

	fmt.Printf("intobfus.Restore(%d, %d)\n", max, key)

	return
}

func genkeys(max uint64, pass uint) (keys []uint64) {
	if pass < 1 {
		log.Fatal("at least 1 pass")
	}
	if pass == 1 {
		keys = []uint64{genkey(max)}
		return
	}

	keys = make([]uint64, pass)
	for p := uint(0); p < pass; p++ {
		key, err := intobfus.GenKey(max)
		if err != nil {
			log.Fatal("cannot generate key: ", err)
		}

		keys[p] = key
	}

	fmt.Println("intofus.MustPipeByKeys(")
	fmt.Printf("\t%d,\n", max)
	for _, k := range keys {
		fmt.Printf("\t%d,\n", k)
	}
	fmt.Println(")")

	return
}

func main() {
	var (
		max  uint64
		pass uint
		bits uint
		show bool
	)
	flag.Uint64Var(&max, "max", 0, "Allows 100 ~ 2^64-1. If your serial number is bwtween 1~255, you should use 255. It overrides bits setting.")
	flag.UintVar(&pass, "pass", 3, "Generates a Pipe() with n nifferent keys")
	flag.UintVar(&bits, "bits", 0, "Allows only 8~64. Set max with bits. -bits=8 euqals to -max=255")
	flag.BoolVar(&show, "show", false, "show some values to stderr to see if it is what you want")
	flag.Parse()

	if max == 0 && bits == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	m := max
	if bits > 0 {
		m = bybit(bits)
	}
	if max >= 100 {
		m = max
	}

	keys := genkeys(m, pass)
	if show {
		showDemo(m, keys)
	}
}

func showDemo(max uint64, keys []uint64) {
	p := intobfus.MustPipeByKey(max, keys...)

	if max >= 500 {
		max = 500
	}

	for i := uint64(1); i < max; i++ {
		log.Printf(
			"Obfuscate(%3d) = %19d",
			i,
			p.Obfuscate(i),
		)
	}
}
