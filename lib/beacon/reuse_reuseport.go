// Copyright (C) 2026 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build aix || darwin || dragonfly || freebsd || netbsd || openbsd

package beacon

import "syscall"

// SO_REUSEADDR alone only permits bindings to different addresses here.
// Identical bindings, such as several 0.0.0.0:21027, need SO_REUSEPORT on all
// of them.
func reuseControl(_, _ string, c syscall.RawConn) error {
	var opErr error
	if err := c.Control(func(fd uintptr) {
		opErr = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
		if opErr == nil {
			opErr = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEPORT, 1)
		}
	}); err != nil {
		return err
	}
	return opErr
}
