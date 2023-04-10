// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5
package oval_def

import (
	"encoding/xml"
	"github.com/gocomply/scap/pkg/scap/models/oval"
	"github.com/gocomply/scap/pkg/scap/models/xml_dsig"
)

// Element
type OvalDefinitions struct {
	XMLName xml.Name `xml:"oval_definitions"`

	Generator oval.GeneratorType `xml:"generator"`

	Definitions *DefinitionsType `xml:"definitions"`

	Tests *TestsType `xml:"tests"`

	Objects *ObjectsType `xml:"objects"`

	States *StatesType `xml:"states"`

	Variables *VariablesType `xml:"variables"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`
}

// Element
type Notes struct {
	XMLName xml.Name `xml:"notes"`

	Note []string `xml:",any"`
}

// Element
type Definition struct {
	XMLName xml.Name `xml:"definition"`

	Id oval.DefinitionIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Class oval.ClassEnumeration `xml:"class,attr"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Metadata MetadataType `xml:"metadata"`

	Notes *oval.NotesType `xml:"notes"`

	Criteria *CriteriaType `xml:"criteria"`
}

// Element
type Test struct {
	XMLName xml.Name `xml:"test"`

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
type Object struct {
	XMLName xml.Name `xml:"object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type Set struct {
	XMLName xml.Name `xml:"set"`

	SetOperator SetOperatorEnumeration `xml:"set_operator,attr,omitempty"`

	Set []Set `xml:"set"`

	ObjectReference []oval.ObjectIDPattern `xml:"object_reference"`

	Filter []Filter `xml:"filter"`
}

// Element
type Filter struct {
	XMLName xml.Name `xml:"filter"`

	Action FilterActionEnumeration `xml:"action,attr,omitempty"`

	Text string `xml:",chardata"`
}

// Element
type State struct {
	XMLName xml.Name `xml:"state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type Variable struct {
	XMLName xml.Name `xml:"variable"`

	Id oval.VariableIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Datatype oval.SimpleDatatypeEnumeration `xml:"datatype,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type ExternalVariable struct {
	XMLName xml.Name `xml:"external_variable"`

	Id oval.VariableIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Datatype oval.SimpleDatatypeEnumeration `xml:"datatype,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type ConstantVariable struct {
	XMLName xml.Name `xml:"constant_variable"`

	Id oval.VariableIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Datatype oval.SimpleDatatypeEnumeration `xml:"datatype,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Value []ValueType `xml:"value"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type LocalVariable struct {
	XMLName xml.Name `xml:"local_variable"`

	Id oval.VariableIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Datatype oval.SimpleDatatypeEnumeration `xml:"datatype,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// XSD ComplexType declarations

type DefinitionsType struct {
	XMLName xml.Name

	Definition []DefinitionType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type DefinitionType struct {
	XMLName xml.Name

	Id oval.DefinitionIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Class oval.ClassEnumeration `xml:"class,attr"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Metadata MetadataType `xml:"metadata"`

	Notes *oval.NotesType `xml:"notes"`

	Criteria *CriteriaType `xml:"criteria"`

	InnerXml string `xml:",innerxml"`
}

type MetadataType struct {
	XMLName xml.Name

	Title string `xml:"title"`

	Affected []AffectedType `xml:"affected"`

	Reference []ReferenceType `xml:"reference"`

	Description string `xml:"description"`

	InnerXml string `xml:",innerxml"`
}

type AffectedType struct {
	XMLName xml.Name

	Family oval.FamilyEnumeration `xml:"family,attr"`

	Platform []string `xml:"platform"`

	Product []string `xml:"product"`

	InnerXml string `xml:",innerxml"`
}

type ReferenceType struct {
	XMLName xml.Name

	Source string `xml:"source,attr"`

	RefId string `xml:"ref_id,attr"`

	RefUrl string `xml:"ref_url,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type CriteriaType struct {
	XMLName xml.Name

	ApplicabilityCheck bool `xml:"applicability_check,attr,omitempty"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Negate bool `xml:"negate,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Criteria []CriteriaType `xml:"criteria"`

	Criterion []CriterionType `xml:"criterion"`

	ExtendDefinition []ExtendDefinitionType `xml:"extend_definition"`

	InnerXml string `xml:",innerxml"`
}

type CriterionType struct {
	XMLName xml.Name

	ApplicabilityCheck bool `xml:"applicability_check,attr,omitempty"`

	TestRef oval.TestIDPattern `xml:"test_ref,attr"`

	Negate bool `xml:"negate,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type ExtendDefinitionType struct {
	XMLName xml.Name

	ApplicabilityCheck bool `xml:"applicability_check,attr,omitempty"`

	DefinitionRef oval.DefinitionIDPattern `xml:"definition_ref,attr"`

	Negate bool `xml:"negate,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type TestsType struct {
	XMLName xml.Name

	Test []TestType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type TestType struct {
	XMLName xml.Name

	Id oval.TestIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	CheckExistence oval.ExistenceEnumeration `xml:"check_existence,attr,omitempty"`

	Check oval.CheckEnumeration `xml:"check,attr"`

	StateOperator oval.OperatorEnumeration `xml:"state_operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type ObjectRefType struct {
	XMLName xml.Name

	ObjectRef oval.ObjectIDPattern `xml:"object_ref,attr"`

	InnerXml string `xml:",innerxml"`
}

type StateRefType struct {
	XMLName xml.Name

	StateRef oval.StateIDPattern `xml:"state_ref,attr"`

	InnerXml string `xml:",innerxml"`
}

type ObjectsType struct {
	XMLName xml.Name

	Object []ObjectType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type ObjectType struct {
	XMLName xml.Name

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type StatesType struct {
	XMLName xml.Name

	State []StateType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type StateType struct {
	XMLName xml.Name

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type VariablesType struct {
	XMLName xml.Name

	Variable []VariableType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type VariableType struct {
	XMLName xml.Name

	Id oval.VariableIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Datatype oval.SimpleDatatypeEnumeration `xml:"datatype,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type PossibleValueType struct {
	XMLName xml.Name

	Hint string `xml:"hint,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type PossibleRestrictionType struct {
	XMLName xml.Name

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Hint string `xml:"hint,attr"`

	Restriction []RestrictionType `xml:"restriction"`

	InnerXml string `xml:",innerxml"`
}

type RestrictionType struct {
	XMLName xml.Name

	Operation oval.OperationEnumeration `xml:"operation,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type ValueType struct {
	XMLName xml.Name

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type LiteralComponentType struct {
	XMLName xml.Name

	Datatype oval.SimpleDatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type ObjectComponentType struct {
	XMLName xml.Name

	ObjectRef oval.ObjectIDPattern `xml:"object_ref,attr"`

	ItemField oval.NonEmptyStringType `xml:"item_field,attr"`

	RecordField oval.NonEmptyStringType `xml:"record_field,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type VariableComponentType struct {
	XMLName xml.Name

	VarRef oval.VariableIDPattern `xml:"var_ref,attr"`

	InnerXml string `xml:",innerxml"`
}

type ArithmeticFunctionType struct {
	XMLName xml.Name

	ArithmeticOperation ArithmeticEnumeration `xml:"arithmetic_operation,attr"`

	InnerXml string `xml:",innerxml"`
}

type BeginFunctionType struct {
	XMLName xml.Name

	Character string `xml:"character,attr"`

	InnerXml string `xml:",innerxml"`
}

type ConcatFunctionType struct {
	XMLName xml.Name

	InnerXml string `xml:",innerxml"`
}

type EndFunctionType struct {
	XMLName xml.Name

	Character string `xml:"character,attr"`

	InnerXml string `xml:",innerxml"`
}

type EscapeRegexFunctionType struct {
	XMLName xml.Name

	InnerXml string `xml:",innerxml"`
}

type SplitFunctionType struct {
	XMLName xml.Name

	Delimiter string `xml:"delimiter,attr"`

	InnerXml string `xml:",innerxml"`
}

type SubstringFunctionType struct {
	XMLName xml.Name

	SubstringStart int `xml:"substring_start,attr"`

	SubstringLength int `xml:"substring_length,attr"`

	InnerXml string `xml:",innerxml"`
}

type TimeDifferenceFunctionType struct {
	XMLName xml.Name

	Format1 DateTimeFormatEnumeration `xml:"format_1,attr,omitempty"`

	Format2 DateTimeFormatEnumeration `xml:"format_2,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type RegexCaptureFunctionType struct {
	XMLName xml.Name

	Pattern string `xml:"pattern,attr"`

	InnerXml string `xml:",innerxml"`
}

type UniqueFunctionType struct {
	XMLName xml.Name

	InnerXml string `xml:",innerxml"`
}

type CountFunctionType struct {
	XMLName xml.Name

	InnerXml string `xml:",innerxml"`
}

type GlobToRegexFunctionType struct {
	XMLName xml.Name

	GlobNoescape bool `xml:"glob_noescape,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntitySimpleBaseType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type EntityComplexBaseType struct {
	XMLName xml.Name

	InnerXml string `xml:",innerxml"`
}

type EntityObjectIPAddressType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectIPAddressStringType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectAnySimpleType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectBinaryType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectBoolType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectFloatType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectIntType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectStringType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectVersionType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectRecordType struct {
	XMLName xml.Name

	Field []EntityObjectFieldType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type EntityObjectFieldType struct {
	XMLName xml.Name

	Name string `xml:"name,attr"`

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type EntityStateSimpleBaseType struct {
	XMLName xml.Name

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type EntityStateComplexBaseType struct {
	XMLName xml.Name

	EntityCheck oval.CheckEnumeration `xml:"entity_check,attr,omitempty"`

	CheckExistence oval.ExistenceEnumeration `xml:"check_existence,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateIPAddressType struct {
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

type EntityStateIPAddressStringType struct {
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

type EntityStateAnySimpleType struct {
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

type EntityStateBinaryType struct {
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

type EntityStateBoolType struct {
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

type EntityStateFloatType struct {
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

type EntityStateIntType struct {
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

type EntityStateEVRStringType struct {
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

type EntityStateDebianEVRStringType struct {
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

type EntityStateVersionType struct {
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

type EntityStateFileSetRevisionType struct {
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

type EntityStateIOSVersionType struct {
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

type EntityStateStringType struct {
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

type EntityStateRecordType struct {
	XMLName xml.Name

	EntityCheck oval.CheckEnumeration `xml:"entity_check,attr,omitempty"`

	CheckExistence oval.ExistenceEnumeration `xml:"check_existence,attr,omitempty"`

	Field []EntityStateFieldType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateFieldType struct {
	XMLName xml.Name

	Name string `xml:"name,attr"`

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations

type ArithmeticEnumeration string

const ArithmeticEnumerationAdd ArithmeticEnumeration = "add"

const ArithmeticEnumerationMultiply ArithmeticEnumeration = "multiply"

type DateTimeFormatEnumeration string

const DateTimeFormatEnumerationYearMonthDay DateTimeFormatEnumeration = "year_month_day"

const DateTimeFormatEnumerationMonthDayYear DateTimeFormatEnumeration = "month_day_year"

const DateTimeFormatEnumerationDayMonthYear DateTimeFormatEnumeration = "day_month_year"

const DateTimeFormatEnumerationWinFiletime DateTimeFormatEnumeration = "win_filetime"

const DateTimeFormatEnumerationSecondsSinceEpoch DateTimeFormatEnumeration = "seconds_since_epoch"

const DateTimeFormatEnumerationCimDatetime DateTimeFormatEnumeration = "cim_datetime"

type FilterActionEnumeration string

const FilterActionEnumerationExclude FilterActionEnumeration = "exclude"

const FilterActionEnumerationInclude FilterActionEnumeration = "include"

type SetOperatorEnumeration string

const SetOperatorEnumerationComplement SetOperatorEnumeration = "COMPLEMENT"

const SetOperatorEnumerationIntersection SetOperatorEnumeration = "INTERSECTION"

const SetOperatorEnumerationUnion SetOperatorEnumeration = "UNION"
