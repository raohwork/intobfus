/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package intobfus

import "strconv"

// ErrRestore indicates provided parameters are invalid
type ErrRestore struct{}

func (e ErrRestore) Error() string {
	return "intobfus: invalid parameters to restore"
}

// ErrCode indicates provided code is invalid
type ErrCode uint64

func (e ErrCode) Error() string {
	return "intobfus: invalid code: " + strconv.FormatUint(uint64(e), 10)
}

// ErrKey indicates provided key value is invalid
type ErrKey uint64

func (e ErrKey) Error() string {
	return "intobfus: key has no modinverse: " + strconv.FormatUint(uint64(e), 10)
}
