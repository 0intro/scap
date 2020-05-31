// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://www.w3.org/2000/09/xmldsig#
package xml_dsig

import (
	"encoding/xml"
)

// Element
type Signature struct {
	XMLName xml.Name `xml:Signature`

	SignedInfo SignedInfo `xml:"SignedInfo"`

	SignatureValue SignatureValue `xml:"SignatureValue"`

	KeyInfo *KeyInfo `xml:"KeyInfo"`

	Object []Object `xml:"Object"`
}

// Element
type SignatureValue struct {
	XMLName xml.Name `xml:SignatureValue`
}

// Element
type SignedInfo struct {
	XMLName xml.Name `xml:SignedInfo`

	CanonicalizationMethod CanonicalizationMethod `xml:"CanonicalizationMethod"`

	SignatureMethod SignatureMethod `xml:"SignatureMethod"`

	Reference []Reference `xml:"Reference"`
}

// Element
type CanonicalizationMethod struct {
	XMLName xml.Name `xml:CanonicalizationMethod`
}

// Element
type SignatureMethod struct {
	XMLName xml.Name `xml:SignatureMethod`

	HMACOutputLength string `xml:"HMACOutputLength"`
}

// Element
type Reference struct {
	XMLName xml.Name `xml:Reference`

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
}

// Element
type DigestMethod struct {
	XMLName xml.Name `xml:DigestMethod`
}

// Element
type DigestValue struct {
	XMLName xml.Name `xml:DigestValue`
}

// Element
type KeyInfo struct {
	XMLName xml.Name `xml:KeyInfo`
}

// Element
type KeyName struct {
	XMLName xml.Name `xml:KeyName`
}

// Element
type MgmtData struct {
	XMLName xml.Name `xml:MgmtData`
}

// Element
type KeyValue struct {
	XMLName xml.Name `xml:KeyValue`
}

// Element
type RetrievalMethod struct {
	XMLName xml.Name `xml:RetrievalMethod`

	Transforms *Transforms `xml:"Transforms"`
}

// Element
type X509Data struct {
	XMLName xml.Name `xml:X509Data`

	X509IssuerSerial X509IssuerSerialType `xml:"X509IssuerSerial"`

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
}

// Element
type Manifest struct {
	XMLName xml.Name `xml:Manifest`

	Reference []Reference `xml:"Reference"`
}

// Element
type SignatureProperties struct {
	XMLName xml.Name `xml:SignatureProperties`

	SignatureProperty []SignatureProperty `xml:"SignatureProperty"`
}

// Element
type SignatureProperty struct {
	XMLName xml.Name `xml:SignatureProperty`
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
	Id string `xml:"Id,attr"`

	SignedInfo SignedInfo `xml:"SignedInfo"`

	SignatureValue SignatureValue `xml:"SignatureValue"`

	KeyInfo *KeyInfo `xml:"KeyInfo"`

	Object []Object `xml:"Object"`
}

type SignatureValueType struct {
	Id string `xml:"Id,attr"`
}

type SignedInfoType struct {
	Id string `xml:"Id,attr"`

	CanonicalizationMethod CanonicalizationMethod `xml:"CanonicalizationMethod"`

	SignatureMethod SignatureMethod `xml:"SignatureMethod"`

	Reference []Reference `xml:"Reference"`
}

type CanonicalizationMethodType struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type SignatureMethodType struct {
	Algorithm string `xml:"Algorithm,attr"`

	HMACOutputLength string `xml:"HMACOutputLength"`
}

type ReferenceType struct {
	Id string `xml:"Id,attr"`

	URI string `xml:"URI,attr"`

	Type string `xml:"Type,attr"`

	Transforms *Transforms `xml:"Transforms"`

	DigestMethod DigestMethod `xml:"DigestMethod"`

	DigestValue DigestValue `xml:"DigestValue"`
}

type TransformsType struct {
	Transform []Transform `xml:"Transform"`
}

type TransformType struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type DigestMethodType struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type KeyInfoType struct {
	Id string `xml:"Id,attr"`
}

type KeyValueType struct {
}

type RetrievalMethodType struct {
	URI string `xml:"URI,attr"`

	Type string `xml:"Type,attr"`

	Transforms *Transforms `xml:"Transforms"`
}

type X509DataType struct {
	X509IssuerSerial X509IssuerSerialType `xml:"X509IssuerSerial"`

	X509SKI string `xml:"X509SKI"`

	X509SubjectName string `xml:"X509SubjectName"`

	X509Certificate string `xml:"X509Certificate"`

	X509CRL string `xml:"X509CRL"`
}

type X509IssuerSerialType struct {
	X509IssuerName string `xml:"X509IssuerName"`

	X509SerialNumber string `xml:"X509SerialNumber"`
}

type PGPDataType struct {
}

type SPKIDataType struct {
	SPKISexp string `xml:"SPKISexp"`
}

type ObjectType struct {
	Id string `xml:"Id,attr"`

	MimeType string `xml:"MimeType,attr"`

	Encoding string `xml:"Encoding,attr"`
}

type ManifestType struct {
	Id string `xml:"Id,attr"`

	Reference []Reference `xml:"Reference"`
}

type SignaturePropertiesType struct {
	Id string `xml:"Id,attr"`

	SignatureProperty []SignatureProperty `xml:"SignatureProperty"`
}

type SignaturePropertyType struct {
	Target string `xml:"Target,attr"`

	Id string `xml:"Id,attr"`
}

type DSAKeyValueType struct {
	G string `xml:"G"`

	Y string `xml:"Y"`

	J string `xml:"J"`
}

type RSAKeyValueType struct {
	Modulus string `xml:"Modulus"`

	Exponent string `xml:"Exponent"`
}
