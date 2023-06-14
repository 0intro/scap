// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
package iosxe_def

import (
	"encoding/xml"
	"github.com/0intro/scap/pkg/scap/models/oval"
	"github.com/0intro/scap/pkg/scap/models/oval_def"
	"github.com/0intro/scap/pkg/scap/models/xml_dsig"
)

// Element
type GlobalTest struct {
	XMLName xml.Name `xml:"global_test"`

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
type GlobalObject struct {
	XMLName xml.Name `xml:"global_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	GlobalCommand *oval_def.EntityObjectStringType `xml:"global_command"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type GlobalState struct {
	XMLName xml.Name `xml:"global_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	GlobalCommand *oval_def.EntityStateStringType `xml:"global_command"`

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

	Platform *oval_def.EntityStateStringType `xml:"platform"`

	Rp *oval_def.EntityStateIntType `xml:"rp"`

	Pkg *oval_def.EntityStateStringType `xml:"pkg"`

	VersionString *oval_def.EntityStateStringType `xml:"version_string"`

	MajorRelease *oval_def.EntityStateIntType `xml:"major_release"`

	Release *oval_def.EntityStateIntType `xml:"release"`

	Rebuild *oval_def.EntityStateIntType `xml:"rebuild"`

	Train *oval_def.EntityStateStringType `xml:"train"`

	IosRelease *oval_def.EntityStateStringType `xml:"ios_release"`

	IosTrain *oval_def.EntityStateStringType `xml:"ios_train"`

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

	IpDirectedBroadcast *oval_def.EntityStateBoolType `xml:"ip_directed_broadcast"`

	ProxyArp *oval_def.EntityStateBoolType `xml:"proxy_arp"`

	Shutdown *oval_def.EntityStateBoolType `xml:"shutdown"`

	HardwareAddr *oval_def.EntityStateStringType `xml:"hardware_addr"`

	Ipv4Address *oval_def.EntityStateIPAddressStringType `xml:"ipv4_address"`

	Ipv6Address *oval_def.EntityStateIPAddressStringType `xml:"ipv6_address"`

	Ipv4AccessList *oval_def.EntityStateStringType `xml:"ipv4_access_list"`

	Ipv6AccessList *oval_def.EntityStateStringType `xml:"ipv6_access_list"`

	CryptoMap *oval_def.EntityStateStringType `xml:"crypto_map"`

	Ipv4UrpfCommand *oval_def.EntityStateStringType `xml:"ipv4_urpf_command"`

	Ipv6UrpfCommand *oval_def.EntityStateStringType `xml:"ipv6_urpf_command"`

	UrpfCommand *oval_def.EntityStateStringType `xml:"urpf_command"`

	SwitchportTrunkEncapsulation *EntityStateTrunkEncapType `xml:"switchport_trunk_encapsulation"`

	SwitchportMode *EntityStateSwitchportModeType `xml:"switchport_mode"`

	SwitchportNativeVlan *InterfaceStateSwitchportNativeVlan `xml:"switchport_native_vlan"`

	SwitchportAccessVlan *InterfaceStateSwitchportAccessVlan `xml:"switchport_access_vlan"`

	SwitchportTrunkedVlans *oval_def.EntityStateStringType `xml:"switchport_trunked_vlans"`

	SwitchportPrunedVlans *oval_def.EntityStateStringType `xml:"switchport_pruned_vlans"`

	SwitchportPortSecurity *oval_def.EntityStateStringType `xml:"switchport_port_security"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type SectionTest struct {
	XMLName xml.Name `xml:"section_test"`

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
type SectionObject struct {
	XMLName xml.Name `xml:"section_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	SectionCommand *oval_def.EntityObjectStringType `xml:"section_command"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type SectionState struct {
	XMLName xml.Name `xml:"section_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	SectionCommand *oval_def.EntityStateStringType `xml:"section_command"`

	SectionConfigLines *oval_def.EntityStateStringType `xml:"section_config_lines"`

	ConfigLine *oval_def.EntityStateStringType `xml:"config_line"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type RouterTest struct {
	XMLName xml.Name `xml:"router_test"`

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
type RouterObject struct {
	XMLName xml.Name `xml:"router_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Protocol *EntityObjectRoutingProtocolType `xml:"protocol"`

	IdElm *oval_def.EntityObjectIntType `xml:"id"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type RouterState struct {
	XMLName xml.Name `xml:"router_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Protocol EntityStateRoutingProtocolType `xml:"protocol"`

	IdElm *oval_def.EntityStateIntType `xml:"id"`

	Network *oval_def.EntityStateStringType `xml:"network"`

	BgpNeighbor *oval_def.EntityStateStringType `xml:"bgp_neighbor"`

	OspfAuthenticationArea *RouterStateOspfAuthenticationArea `xml:"ospf_authentication_area"`

	RouterConfigLines *oval_def.EntityStateStringType `xml:"router_config_lines"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type BgpneighborTest struct {
	XMLName xml.Name `xml:"bgpneighbor_test"`

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
type BgpneighborObject struct {
	XMLName xml.Name `xml:"bgpneighbor_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Neighbor *oval_def.EntityObjectStringType `xml:"neighbor"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type BgpneighborState struct {
	XMLName xml.Name `xml:"bgpneighbor_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Neighbor *oval_def.EntityStateStringType `xml:"neighbor"`

	Password *oval_def.EntityStateStringType `xml:"password"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type RoutingprotocolauthintfTest struct {
	XMLName xml.Name `xml:"routingprotocolauthintf_test"`

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
type RoutingprotocolauthintfObject struct {
	XMLName xml.Name `xml:"routingprotocolauthintf_object"`

	Id oval.ObjectIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Set *oval_def.Set `xml:"set"`

	Interface *oval_def.EntityObjectStringType `xml:"interface"`

	Protocol *EntityObjectRoutingProtocolType `xml:"protocol"`

	Filter []oval_def.Filter `xml:"filter"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type RoutingprotocolauthintfState struct {
	XMLName xml.Name `xml:"routingprotocolauthintf_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Interface *oval_def.EntityStateStringType `xml:"interface"`

	Protocol *EntityStateRoutingProtocolType `xml:"protocol"`

	IdElm *oval_def.EntityStateIntType `xml:"id"`

	AuthType *EntityStateRoutingAuthTypeStringType `xml:"auth_type"`

	OspfArea *RoutingprotocolauthintfStateOspfArea `xml:"ospf_area"`

	KeyChain *oval_def.EntityStateStringType `xml:"key_chain"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

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
type SnmphostTest struct {
	XMLName xml.Name `xml:"snmphost_test"`

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
type SnmphostObject struct {
	XMLName xml.Name `xml:"snmphost_object"`

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
type SnmphostState struct {
	XMLName xml.Name `xml:"snmphost_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Host *oval_def.EntityStateStringType `xml:"host"`

	CommunityOrUser *oval_def.EntityStateStringType `xml:"community_or_user"`

	VersionElm *EntityStateSNMPVersionStringType `xml:"version"`

	Snmpv3SecLevel *EntityStateSNMPSecLevelStringType `xml:"snmpv3_sec_level"`

	Traps *oval_def.EntityStateStringType `xml:"traps"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type SnmpcommunityTest struct {
	XMLName xml.Name `xml:"snmpcommunity_test"`

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
type SnmpcommunityObject struct {
	XMLName xml.Name `xml:"snmpcommunity_object"`

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
type SnmpcommunityState struct {
	XMLName xml.Name `xml:"snmpcommunity_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	View *oval_def.EntityStateStringType `xml:"view"`

	Mode *EntityStateSNMPModeStringType `xml:"mode"`

	Ipv4Acl *oval_def.EntityStateStringType `xml:"ipv4_acl"`

	Ipv6Acl *oval_def.EntityStateStringType `xml:"ipv6_acl"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type SnmpuserTest struct {
	XMLName xml.Name `xml:"snmpuser_test"`

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
type SnmpuserObject struct {
	XMLName xml.Name `xml:"snmpuser_object"`

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
type SnmpuserState struct {
	XMLName xml.Name `xml:"snmpuser_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Group *oval_def.EntityStateStringType `xml:"group"`

	VersionElm *EntityStateSNMPVersionStringType `xml:"version"`

	Ipv4Acl *oval_def.EntityStateStringType `xml:"ipv4_acl"`

	Ipv6Acl *oval_def.EntityStateStringType `xml:"ipv6_acl"`

	Priv *EntityStateSNMPPrivStringType `xml:"priv"`

	Auth *EntityStateSNMPAuthStringType `xml:"auth"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type SnmpgroupTest struct {
	XMLName xml.Name `xml:"snmpgroup_test"`

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
type SnmpgroupObject struct {
	XMLName xml.Name `xml:"snmpgroup_object"`

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
type SnmpgroupState struct {
	XMLName xml.Name `xml:"snmpgroup_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	VersionElm *EntityStateSNMPVersionStringType `xml:"version"`

	Snmpv3SecLevel *EntityStateSNMPSecLevelStringType `xml:"snmpv3_sec_level"`

	Ipv4Acl *oval_def.EntityStateStringType `xml:"ipv4_acl"`

	Ipv6Acl *oval_def.EntityStateStringType `xml:"ipv6_acl"`

	ReadView *oval_def.EntityStateStringType `xml:"read_view"`

	WriteView *oval_def.EntityStateStringType `xml:"write_view"`

	NotifyView *oval_def.EntityStateStringType `xml:"notify_view"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type SnmpviewTest struct {
	XMLName xml.Name `xml:"snmpview_test"`

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
type SnmpviewObject struct {
	XMLName xml.Name `xml:"snmpview_object"`

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
type SnmpviewState struct {
	XMLName xml.Name `xml:"snmpview_state"`

	Id oval.StateIDPattern `xml:"id,attr"`

	Version uint64 `xml:"version,attr"`

	Operator oval.OperatorEnumeration `xml:"operator,attr,omitempty"`

	Comment oval.NonEmptyStringType `xml:"comment,attr,omitempty"`

	Deprecated bool `xml:"deprecated,attr,omitempty"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	MibFamily *oval_def.EntityStateStringType `xml:"mib_family"`

	Include *oval_def.EntityStateBoolType `xml:"include"`

	Signature *xml_dsig.SignatureType `xml:"Signature"`

	Notes *oval.NotesType `xml:"notes"`
}

// Element
type InterfaceStateSwitchportNativeVlan struct {
	XMLName xml.Name `xml:"switchport_native_vlan"`

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`
}

// Element
type InterfaceStateSwitchportAccessVlan struct {
	XMLName xml.Name `xml:"switchport_access_vlan"`

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`
}

// Element
type RouterStateOspfAuthenticationArea struct {
	XMLName xml.Name `xml:"ospf_authentication_area"`

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`
}

// Element
type RoutingprotocolauthintfStateOspfArea struct {
	XMLName xml.Name `xml:"ospf_area"`

	EntityCheck string `xml:"entity_check,attr,omitempty"`

	CheckExistence string `xml:"check_existence,attr,omitempty"`

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`
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

type EntityObjectRoutingProtocolType struct {
	XMLName xml.Name

	Datatype oval.DatatypeEnumeration `xml:"datatype,attr,omitempty"`

	Operation oval.OperationEnumeration `xml:"operation,attr,omitempty"`

	Mask bool `xml:"mask,attr,omitempty"`

	VarRef oval.VariableIDPattern `xml:"var_ref,attr,omitempty"`

	VarCheck oval.CheckEnumeration `xml:"var_check,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type EntityStateTrunkEncapType struct {
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

type EntityStateSwitchportModeType struct {
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

type EntityStateRoutingProtocolType struct {
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

type EntityStateRoutingAuthTypeStringType struct {
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

type EntityStateSNMPModeStringType struct {
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

// XSD SimpleType declarations
