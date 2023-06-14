// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
package ind_def

import (
	"encoding/xml"
	"github.com/0intro/scap/pkg/scap/models/oval"
	"github.com/0intro/scap/pkg/scap/models/oval_def"
	"github.com/0intro/scap/pkg/scap/models/xml_dsig"
)

// Element
type FamilyTest struct {
	XMLName xml.Name `xml:"family_test"`

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
type FamilyObject struct {
	XMLName xml.Name `xml:"family_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type FamilyState struct {
	XMLName xml.Name `xml:"family_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Family *EntityStateFamilyType `xml:"family"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type FilehashTest struct {
	XMLName xml.Name `xml:"filehash_test"`

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
type FilehashObject struct {
	XMLName xml.Name `xml:"filehash_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Behaviors *FileBehaviors `xml:"behaviors"`

	Filepath *oval_def.EntityObjectStringType `xml:"filepath"`

	Path *oval_def.EntityObjectStringType `xml:"path"`

	Filename *oval_def.EntityObjectStringType `xml:"filename"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type FilehashState struct {
	XMLName xml.Name `xml:"filehash_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	Md5 *oval_def.EntityStateStringType `xml:"md5"`

	Sha1 *oval_def.EntityStateStringType `xml:"sha1"`

	WindowsView *EntityStateWindowsViewType `xml:"windows_view"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type Filehash58Test struct {
	XMLName xml.Name `xml:"filehash58_test"`

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
type Filehash58Object struct {
	XMLName xml.Name `xml:"filehash58_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Behaviors *FileBehaviors `xml:"behaviors"`

	HashType *EntityObjectHashTypeType `xml:"hash_type"`

	Filter []oval_def.Filter `xml:"filter"`

	Filepath *oval_def.EntityObjectStringType `xml:"filepath"`

	Path *oval_def.EntityObjectStringType `xml:"path"`

	Filename *oval_def.EntityObjectStringType `xml:"filename"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type Filehash58State struct {
	XMLName xml.Name `xml:"filehash58_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	HashType *EntityStateHashTypeType `xml:"hash_type"`

	Hash *oval_def.EntityStateStringType `xml:"hash"`

	WindowsView *EntityStateWindowsViewType `xml:"windows_view"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type EnvironmentvariableTest struct {
	XMLName xml.Name `xml:"environmentvariable_test"`

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
type EnvironmentvariableObject struct {
	XMLName xml.Name `xml:"environmentvariable_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type EnvironmentvariableState struct {
	XMLName xml.Name `xml:"environmentvariable_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type Environmentvariable58Test struct {
	XMLName xml.Name `xml:"environmentvariable58_test"`

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
type Environmentvariable58Object struct {
	XMLName xml.Name `xml:"environmentvariable58_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Pid *oval_def.EntityObjectIntType `xml:"pid"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type Environmentvariable58State struct {
	XMLName xml.Name `xml:"environmentvariable58_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Pid *oval_def.EntityStateIntType `xml:"pid"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type LdapTest struct {
	XMLName xml.Name `xml:"ldap_test"`

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
type LdapObject struct {
	XMLName xml.Name `xml:"ldap_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Behaviors *LdapBehaviors `xml:"behaviors"`

	Suffix *oval_def.EntityObjectStringType `xml:"suffix"`

	RelativeDn *oval_def.EntityObjectStringType `xml:"relative_dn"`

	Attribute *oval_def.EntityObjectStringType `xml:"attribute"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type LdapState struct {
	XMLName xml.Name `xml:"ldap_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Suffix *oval_def.EntityStateStringType `xml:"suffix"`

	RelativeDn *oval_def.EntityStateStringType `xml:"relative_dn"`

	Attribute *oval_def.EntityStateStringType `xml:"attribute"`

	ObjectClass *oval_def.EntityStateStringType `xml:"object_class"`

	Ldaptype *EntityStateLdaptypeType `xml:"ldaptype"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type Ldap57Test struct {
	XMLName xml.Name `xml:"ldap57_test"`

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
type Ldap57Object struct {
	XMLName xml.Name `xml:"ldap57_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Behaviors *LdapBehaviors `xml:"behaviors"`

	Suffix *oval_def.EntityObjectStringType `xml:"suffix"`

	RelativeDn *oval_def.EntityObjectStringType `xml:"relative_dn"`

	Attribute *oval_def.EntityObjectStringType `xml:"attribute"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type Ldap57State struct {
	XMLName xml.Name `xml:"ldap57_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Suffix *oval_def.EntityStateStringType `xml:"suffix"`

	RelativeDn *oval_def.EntityStateStringType `xml:"relative_dn"`

	Attribute *oval_def.EntityStateStringType `xml:"attribute"`

	ObjectClass *oval_def.EntityStateStringType `xml:"object_class"`

	Ldaptype *EntityStateLdaptypeType `xml:"ldaptype"`

	Value *oval_def.EntityStateRecordType `xml:"value"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type SqlTest struct {
	XMLName xml.Name `xml:"sql_test"`

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
type SqlObject struct {
	XMLName xml.Name `xml:"sql_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Engine *EntityObjectEngineType `xml:"engine"`

	VersionElm *oval_def.EntityObjectStringType `xml:"version"`

	ConnectionString *oval_def.EntityObjectStringType `xml:"connection_string"`

	Sql *oval_def.EntityObjectStringType `xml:"sql"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type SqlState struct {
	XMLName xml.Name `xml:"sql_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Engine *EntityStateEngineType `xml:"engine"`

	VersionElm *oval_def.EntityStateStringType `xml:"version"`

	ConnectionString *oval_def.EntityStateStringType `xml:"connection_string"`

	Sql *oval_def.EntityStateStringType `xml:"sql"`

	Result *oval_def.EntityStateAnySimpleType `xml:"result"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type Sql57Test struct {
	XMLName xml.Name `xml:"sql57_test"`

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
type Sql57Object struct {
	XMLName xml.Name `xml:"sql57_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Engine *EntityObjectEngineType `xml:"engine"`

	VersionElm *oval_def.EntityObjectStringType `xml:"version"`

	ConnectionString *oval_def.EntityObjectStringType `xml:"connection_string"`

	Sql *oval_def.EntityObjectStringType `xml:"sql"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type Sql57State struct {
	XMLName xml.Name `xml:"sql57_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Engine *EntityStateEngineType `xml:"engine"`

	VersionElm *oval_def.EntityStateStringType `xml:"version"`

	ConnectionString *oval_def.EntityStateStringType `xml:"connection_string"`

	Sql *oval_def.EntityStateStringType `xml:"sql"`

	Result *oval_def.EntityStateRecordType `xml:"result"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type Textfilecontent54Test struct {
	XMLName xml.Name `xml:"textfilecontent54_test"`

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
type Textfilecontent54Object struct {
	XMLName xml.Name `xml:"textfilecontent54_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Behaviors *Textfilecontent54Behaviors `xml:"behaviors"`

	Pattern *oval_def.EntityObjectStringType `xml:"pattern"`

	Instance *oval_def.EntityObjectIntType `xml:"instance"`

	Filter []oval_def.Filter `xml:"filter"`

	Filepath *oval_def.EntityObjectStringType `xml:"filepath"`

	Path *oval_def.EntityObjectStringType `xml:"path"`

	Filename *oval_def.EntityObjectStringType `xml:"filename"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type Textfilecontent54State struct {
	XMLName xml.Name `xml:"textfilecontent54_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	Pattern *oval_def.EntityStateStringType `xml:"pattern"`

	Instance *oval_def.EntityStateIntType `xml:"instance"`

	Text *oval_def.EntityStateAnySimpleType `xml:"text"`

	Subexpression *oval_def.EntityStateAnySimpleType `xml:"subexpression"`

	WindowsView *EntityStateWindowsViewType `xml:"windows_view"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type TextfilecontentTest struct {
	XMLName xml.Name `xml:"textfilecontent_test"`

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
type TextfilecontentObject struct {
	XMLName xml.Name `xml:"textfilecontent_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Behaviors *FileBehaviors `xml:"behaviors"`

	Path *oval_def.EntityObjectStringType `xml:"path"`

	Filename *oval_def.EntityObjectStringType `xml:"filename"`

	Line *oval_def.EntityObjectStringType `xml:"line"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type TextfilecontentState struct {
	XMLName xml.Name `xml:"textfilecontent_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	Line *oval_def.EntityStateStringType `xml:"line"`

	Subexpression *oval_def.EntityStateAnySimpleType `xml:"subexpression"`

	WindowsView *EntityStateWindowsViewType `xml:"windows_view"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type UnknownTest struct {
	XMLName xml.Name `xml:"unknown_test"`

	Id oval.TestIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	CheckExistence oval.ExistenceEnumeration `xml:"check_existence,attr,omitempty"`

	Check oval.CheckEnumeration `xml:"check,attr"`

	StateOperator oval.OperatorEnumeration `xml:"state_operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type VariableTest struct {
	XMLName xml.Name `xml:"variable_test"`

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
type VariableObject struct {
	XMLName xml.Name `xml:"variable_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	VarRef *EntityObjectVariableRefType `xml:"var_ref"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type VariableState struct {
	XMLName xml.Name `xml:"variable_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	VarRef *EntityStateVariableRefType `xml:"var_ref"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type XmlfilecontentTest struct {
	XMLName xml.Name `xml:"xmlfilecontent_test"`

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
type XmlfilecontentObject struct {
	XMLName xml.Name `xml:"xmlfilecontent_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Behaviors *FileBehaviors `xml:"behaviors"`

	Xpath *oval_def.EntityObjectStringType `xml:"xpath"`

	Filter []oval_def.Filter `xml:"filter"`

	Filepath *oval_def.EntityObjectStringType `xml:"filepath"`

	Path *oval_def.EntityObjectStringType `xml:"path"`

	Filename *oval_def.EntityObjectStringType `xml:"filename"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type XmlfilecontentState struct {
	XMLName xml.Name `xml:"xmlfilecontent_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	Xpath *oval_def.EntityStateStringType `xml:"xpath"`

	ValueOf *oval_def.EntityStateAnySimpleType `xml:"value_of"`

	WindowsView *EntityStateWindowsViewType `xml:"windows_view"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type YamlfilecontentTest struct {
	XMLName xml.Name `xml:"yamlfilecontent_test"`

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
type YamlfilecontentObject struct {
	XMLName xml.Name `xml:"yamlfilecontent_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Behaviors *FileBehaviors `xml:"behaviors"`

	Yamlpath *oval_def.EntityObjectStringType `xml:"yamlpath"`

	Filter []oval_def.Filter `xml:"filter"`

	Filepath *oval_def.EntityObjectStringType `xml:"filepath"`

	Content *oval_def.EntityObjectStringType `xml:"content"`

	Path *oval_def.EntityObjectStringType `xml:"path"`

	Filename *oval_def.EntityObjectStringType `xml:"filename"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type YamlfilecontentState struct {
	XMLName xml.Name `xml:"yamlfilecontent_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	Content *oval_def.EntityStateStringType `xml:"content"`

	Yamlpath *oval_def.EntityStateStringType `xml:"yamlpath"`

	Value *oval_def.EntityStateRecordType `xml:"value"`

	WindowsView *EntityStateWindowsViewType `xml:"windows_view"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// XSD ComplexType declarations

type LdapBehaviors struct {
	XMLName xml.Name

	Scope string `xml:"scope,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type Textfilecontent54Behaviors struct {
	XMLName xml.Name

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
	XMLName xml.Name

	MaxDepth string `xml:"max_depth,attr,omitempty"`

	Recurse string `xml:"recurse,attr,omitempty"`

	RecurseDirection string `xml:"recurse_direction,attr,omitempty"`

	RecurseFileSystem string `xml:"recurse_file_system,attr,omitempty"`

	WindowsView string `xml:"windows_view,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectEngineType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateEngineType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateFamilyType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectHashTypeType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateHashTypeType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectVariableRefType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateVariableRefType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateLdaptypeType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateWindowsViewType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations
