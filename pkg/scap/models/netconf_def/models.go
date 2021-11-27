// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#netconf
package netconf_def

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_def"
	"github.com/gocomply/scap/pkg/scap/models/xml_dsig"
)

// Element
type ConfigTest struct {
	XMLName xml.Name `xml:"config_test"`

	Id oval.TestIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

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
type ConfigObject struct {
	XMLName xml.Name `xml:"config_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Xpath *oval_def.EntityObjectStringType `xml:"xpath"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type ConfigState struct {
	XMLName xml.Name `xml:"config_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version int `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Xpath *oval_def.EntityStateStringType `xml:"xpath"`

	ValueOf *oval_def.EntityStateAnySimpleType `xml:"value_of"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// XSD ComplexType declarations

// XSD SimpleType declarations
