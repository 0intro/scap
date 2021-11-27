// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5#independent
package ind_sc

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_sc"
)

// Element
type FamilyItem struct {
	XMLName xml.Name `xml:"family_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Family *EntityItemFamilyType `xml:"family"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type FilehashItem struct {
	XMLName xml.Name `xml:"filehash_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Filepath *oval_sc.EntityItemStringType `xml:"filepath"`

	Path *oval_sc.EntityItemStringType `xml:"path"`

	Filename *oval_sc.EntityItemStringType `xml:"filename"`

	Md5 *oval_sc.EntityItemStringType `xml:"md5"`

	Sha1 *oval_sc.EntityItemStringType `xml:"sha1"`

	WindowsView *EntityItemWindowsViewType `xml:"windows_view"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type Filehash58Item struct {
	XMLName xml.Name `xml:"filehash58_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Filepath *oval_sc.EntityItemStringType `xml:"filepath"`

	Path *oval_sc.EntityItemStringType `xml:"path"`

	Filename *oval_sc.EntityItemStringType `xml:"filename"`

	HashType *EntityItemHashTypeType `xml:"hash_type"`

	Hash *oval_sc.EntityItemStringType `xml:"hash"`

	WindowsView *EntityItemWindowsViewType `xml:"windows_view"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type EnvironmentvariableItem struct {
	XMLName xml.Name `xml:"environmentvariable_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Value *oval_sc.EntityItemAnySimpleType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type Environmentvariable58Item struct {
	XMLName xml.Name `xml:"environmentvariable58_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Pid *oval_sc.EntityItemIntType `xml:"pid"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Value *oval_sc.EntityItemAnySimpleType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type LdapItem struct {
	XMLName xml.Name `xml:"ldap_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Suffix *oval_sc.EntityItemStringType `xml:"suffix"`

	RelativeDn *oval_sc.EntityItemStringType `xml:"relative_dn"`

	Attribute *oval_sc.EntityItemStringType `xml:"attribute"`

	ObjectClass *oval_sc.EntityItemStringType `xml:"object_class"`

	Ldaptype *EntityItemLdaptypeType `xml:"ldaptype"`

	Value []oval_sc.EntityItemAnySimpleType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type Ldap57Item struct {
	XMLName xml.Name `xml:"ldap57_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Suffix *oval_sc.EntityItemStringType `xml:"suffix"`

	RelativeDn *oval_sc.EntityItemStringType `xml:"relative_dn"`

	Attribute *oval_sc.EntityItemStringType `xml:"attribute"`

	ObjectClass *oval_sc.EntityItemStringType `xml:"object_class"`

	Ldaptype *EntityItemLdaptypeType `xml:"ldaptype"`

	Value []oval_sc.EntityItemRecordType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SqlItem struct {
	XMLName xml.Name `xml:"sql_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Engine *EntityItemEngineType `xml:"engine"`

	Version *oval_sc.EntityItemStringType `xml:"version"`

	ConnectionString *oval_sc.EntityItemStringType `xml:"connection_string"`

	Sql *oval_sc.EntityItemStringType `xml:"sql"`

	Result []oval_sc.EntityItemAnySimpleType `xml:"result"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type Sql57Item struct {
	XMLName xml.Name `xml:"sql57_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Engine *EntityItemEngineType `xml:"engine"`

	Version *oval_sc.EntityItemStringType `xml:"version"`

	ConnectionString *oval_sc.EntityItemStringType `xml:"connection_string"`

	Sql *oval_sc.EntityItemStringType `xml:"sql"`

	Result []oval_sc.EntityItemRecordType `xml:"result"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type TextfilecontentItem struct {
	XMLName xml.Name `xml:"textfilecontent_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Filepath *oval_sc.EntityItemStringType `xml:"filepath"`

	Path *oval_sc.EntityItemStringType `xml:"path"`

	Filename *oval_sc.EntityItemStringType `xml:"filename"`

	Pattern *oval_sc.EntityItemStringType `xml:"pattern"`

	Instance *oval_sc.EntityItemIntType `xml:"instance"`

	Line *oval_sc.EntityItemStringType `xml:"line"`

	Text *oval_sc.EntityItemAnySimpleType `xml:"text"`

	Subexpression []oval_sc.EntityItemAnySimpleType `xml:"subexpression"`

	WindowsView *EntityItemWindowsViewType `xml:"windows_view"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type VariableItem struct {
	XMLName xml.Name `xml:"variable_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	VarRef *EntityItemVariableRefType `xml:"var_ref"`

	Value []oval_sc.EntityItemAnySimpleType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type XmlfilecontentItem struct {
	XMLName xml.Name `xml:"xmlfilecontent_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Filepath *oval_sc.EntityItemStringType `xml:"filepath"`

	Path *oval_sc.EntityItemStringType `xml:"path"`

	Filename *oval_sc.EntityItemStringType `xml:"filename"`

	Xpath *oval_sc.EntityItemStringType `xml:"xpath"`

	ValueOf []oval_sc.EntityItemAnySimpleType `xml:"value_of"`

	WindowsView *EntityItemWindowsViewType `xml:"windows_view"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type YamlfilecontentItem struct {
	XMLName xml.Name `xml:"yamlfilecontent_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Filepath *oval_sc.EntityItemStringType `xml:"filepath"`

	Path *oval_sc.EntityItemStringType `xml:"path"`

	Filename *oval_sc.EntityItemStringType `xml:"filename"`

	Yamlpath *oval_sc.EntityItemStringType `xml:"yamlpath"`

	Value []oval_sc.EntityItemRecordType `xml:"value"`

	WindowsView *EntityItemWindowsViewType `xml:"windows_view"`

	Message []oval.MessageType `xml:"message"`
}

// XSD ComplexType declarations

type EntityItemEngineType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemFamilyType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemHashTypeType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemVariableRefType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemLdaptypeType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemWindowsViewType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations
