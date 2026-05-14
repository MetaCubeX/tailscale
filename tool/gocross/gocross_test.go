//go:build ignore

// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"testing"

	"github.com/metacubex/tailscale/tstest/deptest"
)

func TestDeps(t *testing.T) {
	deptest.DepChecker{
		BadDeps: map[string]string{
			"github.com/metacubex/tailscale/tailcfg": "circular dependency via go generate",
			"github.com/metacubex/tailscale/version": "circular dependency via go generate",
		},
	}.Check(t)
}
