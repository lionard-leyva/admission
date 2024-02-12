package admission_controller

import (
	"encoding/json"
	"fmt"
	v1 "k8s.io/api/admission/v1"
	admissionv1 "k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type AdmissionHandler struct {
	reader    *DeploymentReader
	annotator Annotator
}

func NewAdmissionHandler(reader *DeploymentReader, annotator Annotator) *AdmissionHandler {
	return &AdmissionHandler{
		reader:    reader,
		annotator: annotator,
	}
}

func (ah *AdmissionHandler) HandleAdmissionRequest(req *v1.AdmissionRequest) (*v1.AdmissionResponse, error) {
	deployment, err := ah.reader.Read(*req)
	if err != nil {
		return nil, fmt.Errorf("error reading deployment: %v", err)
	}

	if err := ah.annotator.Annotate(&deployment); err != nil {
		return nil, fmt.Errorf("error annotating deployment: %v", err)
	}

	rawReq, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	admissionReview := &admissionv1.AdmissionReview{
		Request: &admissionv1.AdmissionRequest{
			Object: runtime.RawExtension{
				Raw: rawReq,
			},
		},
	}

	admissionReview, _ = ah.MyPodMutationHandler(admissionReview)

	rawReq = admissionReview.Request.Object.Raw
	req = &v1.AdmissionRequest{}
	if err := json.Unmarshal(rawReq, req); err != nil {
		return nil, fmt.Errorf("error unmarshaling request: %v", err)
	}
	deployment, err = ah.reader.Read(*req)
	if err != nil {
		return nil, fmt.Errorf("error reading deployment: %v", err)
	}

	marshaledDeployment, err := json.Marshal(deployment)
	if err != nil {
		return nil, fmt.Errorf("error marshaling deployment: %v", err)
	}

	return &v1.AdmissionResponse{
		UID:     req.UID,
		Allowed: true,
		Patch:   marshaledDeployment,
	}, nil
}

func initContainers() ([]corev1.Container, error) {
	// Implementation of the function
	return []corev1.Container{}, nil
}

func containerEnvoy(admissionReview *admissionv1.AdmissionReview, arg2 interface{}) (*corev1.Container, error) {
	// Implementation of the function
	return &corev1.Container{}, nil
}

func containerOpa(secretValueAzConfig interface{}, enviroment *interface{}) (*corev1.Container, error) {
	// Implementation of the function
	return &corev1.Container{}, nil
}

func (ah *AdmissionHandler) MyPodMutationHandler(admissionReview *admissionv1.AdmissionReview) (*admissionv1.AdmissionReview, error) {
	//var secret_value_az_config interface{}
	//var enviroment interface{}

	// Decodificar el pod del AdmissionReview
	//pod := &corev1.Pod{}

	// Agregar contenedores de inicialización
	//	initContainers, _ := initContainers()

	// Agregar contenedor Envoy
	//envoyContainer, _ := containerEnvoy(admissionReview, nil)

	// Agregar contenedor OPA
	//opaContainer, _ := containerOpa(secret_value_az_config, &enviroment)

	return admissionReview, nil
}
