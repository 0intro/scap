// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://checklists.nist.gov/xccdf/1.2
package cdf

import (
	"encoding/xml"
	"github.com/0intro/scap/pkg/scap/models/cpe_language"
)

// Element
type Benchmark struct {
	XMLName xml.Name `xml:"Benchmark"`

	Id BenchmarkIdType `xml:"id,attr"`

	Id2 string `xml:"Id,attr,omitempty"`

	Resolved bool `xml:"resolved,attr,omitempty"`

	Style string `xml:"style,attr,omitempty"`

	StyleHref string `xml:"style-href,attr,omitempty"`

	XmlLang string `xml:"lang,attr"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Title []TextType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Notice []NoticeType `xml:"notice"`

	FrontMatter []HtmlTextWithSubType `xml:"front-matter"`

	RearMatter []HtmlTextWithSubType `xml:"rear-matter"`

	Reference []ReferenceType `xml:"reference"`

	PlainText []PlainTextType `xml:"plain-text"`

	PlatformSpecification *cpe_language.PlatformSpecificationType `xml:"platform-specification"`

	Platform []CPE2IdrefType `xml:"platform"`

	Version VersionType `xml:"version"`

	Metadata []MetadataType `xml:"metadata"`

	Model []Model `xml:"model"`

	Profile []ProfileType `xml:"Profile"`

	Value []ValueType `xml:"Value"`

	TestResult []TestResultType `xml:"TestResult"`

	Signature *SignatureType `xml:"signature"`

	Group []GroupType `xml:"Group"`

	Rule []RuleType `xml:"Rule"`
}

// Element
type Status struct {
	XMLName xml.Name `xml:"status"`

	Date string `xml:"date,attr,omitempty"`

	Text string `xml:",chardata"`
}

// Element
type Model struct {
	XMLName xml.Name `xml:"model"`

	System string `xml:"system,attr"`

	Param []ParamType `xml:",any"`
}

// Element
type Item struct {
	XMLName xml.Name `xml:"Item"`

	Abstract bool `xml:"abstract,attr,omitempty"`

	ClusterId string `xml:"cluster-id,attr,omitempty"`

	Extends string `xml:"extends,attr,omitempty"`

	Hidden bool `xml:"hidden,attr,omitempty"`

	ProhibitChanges bool `xml:"prohibitChanges,attr,omitempty"`

	XmlLang string `xml:"lang,attr"`

	XmlBase string `xml:"base,attr"`

	Id string `xml:"Id,attr,omitempty"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version *VersionType `xml:"version"`

	Title []TextWithSubType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Warning []WarningType `xml:"warning"`

	Question []TextType `xml:"question"`

	Reference []ReferenceType `xml:"reference"`

	Metadata []MetadataType `xml:"metadata"`
}

// Element
type Group struct {
	XMLName xml.Name `xml:"Group"`

	Id string `xml:"id,attr"`

	Selected string `xml:"selected,attr,omitempty"`

	Weight string `xml:"weight,attr,omitempty"`

	Abstract bool `xml:"abstract,attr,omitempty"`

	ClusterId string `xml:"cluster-id,attr,omitempty"`

	Extends string `xml:"extends,attr,omitempty"`

	Hidden bool `xml:"hidden,attr,omitempty"`

	ProhibitChanges bool `xml:"prohibitChanges,attr,omitempty"`

	XmlLang string `xml:"lang,attr"`

	XmlBase string `xml:"base,attr"`

	Value []ValueType `xml:"Value"`

	Signature *SignatureType `xml:"signature"`

	Group []GroupType `xml:"Group"`

	Rule []RuleType `xml:"Rule"`

	Rationale []HtmlTextWithSubType `xml:"rationale"`

	Platform []OverrideableCPE2IdrefType `xml:"platform"`

	Requires []IdrefListType `xml:"requires"`

	Conflicts []IdrefType `xml:"conflicts"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version *VersionType `xml:"version"`

	Title []TextWithSubType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Warning []WarningType `xml:"warning"`

	Question []TextType `xml:"question"`

	Reference []ReferenceType `xml:"reference"`

	Metadata []MetadataType `xml:"metadata"`
}

// Element
type Rule struct {
	XMLName xml.Name `xml:"Rule"`

	Id string `xml:"id,attr"`

	Role string `xml:"role,attr,omitempty"`

	Severity string `xml:"severity,attr,omitempty"`

	Multiple string `xml:"multiple,attr,omitempty"`

	Selected string `xml:"selected,attr,omitempty"`

	Weight string `xml:"weight,attr,omitempty"`

	Abstract bool `xml:"abstract,attr,omitempty"`

	ClusterId string `xml:"cluster-id,attr,omitempty"`

	Extends string `xml:"extends,attr,omitempty"`

	Hidden bool `xml:"hidden,attr,omitempty"`

	ProhibitChanges bool `xml:"prohibitChanges,attr,omitempty"`

	XmlLang string `xml:"lang,attr"`

	XmlBase string `xml:"base,attr"`

	Ident []IdentType `xml:"ident"`

	ImpactMetric string `xml:"impact-metric"`

	ProfileNote []ProfileNoteType `xml:"profile-note"`

	Fixtext []FixTextType `xml:"fixtext"`

	Fix []FixType `xml:"fix"`

	Signature *SignatureType `xml:"signature"`

	Check []CheckType `xml:"check"`

	ComplexCheck *ComplexCheckType `xml:"complex-check"`

	Rationale []HtmlTextWithSubType `xml:"rationale"`

	Platform []OverrideableCPE2IdrefType `xml:"platform"`

	Requires []IdrefListType `xml:"requires"`

	Conflicts []IdrefType `xml:"conflicts"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version *VersionType `xml:"version"`

	Title []TextWithSubType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Warning []WarningType `xml:"warning"`

	Question []TextType `xml:"question"`

	Reference []ReferenceType `xml:"reference"`

	Metadata []MetadataType `xml:"metadata"`
}

// Element
type Value struct {
	XMLName xml.Name `xml:"Value"`

	Id string `xml:"id,attr"`

	Type string `xml:"type,attr,omitempty"`

	Operator string `xml:"operator,attr,omitempty"`

	Interactive string `xml:"interactive,attr,omitempty"`

	InterfaceHint string `xml:"interfaceHint,attr,omitempty"`

	Abstract bool `xml:"abstract,attr,omitempty"`

	ClusterId string `xml:"cluster-id,attr,omitempty"`

	Extends string `xml:"extends,attr,omitempty"`

	Hidden bool `xml:"hidden,attr,omitempty"`

	ProhibitChanges bool `xml:"prohibitChanges,attr,omitempty"`

	XmlLang string `xml:"lang,attr"`

	XmlBase string `xml:"base,attr"`

	Match []SelStringType `xml:"match"`

	LowerBound []SelNumType `xml:"lower-bound"`

	UpperBound []SelNumType `xml:"upper-bound"`

	Choices []SelChoicesType `xml:"choices"`

	Source []UriRefType `xml:"source"`

	Signature *SignatureType `xml:"signature"`

	Value []SelStringType `xml:"value"`

	ComplexValue []SelComplexValueType `xml:"complex-value"`

	Default []SelStringType `xml:"default"`

	ComplexDefault []SelComplexValueType `xml:"complex-default"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version *VersionType `xml:"version"`

	Title []TextWithSubType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Warning []WarningType `xml:"warning"`

	Question []TextType `xml:"question"`

	Reference []ReferenceType `xml:"reference"`

	Metadata []MetadataType `xml:"metadata"`
}

// Element
type Profile struct {
	XMLName xml.Name `xml:"Profile"`

	Id ProfileIdType `xml:"id,attr"`

	ProhibitChanges bool `xml:"prohibitChanges,attr,omitempty"`

	Abstract bool `xml:"abstract,attr,omitempty"`

	NoteTag string `xml:"note-tag,attr,omitempty"`

	Extends string `xml:"extends,attr,omitempty"`

	XmlBase string `xml:"base,attr"`

	Id2 string `xml:"Id,attr,omitempty"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version *VersionType `xml:"version"`

	Title []TextWithSubType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Reference []ReferenceType `xml:"reference"`

	Platform []OverrideableCPE2IdrefType `xml:"platform"`

	Metadata []MetadataType `xml:"metadata"`

	Signature *SignatureType `xml:"signature"`

	Select []ProfileSelectType `xml:"select"`

	SetComplexValue []ProfileSetComplexValueType `xml:"set-complex-value"`

	SetValue []ProfileSetValueType `xml:"set-value"`

	RefineValue []ProfileRefineValueType `xml:"refine-value"`

	RefineRule []ProfileRefineRuleType `xml:"refine-rule"`
}

// Element
type TestResult struct {
	XMLName xml.Name `xml:"TestResult"`

	Id TestresultIdType `xml:"id,attr"`

	StartTime string `xml:"start-time,attr,omitempty"`

	EndTime string `xml:"end-time,attr"`

	TestSystem string `xml:"test-system,attr,omitempty"`

	Version string `xml:"version,attr,omitempty"`

	Id2 string `xml:"Id,attr,omitempty"`

	Benchmark *BenchmarkReferenceType `xml:"benchmark"`

	TailoringFile *TailoringReferenceType `xml:"tailoring-file"`

	Title []TextType `xml:"title"`

	Remark []TextType `xml:"remark"`

	Organization []string `xml:"organization"`

	Identity *IdentityType `xml:"identity"`

	Profile *IdrefType `xml:"profile"`

	Target []string `xml:"target"`

	TargetAddress []string `xml:"target-address"`

	TargetFacts *TargetFactsType `xml:"target-facts"`

	Platform []CPE2IdrefType `xml:"platform"`

	RuleResult []RuleResultType `xml:"rule-result"`

	Score []ScoreType `xml:"score"`

	Metadata []MetadataType `xml:"metadata"`

	Signature *SignatureType `xml:"signature"`

	TargetIdRef []TargetIdRefType `xml:"target-id-ref"`

	SetValue []ProfileSetValueType `xml:"set-value"`

	SetComplexValue []ProfileSetComplexValueType `xml:"set-complex-value"`
}

// Element
type Tailoring struct {
	XMLName xml.Name `xml:"Tailoring"`

	Id TailoringIdType `xml:"id,attr"`

	Id2 string `xml:"Id,attr,omitempty"`

	Benchmark *TailoringBenchmarkReferenceType `xml:"benchmark"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version TailoringVersionType `xml:"version"`

	Metadata []MetadataType `xml:"metadata"`

	Profile []ProfileType `xml:"Profile"`

	Signature *SignatureType `xml:"signature"`
}

// XSD ComplexType declarations

type NoticeType struct {
	XMLName xml.Name

	Id string `xml:"id,attr"`

	XmlBase string `xml:"base,attr"`

	XmlLang string `xml:"lang,attr"`

	InnerXml string `xml:",innerxml"`
}

type DcStatusType struct {
	XMLName xml.Name

	InnerXml string `xml:",innerxml"`
}

type PlainTextType struct {
	XMLName xml.Name

	Id string `xml:"id,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type ReferenceType struct {
	XMLName xml.Name

	Href string `xml:"href,attr"`

	Override bool `xml:"override,attr"`

	InnerXml string `xml:",innerxml"`
}

type SignatureType struct {
	XMLName xml.Name

	InnerXml string `xml:",innerxml"`
}

type MetadataType struct {
	XMLName xml.Name

	InnerXml string `xml:",innerxml"`
}

type ParamType struct {
	XMLName xml.Name

	Name string `xml:"name,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type VersionType struct {
	XMLName xml.Name

	Time string `xml:"time,attr,omitempty"`

	Update string `xml:"update,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type TextType struct {
	XMLName xml.Name

	XmlLang string `xml:"lang,attr"`

	Override bool `xml:"override,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type HtmlTextType struct {
	XMLName xml.Name

	XmlLang string `xml:"lang,attr"`

	Override bool `xml:"override,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type HtmlTextWithSubType struct {
	XMLName xml.Name

	XmlLang string `xml:"lang,attr"`

	Override bool `xml:"override,attr,omitempty"`

	Sub []SubType `xml:"sub"`

	InnerXml string `xml:",innerxml"`
}

type ProfileNoteType struct {
	XMLName xml.Name

	XmlLang string `xml:"lang,attr"`

	Tag string `xml:"tag,attr"`

	Sub []SubType `xml:"sub"`

	InnerXml string `xml:",innerxml"`
}

type TextWithSubType struct {
	XMLName xml.Name

	XmlLang string `xml:"lang,attr"`

	Override bool `xml:"override,attr,omitempty"`

	Sub []SubType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type SubType struct {
	XMLName xml.Name

	Use string `xml:"use,attr,omitempty"`

	Idref string `xml:"idref,attr"`

	InnerXml string `xml:",innerxml"`
}

type IdrefType struct {
	XMLName xml.Name

	Idref string `xml:"idref,attr"`

	InnerXml string `xml:",innerxml"`
}

type IdrefListType struct {
	XMLName xml.Name

	Idref string `xml:"idref,attr"`

	InnerXml string `xml:",innerxml"`
}

type CPE2IdrefType struct {
	XMLName xml.Name

	Idref string `xml:"idref,attr"`

	InnerXml string `xml:",innerxml"`
}

type OverrideableCPE2IdrefType struct {
	XMLName xml.Name

	Override string `xml:"override,attr,omitempty"`

	Idref string `xml:"idref,attr"`

	InnerXml string `xml:",innerxml"`
}

type ItemType struct {
	XMLName xml.Name

	Abstract bool `xml:"abstract,attr,omitempty"`

	ClusterId string `xml:"cluster-id,attr,omitempty"`

	Extends string `xml:"extends,attr,omitempty"`

	Hidden bool `xml:"hidden,attr,omitempty"`

	ProhibitChanges bool `xml:"prohibitChanges,attr,omitempty"`

	XmlLang string `xml:"lang,attr"`

	XmlBase string `xml:"base,attr"`

	Id string `xml:"Id,attr,omitempty"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version *VersionType `xml:"version"`

	Title []TextWithSubType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Warning []WarningType `xml:"warning"`

	Question []TextType `xml:"question"`

	Reference []ReferenceType `xml:"reference"`

	Metadata []MetadataType `xml:"metadata"`

	InnerXml string `xml:",innerxml"`
}

type SelectableItemType struct {
	XMLName xml.Name

	Selected string `xml:"selected,attr,omitempty"`

	Weight string `xml:"weight,attr,omitempty"`

	Abstract bool `xml:"abstract,attr,omitempty"`

	ClusterId string `xml:"cluster-id,attr,omitempty"`

	Extends string `xml:"extends,attr,omitempty"`

	Hidden bool `xml:"hidden,attr,omitempty"`

	ProhibitChanges bool `xml:"prohibitChanges,attr,omitempty"`

	XmlLang string `xml:"lang,attr"`

	XmlBase string `xml:"base,attr"`

	Id string `xml:"Id,attr,omitempty"`

	Rationale []HtmlTextWithSubType `xml:"rationale"`

	Platform []OverrideableCPE2IdrefType `xml:"platform"`

	Requires []IdrefListType `xml:"requires"`

	Conflicts []IdrefType `xml:"conflicts"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version *VersionType `xml:"version"`

	Title []TextWithSubType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Warning []WarningType `xml:"warning"`

	Question []TextType `xml:"question"`

	Reference []ReferenceType `xml:"reference"`

	Metadata []MetadataType `xml:"metadata"`

	InnerXml string `xml:",innerxml"`
}

type GroupType struct {
	XMLName xml.Name

	Id string `xml:"id,attr"`

	Selected string `xml:"selected,attr,omitempty"`

	Weight string `xml:"weight,attr,omitempty"`

	Abstract bool `xml:"abstract,attr,omitempty"`

	ClusterId string `xml:"cluster-id,attr,omitempty"`

	Extends string `xml:"extends,attr,omitempty"`

	Hidden bool `xml:"hidden,attr,omitempty"`

	ProhibitChanges bool `xml:"prohibitChanges,attr,omitempty"`

	XmlLang string `xml:"lang,attr"`

	XmlBase string `xml:"base,attr"`

	Value []ValueType `xml:"Value"`

	Signature *SignatureType `xml:"signature"`

	Group []GroupType `xml:"Group"`

	Rule []RuleType `xml:"Rule"`

	Rationale []HtmlTextWithSubType `xml:"rationale"`

	Platform []OverrideableCPE2IdrefType `xml:"platform"`

	Requires []IdrefListType `xml:"requires"`

	Conflicts []IdrefType `xml:"conflicts"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version *VersionType `xml:"version"`

	Title []TextWithSubType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Warning []WarningType `xml:"warning"`

	Question []TextType `xml:"question"`

	Reference []ReferenceType `xml:"reference"`

	Metadata []MetadataType `xml:"metadata"`

	InnerXml string `xml:",innerxml"`
}

type RuleType struct {
	XMLName xml.Name

	Id string `xml:"id,attr"`

	Role string `xml:"role,attr,omitempty"`

	Severity string `xml:"severity,attr,omitempty"`

	Multiple string `xml:"multiple,attr,omitempty"`

	Selected string `xml:"selected,attr,omitempty"`

	Weight string `xml:"weight,attr,omitempty"`

	Abstract bool `xml:"abstract,attr,omitempty"`

	ClusterId string `xml:"cluster-id,attr,omitempty"`

	Extends string `xml:"extends,attr,omitempty"`

	Hidden bool `xml:"hidden,attr,omitempty"`

	ProhibitChanges bool `xml:"prohibitChanges,attr,omitempty"`

	XmlLang string `xml:"lang,attr"`

	XmlBase string `xml:"base,attr"`

	Ident []IdentType `xml:"ident"`

	ImpactMetric string `xml:"impact-metric"`

	ProfileNote []ProfileNoteType `xml:"profile-note"`

	Fixtext []FixTextType `xml:"fixtext"`

	Fix []FixType `xml:"fix"`

	Signature *SignatureType `xml:"signature"`

	Check []CheckType `xml:"check"`

	ComplexCheck *ComplexCheckType `xml:"complex-check"`

	Rationale []HtmlTextWithSubType `xml:"rationale"`

	Platform []OverrideableCPE2IdrefType `xml:"platform"`

	Requires []IdrefListType `xml:"requires"`

	Conflicts []IdrefType `xml:"conflicts"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version *VersionType `xml:"version"`

	Title []TextWithSubType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Warning []WarningType `xml:"warning"`

	Question []TextType `xml:"question"`

	Reference []ReferenceType `xml:"reference"`

	Metadata []MetadataType `xml:"metadata"`

	InnerXml string `xml:",innerxml"`
}

type IdentType struct {
	XMLName xml.Name

	System string `xml:"system,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type WarningType struct {
	XMLName xml.Name

	Category string `xml:"category,attr,omitempty"`

	XmlLang string `xml:"lang,attr"`

	Override bool `xml:"override,attr,omitempty"`

	Sub []SubType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type FixTextType struct {
	XMLName xml.Name

	Fixref string `xml:"fixref,attr,omitempty"`

	Reboot bool `xml:"reboot,attr,omitempty"`

	Strategy FixStrategyEnumType `xml:"strategy,attr,omitempty"`

	Disruption RatingEnumType `xml:"disruption,attr,omitempty"`

	Complexity RatingEnumType `xml:"complexity,attr,omitempty"`

	XmlLang string `xml:"lang,attr"`

	Override bool `xml:"override,attr,omitempty"`

	Sub []SubType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type FixType struct {
	XMLName xml.Name

	Id string `xml:"id,attr,omitempty"`

	Reboot bool `xml:"reboot,attr,omitempty"`

	Strategy FixStrategyEnumType `xml:"strategy,attr,omitempty"`

	Disruption RatingEnumType `xml:"disruption,attr,omitempty"`

	Complexity RatingEnumType `xml:"complexity,attr,omitempty"`

	System string `xml:"system,attr,omitempty"`

	Platform string `xml:"platform,attr,omitempty"`

	Sub []SubType `xml:"sub"`

	Instance []InstanceFixType `xml:"instance"`

	InnerXml string `xml:",innerxml"`
}

type InstanceFixType struct {
	XMLName xml.Name

	Context string `xml:"context,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type ComplexCheckType struct {
	XMLName xml.Name

	Operator CcOperatorEnumType `xml:"operator,attr"`

	Negate bool `xml:"negate,attr,omitempty"`

	Check []CheckType `xml:"check"`

	ComplexCheck []ComplexCheckType `xml:"complex-check"`

	InnerXml string `xml:",innerxml"`
}

type CheckType struct {
	XMLName xml.Name

	System string `xml:"system,attr"`

	Negate bool `xml:"negate,attr,omitempty"`

	Id string `xml:"id,attr,omitempty"`

	Selector string `xml:"selector,attr,omitempty"`

	MultiCheck bool `xml:"multi-check,attr,omitempty"`

	XmlBase string `xml:"base,attr"`

	CheckImport []CheckImportType `xml:"check-import"`

	CheckExport []CheckExportType `xml:"check-export"`

	CheckContentRef []CheckContentRefType `xml:"check-content-ref"`

	CheckContent *CheckContentType `xml:"check-content"`

	InnerXml string `xml:",innerxml"`
}

type CheckImportType struct {
	XMLName xml.Name

	ImportName string `xml:"import-name,attr"`

	ImportXpath string `xml:"import-xpath,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type CheckExportType struct {
	XMLName xml.Name

	ValueId string `xml:"value-id,attr"`

	ExportName string `xml:"export-name,attr"`

	InnerXml string `xml:",innerxml"`
}

type CheckContentRefType struct {
	XMLName xml.Name

	Href string `xml:"href,attr"`

	Name string `xml:"name,attr"`

	InnerXml string `xml:",innerxml"`
}

type CheckContentType struct {
	XMLName xml.Name

	InnerXml string `xml:",innerxml"`
}

type ValueType struct {
	XMLName xml.Name

	Id string `xml:"id,attr"`

	Type string `xml:"type,attr,omitempty"`

	Operator string `xml:"operator,attr,omitempty"`

	Interactive string `xml:"interactive,attr,omitempty"`

	InterfaceHint string `xml:"interfaceHint,attr,omitempty"`

	Abstract bool `xml:"abstract,attr,omitempty"`

	ClusterId string `xml:"cluster-id,attr,omitempty"`

	Extends string `xml:"extends,attr,omitempty"`

	Hidden bool `xml:"hidden,attr,omitempty"`

	ProhibitChanges bool `xml:"prohibitChanges,attr,omitempty"`

	XmlLang string `xml:"lang,attr"`

	XmlBase string `xml:"base,attr"`

	Match []SelStringType `xml:"match"`

	LowerBound []SelNumType `xml:"lower-bound"`

	UpperBound []SelNumType `xml:"upper-bound"`

	Choices []SelChoicesType `xml:"choices"`

	Source []UriRefType `xml:"source"`

	Signature *SignatureType `xml:"signature"`

	Value []SelStringType `xml:"value"`

	ComplexValue []SelComplexValueType `xml:"complex-value"`

	Default []SelStringType `xml:"default"`

	ComplexDefault []SelComplexValueType `xml:"complex-default"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version *VersionType `xml:"version"`

	Title []TextWithSubType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Warning []WarningType `xml:"warning"`

	Question []TextType `xml:"question"`

	Reference []ReferenceType `xml:"reference"`

	Metadata []MetadataType `xml:"metadata"`

	InnerXml string `xml:",innerxml"`
}

type ComplexValueType struct {
	XMLName xml.Name

	Item []string `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type SelComplexValueType struct {
	XMLName xml.Name

	Selector string `xml:"selector,attr,omitempty"`

	Item []string `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type SelChoicesType struct {
	XMLName xml.Name

	MustMatch bool `xml:"mustMatch,attr,omitempty"`

	Selector string `xml:"selector,attr,omitempty"`

	Choice []string `xml:"choice"`

	ComplexChoice []ComplexValueType `xml:"complex-choice"`

	InnerXml string `xml:",innerxml"`
}

type SelStringType struct {
	XMLName xml.Name

	Selector string `xml:"selector,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type SelNumType struct {
	XMLName xml.Name

	Selector string `xml:"selector,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type UriRefType struct {
	XMLName xml.Name

	Uri string `xml:"uri,attr"`

	InnerXml string `xml:",innerxml"`
}

type ProfileType struct {
	XMLName xml.Name

	Id ProfileIdType `xml:"id,attr"`

	ProhibitChanges bool `xml:"prohibitChanges,attr,omitempty"`

	Abstract bool `xml:"abstract,attr,omitempty"`

	NoteTag string `xml:"note-tag,attr,omitempty"`

	Extends string `xml:"extends,attr,omitempty"`

	XmlBase string `xml:"base,attr"`

	Id2 string `xml:"Id,attr,omitempty"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version *VersionType `xml:"version"`

	Title []TextWithSubType `xml:"title"`

	Description []HtmlTextWithSubType `xml:"description"`

	Reference []ReferenceType `xml:"reference"`

	Platform []OverrideableCPE2IdrefType `xml:"platform"`

	Metadata []MetadataType `xml:"metadata"`

	Signature *SignatureType `xml:"signature"`

	Select []ProfileSelectType `xml:"select"`

	SetComplexValue []ProfileSetComplexValueType `xml:"set-complex-value"`

	SetValue []ProfileSetValueType `xml:"set-value"`

	RefineValue []ProfileRefineValueType `xml:"refine-value"`

	RefineRule []ProfileRefineRuleType `xml:"refine-rule"`

	InnerXml string `xml:",innerxml"`
}

type ProfileSelectType struct {
	XMLName xml.Name

	Idref string `xml:"idref,attr"`

	Selected bool `xml:"selected,attr"`

	Remark []TextType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type ProfileSetValueType struct {
	XMLName xml.Name

	Idref string `xml:"idref,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type ProfileSetComplexValueType struct {
	XMLName xml.Name

	Idref string `xml:"idref,attr"`

	Item []string `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type ProfileRefineValueType struct {
	XMLName xml.Name

	Idref string `xml:"idref,attr"`

	Selector string `xml:"selector,attr,omitempty"`

	Operator ValueOperatorType `xml:"operator,attr,omitempty"`

	Remark []TextType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type ProfileRefineRuleType struct {
	XMLName xml.Name

	Idref string `xml:"idref,attr"`

	Weight WeightType `xml:"weight,attr,omitempty"`

	Selector string `xml:"selector,attr,omitempty"`

	Severity SeverityEnumType `xml:"severity,attr,omitempty"`

	Role RoleEnumType `xml:"role,attr,omitempty"`

	Remark []TextType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type TestResultType struct {
	XMLName xml.Name

	Id TestresultIdType `xml:"id,attr"`

	StartTime string `xml:"start-time,attr,omitempty"`

	EndTime string `xml:"end-time,attr"`

	TestSystem string `xml:"test-system,attr,omitempty"`

	Version string `xml:"version,attr,omitempty"`

	Id2 string `xml:"Id,attr,omitempty"`

	Benchmark *BenchmarkReferenceType `xml:"benchmark"`

	TailoringFile *TailoringReferenceType `xml:"tailoring-file"`

	Title []TextType `xml:"title"`

	Remark []TextType `xml:"remark"`

	Organization []string `xml:"organization"`

	Identity *IdentityType `xml:"identity"`

	Profile *IdrefType `xml:"profile"`

	Target []string `xml:"target"`

	TargetAddress []string `xml:"target-address"`

	TargetFacts *TargetFactsType `xml:"target-facts"`

	Platform []CPE2IdrefType `xml:"platform"`

	RuleResult []RuleResultType `xml:"rule-result"`

	Score []ScoreType `xml:"score"`

	Metadata []MetadataType `xml:"metadata"`

	Signature *SignatureType `xml:"signature"`

	TargetIdRef []TargetIdRefType `xml:"target-id-ref"`

	SetValue []ProfileSetValueType `xml:"set-value"`

	SetComplexValue []ProfileSetComplexValueType `xml:"set-complex-value"`

	InnerXml string `xml:",innerxml"`
}

type BenchmarkReferenceType struct {
	XMLName xml.Name

	Href string `xml:"href,attr"`

	Id string `xml:"id,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type ScoreType struct {
	XMLName xml.Name

	System string `xml:"system,attr,omitempty"`

	Maximum float64 `xml:"maximum,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type TargetFactsType struct {
	XMLName xml.Name

	Fact []FactType `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type TargetIdRefType struct {
	XMLName xml.Name

	System string `xml:"system,attr"`

	Href string `xml:"href,attr"`

	Name string `xml:"name,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type IdentityType struct {
	XMLName xml.Name

	Authenticated bool `xml:"authenticated,attr"`

	Privileged bool `xml:"privileged,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type FactType struct {
	XMLName xml.Name

	Name string `xml:"name,attr"`

	Type ValueTypeType `xml:"type,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type TailoringReferenceType struct {
	XMLName xml.Name

	Href string `xml:"href,attr"`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Time string `xml:"time,attr"`

	InnerXml string `xml:",innerxml"`
}

type RuleResultType struct {
	XMLName xml.Name

	Idref string `xml:"idref,attr"`

	Role RoleEnumType `xml:"role,attr,omitempty"`

	Severity SeverityEnumType `xml:"severity,attr,omitempty"`

	Time string `xml:"time,attr,omitempty"`

	Version string `xml:"version,attr,omitempty"`

	Weight WeightType `xml:"weight,attr,omitempty"`

	Result ResultEnumType `xml:"result"`

	Override []OverrideType `xml:"override"`

	Ident []IdentType `xml:"ident"`

	Metadata []MetadataType `xml:"metadata"`

	Message []MessageType `xml:"message"`

	Instance []InstanceResultType `xml:"instance"`

	Fix []FixType `xml:"fix"`

	Check []CheckType `xml:"check"`

	ComplexCheck *ComplexCheckType `xml:"complex-check"`

	InnerXml string `xml:",innerxml"`
}

type InstanceResultType struct {
	XMLName xml.Name

	Context string `xml:"context,attr,omitempty"`

	ParentContext string `xml:"parentContext,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type OverrideType struct {
	XMLName xml.Name

	Time string `xml:"time,attr"`

	Authority string `xml:"authority,attr"`

	OldResult ResultEnumType `xml:"old-result"`

	NewResult ResultEnumType `xml:"new-result"`

	Remark TextType `xml:"remark"`

	InnerXml string `xml:",innerxml"`
}

type MessageType struct {
	XMLName xml.Name

	Severity MsgSevEnumType `xml:"severity,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type TailoringType struct {
	XMLName xml.Name

	Id TailoringIdType `xml:"id,attr"`

	Id2 string `xml:"Id,attr,omitempty"`

	Benchmark *TailoringBenchmarkReferenceType `xml:"benchmark"`

	Status []Status `xml:"status"`

	DcStatus []DcStatusType `xml:"dc-status"`

	Version TailoringVersionType `xml:"version"`

	Metadata []MetadataType `xml:"metadata"`

	Profile []ProfileType `xml:"Profile"`

	Signature *SignatureType `xml:"signature"`

	InnerXml string `xml:",innerxml"`
}

type TailoringBenchmarkReferenceType struct {
	XMLName xml.Name

	Version string `xml:"version,attr,omitempty"`

	Href string `xml:"href,attr"`

	Id string `xml:"id,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type TailoringVersionType struct {
	XMLName xml.Name

	Time string `xml:"time,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations

type StatusType string

const StatusTypeAccepted StatusType = "accepted"

const StatusTypeDeprecated StatusType = "deprecated"

const StatusTypeDraft StatusType = "draft"

const StatusTypeIncomplete StatusType = "incomplete"

const StatusTypeInterim StatusType = "interim"

type BenchmarkIdType string

type RuleIdType string

type GroupIdType string

type ValueIdType string

type ProfileIdType string

type TestresultIdType string

type TailoringIdType string

type WarningCategoryEnumType string

const WarningCategoryEnumTypeGeneral WarningCategoryEnumType = "general"

const WarningCategoryEnumTypeFunctionality WarningCategoryEnumType = "functionality"

const WarningCategoryEnumTypePerformance WarningCategoryEnumType = "performance"

const WarningCategoryEnumTypeHardware WarningCategoryEnumType = "hardware"

const WarningCategoryEnumTypeLegal WarningCategoryEnumType = "legal"

const WarningCategoryEnumTypeRegulatory WarningCategoryEnumType = "regulatory"

const WarningCategoryEnumTypeManagement WarningCategoryEnumType = "management"

const WarningCategoryEnumTypeAudit WarningCategoryEnumType = "audit"

const WarningCategoryEnumTypeDependency WarningCategoryEnumType = "dependency"

type FixStrategyEnumType string

const FixStrategyEnumTypeUnknown FixStrategyEnumType = "unknown"

const FixStrategyEnumTypeConfigure FixStrategyEnumType = "configure"

const FixStrategyEnumTypeCombination FixStrategyEnumType = "combination"

const FixStrategyEnumTypeDisable FixStrategyEnumType = "disable"

const FixStrategyEnumTypeEnable FixStrategyEnumType = "enable"

const FixStrategyEnumTypePatch FixStrategyEnumType = "patch"

const FixStrategyEnumTypePolicy FixStrategyEnumType = "policy"

const FixStrategyEnumTypeRestrict FixStrategyEnumType = "restrict"

const FixStrategyEnumTypeUpdate FixStrategyEnumType = "update"

type RatingEnumType string

const RatingEnumTypeUnknown RatingEnumType = "unknown"

const RatingEnumTypeLow RatingEnumType = "low"

const RatingEnumTypeMedium RatingEnumType = "medium"

const RatingEnumTypeHigh RatingEnumType = "high"

type CcOperatorEnumType string

const CcOperatorEnumTypeOr CcOperatorEnumType = "OR"

const CcOperatorEnumTypeAnd CcOperatorEnumType = "AND"

type WeightType float64

type ValueTypeType string

const ValueTypeTypeNumber ValueTypeType = "number"

const ValueTypeTypeString ValueTypeType = "string"

const ValueTypeTypeBoolean ValueTypeType = "boolean"

type ValueOperatorType string

const ValueOperatorTypeEquals ValueOperatorType = "equals"

const ValueOperatorTypeNotEqual ValueOperatorType = "not equal"

const ValueOperatorTypeGreaterThan ValueOperatorType = "greater than"

const ValueOperatorTypeLessThan ValueOperatorType = "less than"

const ValueOperatorTypeGreaterThanOrEqual ValueOperatorType = "greater than or equal"

const ValueOperatorTypeLessThanOrEqual ValueOperatorType = "less than or equal"

const ValueOperatorTypePatternMatch ValueOperatorType = "pattern match"

type InterfaceHintType string

const InterfaceHintTypeChoice InterfaceHintType = "choice"

const InterfaceHintTypeTextline InterfaceHintType = "textline"

const InterfaceHintTypeText InterfaceHintType = "text"

const InterfaceHintTypeDate InterfaceHintType = "date"

const InterfaceHintTypeDatetime InterfaceHintType = "datetime"

type MsgSevEnumType string

const MsgSevEnumTypeError MsgSevEnumType = "error"

const MsgSevEnumTypeWarning MsgSevEnumType = "warning"

const MsgSevEnumTypeInfo MsgSevEnumType = "info"

type ResultEnumType string

const ResultEnumTypePass ResultEnumType = "pass"

const ResultEnumTypeFail ResultEnumType = "fail"

const ResultEnumTypeError ResultEnumType = "error"

const ResultEnumTypeUnknown ResultEnumType = "unknown"

const ResultEnumTypeNotapplicable ResultEnumType = "notapplicable"

const ResultEnumTypeNotchecked ResultEnumType = "notchecked"

const ResultEnumTypeNotselected ResultEnumType = "notselected"

const ResultEnumTypeInformational ResultEnumType = "informational"

const ResultEnumTypeFixed ResultEnumType = "fixed"

type SeverityEnumType string

const SeverityEnumTypeUnknown SeverityEnumType = "unknown"

const SeverityEnumTypeInfo SeverityEnumType = "info"

const SeverityEnumTypeLow SeverityEnumType = "low"

const SeverityEnumTypeMedium SeverityEnumType = "medium"

const SeverityEnumTypeHigh SeverityEnumType = "high"

type RoleEnumType string

const RoleEnumTypeFull RoleEnumType = "full"

const RoleEnumTypeUnscored RoleEnumType = "unscored"

const RoleEnumTypeUnchecked RoleEnumType = "unchecked"

type SubUseEnumType string

const SubUseEnumTypeValue SubUseEnumType = "value"

const SubUseEnumTypeTitle SubUseEnumType = "title"

const SubUseEnumTypeLegacy SubUseEnumType = "legacy"
