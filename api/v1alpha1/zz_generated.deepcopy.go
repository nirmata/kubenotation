//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureVerification) DeepCopyInto(out *SignatureVerification) {
	*out = *in
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = make(map[ValidationType]ValidationAction, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureVerification.
func (in *SignatureVerification) DeepCopy() *SignatureVerification {
	if in == nil {
		return nil
	}
	out := new(SignatureVerification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrustDocument) DeepCopyInto(out *TrustDocument) {
	*out = *in
	if in.TrustPolicies != nil {
		in, out := &in.TrustPolicies, &out.TrustPolicies
		*out = make([]TrustPolicyStatement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrustDocument.
func (in *TrustDocument) DeepCopy() *TrustDocument {
	if in == nil {
		return nil
	}
	out := new(TrustDocument)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrustPolicy) DeepCopyInto(out *TrustPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrustPolicy.
func (in *TrustPolicy) DeepCopy() *TrustPolicy {
	if in == nil {
		return nil
	}
	out := new(TrustPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrustPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrustPolicyList) DeepCopyInto(out *TrustPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrustPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrustPolicyList.
func (in *TrustPolicyList) DeepCopy() *TrustPolicyList {
	if in == nil {
		return nil
	}
	out := new(TrustPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrustPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrustPolicyStatement) DeepCopyInto(out *TrustPolicyStatement) {
	*out = *in
	if in.RegistryScopes != nil {
		in, out := &in.RegistryScopes, &out.RegistryScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.SignatureVerification.DeepCopyInto(&out.SignatureVerification)
	if in.TrustStores != nil {
		in, out := &in.TrustStores, &out.TrustStores
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TrustedIdentities != nil {
		in, out := &in.TrustedIdentities, &out.TrustedIdentities
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrustPolicyStatement.
func (in *TrustPolicyStatement) DeepCopy() *TrustPolicyStatement {
	if in == nil {
		return nil
	}
	out := new(TrustPolicyStatement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrustStore) DeepCopyInto(out *TrustStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrustStore.
func (in *TrustStore) DeepCopy() *TrustStore {
	if in == nil {
		return nil
	}
	out := new(TrustStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrustStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrustStoreList) DeepCopyInto(out *TrustStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrustStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrustStoreList.
func (in *TrustStoreList) DeepCopy() *TrustStoreList {
	if in == nil {
		return nil
	}
	out := new(TrustStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrustStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrustStoreSpec) DeepCopyInto(out *TrustStoreSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrustStoreSpec.
func (in *TrustStoreSpec) DeepCopy() *TrustStoreSpec {
	if in == nil {
		return nil
	}
	out := new(TrustStoreSpec)
	in.DeepCopyInto(out)
	return out
}
