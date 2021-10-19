/*
Copyright 2021 the original author or authors.

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

package v1

import (
	diemetav1 "dies.dev/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
)

// +die:object=true
type PersistentVolumeClaim = corev1.PersistentVolumeClaim

// +die
type PersistentVolumeClaimSpec = corev1.PersistentVolumeClaimSpec

// +die
type PersistentVolumeClaimStatus = corev1.PersistentVolumeClaimStatus

func (d *PersistentVolumeClaimStatusDie) ConditionsDie(conditions ...*diemetav1.ConditionDie) *PersistentVolumeClaimStatusDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeClaimStatus) {
		r.Conditions = make([]corev1.PersistentVolumeClaimCondition, len(conditions))
		for i := range conditions {
			c := conditions[i].DieRelease()
			r.Conditions[i] = corev1.PersistentVolumeClaimCondition{
				Type:               corev1.PersistentVolumeClaimConditionType(c.Type),
				Status:             corev1.ConditionStatus(c.Status),
				Reason:             c.Reason,
				Message:            c.Message,
				LastTransitionTime: c.LastTransitionTime,
			}
		}
	})
}