// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build go1.19 && (linux || darwin || freebsd || openbsd) && !ts_omit_bird

package main

import (
	"github.com/metacubex/tailscale/chirp"
	"github.com/metacubex/tailscale/wgengine"
)

func init() {
	createBIRDClient = func(ctlSocket string) (wgengine.BIRDClient, error) {
		return chirp.New(ctlSocket)
	}
}
