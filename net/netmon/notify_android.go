// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build android

package netmon

// NotifyNetworkChange is called by the Android host when the default network
// changes. ifName is the LinkProperties interface name of the new default
// network, or "" when connectivity was lost.
func NotifyNetworkChange(ifName, gateway string) {
	UpdateLastKnownDefaultRouteInterface(ifName)
	UpdateLastKnownDefaultGateway(gateway)
	for _, m := range liveMonitorSnapshot() {
		m.InjectEvent()
	}
}
