package reembed

import (
	baseConstraints "github.com/mikhalytch/eggs/constraints/base"
	"github.com/mikhalytch/eggs/constraints/safecast"
)

func Int32[I safecast.ToInt32, O baseConstraints.Int32](t I) O { return O(t) }
