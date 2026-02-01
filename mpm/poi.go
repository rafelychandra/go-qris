package mpm

import "fmt"

type POI string

const (
	PointOfInitiationMethodStatic  POI = "11"
	PointOfInitiationMethodDynamic POI = "12"
)

var ListPOIMethod = []POI{
	PointOfInitiationMethodStatic,
	PointOfInitiationMethodDynamic,
}

func (p POI) Validate() error {
	for _, v := range ListPOIMethod {
		if v == p {
			return nil
		}
	}
	return fmt.Errorf("point of initiation method")
}
