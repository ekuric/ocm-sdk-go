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

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalOrganization writes a value of the 'organization' type to the given writer.
func MarshalOrganization(object *Organization, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeOrganization(object, stream)
	stream.Flush()
	return stream.Error
}

// writeOrganization writes a value of the 'organization' type to the given stream.
func writeOrganization(object *Organization, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	if count > 0 {
		stream.WriteMore()
	}
	stream.WriteObjectField("kind")
	if object.link {
		stream.WriteString(OrganizationLinkKind)
	} else {
		stream.WriteString(OrganizationKind)
	}
	count++
	if object.id != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		stream.WriteString(*object.id)
		count++
	}
	if object.href != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("href")
		stream.WriteString(*object.href)
		count++
	}
	if object.externalID != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("external_id")
		stream.WriteString(*object.externalID)
		count++
	}
	if object.name != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("name")
		stream.WriteString(*object.name)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalOrganization reads a value of the 'organization' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalOrganization(source interface{}) (object *Organization, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readOrganization(iterator)
	err = iterator.Error
	return
}

// readOrganization reads a value of the 'organization' type from the given iterator.
func readOrganization(iterator *jsoniter.Iterator) *Organization {
	object := &Organization{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "kind":
			value := iterator.ReadString()
			object.link = value == OrganizationLinkKind
		case "id":
			value := iterator.ReadString()
			object.id = &value
		case "href":
			value := iterator.ReadString()
			object.href = &value
		case "external_id":
			value := iterator.ReadString()
			object.externalID = &value
		case "name":
			value := iterator.ReadString()
			object.name = &value
		default:
			iterator.ReadAny()
		}
	}
	return object
}
