// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5#aix
package aix_sc

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_sc"
)

// Element
type InterimFixItem struct {
	XMLName xml.Name `xml:"interim_fix_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Vuid *oval_sc.EntityItemStringType `xml:"vuid"`

	Label *oval_sc.EntityItemStringType `xml:"label"`

	Abstract *oval_sc.EntityItemStringType `xml:"abstract"`

	State *EntityItemInterimFixStateType `xml:"state"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type FilesetItem struct {
	XMLName xml.Name `xml:"fileset_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Flstinst *oval_sc.EntityItemStringType `xml:"flstinst"`

	Level *oval_sc.EntityItemVersionType `xml:"level"`

	State *EntityItemFilesetStateType `xml:"state"`

	Description *oval_sc.EntityItemStringType `xml:"description"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type FixItem struct {
	XMLName xml.Name `xml:"fix_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	AparNumber *oval_sc.EntityItemStringType `xml:"apar_number"`

	Abstract *oval_sc.EntityItemStringType `xml:"abstract"`

	Symptom *oval_sc.EntityItemStringType `xml:"symptom"`

	InstallationStatus *EntityItemFixInstallationStatusType `xml:"installation_status"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type NoItem struct {
	XMLName xml.Name `xml:"no_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	Tunable *oval_sc.EntityItemStringType `xml:"tunable"`

	Value *oval_sc.EntityItemAnySimpleType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type OslevelItem struct {
	XMLName xml.Name `xml:"oslevel_item"`

	Id oval.ItemIDPattern `xml:"id,attr"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	MaintenanceLevel *oval_sc.EntityItemVersionType `xml:"maintenance_level"`

	Message []oval.MessageType `xml:"message"`
}

// XSD ComplexType declarations

type EntityItemFilesetStateType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemFixInstallationStatusType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityItemInterimFixStateType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	Status oval_sc.StatusEnumeration `xml:"status,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations
