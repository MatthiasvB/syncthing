// Copyright (C) 2026 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build !(linux || solaris || aix || darwin || dragonfly || freebsd || netbsd || openbsd || windows)

package beacon

import "syscall"

// No known way to share the port here, so the first instance keeps it, as
// before.
func reuseControl(_, _ string, _ syscall.RawConn) error { return nil }
