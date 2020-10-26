// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
package ind_def

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/oval_def"
	"github.com/gocomply/scap/pkg/scap/models/xml_dsig"
)

// Element
type FamilyTest struct {
	XMLName xml.Name `xml:family_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type FamilyObject struct {
	XMLName xml.Name `xml:family_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type FamilyState struct {
	XMLName xml.Name `xml:family_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Family *EntityStateFamilyType `xml:"family"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type FilehashTest struct {
	XMLName xml.Name `xml:filehash_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type FilehashObject struct {
	XMLName xml.Name `xml:filehash_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type FilehashState struct {
	XMLName xml.Name `xml:filehash_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	Md5 *oval_def.EntityStateStringType `xml:"md5"`

	Sha1 *oval_def.EntityStateStringType `xml:"sha1"`

	WindowsView *EntityStateWindowsViewType `xml:"windows_view"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Filehash58Test struct {
	XMLName xml.Name `xml:filehash58_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Filehash58Object struct {
	XMLName xml.Name `xml:filehash58_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Filehash58State struct {
	XMLName xml.Name `xml:filehash58_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	HashType *EntityStateHashTypeType `xml:"hash_type"`

	Hash *oval_def.EntityStateStringType `xml:"hash"`

	WindowsView *EntityStateWindowsViewType `xml:"windows_view"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type EnvironmentvariableTest struct {
	XMLName xml.Name `xml:environmentvariable_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type EnvironmentvariableObject struct {
	XMLName xml.Name `xml:environmentvariable_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type EnvironmentvariableState struct {
	XMLName xml.Name `xml:environmentvariable_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Environmentvariable58Test struct {
	XMLName xml.Name `xml:environmentvariable58_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Environmentvariable58Object struct {
	XMLName xml.Name `xml:environmentvariable58_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Environmentvariable58State struct {
	XMLName xml.Name `xml:environmentvariable58_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Pid *oval_def.EntityStateIntType `xml:"pid"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type LdapTest struct {
	XMLName xml.Name `xml:ldap_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type LdapObject struct {
	XMLName xml.Name `xml:ldap_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type LdapState struct {
	XMLName xml.Name `xml:ldap_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Suffix *oval_def.EntityStateStringType `xml:"suffix"`

	RelativeDn *oval_def.EntityStateStringType `xml:"relative_dn"`

	Attribute *oval_def.EntityStateStringType `xml:"attribute"`

	ObjectClass *oval_def.EntityStateStringType `xml:"object_class"`

	Ldaptype *EntityStateLdaptypeType `xml:"ldaptype"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Ldap57Test struct {
	XMLName xml.Name `xml:ldap57_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Ldap57Object struct {
	XMLName xml.Name `xml:ldap57_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Ldap57State struct {
	XMLName xml.Name `xml:ldap57_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Suffix *oval_def.EntityStateStringType `xml:"suffix"`

	RelativeDn *oval_def.EntityStateStringType `xml:"relative_dn"`

	Attribute *oval_def.EntityStateStringType `xml:"attribute"`

	ObjectClass *oval_def.EntityStateStringType `xml:"object_class"`

	Ldaptype *EntityStateLdaptypeType `xml:"ldaptype"`

	Value *oval_def.EntityStateRecordType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SqlTest struct {
	XMLName xml.Name `xml:sql_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SqlObject struct {
	XMLName xml.Name `xml:sql_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SqlState struct {
	XMLName xml.Name `xml:sql_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Engine *EntityStateEngineType `xml:"engine"`

	VersionElm *oval_def.EntityStateStringType `xml:"version"`

	ConnectionString *oval_def.EntityStateStringType `xml:"connection_string"`

	Sql *oval_def.EntityStateStringType `xml:"sql"`

	Result *oval_def.EntityStateAnySimpleType `xml:"result"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Sql57Test struct {
	XMLName xml.Name `xml:sql57_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Sql57Object struct {
	XMLName xml.Name `xml:sql57_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Sql57State struct {
	XMLName xml.Name `xml:sql57_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Engine *EntityStateEngineType `xml:"engine"`

	VersionElm *oval_def.EntityStateStringType `xml:"version"`

	ConnectionString *oval_def.EntityStateStringType `xml:"connection_string"`

	Sql *oval_def.EntityStateStringType `xml:"sql"`

	Result *oval_def.EntityStateRecordType `xml:"result"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Textfilecontent54Test struct {
	XMLName xml.Name `xml:textfilecontent54_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Textfilecontent54Object struct {
	XMLName xml.Name `xml:textfilecontent54_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Textfilecontent54State struct {
	XMLName xml.Name `xml:textfilecontent54_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	Pattern *oval_def.EntityStateStringType `xml:"pattern"`

	Instance *oval_def.EntityStateIntType `xml:"instance"`

	Text *oval_def.EntityStateAnySimpleType `xml:"text"`

	Subexpression *oval_def.EntityStateAnySimpleType `xml:"subexpression"`

	WindowsView *EntityStateWindowsViewType `xml:"windows_view"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type TextfilecontentTest struct {
	XMLName xml.Name `xml:textfilecontent_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type TextfilecontentObject struct {
	XMLName xml.Name `xml:textfilecontent_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type TextfilecontentState struct {
	XMLName xml.Name `xml:textfilecontent_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	Line *oval_def.EntityStateStringType `xml:"line"`

	Subexpression *oval_def.EntityStateAnySimpleType `xml:"subexpression"`

	WindowsView *EntityStateWindowsViewType `xml:"windows_view"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type UnknownTest struct {
	XMLName xml.Name `xml:unknown_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VariableTest struct {
	XMLName xml.Name `xml:variable_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VariableObject struct {
	XMLName xml.Name `xml:variable_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VariableState struct {
	XMLName xml.Name `xml:variable_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	VarRef *EntityStateVariableRefType `xml:"var_ref"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type XmlfilecontentTest struct {
	XMLName xml.Name `xml:xmlfilecontent_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type XmlfilecontentObject struct {
	XMLName xml.Name `xml:xmlfilecontent_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type XmlfilecontentState struct {
	XMLName xml.Name `xml:xmlfilecontent_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	Xpath *oval_def.EntityStateStringType `xml:"xpath"`

	ValueOf *oval_def.EntityStateAnySimpleType `xml:"value_of"`

	WindowsView *EntityStateWindowsViewType `xml:"windows_view"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type YamlfilecontentTest struct {
	XMLName xml.Name `xml:yamlfilecontent_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr,omitempty"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type YamlfilecontentObject struct {
	XMLName xml.Name `xml:yamlfilecontent_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type YamlfilecontentState struct {
	XMLName xml.Name `xml:yamlfilecontent_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr,omitempty"`

	Comment string `xml:"comment,attr,omitempty"`

	Deprecated string `xml:"deprecated,attr,omitempty"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	Yamlpath *oval_def.EntityStateStringType `xml:"yamlpath"`

	Value *oval_def.EntityStateRecordType `xml:"value"`

	WindowsView *EntityStateWindowsViewType `xml:"windows_view"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// XSD ComplexType declarations

type LdapBehaviors struct {
	Scope string `xml:"scope,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type Textfilecontent54Behaviors struct {
	IgnoreCase string `xml:"ignore_case,attr,omitempty"`

	Multiline string `xml:"multiline,attr,omitempty"`

	Singleline string `xml:"singleline,attr,omitempty"`

	MaxDepth string `xml:"max_depth,attr,omitempty"`

	Recurse string `xml:"recurse,attr,omitempty"`

	RecurseDirection string `xml:"recurse_direction,attr,omitempty"`

	RecurseFileSystem string `xml:"recurse_file_system,attr,omitempty"`

	WindowsView string `xml:"windows_view,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type FileBehaviors struct {
	MaxDepth string `xml:"max_depth,attr,omitempty"`

	Recurse string `xml:"recurse,attr,omitempty"`

	RecurseDirection string `xml:"recurse_direction,attr,omitempty"`

	RecurseFileSystem string `xml:"recurse_file_system,attr,omitempty"`

	WindowsView string `xml:"windows_view,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectEngineType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateEngineType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateFamilyType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityObjectHashTypeType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateHashTypeType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityObjectVariableRefType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateVariableRefType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateLdaptypeType struct {
	InnerXml string `xml:",innerxml"`
}

type EntityStateWindowsViewType struct {
	InnerXml string `xml:",innerxml"`
}
