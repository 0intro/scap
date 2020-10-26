// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://www.w3.org/2000/09/xmldsig#
package xml_dsig

import (
	"encoding/xml"
)

// Element
type Signature struct {
	XMLName xml.Name `xml:Signature`

	Id string `xml:"Id,attr,omitempty"`

	SignedInfo SignedInfo `xml:"SignedInfo"`

	SignatureValue SignatureValue `xml:"SignatureValue"`

	KeyInfo *KeyInfo `xml:"KeyInfo"`

	Object []Object `xml:"Object"`
}

// Element
type SignatureValue struct {
	XMLName xml.Name `xml:SignatureValue`

	Id string `xml:"Id,attr,omitempty"`

	Text string `xml:",chardata"`
}

// Element
type SignedInfo struct {
	XMLName xml.Name `xml:SignedInfo`

	Id string `xml:"Id,attr,omitempty"`

	CanonicalizationMethod CanonicalizationMethod `xml:"CanonicalizationMethod"`

	SignatureMethod SignatureMethod `xml:"SignatureMethod"`

	Reference []Reference `xml:"Reference"`
}

// Element
type CanonicalizationMethod struct {
	XMLName xml.Name `xml:CanonicalizationMethod`

	Algorithm string `xml:"Algorithm,attr"`
}

// Element
type SignatureMethod struct {
	XMLName xml.Name `xml:SignatureMethod`

	Algorithm string `xml:"Algorithm,attr"`

	HMACOutputLength string `xml:"HMACOutputLength"`
}

// Element
type Reference struct {
	XMLName xml.Name `xml:Reference`

	Id string `xml:"Id,attr,omitempty"`

	URI string `xml:"URI,attr,omitempty"`

	Type string `xml:"Type,attr,omitempty"`

	Transforms *Transforms `xml:"Transforms"`

	DigestMethod DigestMethod `xml:"DigestMethod"`

	DigestValue DigestValue `xml:"DigestValue"`
}

// Element
type Transforms struct {
	XMLName xml.Name `xml:Transforms`

	Transform []Transform `xml:"Transform"`
}

// Element
type Transform struct {
	XMLName xml.Name `xml:Transform`

	Algorithm string `xml:"Algorithm,attr"`

	XPath []string `xml:"XPath"`
}

// Element
type DigestMethod struct {
	XMLName xml.Name `xml:DigestMethod`

	Algorithm string `xml:"Algorithm,attr"`
}

// Element
type DigestValue struct {
	XMLName xml.Name `xml:DigestValue`

	Text string `xml:",chardata"`
}

// Element
type KeyInfo struct {
	XMLName xml.Name `xml:KeyInfo`

	Id string `xml:"Id,attr,omitempty"`

	KeyName []KeyName `xml:"KeyName"`

	KeyValue []KeyValue `xml:"KeyValue"`

	RetrievalMethod []RetrievalMethod `xml:"RetrievalMethod"`

	X509Data []X509Data `xml:"X509Data"`

	PGPData []PGPData `xml:"PGPData"`

	SPKIData []SPKIData `xml:"SPKIData"`

	MgmtData []MgmtData `xml:"MgmtData"`
}

// Element
type KeyName struct {
	XMLName xml.Name `xml:KeyName`

	Text string `xml:",chardata"`
}

// Element
type MgmtData struct {
	XMLName xml.Name `xml:MgmtData`

	Text string `xml:",chardata"`
}

// Element
type KeyValue struct {
	XMLName xml.Name `xml:KeyValue`

	DSAKeyValue *DSAKeyValue `xml:"DSAKeyValue"`

	RSAKeyValue *RSAKeyValue `xml:"RSAKeyValue"`
}

// Element
type RetrievalMethod struct {
	XMLName xml.Name `xml:RetrievalMethod`

	URI string `xml:"URI,attr"`

	Type string `xml:"Type,attr,omitempty"`

	Transforms *Transforms `xml:"Transforms"`
}

// Element
type X509Data struct {
	XMLName xml.Name `xml:X509Data`

	X509IssuerSerial *X509IssuerSerialType `xml:"X509IssuerSerial"`

	X509SKI string `xml:"X509SKI"`

	X509SubjectName string `xml:"X509SubjectName"`

	X509Certificate string `xml:"X509Certificate"`

	X509CRL string `xml:"X509CRL"`
}

// Element
type PGPData struct {
	XMLName xml.Name `xml:PGPData`
}

// Element
type SPKIData struct {
	XMLName xml.Name `xml:SPKIData`

	SPKISexp string `xml:"SPKISexp"`
}

// Element
type Object struct {
	XMLName xml.Name `xml:Object`

	Id string `xml:"Id,attr,omitempty"`

	MimeType string `xml:"MimeType,attr,omitempty"`

	Encoding string `xml:"Encoding,attr,omitempty"`
}

// Element
type Manifest struct {
	XMLName xml.Name `xml:Manifest`

	Id string `xml:"Id,attr,omitempty"`

	Reference []Reference `xml:"Reference"`
}

// Element
type SignatureProperties struct {
	XMLName xml.Name `xml:SignatureProperties`

	Id string `xml:"Id,attr,omitempty"`

	SignatureProperty []SignatureProperty `xml:"SignatureProperty"`
}

// Element
type SignatureProperty struct {
	XMLName xml.Name `xml:SignatureProperty`

	Target string `xml:"Target,attr"`

	Id string `xml:"Id,attr,omitempty"`
}

// Element
type DSAKeyValue struct {
	XMLName xml.Name `xml:DSAKeyValue`

	G string `xml:"G"`

	Y string `xml:"Y"`

	J string `xml:"J"`
}

// Element
type RSAKeyValue struct {
	XMLName xml.Name `xml:RSAKeyValue`

	Modulus string `xml:"Modulus"`

	Exponent string `xml:"Exponent"`
}

// XSD ComplexType declarations

type SignatureType struct {
	Id string `xml:"Id,attr,omitempty"`

	SignedInfo SignedInfo `xml:"SignedInfo"`

	SignatureValue SignatureValue `xml:"SignatureValue"`

	KeyInfo *KeyInfo `xml:"KeyInfo"`

	Object []Object `xml:"Object"`

	InnerXml string `xml:",innerxml"`
}

type SignatureValueType struct {
	Id string `xml:"Id,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type SignedInfoType struct {
	Id string `xml:"Id,attr,omitempty"`

	CanonicalizationMethod CanonicalizationMethod `xml:"CanonicalizationMethod"`

	SignatureMethod SignatureMethod `xml:"SignatureMethod"`

	Reference []Reference `xml:"Reference"`

	InnerXml string `xml:",innerxml"`
}

type CanonicalizationMethodType struct {
	Algorithm string `xml:"Algorithm,attr"`

	InnerXml string `xml:",innerxml"`
}

type SignatureMethodType struct {
	Algorithm string `xml:"Algorithm,attr"`

	HMACOutputLength string `xml:"HMACOutputLength"`

	InnerXml string `xml:",innerxml"`
}

type ReferenceType struct {
	Id string `xml:"Id,attr,omitempty"`

	URI string `xml:"URI,attr,omitempty"`

	Type string `xml:"Type,attr,omitempty"`

	Transforms *Transforms `xml:"Transforms"`

	DigestMethod DigestMethod `xml:"DigestMethod"`

	DigestValue DigestValue `xml:"DigestValue"`

	InnerXml string `xml:",innerxml"`
}

type TransformsType struct {
	Transform []Transform `xml:"Transform"`

	InnerXml string `xml:",innerxml"`
}

type TransformType struct {
	Algorithm string `xml:"Algorithm,attr"`

	XPath []string `xml:"XPath"`

	InnerXml string `xml:",innerxml"`
}

type DigestMethodType struct {
	Algorithm string `xml:"Algorithm,attr"`

	InnerXml string `xml:",innerxml"`
}

type KeyInfoType struct {
	Id string `xml:"Id,attr,omitempty"`

	KeyName []KeyName `xml:"KeyName"`

	KeyValue []KeyValue `xml:"KeyValue"`

	RetrievalMethod []RetrievalMethod `xml:"RetrievalMethod"`

	X509Data []X509Data `xml:"X509Data"`

	PGPData []PGPData `xml:"PGPData"`

	SPKIData []SPKIData `xml:"SPKIData"`

	MgmtData []MgmtData `xml:"MgmtData"`

	InnerXml string `xml:",innerxml"`
}

type KeyValueType struct {
	DSAKeyValue *DSAKeyValue `xml:"DSAKeyValue"`

	RSAKeyValue *RSAKeyValue `xml:"RSAKeyValue"`

	InnerXml string `xml:",innerxml"`
}

type RetrievalMethodType struct {
	URI string `xml:"URI,attr"`

	Type string `xml:"Type,attr,omitempty"`

	Transforms *Transforms `xml:"Transforms"`

	InnerXml string `xml:",innerxml"`
}

type X509DataType struct {
	X509IssuerSerial *X509IssuerSerialType `xml:"X509IssuerSerial"`

	X509SKI string `xml:"X509SKI"`

	X509SubjectName string `xml:"X509SubjectName"`

	X509Certificate string `xml:"X509Certificate"`

	X509CRL string `xml:"X509CRL"`

	InnerXml string `xml:",innerxml"`
}

type X509IssuerSerialType struct {
	X509IssuerName string `xml:"X509IssuerName"`

	X509SerialNumber string `xml:"X509SerialNumber"`

	InnerXml string `xml:",innerxml"`
}

type PGPDataType struct {
	InnerXml string `xml:",innerxml"`
}

type SPKIDataType struct {
	SPKISexp string `xml:"SPKISexp"`

	InnerXml string `xml:",innerxml"`
}

type ObjectType struct {
	Id string `xml:"Id,attr,omitempty"`

	MimeType string `xml:"MimeType,attr,omitempty"`

	Encoding string `xml:"Encoding,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type ManifestType struct {
	Id string `xml:"Id,attr,omitempty"`

	Reference []Reference `xml:"Reference"`

	InnerXml string `xml:",innerxml"`
}

type SignaturePropertiesType struct {
	Id string `xml:"Id,attr,omitempty"`

	SignatureProperty []SignatureProperty `xml:"SignatureProperty"`

	InnerXml string `xml:",innerxml"`
}

type SignaturePropertyType struct {
	Target string `xml:"Target,attr"`

	Id string `xml:"Id,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type DSAKeyValueType struct {
	G string `xml:"G"`

	Y string `xml:"Y"`

	J string `xml:"J"`

	InnerXml string `xml:",innerxml"`
}

type RSAKeyValueType struct {
	Modulus string `xml:"Modulus"`

	Exponent string `xml:"Exponent"`

	InnerXml string `xml:",innerxml"`
}
