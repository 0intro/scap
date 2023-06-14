// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
package asa_def

import (
	"encoding/xml"
	"github.com/0intro/scap/pkg/scap/models/oval"
	"github.com/0intro/scap/pkg/scap/models/oval_def"
	"github.com/0intro/scap/pkg/scap/models/xml_dsig"
)

// Element
type AclTest struct {
	XMLName xml.Name `xml:"acl_test"`

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
type AclObject struct {
	XMLName xml.Name `xml:"acl_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	IpVersion *EntityObjectAccessListIPVersionType `xml:"ip_version"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type AclState struct {
	XMLName xml.Name `xml:"acl_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	IpVersion *EntityStateAccessListIPVersionType `xml:"ip_version"`

	Use *EntityStateAccessListUseType `xml:"use"`

	UsedIn *oval_def.EntityStateStringType `xml:"used_in"`

	InterfaceDirection *EntityStateAccessListInterfaceDirectionType `xml:"interface_direction"`

	AclConfigLines *oval_def.EntityStateStringType `xml:"acl_config_lines"`

	ConfigLine *oval_def.EntityStateStringType `xml:"config_line"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type ClassMapTest struct {
	XMLName xml.Name `xml:"class_map_test"`

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
type ClassMapObject struct {
	XMLName xml.Name `xml:"class_map_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type ClassMapState struct {
	XMLName xml.Name `xml:"class_map_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Type *EntityStateClassMapType `xml:"type"`

	TypeInspect *EntityStateInspectionType `xml:"type_inspect"`

	MatchAllAny *EntityStateMatchType `xml:"match_all_any"`

	Match *oval_def.EntityStateStringType `xml:"match"`

	UsedInClassMap *oval_def.EntityStateStringType `xml:"used_in_class_map"`

	UsedInPolicyMap *oval_def.EntityStateStringType `xml:"used_in_policy_map"`

	PolicyMapAction *oval_def.EntityStateStringType `xml:"policy_map_action"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type InterfaceTest struct {
	XMLName xml.Name `xml:"interface_test"`

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
type InterfaceObject struct {
	XMLName xml.Name `xml:"interface_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type InterfaceState struct {
	XMLName xml.Name `xml:"interface_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	ProxyArp *oval_def.EntityStateBoolType `xml:"proxy_arp"`

	Shutdown *oval_def.EntityStateBoolType `xml:"shutdown"`

	HardwareAddr *oval_def.EntityStateStringType `xml:"hardware_addr"`

	Ipv4Address *oval_def.EntityStateIPAddressStringType `xml:"ipv4_address"`

	Ipv6Address *oval_def.EntityStateIPAddressStringType `xml:"ipv6_address"`

	Ipv4AccessList *oval_def.EntityStateStringType `xml:"ipv4_access_list"`

	Ipv6AccessList *oval_def.EntityStateStringType `xml:"ipv6_access_list"`

	Ipv4V6AccessList *oval_def.EntityStateStringType `xml:"ipv4_v6_access_list"`

	CryptoMap *oval_def.EntityStateStringType `xml:"crypto_map"`

	Ipv4UrpfCommand *oval_def.EntityStateStringType `xml:"ipv4_urpf_command"`

	Ipv6UrpfCommand *oval_def.EntityStateStringType `xml:"ipv6_urpf_command"`

	UrpfCommand *oval_def.EntityStateStringType `xml:"urpf_command"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type LineTest struct {
	XMLName xml.Name `xml:"line_test"`

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
type LineObject struct {
	XMLName xml.Name `xml:"line_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	ShowSubcommand *oval_def.EntityObjectStringType `xml:"show_subcommand"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type LineState struct {
	XMLName xml.Name `xml:"line_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	ShowSubcommand *oval_def.EntityStateStringType `xml:"show_subcommand"`

	ConfigLine *oval_def.EntityStateStringType `xml:"config_line"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type PolicyMapTest struct {
	XMLName xml.Name `xml:"policy_map_test"`

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
type PolicyMapObject struct {
	XMLName xml.Name `xml:"policy_map_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type PolicyMapState struct {
	XMLName xml.Name `xml:"policy_map_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	TypeInspect *EntityStateInspectionType `xml:"type_inspect"`

	Parameters *oval_def.EntityStateStringType `xml:"parameters"`

	MatchAction *oval_def.EntityStateStringType `xml:"match_action"`

	UsedIn *oval_def.EntityStateStringType `xml:"used_in"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type ServicePolicyTest struct {
	XMLName xml.Name `xml:"service_policy_test"`

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
type ServicePolicyObject struct {
	XMLName xml.Name `xml:"service_policy_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type ServicePolicyState struct {
	XMLName xml.Name `xml:"service_policy_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Applied *EntityStateApplyServicePolicyType `xml:"applied"`

	Interface *oval_def.EntityStateStringType `xml:"interface"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type SnmpHostTest struct {
	XMLName xml.Name `xml:"snmp_host_test"`

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
type SnmpHostObject struct {
	XMLName xml.Name `xml:"snmp_host_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Host *oval_def.EntityObjectStringType `xml:"host"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type SnmpHostState struct {
	XMLName xml.Name `xml:"snmp_host_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Interface *oval_def.EntityStateStringType `xml:"interface"`

	Host *oval_def.EntityStateStringType `xml:"host"`

	Snmpv3User *oval_def.EntityStateStringType `xml:"snmpv3_user"`

	VersionElm *EntityStateSNMPVersionStringType `xml:"version"`

	Poll *oval_def.EntityStateBoolType `xml:"poll"`

	Traps *oval_def.EntityStateBoolType `xml:"traps"`

	UdpPort *oval_def.EntityStateIntType `xml:"udp_port"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type SnmpUserTest struct {
	XMLName xml.Name `xml:"snmp_user_test"`

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
type SnmpUserObject struct {
	XMLName xml.Name `xml:"snmp_user_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type SnmpUserState struct {
	XMLName xml.Name `xml:"snmp_user_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Group *oval_def.EntityStateStringType `xml:"group"`

	Priv *EntityStateSNMPPrivStringType `xml:"priv"`

	Auth *EntityStateSNMPAuthStringType `xml:"auth"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type SnmpGroupTest struct {
	XMLName xml.Name `xml:"snmp_group_test"`

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
type SnmpGroupObject struct {
	XMLName xml.Name `xml:"snmp_group_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type SnmpGroupState struct {
	XMLName xml.Name `xml:"snmp_group_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Snmpv3SecLevel *EntityStateSNMPSecLevelStringType `xml:"snmpv3_sec_level"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type TcpMapTest struct {
	XMLName xml.Name `xml:"tcp_map_test"`

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
type TcpMapObject struct {
	XMLName xml.Name `xml:"tcp_map_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Name *oval_def.EntityObjectStringType `xml:"name"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type TcpMapState struct {
	XMLName xml.Name `xml:"tcp_map_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Options *oval_def.EntityStateStringType `xml:"options"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type VersionTest struct {
	XMLName xml.Name `xml:"version_test"`

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
type VersionObject struct {
	XMLName xml.Name `xml:"version_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type VersionState struct {
	XMLName xml.Name `xml:"version_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	AsaRelease *oval_def.EntityStateStringType `xml:"asa_release"`

	AsaMajorRelease *oval_def.EntityStateVersionType `xml:"asa_major_release"`

	AsaMinorRelease *oval_def.EntityStateVersionType `xml:"asa_minor_release"`

	AsaBuild *oval_def.EntityStateIntType `xml:"asa_build"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// XSD ComplexType declarations

type EntityObjectAccessListIPVersionType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateAccessListIPVersionType struct {
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

type EntityStateAccessListUseType struct {
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

type EntityStateAccessListInterfaceDirectionType struct {
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

type EntityStateClassMapType struct {
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

type EntityStateInspectionType struct {
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

type EntityStateApplyServicePolicyType struct {
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

type EntityStateMatchType struct {
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

type EntityStateSNMPVersionStringType struct {
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

type EntityStateSNMPSecLevelStringType struct {
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

type EntityStateSNMPAuthStringType struct {
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

type EntityStateSNMPPrivStringType struct {
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
