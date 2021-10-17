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
	corev1 "k8s.io/api/core/v1"
)

// +die:object=true
type PersistentVolume = corev1.PersistentVolume

// +die
type PersistentVolumeSpec = corev1.PersistentVolumeSpec

func (d *PersistentVolumeSpecDie) GCEPersistentDiskDie(fn func(d *GCEPersistentDiskVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := GCEPersistentDiskVolumeSourceBlank.DieImmutable(false)
		if v := r.GCEPersistentDisk; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			GCEPersistentDisk: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) AWSElasticBlockStoreDie(fn func(d *AWSElasticBlockStoreVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := AWSElasticBlockStoreVolumeSourceBlank.DieImmutable(false)
		if v := r.AWSElasticBlockStore; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			AWSElasticBlockStore: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) HostPathDie(fn func(d *HostPathVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := HostPathVolumeSourceBlank.DieImmutable(false)
		if v := r.HostPath; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			HostPath: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) GlusterfsDie(fn func(d *GlusterfsPersistentVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := GlusterfsPersistentVolumeSourceBlank.DieImmutable(false)
		if v := r.Glusterfs; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			Glusterfs: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) NFSDie(fn func(d *NFSVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := NFSVolumeSourceBlank.DieImmutable(false)
		if v := r.NFS; v != nil {
			d.DieFeed(*r.NFS)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			NFS: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) RBDDie(fn func(d *RBDPersistentVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := RBDPersistentVolumeSourceBlank.DieImmutable(false)
		if v := r.RBD; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			RBD: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) ISCSIDie(fn func(d *ISCSIPersistentVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := ISCSIPersistentVolumeSourceBlank.DieImmutable(false)
		if v := r.ISCSI; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			ISCSI: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) CinderDie(fn func(d *CinderPersistentVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := CinderPersistentVolumeSourceBlank.DieImmutable(false)
		if v := r.Cinder; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			Cinder: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) CephFSDie(fn func(d *CephFSPersistentVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := CephFSPersistentVolumeSourceBlank.DieImmutable(false)
		if v := r.CephFS; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			CephFS: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) FCDie(fn func(d *FCVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := FCVolumeSourceBlank.DieImmutable(false)
		if v := r.FC; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			FC: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) FlockerDie(fn func(d *FlockerVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := FlockerVolumeSourceBlank.DieImmutable(false)
		if v := r.Flocker; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			Flocker: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) FlexVolumeDie(fn func(d *FlexPersistentVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := FlexPersistentVolumeSourceBlank.DieImmutable(false)
		if v := r.FlexVolume; v != nil {
			d.DieFeed(*r.FlexVolume)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			FlexVolume: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) AzureFileDie(fn func(d *AzureFilePersistentVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := AzureFilePersistentVolumeSourceBlank.DieImmutable(false)
		if v := r.AzureFile; v != nil {
			d.DieFeed(*r.AzureFile)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			AzureFile: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) VsphereVolumeDie(fn func(d *VsphereVirtualDiskVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := VsphereVirtualDiskVolumeSourceBlank.DieImmutable(false)
		if v := r.VsphereVolume; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			VsphereVolume: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) QuobyteDie(fn func(d *QuobyteVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := QuobyteVolumeSourceBlank.DieImmutable(false)
		if v := r.Quobyte; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			Quobyte: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) AzureDiskDie(fn func(d *AzureDiskVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := AzureDiskVolumeSourceBlank.DieImmutable(false)
		if v := r.AzureDisk; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			AzureDisk: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) PhotonPersistentDiskDie(fn func(d *PhotonPersistentDiskVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := PhotonPersistentDiskVolumeSourceBlank.DieImmutable(false)
		if v := r.PhotonPersistentDisk; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			PhotonPersistentDisk: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) PortworxVolumeDie(fn func(d *PortworxVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := PortworxVolumeSourceBlank.DieImmutable(false)
		if v := r.PortworxVolume; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			PortworxVolume: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) ScaleIODie(fn func(d *ScaleIOPersistentVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := ScaleIOPersistentVolumeSourceBlank.DieImmutable(false)
		if v := r.ScaleIO; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			ScaleIO: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) LocalDie(fn func(d *LocalVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := LocalVolumeSourceBlank.DieImmutable(false)
		if v := r.Local; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			Local: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) StorageOSDie(fn func(d *StorageOSPersistentVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := StorageOSPersistentVolumeSourceBlank.DieImmutable(false)
		if v := r.StorageOS; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			StorageOS: &v,
		}
	})
}

func (d *PersistentVolumeSpecDie) CSIDie(fn func(d *CSIPersistentVolumeSourceDie)) *PersistentVolumeSpecDie {
	return d.DieStamp(func(r *corev1.PersistentVolumeSpec) {
		d := CSIPersistentVolumeSourceBlank.DieImmutable(false)
		if v := r.CSI; v != nil {
			d.DieFeed(*v)
		}
		fn(d)
		v := d.DieRelease()
		r.PersistentVolumeSource = corev1.PersistentVolumeSource{
			CSI: &v,
		}
	})
}

// +die
type PersistentVolumeStatus = corev1.PersistentVolumeStatus

// +die
type GlusterfsPersistentVolumeSource = corev1.GlusterfsPersistentVolumeSource

// +die
type RBDPersistentVolumeSource = corev1.RBDPersistentVolumeSource

// +die
type ISCSIPersistentVolumeSource = corev1.ISCSIPersistentVolumeSource

// +die
type CinderPersistentVolumeSource = corev1.CinderPersistentVolumeSource

// +die
type CephFSPersistentVolumeSource = corev1.CephFSPersistentVolumeSource

// +die
type FlexPersistentVolumeSource = corev1.FlexPersistentVolumeSource

// +die
type AzureFilePersistentVolumeSource = corev1.AzureFilePersistentVolumeSource

// +die
type ScaleIOPersistentVolumeSource = corev1.ScaleIOPersistentVolumeSource

// +die
type LocalVolumeSource = corev1.LocalVolumeSource

// +die
type StorageOSPersistentVolumeSource = corev1.StorageOSPersistentVolumeSource

// +die
type CSIPersistentVolumeSource = corev1.CSIPersistentVolumeSource
