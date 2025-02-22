// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PrivilegesInitParameters struct {

	// if nulll the privilege will apply to all namespaces. must not be an empty string
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (String) Privilege name
	Privilege *string `json:"privilege,omitempty" tf:"privilege,omitempty"`

	// if null the privilege will apply to all sets. Must be used with namespace. Must not be an emptry string
	Set *string `json:"set,omitempty" tf:"set,omitempty"`
}

type PrivilegesObservation struct {

	// if nulll the privilege will apply to all namespaces. must not be an empty string
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (String) Privilege name
	Privilege *string `json:"privilege,omitempty" tf:"privilege,omitempty"`

	// if null the privilege will apply to all sets. Must be used with namespace. Must not be an emptry string
	Set *string `json:"set,omitempty" tf:"set,omitempty"`
}

type PrivilegesParameters struct {

	// if nulll the privilege will apply to all namespaces. must not be an empty string
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (String) Privilege name
	// +kubebuilder:validation:Optional
	Privilege *string `json:"privilege" tf:"privilege,omitempty"`

	// if null the privilege will apply to all sets. Must be used with namespace. Must not be an emptry string
	// +kubebuilder:validation:Optional
	Set *string `json:"set,omitempty" tf:"set,omitempty"`
}

type RoleInitParameters struct {

	// (Attributes Set) Privilege set, comprised from {privilege="name",namespace="name",set="name"] maps. Namespace and Set are optional (see below for nested schema)
	// Privilege set, comprised from {privilege="name",namespace="name",set="name"] maps. Namespace and Set are optional
	Privileges []PrivilegesInitParameters `json:"privileges,omitempty" tf:"privileges,omitempty"`

	// (Number) Read quota to apply to the role
	// Read quota to apply to the role
	ReadQuota *float64 `json:"readQuota,omitempty" tf:"read_quota,omitempty"`

	// (List of String) A list of IP addresses allowed to connect.
	// A list of IP addresses allowed to connect.
	WhiteList []*string `json:"whiteList,omitempty" tf:"white_list,omitempty"`

	// (Number) write quota to apply to the role
	// write quota to apply to the role
	WriteQuota *float64 `json:"writeQuota,omitempty" tf:"write_quota,omitempty"`
}

type RoleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Attributes Set) Privilege set, comprised from {privilege="name",namespace="name",set="name"] maps. Namespace and Set are optional (see below for nested schema)
	// Privilege set, comprised from {privilege="name",namespace="name",set="name"] maps. Namespace and Set are optional
	Privileges []PrivilegesObservation `json:"privileges,omitempty" tf:"privileges,omitempty"`

	// (Number) Read quota to apply to the role
	// Read quota to apply to the role
	ReadQuota *float64 `json:"readQuota,omitempty" tf:"read_quota,omitempty"`

	// (List of String) A list of IP addresses allowed to connect.
	// A list of IP addresses allowed to connect.
	WhiteList []*string `json:"whiteList,omitempty" tf:"white_list,omitempty"`

	// (Number) write quota to apply to the role
	// write quota to apply to the role
	WriteQuota *float64 `json:"writeQuota,omitempty" tf:"write_quota,omitempty"`
}

type RoleParameters struct {

	// (Attributes Set) Privilege set, comprised from {privilege="name",namespace="name",set="name"] maps. Namespace and Set are optional (see below for nested schema)
	// Privilege set, comprised from {privilege="name",namespace="name",set="name"] maps. Namespace and Set are optional
	// +kubebuilder:validation:Optional
	Privileges []PrivilegesParameters `json:"privileges,omitempty" tf:"privileges,omitempty"`

	// (Number) Read quota to apply to the role
	// Read quota to apply to the role
	// +kubebuilder:validation:Optional
	ReadQuota *float64 `json:"readQuota,omitempty" tf:"read_quota,omitempty"`

	// (List of String) A list of IP addresses allowed to connect.
	// A list of IP addresses allowed to connect.
	// +kubebuilder:validation:Optional
	WhiteList []*string `json:"whiteList,omitempty" tf:"white_list,omitempty"`

	// (Number) write quota to apply to the role
	// write quota to apply to the role
	// +kubebuilder:validation:Optional
	WriteQuota *float64 `json:"writeQuota,omitempty" tf:"write_quota,omitempty"`
}

// RoleSpec defines the desired state of Role
type RoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RoleInitParameters `json:"initProvider,omitempty"`
}

// RoleStatus defines the observed state of Role.
type RoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Role is the Schema for the Roles API. Aerospike Role
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aerospike}
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privileges) || (has(self.initProvider) && has(self.initProvider.privileges))",message="spec.forProvider.privileges is a required parameter"
	Spec   RoleSpec   `json:"spec"`
	Status RoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleList contains a list of Roles
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Role `json:"items"`
}

// Repository type metadata.
var (
	Role_Kind             = "Role"
	Role_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Role_Kind}.String()
	Role_KindAPIVersion   = Role_Kind + "." + CRDGroupVersion.String()
	Role_GroupVersionKind = CRDGroupVersion.WithKind(Role_Kind)
)

func init() {
	SchemeBuilder.Register(&Role{}, &RoleList{})
}
