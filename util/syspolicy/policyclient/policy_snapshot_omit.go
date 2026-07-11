// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !ts_enable_syspolicy

package policyclient

// policySnapshot is a stub when syspolicy is disabled.
type policySnapshot struct{}
