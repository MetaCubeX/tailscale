// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build aix || solaris || illumos

package tstun

import (
	"github.com/metacubex/tailscale-wireguard-go/tun"
	"github.com/metacubex/tailscale/types/logger"
)

func New(logf logger.Logf, tunName string) (tun.Device, string, error) {
	panic("not implemented")
}

func Diagnose(logf logger.Logf, tunName string, err error) {
	panic("not implemented")
}
