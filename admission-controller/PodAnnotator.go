package admission_controller

import (
	"k8s.io/api/admission/v1"
)

type PodAnnotator interface {
	Handle(v1.AdmissionRequest) (*v1.AdmissionResponse, error)
}
