// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VpcPeeringConnectionSpec defines the desired state of VpcPeeringConnection.
//
// Describes a VPC peering connection.
type VPCPeeringConnectionSpec struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `json:"dryRun,omitempty"`
	// The Amazon Web Services account ID of the owner of the accepter VPC.
	//
	// Default: Your Amazon Web Services account ID
	PeerOwnerID *string `json:"peerOwnerID,omitempty"`
	// The Region code for the accepter VPC, if the accepter VPC is located in a
	// Region other than the Region in which you make the request.
	//
	// Default: The Region in which you make the request.
	PeerRegion *string `json:"peerRegion,omitempty"`
	// The ID of the VPC with which you are creating the VPC peering connection.
	// You must specify this parameter in the request.
	PeerVPCID *string `json:"peerVPCID,omitempty"`
	// The tags to assign to the peering connection.
	TagSpecifications []*TagSpecification `json:"tagSpecifications,omitempty"`
	// The ID of the requester VPC. You must specify this parameter in the request.
	VPCID *string `json:"vpcID,omitempty"`
}

// VPCPeeringConnectionStatus defines the observed state of VPCPeeringConnection
type VPCPeeringConnectionStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// Information about the accepter VPC. CIDR block information is only returned
	// when describing an active VPC peering connection.
	// +kubebuilder:validation:Optional
	AccepterVPCInfo *VPCPeeringConnectionVPCInfo `json:"accepterVPCInfo,omitempty"`
	// The time that an unaccepted VPC peering connection will expire.
	// +kubebuilder:validation:Optional
	ExpirationTime *metav1.Time `json:"expirationTime,omitempty"`
	// Information about the requester VPC. CIDR block information is only returned
	// when describing an active VPC peering connection.
	// +kubebuilder:validation:Optional
	RequesterVPCInfo *VPCPeeringConnectionVPCInfo `json:"requesterVPCInfo,omitempty"`
	// The status of the VPC peering connection.
	// +kubebuilder:validation:Optional
	Status *VPCPeeringConnectionStateReason `json:"status,omitempty"`
	// Any tags assigned to the resource.
	// +kubebuilder:validation:Optional
	Tags []*Tag `json:"tags,omitempty"`
	// The ID of the VPC peering connection.
	// +kubebuilder:validation:Optional
	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionID,omitempty"`
}

// VPCPeeringConnection is the Schema for the VPCPeeringConnections API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type VPCPeeringConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCPeeringConnectionSpec   `json:"spec,omitempty"`
	Status            VPCPeeringConnectionStatus `json:"status,omitempty"`
}

// VPCPeeringConnectionList contains a list of VPCPeeringConnection
// +kubebuilder:object:root=true
type VPCPeeringConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCPeeringConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VPCPeeringConnection{}, &VPCPeeringConnectionList{})
}