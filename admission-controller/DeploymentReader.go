package admission_controller

import (
	"encoding/json"
	"fmt"
	v1 "k8s.io/api/admission/v1"
	appsv1 "k8s.io/api/apps/v1"
)

type DeploymentReader struct{}

func (dr *DeploymentReader) Read(req v1.AdmissionRequest) (*appsv1.Deployment, error) {
	deployment := &appsv1.Deployment{}                                 // esto es un ejemplo de como se puede leer un deployment donde appsv1 es el paquete de k8s.io/api/apps/v1
	if err := json.Unmarshal(req.Object.Raw, deployment); err != nil { // se verifica si el objeto es un deployment y se deserializa el objeto en un struct de go
		return nil, fmt.Errorf("failed to unmarshal deployment: %v", err)
	}
	return deployment, nil // se retorna el deployment deserializado en caso de que no haya errores en la deserialización del objeto, nil en caso contrario
}
