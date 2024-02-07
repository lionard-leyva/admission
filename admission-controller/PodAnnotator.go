package admission_controller

import (
	"k8s.io/api/admission/v1"
)

/**
* PodAnnotator interface
* Esta interfaz define el método Handle que recibe un objeto de tipo v1.AdmissionRequest
*y retorna un objeto de tipo v1.AdmissionResponse y un error
 */
type PodAnnotator interface {
	Handle(v1.AdmissionRequest) (*v1.AdmissionResponse, error)
}
