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

package nat_gateway

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	acktypes "github.com/aws-controllers-k8s/runtime/pkg/types"

	svcapitypes "github.com/aws-controllers-k8s/ec2-controller/apis/v1alpha1"
)

// ClearResolvedReferences removes any reference values that were made
// concrete in the spec. It returns a copy of the input AWSResource which
// contains the original *Ref values, but none of their respective concrete
// values.
func (rm *resourceManager) ClearResolvedReferences(res acktypes.AWSResource) acktypes.AWSResource {
	ko := rm.concreteResource(res).ko.DeepCopy()

	if ko.Spec.AllocationRef != nil {
		ko.Spec.AllocationID = nil
	}

	if ko.Spec.SubnetRef != nil {
		ko.Spec.SubnetID = nil
	}

	return &resource{ko}
}

// ResolveReferences finds if there are any Reference field(s) present
// inside AWSResource passed in the parameter and attempts to resolve those
// reference field(s) into their respective target field(s). It returns a
// copy of the input AWSResource with resolved reference(s), a boolean which
// is set to true if the resource contains any references (regardless of if
// they are resolved successfully) and an error if the passed AWSResource's
// reference field(s) could not be resolved.
func (rm *resourceManager) ResolveReferences(
	ctx context.Context,
	apiReader client.Reader,
	res acktypes.AWSResource,
) (acktypes.AWSResource, bool, error) {
	namespace := res.MetaObject().GetNamespace()
	ko := rm.concreteResource(res).ko

	resourceHasReferences := false
	err := validateReferenceFields(ko)
	if fieldHasReferences, err := rm.resolveReferenceForAllocationID(ctx, apiReader, namespace, ko); err != nil {
		return &resource{ko}, (resourceHasReferences || fieldHasReferences), err
	} else {
		resourceHasReferences = resourceHasReferences || fieldHasReferences
	}

	if fieldHasReferences, err := rm.resolveReferenceForSubnetID(ctx, apiReader, namespace, ko); err != nil {
		return &resource{ko}, (resourceHasReferences || fieldHasReferences), err
	} else {
		resourceHasReferences = resourceHasReferences || fieldHasReferences
	}

	return &resource{ko}, resourceHasReferences, err
}

// validateReferenceFields validates the reference field and corresponding
// identifier field.
func validateReferenceFields(ko *svcapitypes.NATGateway) error {

	if ko.Spec.AllocationRef != nil && ko.Spec.AllocationID != nil {
		return ackerr.ResourceReferenceAndIDNotSupportedFor("AllocationID", "AllocationRef")
	}

	if ko.Spec.SubnetRef != nil && ko.Spec.SubnetID != nil {
		return ackerr.ResourceReferenceAndIDNotSupportedFor("SubnetID", "SubnetRef")
	}
	if ko.Spec.SubnetRef == nil && ko.Spec.SubnetID == nil {
		return ackerr.ResourceReferenceOrIDRequiredFor("SubnetID", "SubnetRef")
	}
	return nil
}

// resolveReferenceForAllocationID reads the resource referenced
// from AllocationRef field and sets the AllocationID
// from referenced resource. Returns a boolean indicating whether a reference
// contains references, or an error
func (rm *resourceManager) resolveReferenceForAllocationID(
	ctx context.Context,
	apiReader client.Reader,
	namespace string,
	ko *svcapitypes.NATGateway,
) (hasReferences bool, err error) {
	if ko.Spec.AllocationRef != nil && ko.Spec.AllocationRef.From != nil {
		hasReferences = true
		arr := ko.Spec.AllocationRef.From
		if arr.Name == nil || *arr.Name == "" {
			return hasReferences, fmt.Errorf("provided resource reference is nil or empty: AllocationRef")
		}
		obj := &svcapitypes.ElasticIPAddress{}
		if err := getReferencedResourceState_ElasticIPAddress(ctx, apiReader, obj, *arr.Name, namespace); err != nil {
			return hasReferences, err
		}
		ko.Spec.AllocationID = (*string)(obj.Status.AllocationID)
	}

	return hasReferences, nil
}

// getReferencedResourceState_ElasticIPAddress looks up whether a referenced resource
// exists and is in a ACK.ResourceSynced=True state. If the referenced resource does exist and is
// in a Synced state, returns nil, otherwise returns `ackerr.ResourceReferenceTerminalFor` or
// `ResourceReferenceNotSyncedFor` depending on if the resource is in a Terminal state.
func getReferencedResourceState_ElasticIPAddress(
	ctx context.Context,
	apiReader client.Reader,
	obj *svcapitypes.ElasticIPAddress,
	name string, // the Kubernetes name of the referenced resource
	namespace string, // the Kubernetes namespace of the referenced resource
) error {
	namespacedName := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	err := apiReader.Get(ctx, namespacedName, obj)
	if err != nil {
		return err
	}
	var refResourceSynced, refResourceTerminal bool
	for _, cond := range obj.Status.Conditions {
		if cond.Type == ackv1alpha1.ConditionTypeResourceSynced &&
			cond.Status == corev1.ConditionTrue {
			refResourceSynced = true
		}
		if cond.Type == ackv1alpha1.ConditionTypeTerminal &&
			cond.Status == corev1.ConditionTrue {
			return ackerr.ResourceReferenceTerminalFor(
				"ElasticIPAddress",
				namespace, name)
		}
	}
	if refResourceTerminal {
		return ackerr.ResourceReferenceTerminalFor(
			"ElasticIPAddress",
			namespace, name)
	}
	if !refResourceSynced {
		return ackerr.ResourceReferenceNotSyncedFor(
			"ElasticIPAddress",
			namespace, name)
	}
	if obj.Status.AllocationID == nil {
		return ackerr.ResourceReferenceMissingTargetFieldFor(
			"ElasticIPAddress",
			namespace, name,
			"Status.AllocationID")
	}
	return nil
}

// resolveReferenceForSubnetID reads the resource referenced
// from SubnetRef field and sets the SubnetID
// from referenced resource. Returns a boolean indicating whether a reference
// contains references, or an error
func (rm *resourceManager) resolveReferenceForSubnetID(
	ctx context.Context,
	apiReader client.Reader,
	namespace string,
	ko *svcapitypes.NATGateway,
) (hasReferences bool, err error) {
	if ko.Spec.SubnetRef != nil && ko.Spec.SubnetRef.From != nil {
		hasReferences = true
		arr := ko.Spec.SubnetRef.From
		if arr.Name == nil || *arr.Name == "" {
			return hasReferences, fmt.Errorf("provided resource reference is nil or empty: SubnetRef")
		}
		obj := &svcapitypes.Subnet{}
		if err := getReferencedResourceState_Subnet(ctx, apiReader, obj, *arr.Name, namespace); err != nil {
			return hasReferences, err
		}
		ko.Spec.SubnetID = (*string)(obj.Status.SubnetID)
	}

	return hasReferences, nil
}

// getReferencedResourceState_Subnet looks up whether a referenced resource
// exists and is in a ACK.ResourceSynced=True state. If the referenced resource does exist and is
// in a Synced state, returns nil, otherwise returns `ackerr.ResourceReferenceTerminalFor` or
// `ResourceReferenceNotSyncedFor` depending on if the resource is in a Terminal state.
func getReferencedResourceState_Subnet(
	ctx context.Context,
	apiReader client.Reader,
	obj *svcapitypes.Subnet,
	name string, // the Kubernetes name of the referenced resource
	namespace string, // the Kubernetes namespace of the referenced resource
) error {
	namespacedName := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	err := apiReader.Get(ctx, namespacedName, obj)
	if err != nil {
		return err
	}
	var refResourceSynced, refResourceTerminal bool
	for _, cond := range obj.Status.Conditions {
		if cond.Type == ackv1alpha1.ConditionTypeResourceSynced &&
			cond.Status == corev1.ConditionTrue {
			refResourceSynced = true
		}
		if cond.Type == ackv1alpha1.ConditionTypeTerminal &&
			cond.Status == corev1.ConditionTrue {
			return ackerr.ResourceReferenceTerminalFor(
				"Subnet",
				namespace, name)
		}
	}
	if refResourceTerminal {
		return ackerr.ResourceReferenceTerminalFor(
			"Subnet",
			namespace, name)
	}
	if !refResourceSynced {
		return ackerr.ResourceReferenceNotSyncedFor(
			"Subnet",
			namespace, name)
	}
	if obj.Status.SubnetID == nil {
		return ackerr.ResourceReferenceMissingTargetFieldFor(
			"Subnet",
			namespace, name,
			"Status.SubnetID")
	}
	return nil
}
