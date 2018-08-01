package exporterkit

import (
	"github.com/giantswarm/microerror"
)

var InvalidConfigError = &microerror.Error{
	Kind: "InvalidConfigError",
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == InvalidConfigError
}
