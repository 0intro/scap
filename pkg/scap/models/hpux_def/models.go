// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
package hpux_def

import (
	"encoding/xml"
)

// Element
type GetconfTest struct {
	XMLName xml.Name `xml:getconf_test`
}

// Element
type GetconfObject struct {
	XMLName xml.Name `xml:getconf_object`
}

// Element
type GetconfState struct {
	XMLName xml.Name `xml:getconf_state`
}

// Element
type NddTest struct {
	XMLName xml.Name `xml:ndd_test`
}

// Element
type NddObject struct {
	XMLName xml.Name `xml:ndd_object`
}

// Element
type NddState struct {
	XMLName xml.Name `xml:ndd_state`
}

// Element
type Patch53Test struct {
	XMLName xml.Name `xml:patch53_test`
}

// Element
type Patch53Object struct {
	XMLName xml.Name `xml:patch53_object`
}

// Element
type Patch53State struct {
	XMLName xml.Name `xml:patch53_state`
}

// Element
type PatchTest struct {
	XMLName xml.Name `xml:patch_test`
}

// Element
type PatchObject struct {
	XMLName xml.Name `xml:patch_object`
}

// Element
type PatchState struct {
	XMLName xml.Name `xml:patch_state`
}

// Element
type SwlistTest struct {
	XMLName xml.Name `xml:swlist_test`
}

// Element
type SwlistObject struct {
	XMLName xml.Name `xml:swlist_object`
}

// Element
type SwlistState struct {
	XMLName xml.Name `xml:swlist_state`
}

// Element
type TrustedTest struct {
	XMLName xml.Name `xml:trusted_test`
}

// Element
type TrustedObject struct {
	XMLName xml.Name `xml:trusted_object`
}

// Element
type TrustedState struct {
	XMLName xml.Name `xml:trusted_state`
}

// XSD ComplexType declarations

type Patch53Behaviors struct {
}
