//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Knative Authors

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	apis "knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Channelable) DeepCopyInto(out *Channelable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Channelable.
func (in *Channelable) DeepCopy() *Channelable {
	if in == nil {
		return nil
	}
	out := new(Channelable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Channelable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelableList) DeepCopyInto(out *ChannelableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Channelable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelableList.
func (in *ChannelableList) DeepCopy() *ChannelableList {
	if in == nil {
		return nil
	}
	out := new(ChannelableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChannelableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelableSpec) DeepCopyInto(out *ChannelableSpec) {
	*out = *in
	in.SubscribableSpec.DeepCopyInto(&out.SubscribableSpec)
	if in.Delivery != nil {
		in, out := &in.Delivery, &out.Delivery
		*out = new(DeliverySpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelableSpec.
func (in *ChannelableSpec) DeepCopy() *ChannelableSpec {
	if in == nil {
		return nil
	}
	out := new(ChannelableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelableStatus) DeepCopyInto(out *ChannelableStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	in.SubscribableStatus.DeepCopyInto(&out.SubscribableStatus)
	in.DeliveryStatus.DeepCopyInto(&out.DeliveryStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelableStatus.
func (in *ChannelableStatus) DeepCopy() *ChannelableStatus {
	if in == nil {
		return nil
	}
	out := new(ChannelableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeliverySpec) DeepCopyInto(out *DeliverySpec) {
	*out = *in
	if in.DeadLetterSink != nil {
		in, out := &in.DeadLetterSink, &out.DeadLetterSink
		*out = new(duckv1.Destination)
		(*in).DeepCopyInto(*out)
	}
	if in.Retry != nil {
		in, out := &in.Retry, &out.Retry
		*out = new(int32)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(string)
		**out = **in
	}
	if in.BackoffPolicy != nil {
		in, out := &in.BackoffPolicy, &out.BackoffPolicy
		*out = new(BackoffPolicyType)
		**out = **in
	}
	if in.BackoffDelay != nil {
		in, out := &in.BackoffDelay, &out.BackoffDelay
		*out = new(string)
		**out = **in
	}
	if in.RetryAfterMax != nil {
		in, out := &in.RetryAfterMax, &out.RetryAfterMax
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeliverySpec.
func (in *DeliverySpec) DeepCopy() *DeliverySpec {
	if in == nil {
		return nil
	}
	out := new(DeliverySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeliveryStatus) DeepCopyInto(out *DeliveryStatus) {
	*out = *in
	if in.DeadLetterSinkURI != nil {
		in, out := &in.DeadLetterSinkURI, &out.DeadLetterSinkURI
		*out = new(apis.URL)
		(*in).DeepCopyInto(*out)
	}
	if in.DeadLetterSinkCACerts != nil {
		in, out := &in.DeadLetterSinkCACerts, &out.DeadLetterSinkCACerts
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeliveryStatus.
func (in *DeliveryStatus) DeepCopy() *DeliveryStatus {
	if in == nil {
		return nil
	}
	out := new(DeliveryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subscribable) DeepCopyInto(out *Subscribable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subscribable.
func (in *Subscribable) DeepCopy() *Subscribable {
	if in == nil {
		return nil
	}
	out := new(Subscribable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Subscribable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscribableList) DeepCopyInto(out *SubscribableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Subscribable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscribableList.
func (in *SubscribableList) DeepCopy() *SubscribableList {
	if in == nil {
		return nil
	}
	out := new(SubscribableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubscribableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscribableSpec) DeepCopyInto(out *SubscribableSpec) {
	*out = *in
	if in.Subscribers != nil {
		in, out := &in.Subscribers, &out.Subscribers
		*out = make([]SubscriberSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscribableSpec.
func (in *SubscribableSpec) DeepCopy() *SubscribableSpec {
	if in == nil {
		return nil
	}
	out := new(SubscribableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscribableStatus) DeepCopyInto(out *SubscribableStatus) {
	*out = *in
	if in.Subscribers != nil {
		in, out := &in.Subscribers, &out.Subscribers
		*out = make([]SubscriberStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscribableStatus.
func (in *SubscribableStatus) DeepCopy() *SubscribableStatus {
	if in == nil {
		return nil
	}
	out := new(SubscribableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriberSpec) DeepCopyInto(out *SubscriberSpec) {
	*out = *in
	if in.SubscriberURI != nil {
		in, out := &in.SubscriberURI, &out.SubscriberURI
		*out = new(apis.URL)
		(*in).DeepCopyInto(*out)
	}
	if in.SubscriberCACerts != nil {
		in, out := &in.SubscriberCACerts, &out.SubscriberCACerts
		*out = new(string)
		**out = **in
	}
	if in.ReplyURI != nil {
		in, out := &in.ReplyURI, &out.ReplyURI
		*out = new(apis.URL)
		(*in).DeepCopyInto(*out)
	}
	if in.ReplyCACerts != nil {
		in, out := &in.ReplyCACerts, &out.ReplyCACerts
		*out = new(string)
		**out = **in
	}
	if in.Delivery != nil {
		in, out := &in.Delivery, &out.Delivery
		*out = new(DeliverySpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriberSpec.
func (in *SubscriberSpec) DeepCopy() *SubscriberSpec {
	if in == nil {
		return nil
	}
	out := new(SubscriberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriberStatus) DeepCopyInto(out *SubscriberStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriberStatus.
func (in *SubscriberStatus) DeepCopy() *SubscriberStatus {
	if in == nil {
		return nil
	}
	out := new(SubscriberStatus)
	in.DeepCopyInto(out)
	return out
}
