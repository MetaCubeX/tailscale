// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

package netmon

import "sync"

// Hosts that embed tsnet instead of using libtailscale have no handle on the
// Monitor owned by each tsnet.Server, so keep a registry of live monitors for
// the host to wake when the platform reports a default-network change.
var (
	liveMonitorsMu sync.Mutex
	liveMonitors   = make(map[*Monitor]struct{})
)

func registerLiveMonitor(m *Monitor) {
	liveMonitorsMu.Lock()
	liveMonitors[m] = struct{}{}
	liveMonitorsMu.Unlock()
}

func unregisterLiveMonitor(m *Monitor) {
	liveMonitorsMu.Lock()
	delete(liveMonitors, m)
	liveMonitorsMu.Unlock()
}

func liveMonitorSnapshot() []*Monitor {
	liveMonitorsMu.Lock()
	defer liveMonitorsMu.Unlock()
	ms := make([]*Monitor, 0, len(liveMonitors))
	for m := range liveMonitors {
		ms = append(ms, m)
	}
	return ms
}
