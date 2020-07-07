package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// MacallocationSpec defines the desired state of Macallocation
type MacallocationSpec struct {
  Ipaddress int `json:"ipaddress"`
}

// +kubebuilder:object:root=true

// Macallocation is the Schema for the macallocations API
type Macallocation struct {
  metav1.TypeMeta   `json:",inline"`
  metav1.ObjectMeta `json:"metadata,omitempty"`

  Spec MacallocationSpec `json:"spec"`
}

// +kubebuilder:object:root=true

// MacallocationList contains a list of Macallocation
type MacallocationList struct {
  metav1.TypeMeta `json:",inline"`
  metav1.ListMeta `json:"metadata,omitempty"`

  Items []Macallocation `json:"items"`
}
