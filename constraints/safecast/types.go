package safecast

import baseConstraints "github.com/mikhalytch/eggs/constraints/base"

type (
	ToInt64 interface {
		baseConstraints.Signed |
			baseConstraints.Uint | baseConstraints.Uint8 | baseConstraints.Uint16 | baseConstraints.Uint32
	}
)
