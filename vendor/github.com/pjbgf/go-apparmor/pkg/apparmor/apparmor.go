package apparmor

import (
	"github.com/go-logr/logr"
)

type aa interface {
	// WithLogger sets a logger to be using whilst executing operations.
	WithLogger(logger logr.Logger) aa

	// Enabled checks whether AppArmor is enabled in the kernel.
	Enabled() (bool, error)

	// Enforceable checks whether AppArmor is installed, enabled and that
	// policies are enforceable.
	Enforceable() (bool, error)

	// AppArmorFS returns the path where the AppArmor filesystem
	// was mounted.
	AppArmorFS() (string, error)

	// DeletePolicy removes an AppArmor policy from the kernel.
	DeletePolicy(policyName string) error

	// LoadPolicy loads an AppArmor policy into the kernel.
	LoadPolicy(fileName string) error
}
