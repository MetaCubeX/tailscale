// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

package netmon

import (
	"net"
	"sync/atomic"
	"testing"
	"time"

	"github.com/metacubex/tailscale/util/eventbus"
)

func TestPollWallTimeRetriesAfterWake(t *testing.T) {
	m := &Monitor{
		change:    make(chan bool, 1),
		lastWall:  wallTime().Add(-time.Hour),
		wallTimer: time.NewTimer(time.Hour),
	}
	defer m.wallTimer.Stop()

	m.pollWallTime()
	select {
	case forceCallbacks := <-m.change:
		if !forceCallbacks {
			t.Fatal("initial wake event did not force callbacks")
		}
	default:
		t.Fatal("time jump did not inject a wake event")
	}

	// Simulate debounce consuming the wake event while the physical interface
	// is still down. Windows sometimes does not deliver another interface event
	// after Wi-Fi finishes reconnecting.
	m.mu.Lock()
	m.resetTimeJumpedLocked()
	m.lastWall = wallTime()
	m.mu.Unlock()

	for retry := 0; retry < postWakePollCount; retry++ {
		m.pollWallTime()
		select {
		case forceCallbacks := <-m.change:
			if forceCallbacks {
				t.Fatalf("post-wake retry %d unexpectedly forced callbacks", retry+1)
			}
		default:
			t.Fatalf("post-wake interface poll %d was not scheduled", retry+1)
		}
	}

	m.pollWallTime()
	select {
	case <-m.change:
		t.Fatalf("post-wake polling continued past %v", postWakePollDuration)
	default:
	}
}

func TestWakeRetryObservesRecoveredInterfaceWithoutOSEvent(t *testing.T) {
	var interfaceUp atomic.Bool
	interfaceUp.Store(true)

	ipnet := func(s string) net.Addr {
		ip, prefix, err := net.ParseCIDR(s)
		if err != nil {
			t.Fatal(err)
		}
		prefix.IP = ip
		return prefix
	}
	oldInterfaceGetter := altNetInterfaces
	altNetInterfaces = func() ([]Interface, error) {
		flags := net.FlagBroadcast | net.FlagMulticast
		var addrs []net.Addr
		if interfaceUp.Load() {
			flags |= net.FlagUp | net.FlagRunning
			addrs = []net.Addr{ipnet("192.0.2.10/24")}
		}
		return []Interface{{
			Interface: &net.Interface{Index: 100, MTU: 1500, Name: "wake-test0", Flags: flags},
			AltAddrs:  addrs,
		}}, nil
	}
	t.Cleanup(func() { altNetInterfaces = oldInterfaceGetter })

	bus := eventbus.New()
	defer bus.Close()
	m, err := New(bus, t.Logf)
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()

	states := make(chan bool, 2)
	m.RegisterChangeCallback(func(delta *ChangeDelta) {
		states <- delta.AnyInterfaceUp()
	})
	m.Start()

	interfaceUp.Store(false)
	m.mu.Lock()
	m.lastWall = wallTime().Add(-time.Hour)
	m.mu.Unlock()
	m.pollWallTime()
	select {
	case up := <-states:
		if up {
			t.Fatal("wake snapshot unexpectedly reported an interface up")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for offline wake snapshot")
	}

	// The interface recovers without a Windows address or route event.
	interfaceUp.Store(true)
	m.pollWallTime()
	select {
	case up := <-states:
		if !up {
			t.Fatal("post-wake poll did not observe the recovered interface")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for post-wake interface recovery")
	}
}
