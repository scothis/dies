//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by diegen. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	fmtx "fmt"
	metav1 "github.com/scothis/dies/apis/meta/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

type HorizontalPodAutoscalerDie struct {
	metav1.FrozenObjectMeta
	mutable bool
	r       autoscalingv1.HorizontalPodAutoscaler
}

var HorizontalPodAutoscalerBlank = (&HorizontalPodAutoscalerDie{}).DieFeed(autoscalingv1.HorizontalPodAutoscaler{})

func (d *HorizontalPodAutoscalerDie) DieImmutable(immutable bool) *HorizontalPodAutoscalerDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *HorizontalPodAutoscalerDie) DieFeed(r autoscalingv1.HorizontalPodAutoscaler) *HorizontalPodAutoscalerDie {
	if d.mutable {
		d.FrozenObjectMeta = metav1.FreezeObjectMeta(r.ObjectMeta)
		d.r = r
		return d
	}
	return &HorizontalPodAutoscalerDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

func (d *HorizontalPodAutoscalerDie) DieRelease() autoscalingv1.HorizontalPodAutoscaler {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *HorizontalPodAutoscalerDie) DieStamp(fn func(r *autoscalingv1.HorizontalPodAutoscaler)) *HorizontalPodAutoscalerDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *HorizontalPodAutoscalerDie) DeepCopy() *HorizontalPodAutoscalerDie {
	r := *d.r.DeepCopy()
	return &HorizontalPodAutoscalerDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

func (d *HorizontalPodAutoscalerDie) DeepCopyObject() runtime.Object {
	return d.r.DeepCopy()
}

func (d *HorizontalPodAutoscalerDie) GetObjectKind() schema.ObjectKind {
	r := d.DieRelease()
	return r.GetObjectKind()
}

func (d *HorizontalPodAutoscalerDie) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.r)
}

func (d *HorizontalPodAutoscalerDie) UnmarshalJSON(b []byte) error {
	if d == HorizontalPodAutoscalerBlank {
		return fmtx.Errorf("cannot unmarshal into the root object, create a copy first")
	}
	r := &autoscalingv1.HorizontalPodAutoscaler{}
	err := json.Unmarshal(b, r)
	*d = *d.DieFeed(*r)
	return err
}

func (d *HorizontalPodAutoscalerDie) MetadataDie(fn func(d *metav1.ObjectMetaDie)) *HorizontalPodAutoscalerDie {
	return d.DieStamp(func(r *autoscalingv1.HorizontalPodAutoscaler) {
		d := metav1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

func (d *HorizontalPodAutoscalerDie) SpecDie(fn func(d *HorizontalPodAutoscalerSpecDie)) *HorizontalPodAutoscalerDie {
	return d.DieStamp(func(r *autoscalingv1.HorizontalPodAutoscaler) {
		d := HorizontalPodAutoscalerSpecBlank.DieImmutable(false).DieFeed(r.Spec)
		fn(d)
		r.Spec = d.DieRelease()
	})
}

func (d *HorizontalPodAutoscalerDie) StatusDie(fn func(d *HorizontalPodAutoscalerStatusDie)) *HorizontalPodAutoscalerDie {
	return d.DieStamp(func(r *autoscalingv1.HorizontalPodAutoscaler) {
		d := HorizontalPodAutoscalerStatusBlank.DieImmutable(false).DieFeed(r.Status)
		fn(d)
		r.Status = d.DieRelease()
	})
}

var _ apismetav1.Object = (*HorizontalPodAutoscalerDie)(nil)
var _ apismetav1.ObjectMetaAccessor = (*HorizontalPodAutoscalerDie)(nil)
var _ runtime.Object = (*HorizontalPodAutoscalerDie)(nil)

// behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
func (d *HorizontalPodAutoscalerDie) Spec(v autoscalingv1.HorizontalPodAutoscalerSpec) *HorizontalPodAutoscalerDie {
	return d.DieStamp(func(r *autoscalingv1.HorizontalPodAutoscaler) {
		r.Spec = v
	})
}

// current information about the autoscaler.
func (d *HorizontalPodAutoscalerDie) Status(v autoscalingv1.HorizontalPodAutoscalerStatus) *HorizontalPodAutoscalerDie {
	return d.DieStamp(func(r *autoscalingv1.HorizontalPodAutoscaler) {
		r.Status = v
	})
}

type HorizontalPodAutoscalerSpecDie struct {
	mutable bool
	r       autoscalingv1.HorizontalPodAutoscalerSpec
}

var HorizontalPodAutoscalerSpecBlank = (&HorizontalPodAutoscalerSpecDie{}).DieFeed(autoscalingv1.HorizontalPodAutoscalerSpec{})

func (d *HorizontalPodAutoscalerSpecDie) DieImmutable(immutable bool) *HorizontalPodAutoscalerSpecDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *HorizontalPodAutoscalerSpecDie) DieFeed(r autoscalingv1.HorizontalPodAutoscalerSpec) *HorizontalPodAutoscalerSpecDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &HorizontalPodAutoscalerSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *HorizontalPodAutoscalerSpecDie) DieRelease() autoscalingv1.HorizontalPodAutoscalerSpec {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *HorizontalPodAutoscalerSpecDie) DieStamp(fn func(r *autoscalingv1.HorizontalPodAutoscalerSpec)) *HorizontalPodAutoscalerSpecDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *HorizontalPodAutoscalerSpecDie) DeepCopy() *HorizontalPodAutoscalerSpecDie {
	r := *d.r.DeepCopy()
	return &HorizontalPodAutoscalerSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

// reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
func (d *HorizontalPodAutoscalerSpecDie) ScaleTargetRef(v autoscalingv1.CrossVersionObjectReference) *HorizontalPodAutoscalerSpecDie {
	return d.DieStamp(func(r *autoscalingv1.HorizontalPodAutoscalerSpec) {
		r.ScaleTargetRef = v
	})
}

// minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
func (d *HorizontalPodAutoscalerSpecDie) MinReplicas(v *int32) *HorizontalPodAutoscalerSpecDie {
	return d.DieStamp(func(r *autoscalingv1.HorizontalPodAutoscalerSpec) {
		r.MinReplicas = v
	})
}

// upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
func (d *HorizontalPodAutoscalerSpecDie) MaxReplicas(v int32) *HorizontalPodAutoscalerSpecDie {
	return d.DieStamp(func(r *autoscalingv1.HorizontalPodAutoscalerSpec) {
		r.MaxReplicas = v
	})
}

// target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
func (d *HorizontalPodAutoscalerSpecDie) TargetCPUUtilizationPercentage(v *int32) *HorizontalPodAutoscalerSpecDie {
	return d.DieStamp(func(r *autoscalingv1.HorizontalPodAutoscalerSpec) {
		r.TargetCPUUtilizationPercentage = v
	})
}

type HorizontalPodAutoscalerStatusDie struct {
	mutable bool
	r       autoscalingv1.HorizontalPodAutoscalerStatus
}

var HorizontalPodAutoscalerStatusBlank = (&HorizontalPodAutoscalerStatusDie{}).DieFeed(autoscalingv1.HorizontalPodAutoscalerStatus{})

func (d *HorizontalPodAutoscalerStatusDie) DieImmutable(immutable bool) *HorizontalPodAutoscalerStatusDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

func (d *HorizontalPodAutoscalerStatusDie) DieFeed(r autoscalingv1.HorizontalPodAutoscalerStatus) *HorizontalPodAutoscalerStatusDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &HorizontalPodAutoscalerStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *HorizontalPodAutoscalerStatusDie) DieRelease() autoscalingv1.HorizontalPodAutoscalerStatus {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

func (d *HorizontalPodAutoscalerStatusDie) DieStamp(fn func(r *autoscalingv1.HorizontalPodAutoscalerStatus)) *HorizontalPodAutoscalerStatusDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

func (d *HorizontalPodAutoscalerStatusDie) DeepCopy() *HorizontalPodAutoscalerStatusDie {
	r := *d.r.DeepCopy()
	return &HorizontalPodAutoscalerStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

// most recent generation observed by this autoscaler.
func (d *HorizontalPodAutoscalerStatusDie) ObservedGeneration(v *int64) *HorizontalPodAutoscalerStatusDie {
	return d.DieStamp(func(r *autoscalingv1.HorizontalPodAutoscalerStatus) {
		r.ObservedGeneration = v
	})
}

// last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
func (d *HorizontalPodAutoscalerStatusDie) LastScaleTime(v *apismetav1.Time) *HorizontalPodAutoscalerStatusDie {
	return d.DieStamp(func(r *autoscalingv1.HorizontalPodAutoscalerStatus) {
		r.LastScaleTime = v
	})
}

// current number of replicas of pods managed by this autoscaler.
func (d *HorizontalPodAutoscalerStatusDie) CurrentReplicas(v int32) *HorizontalPodAutoscalerStatusDie {
	return d.DieStamp(func(r *autoscalingv1.HorizontalPodAutoscalerStatus) {
		r.CurrentReplicas = v
	})
}

// desired number of replicas of pods managed by this autoscaler.
func (d *HorizontalPodAutoscalerStatusDie) DesiredReplicas(v int32) *HorizontalPodAutoscalerStatusDie {
	return d.DieStamp(func(r *autoscalingv1.HorizontalPodAutoscalerStatus) {
		r.DesiredReplicas = v
	})
}

// current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
func (d *HorizontalPodAutoscalerStatusDie) CurrentCPUUtilizationPercentage(v *int32) *HorizontalPodAutoscalerStatusDie {
	return d.DieStamp(func(r *autoscalingv1.HorizontalPodAutoscalerStatus) {
		r.CurrentCPUUtilizationPercentage = v
	})
}
