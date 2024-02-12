package admission_controller

import (
	appsv1 "k8s.io/api/apps/v1"
)

type Annotator interface {
	Annotate(deployment **appsv1.Deployment) error
}

type PodAnnotator struct{}

func (pa *PodAnnotator) Annotate(deployment *appsv1.Deployment) error {
	if deployment.Annotations == nil {
		deployment.Annotations = make(map[string]string)
	}

	// Ejemplo de cómo manejar la anotación 'opa-enabled'
	if _, exists := deployment.Annotations["opa-enabled"]; !exists {
		// Si la anotación 'opa-enabled'
		deployment.Annotations["opa-enabled"] = "true"
	}

	// Ejemplo de cómo manejar la anotación 'opa-cfg-file'
	if _, exists := deployment.Annotations["opa-cfg-file"]; !exists {
		// Si la anotación 'opa-cfg-file' no existe, se podría añadir
		deployment.Annotations["opa-cfg-file"] = "config_file_name"
	}

	// Ejemplo de cómo manejar la anotación 'opa-bundle-name'
	if _, exists := deployment.Annotations["opa-bundle-name"]; !exists {
		// Si la anotación 'opa-bundle-name' no existe, se podría añadir
		deployment.Annotations["opa-bundle-name"] = "bundle_name"
	}

	return nil
}
