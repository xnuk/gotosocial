// Code generated by astool. DO NOT EDIT.

package vocab

import "net/url"

// GoToSocialInteractionPolicyPropertyIterator represents a single value for the
// "interactionPolicy" property.
type GoToSocialInteractionPolicyPropertyIterator interface {
	// Get returns the value of this property. When
	// IsGoToSocialInteractionPolicy returns false, Get will return any
	// arbitrary value.
	Get() GoToSocialInteractionPolicy
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return any arbitrary value.
	GetIRI() *url.URL
	// GetType returns the value in this property as a Type. Returns nil if
	// the value is not an ActivityStreams type, such as an IRI or another
	// value.
	GetType() Type
	// HasAny returns true if the value or IRI is set.
	HasAny() bool
	// IsGoToSocialInteractionPolicy returns true if this property is set and
	// not an IRI.
	IsGoToSocialInteractionPolicy() bool
	// IsIRI returns true if this property is an IRI.
	IsIRI() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o GoToSocialInteractionPolicyPropertyIterator) bool
	// Name returns the name of this property: "GoToSocialInteractionPolicy".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() GoToSocialInteractionPolicyPropertyIterator
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() GoToSocialInteractionPolicyPropertyIterator
	// Set sets the value of this property. Calling
	// IsGoToSocialInteractionPolicy afterwards will return true.
	Set(v GoToSocialInteractionPolicy)
	// SetIRI sets the value of this property. Calling IsIRI afterwards will
	// return true.
	SetIRI(v *url.URL)
	// SetType attempts to set the property for the arbitrary type. Returns an
	// error if it is not a valid type to set on this property.
	SetType(t Type) error
}

// InteractionPolicy for an ActivityStreams Object.
//
//   null
type GoToSocialInteractionPolicyProperty interface {
	// AppendGoToSocialInteractionPolicy appends a InteractionPolicy value to
	// the back of a list of the property "interactionPolicy". Invalidates
	// iterators that are traversing using Prev.
	AppendGoToSocialInteractionPolicy(v GoToSocialInteractionPolicy)
	// AppendIRI appends an IRI value to the back of a list of the property
	// "interactionPolicy"
	AppendIRI(v *url.URL)
	// PrependType prepends an arbitrary type value to the front of a list of
	// the property "interactionPolicy". Invalidates iterators that are
	// traversing using Prev. Returns an error if the type is not a valid
	// one to set for this property.
	AppendType(t Type) error
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) GoToSocialInteractionPolicyPropertyIterator
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() GoToSocialInteractionPolicyPropertyIterator
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() GoToSocialInteractionPolicyPropertyIterator
	// InsertGoToSocialInteractionPolicy inserts a InteractionPolicy value at
	// the specified index for a property "interactionPolicy". Existing
	// elements at that index and higher are shifted back once.
	// Invalidates all iterators.
	InsertGoToSocialInteractionPolicy(idx int, v GoToSocialInteractionPolicy)
	// Insert inserts an IRI value at the specified index for a property
	// "interactionPolicy". Existing elements at that index and higher are
	// shifted back once. Invalidates all iterators.
	InsertIRI(idx int, v *url.URL)
	// PrependType prepends an arbitrary type value to the front of a list of
	// the property "interactionPolicy". Invalidates all iterators.
	// Returns an error if the type is not a valid one to set for this
	// property.
	InsertType(idx int, t Type) error
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API method specifically needed only for alternate
	// implementations for go-fed. Applications should not use this
	// method. Panics if the index is out of bounds.
	KindIndex(idx int) int
	// Len returns the number of values that exist for the "interactionPolicy"
	// property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o GoToSocialInteractionPolicyProperty) bool
	// Name returns the name of this property ("interactionPolicy") with any
	// alias.
	Name() string
	// PrependGoToSocialInteractionPolicy prepends a InteractionPolicy value
	// to the front of a list of the property "interactionPolicy".
	// Invalidates all iterators.
	PrependGoToSocialInteractionPolicy(v GoToSocialInteractionPolicy)
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "interactionPolicy".
	PrependIRI(v *url.URL)
	// PrependType prepends an arbitrary type value to the front of a list of
	// the property "interactionPolicy". Invalidates all iterators.
	// Returns an error if the type is not a valid one to set for this
	// property.
	PrependType(t Type) error
	// Remove deletes an element at the specified index from a list of the
	// property "interactionPolicy", regardless of its type. Panics if the
	// index is out of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// Set sets a InteractionPolicy value to be at the specified index for the
	// property "interactionPolicy". Panics if the index is out of bounds.
	// Invalidates all iterators.
	Set(idx int, v GoToSocialInteractionPolicy)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "interactionPolicy". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// SetType sets an arbitrary type value to the specified index of the
	// property "interactionPolicy". Invalidates all iterators. Returns an
	// error if the type is not a valid one to set for this property.
	// Panics if the index is out of bounds.
	SetType(idx int, t Type) error
	// Swap swaps the location of values at two indices for the
	// "interactionPolicy" property.
	Swap(i, j int)
}
