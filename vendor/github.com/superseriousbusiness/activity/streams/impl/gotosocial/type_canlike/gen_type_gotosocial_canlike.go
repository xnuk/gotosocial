// Code generated by astool. DO NOT EDIT.

package typecanlike

import vocab "github.com/superseriousbusiness/activity/streams/vocab"

type GoToSocialCanLike struct {
	GoToSocialAlways           vocab.GoToSocialAlwaysProperty
	GoToSocialApprovalRequired vocab.GoToSocialApprovalRequiredProperty
	JSONLDId                   vocab.JSONLDIdProperty
	alias                      string
	unknown                    map[string]interface{}
}

// CanLikeIsDisjointWith returns true if the other provided type is disjoint with
// the CanLike type.
func CanLikeIsDisjointWith(other vocab.Type) bool {
	// Shortcut implementation: is not disjoint with anything.
	return false
}

// CanLikeIsExtendedBy returns true if the other provided type extends from the
// CanLike type. Note that it returns false if the types are the same; see the
// "IsOrExtendsCanLike" variant instead.
func CanLikeIsExtendedBy(other vocab.Type) bool {
	// Shortcut implementation: is not extended by anything.
	return false
}

// DeserializeCanLike creates a CanLike from a map representation that has been
// unmarshalled from a text or binary format.
func DeserializeCanLike(m map[string]interface{}, aliasMap map[string]string) (*GoToSocialCanLike, error) {
	alias := ""
	if a, ok := aliasMap["https://gotosocial.org/ns"]; ok {
		alias = a
	}
	this := &GoToSocialCanLike{
		alias:   alias,
		unknown: make(map[string]interface{}),
	}

	// Begin: Known property deserialization
	if p, err := mgr.DeserializeAlwaysPropertyGoToSocial()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.GoToSocialAlways = p
	}
	if p, err := mgr.DeserializeApprovalRequiredPropertyGoToSocial()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.GoToSocialApprovalRequired = p
	}
	if p, err := mgr.DeserializeIdPropertyJSONLD()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.JSONLDId = p
	}
	// End: Known property deserialization

	// Begin: Unknown deserialization
	for k, v := range m {
		// Begin: Code that ensures a property name is unknown
		if k == "always" {
			continue
		} else if k == "approvalRequired" {
			continue
		} else if k == "id" {
			continue
		} // End: Code that ensures a property name is unknown

		this.unknown[k] = v
	}
	// End: Unknown deserialization

	return this, nil
}

// GoToSocialCanLikeExtends returns true if the CanLike type extends from the
// other type.
func GoToSocialCanLikeExtends(other vocab.Type) bool {
	// Shortcut implementation: this does not extend anything.
	return false
}

// IsOrExtendsCanLike returns true if the other provided type is the CanLike type
// or extends from the CanLike type.
func IsOrExtendsCanLike(other vocab.Type) bool {
	if other.GetTypeName() == "CanLike" {
		return true
	}
	return CanLikeIsExtendedBy(other)
}

// NewGoToSocialCanLike creates a new CanLike type
func NewGoToSocialCanLike() *GoToSocialCanLike {
	return &GoToSocialCanLike{
		alias:   "",
		unknown: make(map[string]interface{}),
	}
}

// GetGoToSocialAlways returns the "always" property if it exists, and nil
// otherwise.
func (this GoToSocialCanLike) GetGoToSocialAlways() vocab.GoToSocialAlwaysProperty {
	return this.GoToSocialAlways
}

// GetGoToSocialApprovalRequired returns the "approvalRequired" property if it
// exists, and nil otherwise.
func (this GoToSocialCanLike) GetGoToSocialApprovalRequired() vocab.GoToSocialApprovalRequiredProperty {
	return this.GoToSocialApprovalRequired
}

// GetJSONLDId returns the "id" property if it exists, and nil otherwise.
func (this GoToSocialCanLike) GetJSONLDId() vocab.JSONLDIdProperty {
	return this.JSONLDId
}

// GetTypeName returns the name of this type.
func (this GoToSocialCanLike) GetTypeName() string {
	return "CanLike"
}

// GetUnknownProperties returns the unknown properties for the CanLike type. Note
// that this should not be used by app developers. It is only used to help
// determine which implementation is LessThan the other. Developers who are
// creating a different implementation of this type's interface can use this
// method in their LessThan implementation, but routine ActivityPub
// applications should not use this to bypass the code generation tool.
func (this GoToSocialCanLike) GetUnknownProperties() map[string]interface{} {
	return this.unknown
}

// IsExtending returns true if the CanLike type extends from the other type.
func (this GoToSocialCanLike) IsExtending(other vocab.Type) bool {
	return GoToSocialCanLikeExtends(other)
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// type and the specific properties that are set. The value in the map is the
// alias used to import the type and its properties.
func (this GoToSocialCanLike) JSONLDContext() map[string]string {
	m := map[string]string{"https://gotosocial.org/ns": this.alias}
	m = this.helperJSONLDContext(this.GoToSocialAlways, m)
	m = this.helperJSONLDContext(this.GoToSocialApprovalRequired, m)
	m = this.helperJSONLDContext(this.JSONLDId, m)

	return m
}

// LessThan computes if this CanLike is lesser, with an arbitrary but stable
// determination.
func (this GoToSocialCanLike) LessThan(o vocab.GoToSocialCanLike) bool {
	// Begin: Compare known properties
	// Compare property "always"
	if lhs, rhs := this.GoToSocialAlways, o.GetGoToSocialAlways(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "approvalRequired"
	if lhs, rhs := this.GoToSocialApprovalRequired, o.GetGoToSocialApprovalRequired(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "id"
	if lhs, rhs := this.JSONLDId, o.GetJSONLDId(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// End: Compare known properties

	// Begin: Compare unknown properties (only by number of them)
	if len(this.unknown) < len(o.GetUnknownProperties()) {
		return true
	} else if len(o.GetUnknownProperties()) < len(this.unknown) {
		return false
	} // End: Compare unknown properties (only by number of them)

	// All properties are the same.
	return false
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format.
func (this GoToSocialCanLike) Serialize() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	// Begin: Serialize known properties
	// Maybe serialize property "always"
	if this.GoToSocialAlways != nil {
		if i, err := this.GoToSocialAlways.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.GoToSocialAlways.Name()] = i
		}
	}
	// Maybe serialize property "approvalRequired"
	if this.GoToSocialApprovalRequired != nil {
		if i, err := this.GoToSocialApprovalRequired.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.GoToSocialApprovalRequired.Name()] = i
		}
	}
	// Maybe serialize property "id"
	if this.JSONLDId != nil {
		if i, err := this.JSONLDId.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.JSONLDId.Name()] = i
		}
	}
	// End: Serialize known properties

	// Begin: Serialize unknown properties
	for k, v := range this.unknown {
		// To be safe, ensure we aren't overwriting a known property
		if _, has := m[k]; !has {
			m[k] = v
		}
	}
	// End: Serialize unknown properties

	return m, nil
}

// SetGoToSocialAlways sets the "always" property.
func (this *GoToSocialCanLike) SetGoToSocialAlways(i vocab.GoToSocialAlwaysProperty) {
	this.GoToSocialAlways = i
}

// SetGoToSocialApprovalRequired sets the "approvalRequired" property.
func (this *GoToSocialCanLike) SetGoToSocialApprovalRequired(i vocab.GoToSocialApprovalRequiredProperty) {
	this.GoToSocialApprovalRequired = i
}

// SetJSONLDId sets the "id" property.
func (this *GoToSocialCanLike) SetJSONLDId(i vocab.JSONLDIdProperty) {
	this.JSONLDId = i
}

// VocabularyURI returns the vocabulary's URI as a string.
func (this GoToSocialCanLike) VocabularyURI() string {
	return "https://gotosocial.org/ns"
}

// helperJSONLDContext obtains the context uris and their aliases from a property,
// if it is not nil.
func (this GoToSocialCanLike) helperJSONLDContext(i jsonldContexter, toMerge map[string]string) map[string]string {
	if i == nil {
		return toMerge
	}
	for k, v := range i.JSONLDContext() {
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		toMerge[k] = v
	}
	return toMerge
}
