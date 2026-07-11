// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build ts_enable_syspolicy

package policyclient

import "github.com/metacubex/tailscale/util/syspolicy/setting"

// policySnapshot is an alias for [setting.Snapshot] when syspolicy is enabled.
type policySnapshot = setting.Snapshot
