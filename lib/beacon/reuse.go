// Copyright (C) 2026 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

package beacon

// The IPv4 broadcast listener binds 0.0.0.0:21027. By default only one
// process per host can do that, so with several Syncthing instances on one
// host (one per user, say) only the first gets IPv4 local discovery.
//
// reuseControl lets every instance bind the port, with the same socket options
// Go's net package already sets on multicast listeners
// (setDefaultMulticastSockopts). The IPv6 multicast listener has therefore
// always been shareable, and this gives the IPv4 listener the same behavior on
// every platform. The kernel delivers each broadcast to every socket bound to
// the port. Local discovery receives nothing but broadcasts, which every host
// on the LAN can read anyway, so sharing the port exposes nothing new.
