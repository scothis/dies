//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021-2022 the original author or authors.

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
	metav1 "dies.dev/apis/meta/v1"
	json "encoding/json"
	fmtx "fmt"
	authenticationv1 "k8s.io/api/authentication/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	jsonpath "k8s.io/client-go/util/jsonpath"
	osx "os"
	reflectx "reflect"
	yaml "sigs.k8s.io/yaml"
)

var TokenReviewBlank = (&TokenReviewDie{}).DieFeed(authenticationv1.TokenReview{})

type TokenReviewDie struct {
	metav1.FrozenObjectMeta
	mutable bool
	r       authenticationv1.TokenReview
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *TokenReviewDie) DieImmutable(immutable bool) *TokenReviewDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *TokenReviewDie) DieFeed(r authenticationv1.TokenReview) *TokenReviewDie {
	if d.mutable {
		d.FrozenObjectMeta = metav1.FreezeObjectMeta(r.ObjectMeta)
		d.r = r
		return d
	}
	return &TokenReviewDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *TokenReviewDie) DieFeedPtr(r *authenticationv1.TokenReview) *TokenReviewDie {
	if r == nil {
		r = &authenticationv1.TokenReview{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *TokenReviewDie) DieFeedJSON(j []byte) *TokenReviewDie {
	r := authenticationv1.TokenReview{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *TokenReviewDie) DieFeedYAML(y []byte) *TokenReviewDie {
	r := authenticationv1.TokenReview{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *TokenReviewDie) DieFeedYAMLFile(name string) *TokenReviewDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *TokenReviewDie) DieFeedRawExtension(raw runtime.RawExtension) *TokenReviewDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *TokenReviewDie) DieRelease() authenticationv1.TokenReview {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *TokenReviewDie) DieReleasePtr() *authenticationv1.TokenReview {
	r := d.DieRelease()
	return &r
}

// DieReleaseUnstructured returns the resource managed by the die as an unstructured object. Panics on error.
func (d *TokenReviewDie) DieReleaseUnstructured() *unstructured.Unstructured {
	r := d.DieReleasePtr()
	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(r)
	if err != nil {
		panic(err)
	}
	return &unstructured.Unstructured{
		Object: u,
	}
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *TokenReviewDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *TokenReviewDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *TokenReviewDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *TokenReviewDie) DieStamp(fn func(r *authenticationv1.TokenReview)) *TokenReviewDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type or a pointer to that type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *TokenReviewDie) DieStampAt(jp string, fn interface{}) *TokenReviewDie {
	return d.DieStamp(func(r *authenticationv1.TokenReview) {
		if ni := reflectx.ValueOf(fn).Type().NumIn(); ni != 1 {
			panic(fmtx.Errorf("callback function must have 1 input parameters, found %d", ni))
		}
		if no := reflectx.ValueOf(fn).Type().NumOut(); no != 0 {
			panic(fmtx.Errorf("callback function must have 0 output parameters, found %d", no))
		}

		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			arg0t := reflectx.ValueOf(fn).Type().In(0)

			var args []reflectx.Value
			if cv.Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv}
			} else if cv.CanAddr() && cv.Addr().Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv.Addr()}
			} else {
				panic(fmtx.Errorf("callback function must accept value of type %q, found type %q", cv.Type(), arg0t))
			}

			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DieWith returns a new die after passing the current die to the callback function. The passed die is mutable.
func (d *TokenReviewDie) DieWith(fns ...func(d *TokenReviewDie)) *TokenReviewDie {
	nd := TokenReviewBlank.DieFeed(d.DieRelease()).DieImmutable(false)
	for _, fn := range fns {
		if fn != nil {
			fn(nd)
		}
	}
	return d.DieFeed(nd.DieRelease())
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *TokenReviewDie) DeepCopy() *TokenReviewDie {
	r := *d.r.DeepCopy()
	return &TokenReviewDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

var _ runtime.Object = (*TokenReviewDie)(nil)

func (d *TokenReviewDie) DeepCopyObject() runtime.Object {
	return d.r.DeepCopy()
}

func (d *TokenReviewDie) GetObjectKind() schema.ObjectKind {
	r := d.DieRelease()
	return r.GetObjectKind()
}

func (d *TokenReviewDie) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.r)
}

func (d *TokenReviewDie) UnmarshalJSON(b []byte) error {
	if d == TokenReviewBlank {
		return fmtx.Errorf("cannot unmarshal into the blank die, create a copy first")
	}
	if !d.mutable {
		return fmtx.Errorf("cannot unmarshal into immutable dies, create a mutable version first")
	}
	r := &authenticationv1.TokenReview{}
	err := json.Unmarshal(b, r)
	*d = *d.DieFeed(*r)
	return err
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (d *TokenReviewDie) APIVersion(v string) *TokenReviewDie {
	return d.DieStamp(func(r *authenticationv1.TokenReview) {
		r.APIVersion = v
	})
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (d *TokenReviewDie) Kind(v string) *TokenReviewDie {
	return d.DieStamp(func(r *authenticationv1.TokenReview) {
		r.Kind = v
	})
}

// MetadataDie stamps the resource's ObjectMeta field with a mutable die.
func (d *TokenReviewDie) MetadataDie(fn func(d *metav1.ObjectMetaDie)) *TokenReviewDie {
	return d.DieStamp(func(r *authenticationv1.TokenReview) {
		d := metav1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

// Spec holds information about the request being evaluated
func (d *TokenReviewDie) Spec(v authenticationv1.TokenReviewSpec) *TokenReviewDie {
	return d.DieStamp(func(r *authenticationv1.TokenReview) {
		r.Spec = v
	})
}

// Status is filled in by the server and indicates whether the request can be authenticated.
func (d *TokenReviewDie) Status(v authenticationv1.TokenReviewStatus) *TokenReviewDie {
	return d.DieStamp(func(r *authenticationv1.TokenReview) {
		r.Status = v
	})
}

var TokenRequestSpecBlank = (&TokenRequestSpecDie{}).DieFeed(authenticationv1.TokenRequestSpec{})

type TokenRequestSpecDie struct {
	mutable bool
	r       authenticationv1.TokenRequestSpec
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *TokenRequestSpecDie) DieImmutable(immutable bool) *TokenRequestSpecDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *TokenRequestSpecDie) DieFeed(r authenticationv1.TokenRequestSpec) *TokenRequestSpecDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &TokenRequestSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *TokenRequestSpecDie) DieFeedPtr(r *authenticationv1.TokenRequestSpec) *TokenRequestSpecDie {
	if r == nil {
		r = &authenticationv1.TokenRequestSpec{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *TokenRequestSpecDie) DieFeedJSON(j []byte) *TokenRequestSpecDie {
	r := authenticationv1.TokenRequestSpec{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *TokenRequestSpecDie) DieFeedYAML(y []byte) *TokenRequestSpecDie {
	r := authenticationv1.TokenRequestSpec{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *TokenRequestSpecDie) DieFeedYAMLFile(name string) *TokenRequestSpecDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *TokenRequestSpecDie) DieFeedRawExtension(raw runtime.RawExtension) *TokenRequestSpecDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *TokenRequestSpecDie) DieRelease() authenticationv1.TokenRequestSpec {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *TokenRequestSpecDie) DieReleasePtr() *authenticationv1.TokenRequestSpec {
	r := d.DieRelease()
	return &r
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *TokenRequestSpecDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *TokenRequestSpecDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *TokenRequestSpecDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *TokenRequestSpecDie) DieStamp(fn func(r *authenticationv1.TokenRequestSpec)) *TokenRequestSpecDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type or a pointer to that type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *TokenRequestSpecDie) DieStampAt(jp string, fn interface{}) *TokenRequestSpecDie {
	return d.DieStamp(func(r *authenticationv1.TokenRequestSpec) {
		if ni := reflectx.ValueOf(fn).Type().NumIn(); ni != 1 {
			panic(fmtx.Errorf("callback function must have 1 input parameters, found %d", ni))
		}
		if no := reflectx.ValueOf(fn).Type().NumOut(); no != 0 {
			panic(fmtx.Errorf("callback function must have 0 output parameters, found %d", no))
		}

		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			arg0t := reflectx.ValueOf(fn).Type().In(0)

			var args []reflectx.Value
			if cv.Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv}
			} else if cv.CanAddr() && cv.Addr().Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv.Addr()}
			} else {
				panic(fmtx.Errorf("callback function must accept value of type %q, found type %q", cv.Type(), arg0t))
			}

			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DieWith returns a new die after passing the current die to the callback function. The passed die is mutable.
func (d *TokenRequestSpecDie) DieWith(fns ...func(d *TokenRequestSpecDie)) *TokenRequestSpecDie {
	nd := TokenRequestSpecBlank.DieFeed(d.DieRelease()).DieImmutable(false)
	for _, fn := range fns {
		if fn != nil {
			fn(nd)
		}
	}
	return d.DieFeed(nd.DieRelease())
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *TokenRequestSpecDie) DeepCopy() *TokenRequestSpecDie {
	r := *d.r.DeepCopy()
	return &TokenRequestSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

// Audiences are the intendend audiences of the token. A recipient of a
//
// token must identify themself with an identifier in the list of
//
// audiences of the token, and otherwise should reject the token. A
//
// token issued for multiple audiences may be used to authenticate
//
// against any of the audiences listed but implies a high degree of
//
// trust between the target audiences.
func (d *TokenRequestSpecDie) Audiences(v ...string) *TokenRequestSpecDie {
	return d.DieStamp(func(r *authenticationv1.TokenRequestSpec) {
		r.Audiences = v
	})
}

// ExpirationSeconds is the requested duration of validity of the request. The
//
// token issuer may return a token with a different validity duration so a
//
// client needs to check the 'expiration' field in a response.
func (d *TokenRequestSpecDie) ExpirationSeconds(v *int64) *TokenRequestSpecDie {
	return d.DieStamp(func(r *authenticationv1.TokenRequestSpec) {
		r.ExpirationSeconds = v
	})
}

// BoundObjectRef is a reference to an object that the token will be bound to.
//
// The token will only be valid for as long as the bound object exists.
//
// NOTE: The API server's TokenReview endpoint will validate the
//
// BoundObjectRef, but other audiences may not. Keep ExpirationSeconds
//
// small if you want prompt revocation.
func (d *TokenRequestSpecDie) BoundObjectRef(v *authenticationv1.BoundObjectReference) *TokenRequestSpecDie {
	return d.DieStamp(func(r *authenticationv1.TokenRequestSpec) {
		r.BoundObjectRef = v
	})
}

var BoundObjectReferenceBlank = (&BoundObjectReferenceDie{}).DieFeed(authenticationv1.BoundObjectReference{})

type BoundObjectReferenceDie struct {
	mutable bool
	r       authenticationv1.BoundObjectReference
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *BoundObjectReferenceDie) DieImmutable(immutable bool) *BoundObjectReferenceDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *BoundObjectReferenceDie) DieFeed(r authenticationv1.BoundObjectReference) *BoundObjectReferenceDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &BoundObjectReferenceDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *BoundObjectReferenceDie) DieFeedPtr(r *authenticationv1.BoundObjectReference) *BoundObjectReferenceDie {
	if r == nil {
		r = &authenticationv1.BoundObjectReference{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *BoundObjectReferenceDie) DieFeedJSON(j []byte) *BoundObjectReferenceDie {
	r := authenticationv1.BoundObjectReference{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *BoundObjectReferenceDie) DieFeedYAML(y []byte) *BoundObjectReferenceDie {
	r := authenticationv1.BoundObjectReference{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *BoundObjectReferenceDie) DieFeedYAMLFile(name string) *BoundObjectReferenceDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *BoundObjectReferenceDie) DieFeedRawExtension(raw runtime.RawExtension) *BoundObjectReferenceDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *BoundObjectReferenceDie) DieRelease() authenticationv1.BoundObjectReference {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *BoundObjectReferenceDie) DieReleasePtr() *authenticationv1.BoundObjectReference {
	r := d.DieRelease()
	return &r
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *BoundObjectReferenceDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *BoundObjectReferenceDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *BoundObjectReferenceDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *BoundObjectReferenceDie) DieStamp(fn func(r *authenticationv1.BoundObjectReference)) *BoundObjectReferenceDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type or a pointer to that type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *BoundObjectReferenceDie) DieStampAt(jp string, fn interface{}) *BoundObjectReferenceDie {
	return d.DieStamp(func(r *authenticationv1.BoundObjectReference) {
		if ni := reflectx.ValueOf(fn).Type().NumIn(); ni != 1 {
			panic(fmtx.Errorf("callback function must have 1 input parameters, found %d", ni))
		}
		if no := reflectx.ValueOf(fn).Type().NumOut(); no != 0 {
			panic(fmtx.Errorf("callback function must have 0 output parameters, found %d", no))
		}

		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			arg0t := reflectx.ValueOf(fn).Type().In(0)

			var args []reflectx.Value
			if cv.Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv}
			} else if cv.CanAddr() && cv.Addr().Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv.Addr()}
			} else {
				panic(fmtx.Errorf("callback function must accept value of type %q, found type %q", cv.Type(), arg0t))
			}

			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DieWith returns a new die after passing the current die to the callback function. The passed die is mutable.
func (d *BoundObjectReferenceDie) DieWith(fns ...func(d *BoundObjectReferenceDie)) *BoundObjectReferenceDie {
	nd := BoundObjectReferenceBlank.DieFeed(d.DieRelease()).DieImmutable(false)
	for _, fn := range fns {
		if fn != nil {
			fn(nd)
		}
	}
	return d.DieFeed(nd.DieRelease())
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *BoundObjectReferenceDie) DeepCopy() *BoundObjectReferenceDie {
	r := *d.r.DeepCopy()
	return &BoundObjectReferenceDie{
		mutable: d.mutable,
		r:       r,
	}
}

// Kind of the referent. Valid kinds are 'Pod' and 'Secret'.
func (d *BoundObjectReferenceDie) Kind(v string) *BoundObjectReferenceDie {
	return d.DieStamp(func(r *authenticationv1.BoundObjectReference) {
		r.Kind = v
	})
}

// API version of the referent.
func (d *BoundObjectReferenceDie) APIVersion(v string) *BoundObjectReferenceDie {
	return d.DieStamp(func(r *authenticationv1.BoundObjectReference) {
		r.APIVersion = v
	})
}

// Name of the referent.
func (d *BoundObjectReferenceDie) Name(v string) *BoundObjectReferenceDie {
	return d.DieStamp(func(r *authenticationv1.BoundObjectReference) {
		r.Name = v
	})
}

// UID of the referent.
func (d *BoundObjectReferenceDie) UID(v types.UID) *BoundObjectReferenceDie {
	return d.DieStamp(func(r *authenticationv1.BoundObjectReference) {
		r.UID = v
	})
}

var TokenRequestStatusBlank = (&TokenRequestStatusDie{}).DieFeed(authenticationv1.TokenRequestStatus{})

type TokenRequestStatusDie struct {
	mutable bool
	r       authenticationv1.TokenRequestStatus
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *TokenRequestStatusDie) DieImmutable(immutable bool) *TokenRequestStatusDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *TokenRequestStatusDie) DieFeed(r authenticationv1.TokenRequestStatus) *TokenRequestStatusDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &TokenRequestStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *TokenRequestStatusDie) DieFeedPtr(r *authenticationv1.TokenRequestStatus) *TokenRequestStatusDie {
	if r == nil {
		r = &authenticationv1.TokenRequestStatus{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *TokenRequestStatusDie) DieFeedJSON(j []byte) *TokenRequestStatusDie {
	r := authenticationv1.TokenRequestStatus{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *TokenRequestStatusDie) DieFeedYAML(y []byte) *TokenRequestStatusDie {
	r := authenticationv1.TokenRequestStatus{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *TokenRequestStatusDie) DieFeedYAMLFile(name string) *TokenRequestStatusDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *TokenRequestStatusDie) DieFeedRawExtension(raw runtime.RawExtension) *TokenRequestStatusDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *TokenRequestStatusDie) DieRelease() authenticationv1.TokenRequestStatus {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *TokenRequestStatusDie) DieReleasePtr() *authenticationv1.TokenRequestStatus {
	r := d.DieRelease()
	return &r
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *TokenRequestStatusDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *TokenRequestStatusDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *TokenRequestStatusDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *TokenRequestStatusDie) DieStamp(fn func(r *authenticationv1.TokenRequestStatus)) *TokenRequestStatusDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type or a pointer to that type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *TokenRequestStatusDie) DieStampAt(jp string, fn interface{}) *TokenRequestStatusDie {
	return d.DieStamp(func(r *authenticationv1.TokenRequestStatus) {
		if ni := reflectx.ValueOf(fn).Type().NumIn(); ni != 1 {
			panic(fmtx.Errorf("callback function must have 1 input parameters, found %d", ni))
		}
		if no := reflectx.ValueOf(fn).Type().NumOut(); no != 0 {
			panic(fmtx.Errorf("callback function must have 0 output parameters, found %d", no))
		}

		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			arg0t := reflectx.ValueOf(fn).Type().In(0)

			var args []reflectx.Value
			if cv.Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv}
			} else if cv.CanAddr() && cv.Addr().Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv.Addr()}
			} else {
				panic(fmtx.Errorf("callback function must accept value of type %q, found type %q", cv.Type(), arg0t))
			}

			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DieWith returns a new die after passing the current die to the callback function. The passed die is mutable.
func (d *TokenRequestStatusDie) DieWith(fns ...func(d *TokenRequestStatusDie)) *TokenRequestStatusDie {
	nd := TokenRequestStatusBlank.DieFeed(d.DieRelease()).DieImmutable(false)
	for _, fn := range fns {
		if fn != nil {
			fn(nd)
		}
	}
	return d.DieFeed(nd.DieRelease())
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *TokenRequestStatusDie) DeepCopy() *TokenRequestStatusDie {
	r := *d.r.DeepCopy()
	return &TokenRequestStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

// Token is the opaque bearer token.
func (d *TokenRequestStatusDie) Token(v string) *TokenRequestStatusDie {
	return d.DieStamp(func(r *authenticationv1.TokenRequestStatus) {
		r.Token = v
	})
}

// ExpirationTimestamp is the time of expiration of the returned token.
func (d *TokenRequestStatusDie) ExpirationTimestamp(v apismetav1.Time) *TokenRequestStatusDie {
	return d.DieStamp(func(r *authenticationv1.TokenRequestStatus) {
		r.ExpirationTimestamp = v
	})
}

var UserInfoBlank = (&UserInfoDie{}).DieFeed(authenticationv1.UserInfo{})

type UserInfoDie struct {
	mutable bool
	r       authenticationv1.UserInfo
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *UserInfoDie) DieImmutable(immutable bool) *UserInfoDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *UserInfoDie) DieFeed(r authenticationv1.UserInfo) *UserInfoDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &UserInfoDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *UserInfoDie) DieFeedPtr(r *authenticationv1.UserInfo) *UserInfoDie {
	if r == nil {
		r = &authenticationv1.UserInfo{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *UserInfoDie) DieFeedJSON(j []byte) *UserInfoDie {
	r := authenticationv1.UserInfo{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *UserInfoDie) DieFeedYAML(y []byte) *UserInfoDie {
	r := authenticationv1.UserInfo{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *UserInfoDie) DieFeedYAMLFile(name string) *UserInfoDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *UserInfoDie) DieFeedRawExtension(raw runtime.RawExtension) *UserInfoDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *UserInfoDie) DieRelease() authenticationv1.UserInfo {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *UserInfoDie) DieReleasePtr() *authenticationv1.UserInfo {
	r := d.DieRelease()
	return &r
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *UserInfoDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *UserInfoDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *UserInfoDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *UserInfoDie) DieStamp(fn func(r *authenticationv1.UserInfo)) *UserInfoDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type or a pointer to that type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *UserInfoDie) DieStampAt(jp string, fn interface{}) *UserInfoDie {
	return d.DieStamp(func(r *authenticationv1.UserInfo) {
		if ni := reflectx.ValueOf(fn).Type().NumIn(); ni != 1 {
			panic(fmtx.Errorf("callback function must have 1 input parameters, found %d", ni))
		}
		if no := reflectx.ValueOf(fn).Type().NumOut(); no != 0 {
			panic(fmtx.Errorf("callback function must have 0 output parameters, found %d", no))
		}

		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			arg0t := reflectx.ValueOf(fn).Type().In(0)

			var args []reflectx.Value
			if cv.Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv}
			} else if cv.CanAddr() && cv.Addr().Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv.Addr()}
			} else {
				panic(fmtx.Errorf("callback function must accept value of type %q, found type %q", cv.Type(), arg0t))
			}

			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DieWith returns a new die after passing the current die to the callback function. The passed die is mutable.
func (d *UserInfoDie) DieWith(fns ...func(d *UserInfoDie)) *UserInfoDie {
	nd := UserInfoBlank.DieFeed(d.DieRelease()).DieImmutable(false)
	for _, fn := range fns {
		if fn != nil {
			fn(nd)
		}
	}
	return d.DieFeed(nd.DieRelease())
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *UserInfoDie) DeepCopy() *UserInfoDie {
	r := *d.r.DeepCopy()
	return &UserInfoDie{
		mutable: d.mutable,
		r:       r,
	}
}

// The name that uniquely identifies this user among all active users.
func (d *UserInfoDie) Username(v string) *UserInfoDie {
	return d.DieStamp(func(r *authenticationv1.UserInfo) {
		r.Username = v
	})
}

// A unique value that identifies this user across time. If this user is
//
// deleted and another user by the same name is added, they will have
//
// different UIDs.
func (d *UserInfoDie) UID(v string) *UserInfoDie {
	return d.DieStamp(func(r *authenticationv1.UserInfo) {
		r.UID = v
	})
}

// The names of groups this user is a part of.
func (d *UserInfoDie) Groups(v ...string) *UserInfoDie {
	return d.DieStamp(func(r *authenticationv1.UserInfo) {
		r.Groups = v
	})
}
