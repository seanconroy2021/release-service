//go:build !ignore_autogenerated

/*
Copyright 2022.

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

package v1alpha1

import (
	"github.com/konflux-ci/release-service/tekton/utils"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttributionInfo) DeepCopyInto(out *AttributionInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttributionInfo.
func (in *AttributionInfo) DeepCopy() *AttributionInfo {
	if in == nil {
		return nil
	}
	out := new(AttributionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorItem) DeepCopyInto(out *CollectorItem) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]Param, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorItem.
func (in *CollectorItem) DeepCopy() *CollectorItem {
	if in == nil {
		return nil
	}
	out := new(CollectorItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Collectors) DeepCopyInto(out *Collectors) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CollectorItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Collectors.
func (in *Collectors) DeepCopy() *Collectors {
	if in == nil {
		return nil
	}
	out := new(Collectors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorsInfo) DeepCopyInto(out *CollectorsInfo) {
	*out = *in
	in.ManagedCollectorsProcessing.DeepCopyInto(&out.ManagedCollectorsProcessing)
	in.TenantCollectorsProcessing.DeepCopyInto(&out.TenantCollectorsProcessing)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorsInfo.
func (in *CollectorsInfo) DeepCopy() *CollectorsInfo {
	if in == nil {
		return nil
	}
	out := new(CollectorsInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmptyDirOverrides) DeepCopyInto(out *EmptyDirOverrides) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmptyDirOverrides.
func (in *EmptyDirOverrides) DeepCopy() *EmptyDirOverrides {
	if in == nil {
		return nil
	}
	out := new(EmptyDirOverrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchedReleasePlan) DeepCopyInto(out *MatchedReleasePlan) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchedReleasePlan.
func (in *MatchedReleasePlan) DeepCopy() *MatchedReleasePlan {
	if in == nil {
		return nil
	}
	out := new(MatchedReleasePlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchedReleasePlanAdmission) DeepCopyInto(out *MatchedReleasePlanAdmission) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchedReleasePlanAdmission.
func (in *MatchedReleasePlanAdmission) DeepCopy() *MatchedReleasePlanAdmission {
	if in == nil {
		return nil
	}
	out := new(MatchedReleasePlanAdmission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Param) DeepCopyInto(out *Param) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Param.
func (in *Param) DeepCopy() *Param {
	if in == nil {
		return nil
	}
	out := new(Param)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineInfo) DeepCopyInto(out *PipelineInfo) {
	*out = *in
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineInfo.
func (in *PipelineInfo) DeepCopy() *PipelineInfo {
	if in == nil {
		return nil
	}
	out := new(PipelineInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Release) DeepCopyInto(out *Release) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Release.
func (in *Release) DeepCopy() *Release {
	if in == nil {
		return nil
	}
	out := new(Release)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Release) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseList) DeepCopyInto(out *ReleaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Release, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseList.
func (in *ReleaseList) DeepCopy() *ReleaseList {
	if in == nil {
		return nil
	}
	out := new(ReleaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePlan) DeepCopyInto(out *ReleasePlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePlan.
func (in *ReleasePlan) DeepCopy() *ReleasePlan {
	if in == nil {
		return nil
	}
	out := new(ReleasePlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleasePlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePlanAdmission) DeepCopyInto(out *ReleasePlanAdmission) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePlanAdmission.
func (in *ReleasePlanAdmission) DeepCopy() *ReleasePlanAdmission {
	if in == nil {
		return nil
	}
	out := new(ReleasePlanAdmission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleasePlanAdmission) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePlanAdmissionList) DeepCopyInto(out *ReleasePlanAdmissionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReleasePlanAdmission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePlanAdmissionList.
func (in *ReleasePlanAdmissionList) DeepCopy() *ReleasePlanAdmissionList {
	if in == nil {
		return nil
	}
	out := new(ReleasePlanAdmissionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleasePlanAdmissionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePlanAdmissionSpec) DeepCopyInto(out *ReleasePlanAdmissionSpec) {
	*out = *in
	if in.Applications != nil {
		in, out := &in.Applications, &out.Applications
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Collectors != nil {
		in, out := &in.Collectors, &out.Collectors
		*out = new(Collectors)
		(*in).DeepCopyInto(*out)
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Pipeline != nil {
		in, out := &in.Pipeline, &out.Pipeline
		*out = new(utils.Pipeline)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePlanAdmissionSpec.
func (in *ReleasePlanAdmissionSpec) DeepCopy() *ReleasePlanAdmissionSpec {
	if in == nil {
		return nil
	}
	out := new(ReleasePlanAdmissionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePlanAdmissionStatus) DeepCopyInto(out *ReleasePlanAdmissionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReleasePlans != nil {
		in, out := &in.ReleasePlans, &out.ReleasePlans
		*out = make([]MatchedReleasePlan, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePlanAdmissionStatus.
func (in *ReleasePlanAdmissionStatus) DeepCopy() *ReleasePlanAdmissionStatus {
	if in == nil {
		return nil
	}
	out := new(ReleasePlanAdmissionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePlanList) DeepCopyInto(out *ReleasePlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReleasePlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePlanList.
func (in *ReleasePlanList) DeepCopy() *ReleasePlanList {
	if in == nil {
		return nil
	}
	out := new(ReleasePlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleasePlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePlanSpec) DeepCopyInto(out *ReleasePlanSpec) {
	*out = *in
	if in.Collectors != nil {
		in, out := &in.Collectors, &out.Collectors
		*out = new(Collectors)
		(*in).DeepCopyInto(*out)
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.TenantPipeline != nil {
		in, out := &in.TenantPipeline, &out.TenantPipeline
		*out = new(utils.ParameterizedPipeline)
		(*in).DeepCopyInto(*out)
	}
	if in.FinalPipeline != nil {
		in, out := &in.FinalPipeline, &out.FinalPipeline
		*out = new(utils.ParameterizedPipeline)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePlanSpec.
func (in *ReleasePlanSpec) DeepCopy() *ReleasePlanSpec {
	if in == nil {
		return nil
	}
	out := new(ReleasePlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePlanStatus) DeepCopyInto(out *ReleasePlanStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ReleasePlanAdmission = in.ReleasePlanAdmission
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePlanStatus.
func (in *ReleasePlanStatus) DeepCopy() *ReleasePlanStatus {
	if in == nil {
		return nil
	}
	out := new(ReleasePlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseServiceConfig) DeepCopyInto(out *ReleaseServiceConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseServiceConfig.
func (in *ReleaseServiceConfig) DeepCopy() *ReleaseServiceConfig {
	if in == nil {
		return nil
	}
	out := new(ReleaseServiceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseServiceConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseServiceConfigList) DeepCopyInto(out *ReleaseServiceConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReleaseServiceConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseServiceConfigList.
func (in *ReleaseServiceConfigList) DeepCopy() *ReleaseServiceConfigList {
	if in == nil {
		return nil
	}
	out := new(ReleaseServiceConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseServiceConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseServiceConfigSpec) DeepCopyInto(out *ReleaseServiceConfigSpec) {
	*out = *in
	in.DefaultTimeouts.DeepCopyInto(&out.DefaultTimeouts)
	if in.EmptyDirOverrides != nil {
		in, out := &in.EmptyDirOverrides, &out.EmptyDirOverrides
		*out = make([]EmptyDirOverrides, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseServiceConfigSpec.
func (in *ReleaseServiceConfigSpec) DeepCopy() *ReleaseServiceConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ReleaseServiceConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseServiceConfigStatus) DeepCopyInto(out *ReleaseServiceConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseServiceConfigStatus.
func (in *ReleaseServiceConfigStatus) DeepCopy() *ReleaseServiceConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseServiceConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseSpec) DeepCopyInto(out *ReleaseSpec) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseSpec.
func (in *ReleaseSpec) DeepCopy() *ReleaseSpec {
	if in == nil {
		return nil
	}
	out := new(ReleaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseStatus) DeepCopyInto(out *ReleaseStatus) {
	*out = *in
	if in.Artifacts != nil {
		in, out := &in.Artifacts, &out.Artifacts
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	out.Attribution = in.Attribution
	if in.Collectors != nil {
		in, out := &in.Collectors, &out.Collectors
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.CollectorsProcessing.DeepCopyInto(&out.CollectorsProcessing)
	in.FinalProcessing.DeepCopyInto(&out.FinalProcessing)
	in.ManagedProcessing.DeepCopyInto(&out.ManagedProcessing)
	in.TenantProcessing.DeepCopyInto(&out.TenantProcessing)
	in.Validation.DeepCopyInto(&out.Validation)
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.ExpirationTime != nil {
		in, out := &in.ExpirationTime, &out.ExpirationTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseStatus.
func (in *ReleaseStatus) DeepCopy() *ReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationInfo) DeepCopyInto(out *ValidationInfo) {
	*out = *in
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationInfo.
func (in *ValidationInfo) DeepCopy() *ValidationInfo {
	if in == nil {
		return nil
	}
	out := new(ValidationInfo)
	in.DeepCopyInto(out)
	return out
}
