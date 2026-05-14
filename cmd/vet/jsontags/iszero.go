// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

package jsontags

import (
	"go/types"
	"reflect"

	"github.com/metacubex/tailscale/util/set"
)

var _ = reflect.Value.IsZero // refer for hot-linking purposes

var pureIsZeroMethods map[string]set.Set[string]

// hasPureIsZeroMethod reports whether the IsZero method is truly
// identical to [reflect.Value.IsZero].
func hasPureIsZeroMethod(t types.Type) bool {
	// TODO: Detect this automatically by checking the method AST?
	path, name := typeName(t)
	return pureIsZeroMethods[path].Contains(name)
}

// PureIsZeroMethodsInTailscaleModule is a list of known IsZero methods
// in the "tailscale.com" module that are pure.
var PureIsZeroMethodsInTailscaleModule = map[string]set.Set[string]{
	"github.com/metacubex/tailscale/net/packet": set.Of(
		"TailscaleRejectReason",
	),
	"github.com/metacubex/tailscale/tailcfg": set.Of(
		"UserID",
		"LoginID",
		"NodeID",
		"StableNodeID",
	),
	"github.com/metacubex/tailscale/tka": set.Of(
		"AUMHash",
	),
	"github.com/metacubex/tailscale/types/geo": set.Of(
		"Point",
	),
	"github.com/metacubex/tailscale/tstime/mono": set.Of(
		"Time",
	),
	"github.com/metacubex/tailscale/types/key": set.Of(
		"NLPrivate",
		"NLPublic",
		"DERPMesh",
		"MachinePrivate",
		"MachinePublic",
		"ControlPrivate",
		"DiscoPrivate",
		"DiscoPublic",
		"DiscoShared",
		"HardwareAttestationPublic",
		"ChallengePublic",
		"NodePrivate",
		"NodePublic",
	),
	"github.com/metacubex/tailscale/types/netlogtype": set.Of(
		"Connection",
		"Counts",
	),
}

// RegisterPureIsZeroMethods specifies a list of pure IsZero methods
// where it is identical to calling [reflect.Value.IsZero] on the receiver.
// This is not strictly necessary, but allows for more accurate
// detection of improper use of `json` tags.
//
// This must be called at init and the input must not be mutated.
func RegisterPureIsZeroMethods(methods map[string]set.Set[string]) {
	pureIsZeroMethods = methods
}
