package go_qris

type POI string

const (
	PointOfInitiationMethodStatic  POI = "11"
	PointOfInitiationMethodDynamic POI = "12"
)

var ListPOIMethod = []POI{
	PointOfInitiationMethodStatic,
	PointOfInitiationMethodDynamic,
}
