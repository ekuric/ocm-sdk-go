/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// AWSInfrastructureAccessRoleKind is the name of the type used to represent objects
// of type 'AWS_infrastructure_access_role'.
const AWSInfrastructureAccessRoleKind = "AWSInfrastructureAccessRole"

// AWSInfrastructureAccessRoleLinkKind is the name of the type used to represent links
// to objects of type 'AWS_infrastructure_access_role'.
const AWSInfrastructureAccessRoleLinkKind = "AWSInfrastructureAccessRoleLink"

// AWSInfrastructureAccessRoleNilKind is the name of the type used to nil references
// to objects of type 'AWS_infrastructure_access_role'.
const AWSInfrastructureAccessRoleNilKind = "AWSInfrastructureAccessRoleNil"

// AWSInfrastructureAccessRole represents the values of the 'AWS_infrastructure_access_role' type.
//
// A set of acces permissions for AWS resources
type AWSInfrastructureAccessRole struct {
	id          *string
	href        *string
	link        bool
	description *string
	displayName *string
}

// Kind returns the name of the type of the object.
func (o *AWSInfrastructureAccessRole) Kind() string {
	if o == nil {
		return AWSInfrastructureAccessRoleNilKind
	}
	if o.link {
		return AWSInfrastructureAccessRoleLinkKind
	}
	return AWSInfrastructureAccessRoleKind
}

// ID returns the identifier of the object.
func (o *AWSInfrastructureAccessRole) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *AWSInfrastructureAccessRole) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *AWSInfrastructureAccessRole) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *AWSInfrastructureAccessRole) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *AWSInfrastructureAccessRole) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *AWSInfrastructureAccessRole) Empty() bool {
	return o == nil || (o.id == nil &&
		o.description == nil &&
		o.displayName == nil &&
		true)
}

// Description returns the value of the 'description' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Description of the role.
func (o *AWSInfrastructureAccessRole) Description() string {
	if o != nil && o.description != nil {
		return *o.description
	}
	return ""
}

// GetDescription returns the value of the 'description' attribute and
// a flag indicating if the attribute has a value.
//
// Description of the role.
func (o *AWSInfrastructureAccessRole) GetDescription() (value string, ok bool) {
	ok = o != nil && o.description != nil
	if ok {
		value = *o.description
	}
	return
}

// DisplayName returns the value of the 'display_name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Human friendly identifier of the role, for example `Read only`.
func (o *AWSInfrastructureAccessRole) DisplayName() string {
	if o != nil && o.displayName != nil {
		return *o.displayName
	}
	return ""
}

// GetDisplayName returns the value of the 'display_name' attribute and
// a flag indicating if the attribute has a value.
//
// Human friendly identifier of the role, for example `Read only`.
func (o *AWSInfrastructureAccessRole) GetDisplayName() (value string, ok bool) {
	ok = o != nil && o.displayName != nil
	if ok {
		value = *o.displayName
	}
	return
}

// AWSInfrastructureAccessRoleListKind is the name of the type used to represent list of objects of
// type 'AWS_infrastructure_access_role'.
const AWSInfrastructureAccessRoleListKind = "AWSInfrastructureAccessRoleList"

// AWSInfrastructureAccessRoleListLinkKind is the name of the type used to represent links to list
// of objects of type 'AWS_infrastructure_access_role'.
const AWSInfrastructureAccessRoleListLinkKind = "AWSInfrastructureAccessRoleListLink"

// AWSInfrastructureAccessRoleNilKind is the name of the type used to nil lists of objects of
// type 'AWS_infrastructure_access_role'.
const AWSInfrastructureAccessRoleListNilKind = "AWSInfrastructureAccessRoleListNil"

// AWSInfrastructureAccessRoleList is a list of values of the 'AWS_infrastructure_access_role' type.
type AWSInfrastructureAccessRoleList struct {
	href  *string
	link  bool
	items []*AWSInfrastructureAccessRole
}

// Kind returns the name of the type of the object.
func (l *AWSInfrastructureAccessRoleList) Kind() string {
	if l == nil {
		return AWSInfrastructureAccessRoleListNilKind
	}
	if l.link {
		return AWSInfrastructureAccessRoleListLinkKind
	}
	return AWSInfrastructureAccessRoleListKind
}

// Link returns true iif this is a link.
func (l *AWSInfrastructureAccessRoleList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *AWSInfrastructureAccessRoleList) HREF() string {
	if l != nil && l.href != nil {
		return *l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *AWSInfrastructureAccessRoleList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != nil
	if ok {
		value = *l.href
	}
	return
}

// Len returns the length of the list.
func (l *AWSInfrastructureAccessRoleList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *AWSInfrastructureAccessRoleList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *AWSInfrastructureAccessRoleList) Get(i int) *AWSInfrastructureAccessRole {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *AWSInfrastructureAccessRoleList) Slice() []*AWSInfrastructureAccessRole {
	var slice []*AWSInfrastructureAccessRole
	if l == nil {
		slice = make([]*AWSInfrastructureAccessRole, 0)
	} else {
		slice = make([]*AWSInfrastructureAccessRole, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *AWSInfrastructureAccessRoleList) Each(f func(item *AWSInfrastructureAccessRole) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *AWSInfrastructureAccessRoleList) Range(f func(index int, item *AWSInfrastructureAccessRole) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
