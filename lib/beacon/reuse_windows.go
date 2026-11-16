// Copyright (C) 2026 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

package beacon

import "syscall"

// SO_REUSEADDR on our own socket doesn't make it easier to take the port away
// from us: on Windows that depends on the other socket's options, and is only
// prevented by SO_EXCLUSIVEADDRUSE, which Syncthing has never set.
func reuseControl(_, _ string, c syscall.RawConn) error {
	var opErr error
	if err := c.Control(func(fd uintptr) {
		opErr = syscall.SetsockoptInt(syscall.Handle(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	}); err != nil {
		return err
	}
	return opErr
}
