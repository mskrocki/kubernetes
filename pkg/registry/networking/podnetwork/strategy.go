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

package podnetwork

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/networking"
	"k8s.io/kubernetes/pkg/apis/networking/validation"
)

// podNetworkStrategy implements verification logic for PodNetworks.
type podNetworkStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

// Strategy is the default logic that applies when creating and updating PodNetwork objects.
var Strategy = podNetworkStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

// NamespaceScoped returns false because all podNetworks do not need to be within a namespace.
func (podNetworkStrategy) NamespaceScoped() bool {
	return false
}

func (podNetworkStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {}

func (podNetworkStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {}

// Validate validates a new PodNetworks.
func (podNetworkStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	podNetwork := obj.(*networking.PodNetwork)
	return validation.ValidatePodNetwork(podNetwork)
}

// WarningsOnCreate returns warnings for the creation of the given object.
func (podNetworkStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

// Canonicalize normalizes the object after validation.
func (podNetworkStrategy) Canonicalize(obj runtime.Object) {}

// AllowCreateOnUpdate is false for PodNetwork; this means POST is needed to create one.
func (podNetworkStrategy) AllowCreateOnUpdate() bool {
	return false
}

// ValidateUpdate is the default update validation for an end user.
func (podNetworkStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	validationErrorList := validation.ValidatePodNetwork(obj.(*networking.PodNetwork))
	updateErrorList := validation.ValidatePodNetworkUpdate(obj.(*networking.PodNetwork), old.(*networking.PodNetwork))
	return append(validationErrorList, updateErrorList...)
}

// WarningsOnUpdate returns warnings for the given update.
func (podNetworkStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}

// AllowUnconditionalUpdate is the default update policy for PodNetwork objects.
func (podNetworkStrategy) AllowUnconditionalUpdate() bool {
	return true
}
