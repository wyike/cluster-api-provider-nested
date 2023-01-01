/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha4

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	addonv1alpha1 "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NestedschedulerSpec defines the desired state of Nestedscheduler
type NestedschedulerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Nestedscheduler. Edit nestedscheduler_types.go to remove/update
	NestedComponentSpec `json:",inline"`
}

// NestedschedulerStatus defines the observed state of Nestedscheduler
type NestedschedulerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	addonv1alpha1.CommonStatus `json:",inline"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Nestedscheduler is the Schema for the nestedschedulers API
type Nestedscheduler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NestedschedulerSpec   `json:"spec,omitempty"`
	Status NestedschedulerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NestedschedulerList contains a list of Nestedscheduler
type NestedschedulerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Nestedscheduler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Nestedscheduler{}, &NestedschedulerList{})
}

var _ addonv1alpha1.CommonObject = &Nestedscheduler{}
var _ addonv1alpha1.Patchable = &Nestedscheduler{}

// ComponentName returns the name of the component for use with
// addonv1alpha1.CommonObject.
func (c *Nestedscheduler) ComponentName() string {
	return string(ControllerManager)
}

// CommonSpec returns the addons spec of the object allowing common funcs like
// Channel & Version to be usable.
func (c *Nestedscheduler) CommonSpec() addonv1alpha1.CommonSpec {
	return c.Spec.CommonSpec
}

// GetCommonStatus will return the common status for checking is a component
// was successfully deployed.
func (c *Nestedscheduler) GetCommonStatus() addonv1alpha1.CommonStatus {
	return c.Status.CommonStatus
}

// SetCommonStatus will set the status so that abstract representations can set
// Ready and Phases.
func (c *Nestedscheduler) SetCommonStatus(s addonv1alpha1.CommonStatus) {
	c.Status.CommonStatus = s
}

// PatchSpec returns the patches to be applied.
func (c *Nestedscheduler) PatchSpec() addonv1alpha1.PatchSpec {
	return c.Spec.PatchSpec
}
