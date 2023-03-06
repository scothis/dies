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

package sandbox

import (
	json "encoding/json"
	runtime "k8s.io/apimachinery/pkg/runtime"
	osx "os"
	yaml "sigs.k8s.io/yaml"
)

var DirectBlank = (&DirectDie{}).DieFeed(Direct{})

type DirectDie struct {
	mutable bool
	r       Direct
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *DirectDie) DieImmutable(immutable bool) *DirectDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *DirectDie) DieFeed(r Direct) *DirectDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &DirectDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *DirectDie) DieFeedPtr(r *Direct) *DirectDie {
	if r == nil {
		r = &Direct{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *DirectDie) DieFeedJSON(j []byte) *DirectDie {
	r := Direct{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *DirectDie) DieFeedYAML(y []byte) *DirectDie {
	r := Direct{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *DirectDie) DieFeedYAMLFile(name string) *DirectDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *DirectDie) DieFeedRawExtension(raw runtime.RawExtension) *DirectDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *DirectDie) DieRelease() Direct {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *DirectDie) DieReleasePtr() *Direct {
	r := d.DieRelease()
	return &r
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *DirectDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *DirectDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *DirectDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *DirectDie) DieStamp(fn func(r *Direct)) *DirectDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *DirectDie) DeepCopy() *DirectDie {
	r := *d.r.DeepCopy()
	return &DirectDie{
		mutable: d.mutable,
		r:       r,
	}
}

func (d *DirectDie) Hello(v string) *DirectDie {
	return d.DieStamp(func(r *Direct) {
		r.Hello = v
	})
}

func (d *DirectDie) World(v string) *DirectDie {
	return d.DieStamp(func(r *Direct) {
		r.World = v
	})
}
