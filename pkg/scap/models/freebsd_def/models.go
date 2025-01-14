// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#freebsd
package freebsd_def

import (
	"encoding/xml"
	"github.com/0intro/scap/pkg/scap/models/oval"
	"github.com/0intro/scap/pkg/scap/models/oval_def"
	"github.com/0intro/scap/pkg/scap/models/xml_dsig"
)

// Element
type PortinfoTest struct {
	XMLName xml.Name `xml:"portinfo_test"`

	Id oval.TestIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	CheckExistence oval.ExistenceEnumeration `xml:"check_existence,attr,omitempty"`

	Check oval.CheckEnumeration `xml:"check,attr"`

	StateOperator oval.OperatorEnumeration `xml:"state_operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type PortinfoObject struct {
	XMLName xml.Name `xml:"portinfo_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Pkginst *oval_def.EntityObjectStringType `xml:"pkginst"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type PortinfoState struct {
	XMLName xml.Name `xml:"portinfo_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Pkginst *oval_def.EntityStateStringType `xml:"pkginst"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Category *oval_def.EntityStateStringType `xml:"category"`

	VersionElm *PortinfoStateVersion `xml:"version"`

	Vendor *oval_def.EntityStateStringType `xml:"vendor"`

	Description *oval_def.EntityStateStringType `xml:"description"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type PortinfoStateVersion struct {
	XMLName xml.Name `xml:"version"`

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`
}

// XSD ComplexType declarations

// XSD SimpleType declarations
