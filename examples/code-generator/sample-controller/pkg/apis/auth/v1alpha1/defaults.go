package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

func SetDefaults_FooStatus(obj *FooStatus) {
	if obj.AvailableReplicas != 0 {
		obj.AvailableReplicas = 0
	}
}
