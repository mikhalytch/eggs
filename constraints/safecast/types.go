package safecast

import baseConstraints "github.com/mikhalytch/eggs/constraints/base"

type (
	ToInt64 interface {
		baseConstraints.Signed |
			baseConstraints.Uint8 | baseConstraints.Uint16 | baseConstraints.Uint32
	}
	ToInt32 interface {
		baseConstraints.Int32 | baseConstraints.Int16 | baseConstraints.Int8 |
			baseConstraints.Uint8 | baseConstraints.Uint16
	}
)
