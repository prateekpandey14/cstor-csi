// +build !ignore_autogenerated

/*
Copyright 2019 The OpenEBS Authors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIVolume) DeepCopyInto(out *CSIVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIVolume.
func (in *CSIVolume) DeepCopy() *CSIVolume {
	if in == nil {
		return nil
	}
	out := new(CSIVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CSIVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIVolumeList) DeepCopyInto(out *CSIVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CSIVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIVolumeList.
func (in *CSIVolumeList) DeepCopy() *CSIVolumeList {
	if in == nil {
		return nil
	}
	out := new(CSIVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CSIVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIVolumeSpec) DeepCopyInto(out *CSIVolumeSpec) {
	*out = *in
	in.Volume.DeepCopyInto(&out.Volume)
	out.ISCSI = in.ISCSI
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIVolumeSpec.
func (in *CSIVolumeSpec) DeepCopy() *CSIVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(CSIVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ISCSIInfo) DeepCopyInto(out *ISCSIInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ISCSIInfo.
func (in *ISCSIInfo) DeepCopy() *ISCSIInfo {
	if in == nil {
		return nil
	}
	out := new(ISCSIInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeInfo) DeepCopyInto(out *VolumeInfo) {
	*out = *in
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MountOptions != nil {
		in, out := &in.MountOptions, &out.MountOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeInfo.
func (in *VolumeInfo) DeepCopy() *VolumeInfo {
	if in == nil {
		return nil
	}
	out := new(VolumeInfo)
	in.DeepCopyInto(out)
	return out
}
