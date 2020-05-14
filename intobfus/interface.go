/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package intobfus

// Obfuscator encodes/decodes serial numbers so they "look like" random number
//
// As it is designed to obfuscate id numbers in particular, 0 is not handled:
// Encode(0) = Decode(0) = 0
type Obfuscator interface {
	// Obfuscate encodes serial number to obfuscated number
	//
	// There's an example to show you real usage
	Obfuscate(serial uint64) (ret uint64)
	// Explain decodes obfuscated code to serial number
	//
	// It will return ErrCode if code > max
	Explain(code uint64) (ret uint64, err error)
}
