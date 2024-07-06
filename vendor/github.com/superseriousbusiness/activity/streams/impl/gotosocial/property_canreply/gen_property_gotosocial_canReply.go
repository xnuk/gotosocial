// Code generated by astool. DO NOT EDIT.

package propertycanreply

import (
	"fmt"
	vocab "github.com/superseriousbusiness/activity/streams/vocab"
	"net/url"
)

// GoToSocialCanReplyPropertyIterator is an iterator for a property. It is
// permitted to be a single nilable value type.
type GoToSocialCanReplyPropertyIterator struct {
	gotosocialCanReplyMember vocab.GoToSocialCanReply
	unknown                  interface{}
	iri                      *url.URL
	alias                    string
	myIdx                    int
	parent                   vocab.GoToSocialCanReplyProperty
}

// NewGoToSocialCanReplyPropertyIterator creates a new GoToSocialCanReply property.
func NewGoToSocialCanReplyPropertyIterator() *GoToSocialCanReplyPropertyIterator {
	return &GoToSocialCanReplyPropertyIterator{alias: ""}
}

// deserializeGoToSocialCanReplyPropertyIterator creates an iterator from an
// element that has been unmarshalled from a text or binary format.
func deserializeGoToSocialCanReplyPropertyIterator(i interface{}, aliasMap map[string]string) (*GoToSocialCanReplyPropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://gotosocial.org/ns"]; ok {
		alias = a
	}
	if s, ok := i.(string); ok {
		u, err := url.Parse(s)
		// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
		// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
		if err == nil && len(u.Scheme) > 0 {
			this := &GoToSocialCanReplyPropertyIterator{
				alias: alias,
				iri:   u,
			}
			return this, nil
		}
	}
	if m, ok := i.(map[string]interface{}); ok {
		if v, err := mgr.DeserializeCanReplyGoToSocial()(m, aliasMap); err == nil {
			this := &GoToSocialCanReplyPropertyIterator{
				alias:                    alias,
				gotosocialCanReplyMember: v,
			}
			return this, nil
		}
	}
	this := &GoToSocialCanReplyPropertyIterator{
		alias:   alias,
		unknown: i,
	}
	return this, nil
}

// Get returns the value of this property. When IsGoToSocialCanReply returns
// false, Get will return any arbitrary value.
func (this GoToSocialCanReplyPropertyIterator) Get() vocab.GoToSocialCanReply {
	return this.gotosocialCanReplyMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this GoToSocialCanReplyPropertyIterator) GetIRI() *url.URL {
	return this.iri
}

// GetType returns the value in this property as a Type. Returns nil if the value
// is not an ActivityStreams type, such as an IRI or another value.
func (this GoToSocialCanReplyPropertyIterator) GetType() vocab.Type {
	if this.IsGoToSocialCanReply() {
		return this.Get()
	}

	return nil
}

// HasAny returns true if the value or IRI is set.
func (this GoToSocialCanReplyPropertyIterator) HasAny() bool {
	return this.IsGoToSocialCanReply() || this.iri != nil
}

// IsGoToSocialCanReply returns true if this property is set and not an IRI.
func (this GoToSocialCanReplyPropertyIterator) IsGoToSocialCanReply() bool {
	return this.gotosocialCanReplyMember != nil
}

// IsIRI returns true if this property is an IRI.
func (this GoToSocialCanReplyPropertyIterator) IsIRI() bool {
	return this.iri != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this GoToSocialCanReplyPropertyIterator) JSONLDContext() map[string]string {
	m := map[string]string{"https://gotosocial.org/ns": this.alias}
	var child map[string]string
	if this.IsGoToSocialCanReply() {
		child = this.Get().JSONLDContext()
	}
	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this GoToSocialCanReplyPropertyIterator) KindIndex() int {
	if this.IsGoToSocialCanReply() {
		return 0
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this GoToSocialCanReplyPropertyIterator) LessThan(o vocab.GoToSocialCanReplyPropertyIterator) bool {
	// LessThan comparison for if either or both are IRIs.
	if this.IsIRI() && o.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	} else if this.IsIRI() {
		// IRIs are always less than other values, none, or unknowns
		return true
	} else if o.IsIRI() {
		// This other, none, or unknown value is always greater than IRIs
		return false
	}
	// LessThan comparison for the single value or unknown value.
	if !this.IsGoToSocialCanReply() && !o.IsGoToSocialCanReply() {
		// Both are unknowns.
		return false
	} else if this.IsGoToSocialCanReply() && !o.IsGoToSocialCanReply() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsGoToSocialCanReply() && o.IsGoToSocialCanReply() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return this.Get().LessThan(o.Get())
	}
}

// Name returns the name of this property: "GoToSocialCanReply".
func (this GoToSocialCanReplyPropertyIterator) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "GoToSocialCanReply"
	} else {
		return "GoToSocialCanReply"
	}
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this GoToSocialCanReplyPropertyIterator) Next() vocab.GoToSocialCanReplyPropertyIterator {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this GoToSocialCanReplyPropertyIterator) Prev() vocab.GoToSocialCanReplyPropertyIterator {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// Set sets the value of this property. Calling IsGoToSocialCanReply afterwards
// will return true.
func (this *GoToSocialCanReplyPropertyIterator) Set(v vocab.GoToSocialCanReply) {
	this.clear()
	this.gotosocialCanReplyMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *GoToSocialCanReplyPropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.iri = v
}

// SetType attempts to set the property for the arbitrary type. Returns an error
// if it is not a valid type to set on this property.
func (this *GoToSocialCanReplyPropertyIterator) SetType(t vocab.Type) error {
	if v, ok := t.(vocab.GoToSocialCanReply); ok {
		this.Set(v)
		return nil
	}

	return fmt.Errorf("illegal type to set on GoToSocialCanReply property: %T", t)
}

// clear ensures no value of this property is set. Calling IsGoToSocialCanReply
// afterwards will return false.
func (this *GoToSocialCanReplyPropertyIterator) clear() {
	this.unknown = nil
	this.iri = nil
	this.gotosocialCanReplyMember = nil
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this GoToSocialCanReplyPropertyIterator) serialize() (interface{}, error) {
	if this.IsGoToSocialCanReply() {
		return this.Get().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// GoToSocialCanReplyProperty is the non-functional property "canReply". It is
// permitted to have one or more values, and of different value types.
type GoToSocialCanReplyProperty struct {
	properties []*GoToSocialCanReplyPropertyIterator
	alias      string
}

// DeserializeCanReplyProperty creates a "canReply" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeCanReplyProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.GoToSocialCanReplyProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://gotosocial.org/ns"]; ok {
		alias = a
	}
	propName := "canReply"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "canReply")
	}
	i, ok := m[propName]

	if ok {
		this := &GoToSocialCanReplyProperty{
			alias:      alias,
			properties: []*GoToSocialCanReplyPropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeGoToSocialCanReplyPropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeGoToSocialCanReplyPropertyIterator(i, aliasMap); err != nil {
				return this, err
			} else if p != nil {
				this.properties = append(this.properties, p)
			}
		}
		// Set up the properties for iteration.
		for idx, ele := range this.properties {
			ele.parent = this
			ele.myIdx = idx
		}
		return this, nil
	}
	return nil, nil
}

// NewGoToSocialCanReplyProperty creates a new canReply property.
func NewGoToSocialCanReplyProperty() *GoToSocialCanReplyProperty {
	return &GoToSocialCanReplyProperty{alias: ""}
}

// AppendGoToSocialCanReply appends a CanReply value to the back of a list of the
// property "canReply". Invalidates iterators that are traversing using Prev.
func (this *GoToSocialCanReplyProperty) AppendGoToSocialCanReply(v vocab.GoToSocialCanReply) {
	this.properties = append(this.properties, &GoToSocialCanReplyPropertyIterator{
		alias:                    this.alias,
		gotosocialCanReplyMember: v,
		myIdx:                    this.Len(),
		parent:                   this,
	})
}

// AppendIRI appends an IRI value to the back of a list of the property "canReply"
func (this *GoToSocialCanReplyProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &GoToSocialCanReplyPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  this.Len(),
		parent: this,
	})
}

// PrependType prepends an arbitrary type value to the front of a list of the
// property "canReply". Invalidates iterators that are traversing using Prev.
// Returns an error if the type is not a valid one to set for this property.
func (this *GoToSocialCanReplyProperty) AppendType(t vocab.Type) error {
	n := &GoToSocialCanReplyPropertyIterator{
		alias:  this.alias,
		myIdx:  this.Len(),
		parent: this,
	}
	if err := n.SetType(t); err != nil {
		return err
	}
	this.properties = append(this.properties, n)
	return nil
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this GoToSocialCanReplyProperty) At(index int) vocab.GoToSocialCanReplyPropertyIterator {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this GoToSocialCanReplyProperty) Begin() vocab.GoToSocialCanReplyPropertyIterator {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this GoToSocialCanReplyProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this GoToSocialCanReplyProperty) End() vocab.GoToSocialCanReplyPropertyIterator {
	return nil
}

// InsertGoToSocialCanReply inserts a CanReply value at the specified index for a
// property "canReply". Existing elements at that index and higher are shifted
// back once. Invalidates all iterators.
func (this *GoToSocialCanReplyProperty) InsertGoToSocialCanReply(idx int, v vocab.GoToSocialCanReply) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &GoToSocialCanReplyPropertyIterator{
		alias:                    this.alias,
		gotosocialCanReplyMember: v,
		myIdx:                    idx,
		parent:                   this,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Insert inserts an IRI value at the specified index for a property "canReply".
// Existing elements at that index and higher are shifted back once.
// Invalidates all iterators.
func (this *GoToSocialCanReplyProperty) InsertIRI(idx int, v *url.URL) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &GoToSocialCanReplyPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependType prepends an arbitrary type value to the front of a list of the
// property "canReply". Invalidates all iterators. Returns an error if the
// type is not a valid one to set for this property.
func (this *GoToSocialCanReplyProperty) InsertType(idx int, t vocab.Type) error {
	n := &GoToSocialCanReplyPropertyIterator{
		alias:  this.alias,
		myIdx:  idx,
		parent: this,
	}
	if err := n.SetType(t); err != nil {
		return err
	}
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = n
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
	return nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this GoToSocialCanReplyProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://gotosocial.org/ns": this.alias}
	for _, elem := range this.properties {
		child := elem.JSONLDContext()
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		for k, v := range child {
			m[k] = v
		}
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API method specifically needed only for alternate implementations
// for go-fed. Applications should not use this method. Panics if the index is
// out of bounds.
func (this GoToSocialCanReplyProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "canReply" property.
func (this GoToSocialCanReplyProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this GoToSocialCanReplyProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].Get()
			rhs := this.properties[j].Get()
			return lhs.LessThan(rhs)
		} else if idx1 == -2 {
			lhs := this.properties[i].GetIRI()
			rhs := this.properties[j].GetIRI()
			return lhs.String() < rhs.String()
		}
	}
	return false
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this GoToSocialCanReplyProperty) LessThan(o vocab.GoToSocialCanReplyProperty) bool {
	l1 := this.Len()
	l2 := o.Len()
	l := l1
	if l2 < l1 {
		l = l2
	}
	for i := 0; i < l; i++ {
		if this.properties[i].LessThan(o.At(i)) {
			return true
		} else if o.At(i).LessThan(this.properties[i]) {
			return false
		}
	}
	return l1 < l2
}

// Name returns the name of this property ("canReply") with any alias.
func (this GoToSocialCanReplyProperty) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "canReply"
	} else {
		return "canReply"
	}
}

// PrependGoToSocialCanReply prepends a CanReply value to the front of a list of
// the property "canReply". Invalidates all iterators.
func (this *GoToSocialCanReplyProperty) PrependGoToSocialCanReply(v vocab.GoToSocialCanReply) {
	this.properties = append([]*GoToSocialCanReplyPropertyIterator{{
		alias:                    this.alias,
		gotosocialCanReplyMember: v,
		myIdx:                    0,
		parent:                   this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependIRI prepends an IRI value to the front of a list of the property
// "canReply".
func (this *GoToSocialCanReplyProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*GoToSocialCanReplyPropertyIterator{{
		alias:  this.alias,
		iri:    v,
		myIdx:  0,
		parent: this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependType prepends an arbitrary type value to the front of a list of the
// property "canReply". Invalidates all iterators. Returns an error if the
// type is not a valid one to set for this property.
func (this *GoToSocialCanReplyProperty) PrependType(t vocab.Type) error {
	n := &GoToSocialCanReplyPropertyIterator{
		alias:  this.alias,
		myIdx:  0,
		parent: this,
	}
	if err := n.SetType(t); err != nil {
		return err
	}
	this.properties = append([]*GoToSocialCanReplyPropertyIterator{n}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
	return nil
}

// Remove deletes an element at the specified index from a list of the property
// "canReply", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *GoToSocialCanReplyProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &GoToSocialCanReplyPropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this GoToSocialCanReplyProperty) Serialize() (interface{}, error) {
	s := make([]interface{}, 0, len(this.properties))
	for _, iterator := range this.properties {
		if b, err := iterator.serialize(); err != nil {
			return s, err
		} else {
			s = append(s, b)
		}
	}
	// Shortcut: if serializing one value, don't return an array -- pretty sure other Fediverse software would choke on a "type" value with array, for example.
	if len(s) == 1 {
		return s[0], nil
	}
	return s, nil
}

// Set sets a CanReply value to be at the specified index for the property
// "canReply". Panics if the index is out of bounds. Invalidates all iterators.
func (this *GoToSocialCanReplyProperty) Set(idx int, v vocab.GoToSocialCanReply) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &GoToSocialCanReplyPropertyIterator{
		alias:                    this.alias,
		gotosocialCanReplyMember: v,
		myIdx:                    idx,
		parent:                   this,
	}
}

// SetIRI sets an IRI value to be at the specified index for the property
// "canReply". Panics if the index is out of bounds.
func (this *GoToSocialCanReplyProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &GoToSocialCanReplyPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
}

// SetType sets an arbitrary type value to the specified index of the property
// "canReply". Invalidates all iterators. Returns an error if the type is not
// a valid one to set for this property. Panics if the index is out of bounds.
func (this *GoToSocialCanReplyProperty) SetType(idx int, t vocab.Type) error {
	n := &GoToSocialCanReplyPropertyIterator{
		alias:  this.alias,
		myIdx:  idx,
		parent: this,
	}
	if err := n.SetType(t); err != nil {
		return err
	}
	(this.properties)[idx] = n
	return nil
}

// Swap swaps the location of values at two indices for the "canReply" property.
func (this GoToSocialCanReplyProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}
