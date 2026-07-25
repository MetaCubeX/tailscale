// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package jsdeps is a just a list of the packages we import in the
// JavaScript/WASM build, to let us test that our transitive closure of
// dependencies doesn't accidentally grow too large, since binary size
// is more of a concern.
package jsdeps

import (
	_ "bytes"
	_ "context"
	_ "encoding/hex"
	_ "encoding/json"
	_ "fmt"
	_ "log"
	_ "github.com/metacubex/randv2"
	_ "net"
	_ "strings"
	_ "time"

	_ "golang.org/x/crypto/ssh"
	_ "github.com/metacubex/tailscale/control/controlclient"
	_ "github.com/metacubex/tailscale/ipn"
	_ "github.com/metacubex/tailscale/ipn/ipnserver"
	_ "github.com/metacubex/tailscale/net/netaddr"
	_ "github.com/metacubex/tailscale/net/netns"
	_ "github.com/metacubex/tailscale/net/tsdial"
	_ "github.com/metacubex/tailscale/safesocket"
	_ "github.com/metacubex/tailscale/tailcfg"
	_ "github.com/metacubex/tailscale/types/logger"
	_ "github.com/metacubex/tailscale/wgengine"
	_ "github.com/metacubex/tailscale/wgengine/netstack"
	_ "github.com/metacubex/tailscale/words"
)
