package main

import (
	"encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	PLATFORM  = "platform"
	PROJECT   = "project"
	CLUSTER   = "cluster"
	NAMESPACE = "namespace"

	QuotaCPU      = "cpu"
	QuotaMEM      = "memory"
	QuotaStorage  = "storage"
	QuotaPODS     = "pods"
	QuotaPVC      = "persistentvolumeclaims"
	QuotaEStorage = "ephemeral-storage"

	PrefixRequests = "requests"
	PrefixLimits   = "limits"
)

var (
	RequestsCPU     = fmt.Sprintf("%s.%s", PrefixRequests, QuotaCPU)
	RequestsMEM     = fmt.Sprintf("%s.%s", PrefixRequests, QuotaMEM)
	RequestsStorage = fmt.Sprintf("%s.%s", PrefixRequests, QuotaStorage)
	LimitsCPU       = fmt.Sprintf("%s.%s", PrefixLimits, QuotaCPU)
	LimitsMEM       = fmt.Sprintf("%s.%s", PrefixLimits, QuotaMEM)
)

var SupportedQuotaTypes = map[corev1.ResourceName]bool{
	corev1.ResourceName(RequestsCPU):     true,
	corev1.ResourceName(LimitsCPU):       true,
	corev1.ResourceName(RequestsMEM):     true,
	corev1.ResourceName(LimitsMEM):       true,
	corev1.ResourceName(RequestsStorage): true,
	corev1.ResourceName(QuotaPODS):       true,
	corev1.ResourceName(QuotaPVC):        true,
}

func formatResourceName(source corev1.ResourceName) []corev1.ResourceName {
	/*
		ephemeral-storage  -> requests.storage
		cpu                -> requests.cpu and limits.cpu
		memory             -> requests.memory and limits.memory
	*/
	result := make([]corev1.ResourceName, 0)
	switch string(source) {
	case QuotaEStorage:
		result = append(result, corev1.ResourceName(RequestsStorage))
	case QuotaCPU:
		result = append(result, corev1.ResourceName(RequestsCPU))
		result = append(result, corev1.ResourceName(LimitsCPU))
	case QuotaMEM:
		result = append(result, corev1.ResourceName(RequestsMEM))
		result = append(result, corev1.ResourceName(LimitsMEM))
	default:
		result = append(result, source)
	}
	return result
}

func formatResourceList(source corev1.ResourceList) corev1.ResourceList {
	result := corev1.ResourceList{}

	for key, value := range source {
		quantity := *value.Copy()
		resourceNames := formatResourceName(key)
		fmt.Println("resourceNames: ", resourceNames)
		for _, resourceName := range resourceNames {
			if _, supported := SupportedQuotaTypes[resourceName]; supported {
				result[resourceName] = quantity
			}
		}
	}

	return result
}

func main() {
	// precision: 15 - int32(len(num)) - int32(float32(exponent)*3/10) - 1
	resourcesMap := map[string]string{
		// "EiM": "7Ei",                // {"exponent": 60, "precision": -5, "num": 2 << 3 - 1}
		// "PiM": "8191Pi",             // {"exponent": 50, "precision": -5, "num": 2 << 13 - 1}
		// "TiM": "8388607Ti",          // {"exponent": 40, "precision": -5, "num": 2 << 23 - 1}
		// "GiM": "8589934591Gi",       // {"exponent": 30, "precision": -5, "num": 2 << 33 - 1}
		// "MiM": "8796093022207Mi",    // {"exponent": 20, "precision": -5, "num": 2 << 43 - 1}
		// "KiM": "9007199254740991Ki", // {"exponent": 10, "precision": -5, "num": 2 << 53 - 1}
		// "EiN": "1Ei",                // {"exponent": 60, "precision": -5, "num": 1}
		// "PiN": "1Pi",                // {"exponent": 50, "precision": -2, "num": 1}
		// "TiN": "99Ti",               // {"exponent": 40, "precision": 0,  "num": 99}
		// "GiN": "99999Gi",            // {"exponent": 30, "precision": 0,  "num": 99999}
		// "MiN": "99999999Mi",         // {"exponent": 20, "precision": 0,  "num": 99999999}
		// "KiN": "99999999999Ki",      // {"exponent": 10, "precision": 0,  "num": 99999999999}
		"cpu":                    "18",
		"ephemeral-storage":      "257501508Ki",
		"hugepages-1Gi":          "0",
		"hugepages-2Mi":          "0",
		"memory":                 "36166196Ki",
		"pods":                   "",
		"persistentvolumeclaims": "1000",
	}

	bytes, _ := json.Marshal(resourcesMap)

	resources := corev1.ResourceList{}
	json.Unmarshal(bytes, &resources)

	for name, quantity := range resources {
		fmt.Println(fmt.Sprintf("name: %s, quantity: %+v", name, quantity.String()))
	}

	fmt.Println("SupportedQuota: ", SupportedQuotaTypes)

	result := formatResourceList(resources)
	fmt.Println(result)

	var projectCluster struct {
		// Name of the cluster
		Name string `json:"name"`
		// Type of the cluster
		Type string `json:"type"`
		// Quota store the quota info for Project
		Quota map[string]*json.RawMessage `json:"quota"`
	}

	cluster := `{
		"name":  "alazyer",
	}`

	json.Unmarshal([]byte(cluster), &projectCluster)
	fmt.Printf("projectCluster: %+v", projectCluster)

	a, b := json.Marshal(resource.Quantity{})

	fmt.Printf("empty quantity: %v, %v", string(a), b)
}
