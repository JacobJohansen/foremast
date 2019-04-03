package v1alpha1

import (
	"fmt"
	"foremast.ai/foremast/foremast-barrelman/pkg/apis/deployment"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: deployment.GroupName, Version: "v1alpha1"}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	// localSchemeBuilder and AddToScheme will stay in k8s.io/kubernetes.
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	// We only register manually written functions here. The registration of the
	// generated functions takes place in the generated files. The separation
	// makes the code compile even when the generated files are missing.
	localSchemeBuilder.Register(addKnownTypes)
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, &DeploymentMonitor{},
		&DeploymentMonitorList{}, &DeploymentMetadata{}, &DeploymentMetadataList{})

	scheme.AddFieldLabelConversionFunc(
		schema.GroupVersionKind{
			Group:   deployment.GroupName,
			Version: "v1alpha1",
			Kind:    "DeploymentMonitor",
		},
		func(label, value string) (string, string, error) {
			switch label {
			case "metadata.annotations",
				"metadata.labels",
				"metadata.name",
				"metadata.namespace",
				"metadata.uid",
				"spec.StartTime",
				"spec.watchUntil",
				"status.jobId",
				"status.phase":
				return label, value, nil
			default:
				return "", "", fmt.Errorf("field label not supported: %s", label)
			}
		})

	scheme.AddFieldLabelConversionFunc(
		schema.GroupVersionKind{
			Group:   deployment.GroupName,
			Version: "v1alpha1",
			Kind:    "DeploymentMetadata",
		},
		func(label, value string) (string, string, error) {
			switch label {
			case "metadata.annotations",
				"metadata.labels",
				"metadata.name",
				"metadata.namespace",
				"metadata.uid":
				return label, value, nil
			default:
				return "", "", fmt.Errorf("field label not supported: %s", label)
			}
		})

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)

	return nil
}
