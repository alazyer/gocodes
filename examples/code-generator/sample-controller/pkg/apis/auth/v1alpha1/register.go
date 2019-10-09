package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	SchemeBuilder      = runtime.NewSchemeBuilder(addDefaultingFuncs)
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	localSchemeBuilder.Register(addDefaultingFuncs)
}
