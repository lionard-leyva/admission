package admission_controller_test

import (
	"admission-controller"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/admission/v1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"testing"
)

func TestAnnotateAddsOpaEnabledWhenNotPresent(t *testing.T) {
	annotator := &admission_controller.PodAnnotator{}
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{},
		},
	}

	err := annotator.Annotate(deployment)

	assert.Nil(t, err)
	assert.Equal(t, "true", deployment.Annotations["opa-enabled"])
}

func TestAnnotateDoesNotOverrideExistingOpaEnabled(t *testing.T) {
	annotator := &admission_controller.PodAnnotator{}
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"opa-enabled": "false",
			},
		},
	}

	err := annotator.Annotate(deployment)

	assert.Nil(t, err)
	assert.Equal(t, "false", deployment.Annotations["opa-enabled"])
}

func TestReadReturnsErrorWhenUnmarshallingFails(t *testing.T) {
	reader := &admission_controller.DeploymentReader{}
	req := v1.AdmissionRequest{
		Object: runtime.RawExtension{
			Raw: []byte("invalid-json"),
		},
	}

	_, err := reader.Read(req)

	assert.NotNil(t, err)
}

func TestReadReturnsDeploymentWhenUnmarshallingSucceeds(t *testing.T) {
	reader := &admission_controller.DeploymentReader{}
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-deployment",
		},
	}
	rawDeployment, _ := json.Marshal(deployment)
	req := v1.AdmissionRequest{
		Object: runtime.RawExtension{
			Raw: rawDeployment,
		},
	}

	returnedDeployment, err := reader.Read(req)

	assert.Nil(t, err)
	assert.Equal(t, deployment.Name, returnedDeployment.Name)
}
