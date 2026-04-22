/*
Copyright 2026 the original author or authors.

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

package v1alpha2

import (
	schedulingv1alpha2 "k8s.io/api/scheduling/v1alpha2"
)

// +die:object=true,apiVersion=scheduling.k8s.io/v1alpha2,kind=Workload
type _ = schedulingv1alpha2.Workload

// +die
// +die:field:name=ControllerRef,die=TypedLocalObjectReferenceDie,pointer=true
// +die:field:name=PodGroupTemplates,die=PodGroupTemplateDie,listType=map
type _ = schedulingv1alpha2.WorkloadSpec

// +die
type _ = schedulingv1alpha2.TypedLocalObjectReference

// +die
// +die:field:name=SchedulingPolicy,die=PodGroupSchedulingPolicyDie
// +die:field:name=SchedulingConstraints,die=PodGroupSchedulingConstraintsDie,pointer=true
// +die:field:name=ResourceClaims,die=PodGroupResourceClaimDie,listType=map
type _ = schedulingv1alpha2.PodGroupTemplate
