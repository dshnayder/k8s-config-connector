// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigLakeCatalogGVK = GroupVersion.WithKind("BigLakeCatalog")

type Parent struct {
	// Required. The location of the application.
	Location string `json:"location,omitempty"`

	// Required. The host project of the application.
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`
}

// BigLakeCatalogSpec defines the desired state of BigLakeCatalog
// +kcc:proto=google.cloud.bigquery.biglake.v1.Catalog
type BigLakeCatalogSpec struct {
	// Required. Defines the parent path of the resource.
	*Parent `json:",inline"`

	// The BigLakeCatalog name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// BigLakeCatalogStatus defines the config connector machine state of BigLakeCatalog
type BigLakeCatalogStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigLakeCatalog resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigLakeCatalogObservedState `json:"observedState,omitempty"`
}

// BigLakeCatalogObservedState is the state of the BigLakeCatalog resource as most recently observed in GCP.
// +kcc:proto=google.cloud.bigquery.biglake.v1.Catalog
type BigLakeCatalogObservedState struct {
	// Output only. The creation time of the catalog.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Catalog.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last modification time of the catalog.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Catalog.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbiglakecatalog;gcpbiglakecatalogs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigLakeCatalog is the Schema for the BigLakeCatalog API
// +k8s:openapi-gen=true
type BigLakeCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigLakeCatalogSpec   `json:"spec,omitempty"`
	Status BigLakeCatalogStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigLakeCatalogList contains a list of BigLakeCatalog
type BigLakeCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigLakeCatalog `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigLakeCatalog{}, &BigLakeCatalogList{})
}
