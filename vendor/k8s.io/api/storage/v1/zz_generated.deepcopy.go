// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	core_v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AttachedVolumeSource).DeepCopyInto(out.(*AttachedVolumeSource))
			return nil
		}, InType: reflect.TypeOf(&AttachedVolumeSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StorageClass).DeepCopyInto(out.(*StorageClass))
			return nil
		}, InType: reflect.TypeOf(&StorageClass{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StorageClassList).DeepCopyInto(out.(*StorageClassList))
			return nil
		}, InType: reflect.TypeOf(&StorageClassList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*VolumeAttachment).DeepCopyInto(out.(*VolumeAttachment))
			return nil
		}, InType: reflect.TypeOf(&VolumeAttachment{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*VolumeAttachmentList).DeepCopyInto(out.(*VolumeAttachmentList))
			return nil
		}, InType: reflect.TypeOf(&VolumeAttachmentList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*VolumeAttachmentSpec).DeepCopyInto(out.(*VolumeAttachmentSpec))
			return nil
		}, InType: reflect.TypeOf(&VolumeAttachmentSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*VolumeAttachmentStatus).DeepCopyInto(out.(*VolumeAttachmentStatus))
			return nil
		}, InType: reflect.TypeOf(&VolumeAttachmentStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*VolumeError).DeepCopyInto(out.(*VolumeError))
			return nil
		}, InType: reflect.TypeOf(&VolumeError{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachedVolumeSource) DeepCopyInto(out *AttachedVolumeSource) {
	*out = *in
	if in.PersistentVolumeName != nil {
		in, out := &in.PersistentVolumeName, &out.PersistentVolumeName
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachedVolumeSource.
func (in *AttachedVolumeSource) DeepCopy() *AttachedVolumeSource {
	if in == nil {
		return nil
	}
	out := new(AttachedVolumeSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClass) DeepCopyInto(out *StorageClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReclaimPolicy != nil {
		in, out := &in.ReclaimPolicy, &out.ReclaimPolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PersistentVolumeReclaimPolicy)
			**out = **in
		}
	}
	if in.MountOptions != nil {
		in, out := &in.MountOptions, &out.MountOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowVolumeExpansion != nil {
		in, out := &in.AllowVolumeExpansion, &out.AllowVolumeExpansion
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClass.
func (in *StorageClass) DeepCopy() *StorageClass {
	if in == nil {
		return nil
	}
	out := new(StorageClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClassList) DeepCopyInto(out *StorageClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClassList.
func (in *StorageClassList) DeepCopy() *StorageClassList {
	if in == nil {
		return nil
	}
	out := new(StorageClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeAttachment) DeepCopyInto(out *VolumeAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeAttachment.
func (in *VolumeAttachment) DeepCopy() *VolumeAttachment {
	if in == nil {
		return nil
	}
	out := new(VolumeAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeAttachmentList) DeepCopyInto(out *VolumeAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VolumeAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeAttachmentList.
func (in *VolumeAttachmentList) DeepCopy() *VolumeAttachmentList {
	if in == nil {
		return nil
	}
	out := new(VolumeAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeAttachmentSpec) DeepCopyInto(out *VolumeAttachmentSpec) {
	*out = *in
	in.AttachedVolumeSource.DeepCopyInto(&out.AttachedVolumeSource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeAttachmentSpec.
func (in *VolumeAttachmentSpec) DeepCopy() *VolumeAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeAttachmentStatus) DeepCopyInto(out *VolumeAttachmentStatus) {
	*out = *in
	if in.AttachmentMetadata != nil {
		in, out := &in.AttachmentMetadata, &out.AttachmentMetadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AttachError != nil {
		in, out := &in.AttachError, &out.AttachError
		if *in == nil {
			*out = nil
		} else {
			*out = new(VolumeError)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.DetachError != nil {
		in, out := &in.DetachError, &out.DetachError
		if *in == nil {
			*out = nil
		} else {
			*out = new(VolumeError)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeAttachmentStatus.
func (in *VolumeAttachmentStatus) DeepCopy() *VolumeAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(VolumeAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeError) DeepCopyInto(out *VolumeError) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeError.
func (in *VolumeError) DeepCopy() *VolumeError {
	if in == nil {
		return nil
	}
	out := new(VolumeError)
	in.DeepCopyInto(out)
	return out
}
