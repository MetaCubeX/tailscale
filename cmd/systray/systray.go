// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build cgo || !darwin

// systray is a minimal Tailscale systray application.
package main

import (
	"flag"

	"github.com/metacubex/tailscale/client/local"
	"github.com/metacubex/tailscale/client/systray"
	"github.com/metacubex/tailscale/paths"
)

var socket = flag.String("socket", paths.DefaultTailscaledSocket(), "path to tailscaled socket")
var theme = flag.String("theme", "dark", "color theme for Tailscale icon: dark, dark:nobg, light, light:nobg")

func main() {
	flag.Parse()
	lc := &local.Client{Socket: *socket}
	systray.SetTheme(*theme)
	new(systray.Menu).Run(lc)
}
