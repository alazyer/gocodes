package main

import (
	"fmt"

	"github.com/spf13/pflag"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/component-base/featuregate"
)

const (
	PLATFORM  featuregate.Feature = "platform"
	PROJECT   featuregate.Feature = "project"
	CLUSTER   featuregate.Feature = "cluster"
	NAMESPACE featuregate.Feature = "namespace"
)

func Prepare() error {
	return utilfeature.DefaultMutableFeatureGate.Add(map[featuregate.Feature]featuregate.FeatureSpec{
		PLATFORM:  {Default: true, PreRelease: featuregate.Deprecated},
		CLUSTER:   {Default: false, PreRelease: featuregate.Alpha},
		PROJECT:   {Default: true, PreRelease: featuregate.Beta},
		NAMESPACE: {Default: true, PreRelease: featuregate.GA},
	})
}

func Close() {
	// use empty flagset to close the DefaultMetableFeatureGate
	fs := &pflag.FlagSet{}
	utilfeature.DefaultMutableFeatureGate.AddFlag(fs)
}

func Change(feature featuregate.Feature, value bool) error {
	return utilfeature.DefaultMutableFeatureGate.Set(fmt.Sprintf("%s=%t", feature, value))
}

func Enabled(feature featuregate.Feature) bool {
	return utilfeature.DefaultFeatureGate.Enabled(feature)
}

func main() {

	Prepare()
	Close()
	CheckAddToClosed()

	CheckAlpha()
	CheckBeta()
	CheckGA()
}

func CheckAlpha() {
	fmt.Println()
	Check(CLUSTER, true)
	fmt.Println()
}

func CheckBeta() {
	fmt.Println()
	Check(PROJECT, false)
	fmt.Println()
}

func CheckGA() {
	fmt.Println()
	Check(NAMESPACE, false)
	fmt.Println()
}

func Check(feature featuregate.Feature, newValue bool) {
	value := Enabled(feature)

	fmt.Printf("%s Feature default: %t", feature, value)
	fmt.Println()

	fmt.Printf("Change %s feature to %t", feature, newValue)
	fmt.Println()
	err := Change(feature, newValue)
	if err != nil {
		fmt.Printf("Change %s Feature with err: %v", feature, err)
		fmt.Println()
	}

	value = Enabled(feature)

	fmt.Printf("%s Feature after changed: %t", feature, value)
	fmt.Println()
}

func CheckAddToClosed() {
	Close()
	// add feature to closed feature gate, error occurs
	err := utilfeature.DefaultMutableFeatureGate.Add(map[featuregate.Feature]featuregate.FeatureSpec{
		CLUSTER:   {Default: false, PreRelease: featuregate.Alpha},
		PROJECT:   {Default: true, PreRelease: featuregate.Beta},
		NAMESPACE: {Default: true, PreRelease: featuregate.GA},
	})
	if err != nil {
		fmt.Println("Error occured when add new feature: ", err)
	}
}
