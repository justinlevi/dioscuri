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
func (in *IngressMigrate) DeepCopyInto(out *IngressMigrate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressMigrate.
func (in *IngressMigrate) DeepCopy() *IngressMigrate {
	if in == nil {
		return nil
	}
	out := new(IngressMigrate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressMigrate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressMigrateList) DeepCopyInto(out *IngressMigrateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngressMigrate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressMigrateList.
func (in *IngressMigrateList) DeepCopy() *IngressMigrateList {
	if in == nil {
		return nil
	}
	out := new(IngressMigrateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressMigrateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressMigrateSpec) DeepCopyInto(out *IngressMigrateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressMigrateSpec.
func (in *IngressMigrateSpec) DeepCopy() *IngressMigrateSpec {
	if in == nil {
		return nil
	}
	out := new(IngressMigrateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressMigrateStatus) DeepCopyInto(out *IngressMigrateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressMigrateStatus.
func (in *IngressMigrateStatus) DeepCopy() *IngressMigrateStatus {
	if in == nil {
		return nil
	}
	out := new(IngressMigrateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteMigrate) DeepCopyInto(out *RouteMigrate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteMigrate.
func (in *RouteMigrate) DeepCopy() *RouteMigrate {
	if in == nil {
		return nil
	}
	out := new(RouteMigrate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteMigrate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteMigrateConditions) DeepCopyInto(out *RouteMigrateConditions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteMigrateConditions.
func (in *RouteMigrateConditions) DeepCopy() *RouteMigrateConditions {
	if in == nil {
		return nil
	}
	out := new(RouteMigrateConditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteMigrateList) DeepCopyInto(out *RouteMigrateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RouteMigrate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteMigrateList.
func (in *RouteMigrateList) DeepCopy() *RouteMigrateList {
	if in == nil {
		return nil
	}
	out := new(RouteMigrateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteMigrateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteMigrateRoutes) DeepCopyInto(out *RouteMigrateRoutes) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteMigrateRoutes.
func (in *RouteMigrateRoutes) DeepCopy() *RouteMigrateRoutes {
	if in == nil {
		return nil
	}
	out := new(RouteMigrateRoutes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteMigrateSpec) DeepCopyInto(out *RouteMigrateSpec) {
	*out = *in
	out.Routes = in.Routes
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteMigrateSpec.
func (in *RouteMigrateSpec) DeepCopy() *RouteMigrateSpec {
	if in == nil {
		return nil
	}
	out := new(RouteMigrateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteMigrateStatus) DeepCopyInto(out *RouteMigrateStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RouteMigrateConditions, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteMigrateStatus.
func (in *RouteMigrateStatus) DeepCopy() *RouteMigrateStatus {
	if in == nil {
		return nil
	}
	out := new(RouteMigrateStatus)
	in.DeepCopyInto(out)
	return out
}
