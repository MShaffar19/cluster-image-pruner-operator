// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/openshift/cluster-prune-operator/pkg/apis/prune/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePruneServices implements PruneServiceInterface
type FakePruneServices struct {
	Fake *FakePruneV1alpha1
	ns   string
}

var pruneservicesResource = schema.GroupVersionResource{Group: "prune.operator.openshift.io", Version: "v1alpha1", Resource: "pruneservices"}

var pruneservicesKind = schema.GroupVersionKind{Group: "prune.operator.openshift.io", Version: "v1alpha1", Kind: "PruneService"}

// Get takes name of the pruneService, and returns the corresponding pruneService object, and an error if there is any.
func (c *FakePruneServices) Get(name string, options v1.GetOptions) (result *v1alpha1.PruneService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pruneservicesResource, c.ns, name), &v1alpha1.PruneService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PruneService), err
}

// List takes label and field selectors, and returns the list of PruneServices that match those selectors.
func (c *FakePruneServices) List(opts v1.ListOptions) (result *v1alpha1.PruneServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pruneservicesResource, pruneservicesKind, c.ns, opts), &v1alpha1.PruneServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PruneServiceList{ListMeta: obj.(*v1alpha1.PruneServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.PruneServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pruneServices.
func (c *FakePruneServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pruneservicesResource, c.ns, opts))

}

// Create takes the representation of a pruneService and creates it.  Returns the server's representation of the pruneService, and an error, if there is any.
func (c *FakePruneServices) Create(pruneService *v1alpha1.PruneService) (result *v1alpha1.PruneService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pruneservicesResource, c.ns, pruneService), &v1alpha1.PruneService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PruneService), err
}

// Update takes the representation of a pruneService and updates it. Returns the server's representation of the pruneService, and an error, if there is any.
func (c *FakePruneServices) Update(pruneService *v1alpha1.PruneService) (result *v1alpha1.PruneService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pruneservicesResource, c.ns, pruneService), &v1alpha1.PruneService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PruneService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePruneServices) UpdateStatus(pruneService *v1alpha1.PruneService) (*v1alpha1.PruneService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pruneservicesResource, "status", c.ns, pruneService), &v1alpha1.PruneService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PruneService), err
}

// Delete takes name of the pruneService and deletes it. Returns an error if one occurs.
func (c *FakePruneServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pruneservicesResource, c.ns, name), &v1alpha1.PruneService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePruneServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pruneservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PruneServiceList{})
	return err
}

// Patch applies the patch and returns the patched pruneService.
func (c *FakePruneServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PruneService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pruneservicesResource, c.ns, name, data, subresources...), &v1alpha1.PruneService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PruneService), err
}
