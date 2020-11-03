// +build !ignore_autogenerated

/*


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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostNetwork) DeepCopyInto(out *HostNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostNetwork.
func (in *HostNetwork) DeepCopy() *HostNetwork {
	if in == nil {
		return nil
	}
	out := new(HostNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostNetworkList) DeepCopyInto(out *HostNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HostNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostNetworkList.
func (in *HostNetworkList) DeepCopy() *HostNetworkList {
	if in == nil {
		return nil
	}
	out := new(HostNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostNetworkSpec) DeepCopyInto(out *HostNetworkSpec) {
	*out = *in
	if in.InterfaceStatus != nil {
		in, out := &in.InterfaceStatus, &out.InterfaceStatus
		*out = make([]*InterfaceStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(InterfaceStatus)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Sysctl != nil {
		in, out := &in.Sysctl, &out.Sysctl
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostNetworkSpec.
func (in *HostNetworkSpec) DeepCopy() *HostNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(HostNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostNetworkStatus) DeepCopyInto(out *HostNetworkStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostNetworkStatus.
func (in *HostNetworkStatus) DeepCopy() *HostNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(HostNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostNetworkTemplate) DeepCopyInto(out *HostNetworkTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostNetworkTemplate.
func (in *HostNetworkTemplate) DeepCopy() *HostNetworkTemplate {
	if in == nil {
		return nil
	}
	out := new(HostNetworkTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostNetworkTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostNetworkTemplateList) DeepCopyInto(out *HostNetworkTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HostNetworkTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostNetworkTemplateList.
func (in *HostNetworkTemplateList) DeepCopy() *HostNetworkTemplateList {
	if in == nil {
		return nil
	}
	out := new(HostNetworkTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostNetworkTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostNetworkTemplateSpec) DeepCopyInto(out *HostNetworkTemplateSpec) {
	*out = *in
	if in.SriovConfig != nil {
		in, out := &in.SriovConfig, &out.SriovConfig
		*out = make([]SriovConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostNetworkTemplateSpec.
func (in *HostNetworkTemplateSpec) DeepCopy() *HostNetworkTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(HostNetworkTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostNetworkTemplateStatus) DeepCopyInto(out *HostNetworkTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostNetworkTemplateStatus.
func (in *HostNetworkTemplateStatus) DeepCopy() *HostNetworkTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(HostNetworkTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceStatus) DeepCopyInto(out *InterfaceStatus) {
	*out = *in
	if in.SriovStatus != nil {
		in, out := &in.SriovStatus, &out.SriovStatus
		*out = new(SriovStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceStatus.
func (in *InterfaceStatus) DeepCopy() *InterfaceStatus {
	if in == nil {
		return nil
	}
	out := new(InterfaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovConfig) DeepCopyInto(out *SriovConfig) {
	*out = *in
	if in.PfName != nil {
		in, out := &in.PfName, &out.PfName
		*out = new(string)
		**out = **in
	}
	if in.PciAddr != nil {
		in, out := &in.PciAddr, &out.PciAddr
		*out = new(string)
		**out = **in
	}
	if in.VendorId != nil {
		in, out := &in.VendorId, &out.VendorId
		*out = new(string)
		**out = **in
	}
	if in.DeviceId != nil {
		in, out := &in.DeviceId, &out.DeviceId
		*out = new(string)
		**out = **in
	}
	if in.NumVfs != nil {
		in, out := &in.NumVfs, &out.NumVfs
		*out = new(int)
		**out = **in
	}
	if in.MTU != nil {
		in, out := &in.MTU, &out.MTU
		*out = new(int)
		**out = **in
	}
	if in.VfDriver != nil {
		in, out := &in.VfDriver, &out.VfDriver
		*out = new(string)
		**out = **in
	}
	if in.PfDriver != nil {
		in, out := &in.PfDriver, &out.PfDriver
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovConfig.
func (in *SriovConfig) DeepCopy() *SriovConfig {
	if in == nil {
		return nil
	}
	out := new(SriovConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovStatus) DeepCopyInto(out *SriovStatus) {
	*out = *in
	if in.Vfs != nil {
		in, out := &in.Vfs, &out.Vfs
		*out = make([]*VfInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(VfInfo)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovStatus.
func (in *SriovStatus) DeepCopy() *SriovStatus {
	if in == nil {
		return nil
	}
	out := new(SriovStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VfInfo) DeepCopyInto(out *VfInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VfInfo.
func (in *VfInfo) DeepCopy() *VfInfo {
	if in == nil {
		return nil
	}
	out := new(VfInfo)
	in.DeepCopyInto(out)
	return out
}