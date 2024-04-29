package entity_impl

import "fmt"

const (
	ErrInvalidTreeHeight       = "invalid tree height"
	ErrEstateOrDroneIsNil      = "estate or drone is nil"
	ErrEstateIsNil             = "estate is nil"
	ErrDroneIsNil              = "drone is nil"
	ErrPathProviderIsNil       = "path provider is nil"
	ErrNavigatorIsNil          = "navigator is nil"
	ErrInvalidEstateDimension  = "invalid estate dimension"
	ErrInvalidCoordinates      = "invalid coordinates"
	ErrNavigatorAlreadyStarted = "navigator already started"
	ErrNavigatorNotStarted     = "navigator not started"
	ErrNavigatorFailedToStart  = "navigator failed to start"
	ErrDroneInvalidMaxDistance = "invalid drone max distance"
)

func NewErr(errMsg string) error {
	return fmt.Errorf(errMsg)
}
