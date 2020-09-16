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
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlicloudCluster) DeepCopyInto(out *AlicloudCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlicloudCluster.
func (in *AlicloudCluster) DeepCopy() *AlicloudCluster {
	if in == nil {
		return nil
	}
	out := new(AlicloudCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlicloudCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlicloudClusterList) DeepCopyInto(out *AlicloudClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlicloudCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlicloudClusterList.
func (in *AlicloudClusterList) DeepCopy() *AlicloudClusterList {
	if in == nil {
		return nil
	}
	out := new(AlicloudClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlicloudClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlicloudClusterSpec) DeepCopyInto(out *AlicloudClusterSpec) {
	*out = *in
	out.Vpc = in.Vpc
	out.VSwitch = in.VSwitch
	in.Cluster.DeepCopyInto(&out.Cluster)
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlicloudClusterSpec.
func (in *AlicloudClusterSpec) DeepCopy() *AlicloudClusterSpec {
	if in == nil {
		return nil
	}
	out := new(AlicloudClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlicloudClusterStatus) DeepCopyInto(out *AlicloudClusterStatus) {
	*out = *in
	out.Vpc = in.Vpc
	out.VSwitch = in.VSwitch
	out.Cluster = in.Cluster
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlicloudClusterStatus.
func (in *AlicloudClusterStatus) DeepCopy() *AlicloudClusterStatus {
	if in == nil {
		return nil
	}
	out := new(AlicloudClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = make([]Addon, len(*in))
		copy(*out, *in)
	}
	if in.MasterVswitchIds != nil {
		in, out := &in.MasterVswitchIds, &out.MasterVswitchIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MasterInstanceTypes != nil {
		in, out := &in.MasterInstanceTypes, &out.MasterInstanceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Runtime = in.Runtime
	if in.WorkerInstanceTypes != nil {
		in, out := &in.WorkerInstanceTypes, &out.WorkerInstanceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WorkerVswitchIds != nil {
		in, out := &in.WorkerVswitchIds, &out.WorkerVswitchIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeType) DeepCopyInto(out *RuntimeType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeType.
func (in *RuntimeType) DeepCopy() *RuntimeType {
	if in == nil {
		return nil
	}
	out := new(RuntimeType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleSpec) DeepCopyInto(out *ScaleSpec) {
	*out = *in
	if in.WorkerDataDisks != nil {
		in, out := &in.WorkerDataDisks, &out.WorkerDataDisks
		*out = make([]WorkerDataDisk, len(*in))
		copy(*out, *in)
	}
	if in.WorkerInstanceTypes != nil {
		in, out := &in.WorkerInstanceTypes, &out.WorkerInstanceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VSwitchIds != nil {
		in, out := &in.VSwitchIds, &out.VSwitchIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleSpec.
func (in *ScaleSpec) DeepCopy() *ScaleSpec {
	if in == nil {
		return nil
	}
	out := new(ScaleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSwitch) DeepCopyInto(out *VSwitch) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSwitch.
func (in *VSwitch) DeepCopy() *VSwitch {
	if in == nil {
		return nil
	}
	out := new(VSwitch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSwitchSpec) DeepCopyInto(out *VSwitchSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSwitchSpec.
func (in *VSwitchSpec) DeepCopy() *VSwitchSpec {
	if in == nil {
		return nil
	}
	out := new(VSwitchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vpc) DeepCopyInto(out *Vpc) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vpc.
func (in *Vpc) DeepCopy() *Vpc {
	if in == nil {
		return nil
	}
	out := new(Vpc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcSpec) DeepCopyInto(out *VpcSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcSpec.
func (in *VpcSpec) DeepCopy() *VpcSpec {
	if in == nil {
		return nil
	}
	out := new(VpcSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerDataDisk) DeepCopyInto(out *WorkerDataDisk) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerDataDisk.
func (in *WorkerDataDisk) DeepCopy() *WorkerDataDisk {
	if in == nil {
		return nil
	}
	out := new(WorkerDataDisk)
	in.DeepCopyInto(out)
	return out
}
