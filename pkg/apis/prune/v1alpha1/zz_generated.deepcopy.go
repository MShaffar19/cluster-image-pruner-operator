// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageReference) DeepCopyInto(out *ImageReference) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageReference.
func (in *ImageReference) DeepCopy() *ImageReference {
	if in == nil {
		return nil
	}
	out := new(ImageReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageReference) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageReferenceList) DeepCopyInto(out *ImageReferenceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageReferenceList.
func (in *ImageReferenceList) DeepCopy() *ImageReferenceList {
	if in == nil {
		return nil
	}
	out := new(ImageReferenceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageReferenceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageReferenceSpec) DeepCopyInto(out *ImageReferenceSpec) {
	*out = *in
	if in.ImageFieldSelectors != nil {
		in, out := &in.ImageFieldSelectors, &out.ImageFieldSelectors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageReferenceSpec.
func (in *ImageReferenceSpec) DeepCopy() *ImageReferenceSpec {
	if in == nil {
		return nil
	}
	out := new(ImageReferenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PruneService) DeepCopyInto(out *PruneService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PruneService.
func (in *PruneService) DeepCopy() *PruneService {
	if in == nil {
		return nil
	}
	out := new(PruneService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PruneService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PruneServiceImages) DeepCopyInto(out *PruneServiceImages) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PruneServiceImages.
func (in *PruneServiceImages) DeepCopy() *PruneServiceImages {
	if in == nil {
		return nil
	}
	out := new(PruneServiceImages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PruneServiceList) DeepCopyInto(out *PruneServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PruneService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PruneServiceList.
func (in *PruneServiceList) DeepCopy() *PruneServiceList {
	if in == nil {
		return nil
	}
	out := new(PruneServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PruneServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PruneServiceSpec) DeepCopyInto(out *PruneServiceSpec) {
	*out = *in
	out.Images = in.Images
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PruneServiceSpec.
func (in *PruneServiceSpec) DeepCopy() *PruneServiceSpec {
	if in == nil {
		return nil
	}
	out := new(PruneServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PruneServiceStatus) DeepCopyInto(out *PruneServiceStatus) {
	*out = *in
	in.OperatorStatus.DeepCopyInto(&out.OperatorStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PruneServiceStatus.
func (in *PruneServiceStatus) DeepCopy() *PruneServiceStatus {
	if in == nil {
		return nil
	}
	out := new(PruneServiceStatus)
	in.DeepCopyInto(out)
	return out
}
