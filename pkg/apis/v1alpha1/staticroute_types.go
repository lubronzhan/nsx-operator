/* Copyright © 2022 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: Apache-2.0 */

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StaticRouteStatusCondition string

// StaticRouteCondition defines condition of StaticRoute.
type StaticRouteCondition Condition

// StaticRouteSpec defines static routes configuration on VPC.
type StaticRouteSpec struct {
	// Specify network address in CIDR format.
	// +kubebuilder:validation:Format=cidr
	Network string `json:"network"`
	// Next hop gateway
	// +kubebuilder:validation:MinItems=1
	NextHops []NextHop `json:"nextHops"`
}

// NextHop defines next hop configuration for network.
type NextHop struct {
	// Next hop gateway IP address.
	// +kubebuilder:validation:Format=ip
	IPAddress string `json:"ipAddress"`
}

// StaticRouteStatus defines the observed state of StaticRoute.
type StaticRouteStatus struct {
	Conditions []StaticRouteCondition `json:"conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion

// StaticRoute is the Schema for the staticroutes API.
type StaticRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StaticRouteSpec   `json:"spec,omitempty"`
	Status StaticRouteStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// StaticRouteList contains a list of StaticRoute.
type StaticRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StaticRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StaticRoute{}, &StaticRouteList{})
}
