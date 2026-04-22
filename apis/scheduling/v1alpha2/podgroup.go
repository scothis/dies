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

// +die:object=true,apiVersion=scheduling.k8s.io/v1alpha2,kind=PodGroup
type _ = schedulingv1alpha2.PodGroup

// +die
// +die:field:name=PodGroupTemplateRef,die=PodGroupTemplateReferenceDie,pointer=true
// +die:field:name=SchedulingPolicy,die=PodGroupSchedulingPolicyDie
// +die:field:name=SchedulingConstraints,die=PodGroupSchedulingConstraintsDie,pointer=true
// +die:field:name=ResourceClaims,die=PodGroupResourceClaimDie,listType=map
type _ = schedulingv1alpha2.PodGroupSpec

// +die
// +die:field:name=Workload,die=WorkloadPodGroupTemplateReferenceDie,pointer=true
type _ = schedulingv1alpha2.PodGroupTemplateReference

// +die
type _ = schedulingv1alpha2.WorkloadPodGroupTemplateReference

// +die
// +die:field:name=Basic,die=BasicSchedulingPolicyDie,pointer=true
// +die:field:name=Gang,die=GangSchedulingPolicyDie,pointer=true
type _ = schedulingv1alpha2.PodGroupSchedulingPolicy

// +die
type _ = schedulingv1alpha2.BasicSchedulingPolicy

// +die
type _ = schedulingv1alpha2.GangSchedulingPolicy

// +die
// +die:field:name=Topology,die=TopologyConstraintDie,listType=atomic
type _ = schedulingv1alpha2.PodGroupSchedulingConstraints

// +die
type _ = schedulingv1alpha2.TopologyConstraint

// +die
type _ = schedulingv1alpha2.PodGroupResourceClaim

// +die
// +die:field:name=Conditions,package=_/meta/v1,die=ConditionDie,listType=map,listMapKey=Type
// +die:field:name=ResourceClaimStatuses,die=PodGroupResourceClaimStatusDie,listType=map
type _ = schedulingv1alpha2.PodGroupStatus

// +die
type _ = schedulingv1alpha2.PodGroupResourceClaimStatus
