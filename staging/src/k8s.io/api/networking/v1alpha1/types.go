/*
Copyright 2022 The Kubernetes Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.27

// IPAddress represents a single IP of a single IP Family. The object is designed to be used by APIs
// that operate on IP addresses. The object is used by the Service core API for allocation of IP addresses.
// An IP address can be represented in different formats, to guarantee the uniqueness of the IP,
// the name of the object is the IP address in canonical format, four decimal digits separated
// by dots suppressing leading zeros for IPv4 and the representation defined by RFC 5952 for IPv6.
// Valid: 192.168.1.5 or 2001:db8::1 or 2001:db8:aaaa:bbbb:cccc:dddd:eeee:1
// Invalid: 10.01.2.3 or 2001:db8:0:0:0::1
type IPAddress struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// spec is the desired state of the IPAddress.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Spec IPAddressSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// IPAddressSpec describe the attributes in an IP Address.
type IPAddressSpec struct {
	// ParentRef references the resource that an IPAddress is attached to.
	// An IPAddress must reference a parent object.
	// +required
	ParentRef *ParentReference `json:"parentRef,omitempty" protobuf:"bytes,1,opt,name=parentRef"`
}

// ParentReference describes a reference to a parent object.
type ParentReference struct {
	// Group is the group of the object being referenced.
	// +optional
	Group string `json:"group,omitempty" protobuf:"bytes,1,opt,name=group"`
	// Resource is the resource of the object being referenced.
	// +required
	Resource string `json:"resource,omitempty" protobuf:"bytes,2,opt,name=resource"`
	// Namespace is the namespace of the object being referenced.
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace"`
	// Name is the name of the object being referenced.
	// +required
	Name string `json:"name,omitempty" protobuf:"bytes,4,opt,name=name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.27

// IPAddressList contains a list of IPAddress.
type IPAddressList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// items is the list of IPAddresses.
	Items []IPAddress `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.27

// ServiceCIDR defines a range of IP addresses using CIDR format (e.g. 192.168.0.0/24 or 2001:db2::/64).
// This range is used to allocate ClusterIPs to Service objects.
type ServiceCIDR struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// spec is the desired state of the ServiceCIDR.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Spec ServiceCIDRSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	// status represents the current state of the ServiceCIDR.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Status ServiceCIDRStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// ServiceCIDRSpec define the CIDRs the user wants to use for allocating ClusterIPs for Services.
type ServiceCIDRSpec struct {
	// CIDRs defines the IP blocks in CIDR notation (e.g. "192.168.0.0/24" or "2001:db8::/64")
	// from which to assign service cluster IPs. Max of two CIDRs is allowed, one of each IP family.
	// This field is immutable.
	// +optional
	CIDRs []string `json:"cidrs,omitempty" protobuf:"bytes,1,opt,name=cidrs"`
}

const (
	// ServiceCIDRConditionReady represents status of a ServiceCIDR that is ready to be used by the
	// apiserver to allocate ClusterIPs for Services.
	ServiceCIDRConditionReady = "Ready"
	// ServiceCIDRReasonTerminating represents a reason where a ServiceCIDR is not ready because it is
	// being deleted.
	ServiceCIDRReasonTerminating = "Terminating"
)

// ServiceCIDRStatus describes the current state of the ServiceCIDR.
type ServiceCIDRStatus struct {
	// conditions holds an array of metav1.Condition that describe the state of the ServiceCIDR.
	// Current service state
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.27

// ServiceCIDRList contains a list of ServiceCIDR objects.
type ServiceCIDRList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// items is the list of ServiceCIDRs.
	Items []ServiceCIDR `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.30

// PodNetwork represents a logical network in the K8s Cluster.
// This logical network depends on the host networking setup on cluster nodes.
type PodNetwork struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the behavior of a PodNetwork.
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Spec PodNetworkSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Most recently observed status of the PodNetwork.
	// Populated by the system.
	// Read-only.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Status PodNetworkStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// PodNetworkSpec contains the specifications for podNetwork object
type PodNetworkSpec struct {

	// Enabled is used to administratively enable/disable a PodNetwork.
	// When set to false, PodNetwork Ready condition will be set to False.
	// Defaults to True.
	//
	// +optional
	Enabled bool `json:"enabled,omitempty" protobuf:"bytes,1,opt,name=enabled"`

	// ParametersRef points to the vendor or implementation specific params for the
	// podNetwork.
	// +optional
	ParametersRef *ParametersRef `json:"parametersRef,omitempty" protobuf:"bytes,2,opt,name=parametersRef"`

	// Provider specifies the provider implementing this PodNetwork.
	// +optional
	Provider string `json:"provider,omitempty" protobuf:"bytes,3,opt,name=provider"`
}

// ParametersRef defines a custom resource containing additional parameters for the
// PodNetwork.
type ParametersRef struct {
	// Group is the group of the object being referenced.
	Group string `json:"group" protobuf:"bytes,1,opt,name=group"`
	// Kind is the resource of the object being referenced.
	Kind string `json:"kind" protobuf:"bytes,2,opt,name=kind"`
	// Namespace is the namespace of the object being referenced.
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace"`
	// Name is the name of the object being referenced.
	// +required
	Name string `json:"name" protobuf:"bytes,4,opt,name=name"`
}

// PodNetworkConditionType is the type for status conditions on
// a PodNetwork. This type should be used with the
// PodNetworkStatus.Conditions field.
type PodNetworkConditionType string

const (
	// PodNetworkConditionStatusReady represents that the PodNetwork object is
	// correct (validated) and all other conditions are set to “true”. This
	// condition will switch back to “false” if any of the other conditions are
	// “false”. This condition does not indicate readiness of specific PodNetwork
	// on a per Node-basis.
	PodNetworkConditionStatusReady PodNetworkConditionType = "Ready"

	// PodNetworkConditionStatusParamsReady represents that object specified in
	// the “parametersRef” field is ready for use. The owner of the specified
	// object is responsible for handling this condition. The “Ready” condition is
	// dependent on the value of this condition when the “parametersRef” field is
	// not empty. The available “reasons” for this condition are implementation
	// specific.
	PodNetworkConditionStatusParamsReady PodNetworkConditionType = "ParamsReady"
)

// PodNetworkConditionReason defines the set of reasons that explain why a
// particular PodNetwork condition type has been raised.
type PodNetworkConditionReason string

const (
	// PodNetworkConditionReasonParamsNotReady represents a reason where the
	// ParamsReady condition is not present or has “false” value. This can only
	// happen when the “parametersRef” field has a value.
	PodNetworkConditionReasonParamsNotReady PodNetworkConditionReason = "ParamsNotReady"

	// PodNetworkConditionReasonAdministrativelyDisabled represents a reason where
	// the PodNetwork object's Enabled field is set to false.
	PodNetworkConditionReasonAdministrativelyDisabled PodNetworkConditionReason = "AdministrativelyDisabled"
)

// PodNetworkStatus contains the status information related to the PodNetwork.
type PodNetworkStatus struct {
	// Conditions describe the current state of the PodNetwork.
	//
	// Known condition types are:
	//
	// * "Ready"
	// * "ParamsReady"
	//
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.27

// PodNetworkList contains a list of PodNetwork.
type PodNetworkList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// items is the list of PodNetworks.
	Items []PodNetwork `json:"items" protobuf:"bytes,2,rep,name=items"`
}
