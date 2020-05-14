/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

// Package intobfus provides a simple obfuscator for serial numbers.
//
// It uses multiplicative modular to make serial numbers "look like" randomly
// generated. Take a look to examples about how to use it.
//
// It is suggested to use Pipe() with 2 or more different keys, which can make it
// "more like" random numbers.
package intobfus
