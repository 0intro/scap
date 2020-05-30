// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5#ios
package ios_sc

import (
	"encoding/xml"
)

// Element
type AclItem struct {
	XMLName xml.Name `xml:acl_item`
}

// Element
type BgpneighborItem struct {
	XMLName xml.Name `xml:bgpneighbor_item`
}

// Element
type GlobalItem struct {
	XMLName xml.Name `xml:global_item`
}

// Element
type InterfaceItem struct {
	XMLName xml.Name `xml:interface_item`
}

// Element
type LineItem struct {
	XMLName xml.Name `xml:line_item`
}

// Element
type RouterItem struct {
	XMLName xml.Name `xml:router_item`
}

// Element
type RoutingprotocolauthintfItem struct {
	XMLName xml.Name `xml:routingprotocolauthintf_item`
}

// Element
type SectionItem struct {
	XMLName xml.Name `xml:section_item`
}

// Element
type SnmpItem struct {
	XMLName xml.Name `xml:snmp_item`
}

// Element
type SnmpcommunityItem struct {
	XMLName xml.Name `xml:snmpcommunity_item`
}

// Element
type SnmpgroupItem struct {
	XMLName xml.Name `xml:snmpgroup_item`
}

// Element
type SnmphostItem struct {
	XMLName xml.Name `xml:snmphost_item`
}

// Element
type SnmpuserItem struct {
	XMLName xml.Name `xml:snmpuser_item`
}

// Element
type SnmpviewItem struct {
	XMLName xml.Name `xml:snmpview_item`
}

// Element
type TclshItem struct {
	XMLName xml.Name `xml:tclsh_item`
}

// Element
type VersionItem struct {
	XMLName xml.Name `xml:version_item`
}

// XSD ComplexType declarations

type EntityItemAccessListInterfaceDirectionType struct {
}

type EntityItemAccessListIPVersionType struct {
}

type EntityItemAccessListUseType struct {
}

type EntityItemRoutingAuthTypeStringType struct {
}

type EntityItemRoutingProtocolType struct {
}

type EntityItemSNMPVersionStringType struct {
}

type EntityItemSNMPSecLevelStringType struct {
}

type EntityItemSNMPModeStringType struct {
}

type EntityItemSNMPAuthStringType struct {
}

type EntityItemSNMPPrivStringType struct {
}

type EntityItemSwitchportModeType struct {
}

type EntityItemTrunkEncapType struct {
}
