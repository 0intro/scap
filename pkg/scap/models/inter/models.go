// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://scap.nist.gov/schema/ocil/2.0
package inter

import (
	"encoding/xml"
)

// Element
type Ocil struct {
	XMLName xml.Name `xml:ocil`

	Generator GeneratorType `xml:"generator"`

	Document *DocumentType `xml:"document"`

	Questionnaires QuestionnairesType `xml:"questionnaires"`

	TestActions TestActionsType `xml:"test_actions"`

	Questions QuestionsType `xml:"questions"`

	Artifacts *ArtifactsType `xml:"artifacts"`

	Variables *VariablesType `xml:"variables"`

	Results *ResultsType `xml:"results"`
}

// Element
type TestAction struct {
	XMLName xml.Name `xml:test_action`

	Revision string `xml:"revision,attr"`

	Notes []string `xml:"notes"`
}

// Element
type QuestionTestAction struct {
	XMLName xml.Name `xml:question_test_action`

	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`
}

// Element
type BooleanQuestionTestAction struct {
	XMLName xml.Name `xml:boolean_question_test_action`

	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	WhenTrue TestActionConditionType `xml:"when_true"`

	WhenFalse TestActionConditionType `xml:"when_false"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`
}

// Element
type ChoiceQuestionTestAction struct {
	XMLName xml.Name `xml:choice_question_test_action`

	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	WhenChoice []ChoiceTestActionConditionType `xml:"when_choice"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`
}

// Element
type NumericQuestionTestAction struct {
	XMLName xml.Name `xml:numeric_question_test_action`

	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`
}

// Element
type StringQuestionTestAction struct {
	XMLName xml.Name `xml:string_question_test_action`

	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	WhenPattern []PatternTestActionConditionType `xml:"when_pattern"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`
}

// Element
type Question struct {
	XMLName xml.Name `xml:question`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`
}

// Element
type BooleanQuestion struct {
	XMLName xml.Name `xml:boolean_question`

	DefaultAnswer string `xml:"default_answer,attr,omitempty"`

	Model string `xml:"model,attr,omitempty"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`
}

// Element
type ChoiceQuestion struct {
	XMLName xml.Name `xml:choice_question`

	DefaultAnswerRef string `xml:"default_answer_ref,attr,omitempty"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`
}

// Element
type NumericQuestion struct {
	XMLName xml.Name `xml:numeric_question`

	DefaultAnswer string `xml:"default_answer,attr,omitempty"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`
}

// Element
type StringQuestion struct {
	XMLName xml.Name `xml:string_question`

	DefaultAnswer string `xml:"default_answer,attr,omitempty"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`
}

// Element
type Variable struct {
	XMLName xml.Name `xml:variable`

	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`
}

// Element
type ConstantVariable struct {
	XMLName xml.Name `xml:constant_variable`

	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Value string `xml:"value"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`
}

// Element
type LocalVariable struct {
	XMLName xml.Name `xml:local_variable`

	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Set string `xml:"set"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`
}

// Element
type ExternalVariable struct {
	XMLName xml.Name `xml:external_variable`

	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`
}

// Element
type Target struct {
	XMLName xml.Name `xml:target`

	Revision string `xml:"revision,attr"`

	Name string `xml:"name"`

	Notes []string `xml:"notes"`
}

// Element
type User struct {
	XMLName xml.Name `xml:user`

	Revision string `xml:"revision,attr"`

	Organization []string `xml:"organization"`

	Position []string `xml:"position"`

	Email []string `xml:"email"`

	Name string `xml:"name"`

	Notes []string `xml:"notes"`
}

// Element
type System struct {
	XMLName xml.Name `xml:system`

	Revision string `xml:"revision,attr"`

	Organization string `xml:"organization"`

	Ipaddress []string `xml:"ipaddress"`

	Description *TextType `xml:"description"`

	Name string `xml:"name"`

	Notes []string `xml:"notes"`
}

// Element
type QuestionResult struct {
	XMLName xml.Name `xml:question_result`

	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr,omitempty"`
}

// Element
type BooleanQuestionResult struct {
	XMLName xml.Name `xml:boolean_question_result`

	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr,omitempty"`

	Answer bool `xml:"answer"`
}

// Element
type ChoiceQuestionResult struct {
	XMLName xml.Name `xml:choice_question_result`

	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr,omitempty"`

	Answer ChoiceAnswerType `xml:"answer"`
}

// Element
type NumericQuestionResult struct {
	XMLName xml.Name `xml:numeric_question_result`

	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr,omitempty"`

	Answer float64 `xml:"answer"`
}

// Element
type StringQuestionResult struct {
	XMLName xml.Name `xml:string_question_result`

	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr,omitempty"`

	Answer string `xml:"answer"`
}

// Element
type ArtifactValue struct {
	XMLName xml.Name `xml:artifact_value`
}

// Element
type TextArtifactValue struct {
	XMLName xml.Name `xml:text_artifact_value`

	MimeType string `xml:"mime_type,attr"`

	Data string `xml:"data"`
}

// Element
type BinaryArtifactValue struct {
	XMLName xml.Name `xml:binary_artifact_value`

	MimeType string `xml:"mime_type,attr"`

	Data string `xml:"data"`
}

// Element
type ReferenceArtifactValue struct {
	XMLName xml.Name `xml:reference_artifact_value`

	Reference Reference `xml:"reference"`
}

// Element
type Expression struct {
	XMLName xml.Name `xml:expression`

	Value string `xml:"value"`
}

// Element
type WhenPattern struct {
	XMLName xml.Name `xml:when_pattern`

	Pattern string `xml:"pattern,attr"`

	Value string `xml:"value"`
}

// Element
type WhenChoice struct {
	XMLName xml.Name `xml:when_choice`

	ChoiceRef string `xml:"choice_ref,attr"`

	Value string `xml:"value"`
}

// Element
type WhenRange struct {
	XMLName xml.Name `xml:when_range`

	Min string `xml:"min,attr"`

	Max string `xml:"max,attr"`

	Value string `xml:"value"`
}

// Element
type WhenBoolean struct {
	XMLName xml.Name `xml:when_boolean`

	Value string `xml:"value,attr"`

	ValueElm string `xml:"value"`
}

// Element
type Reference struct {
	XMLName xml.Name `xml:reference`

	Href string `xml:"href,attr"`
}

// XSD ComplexType declarations

type OCILType struct {
	Generator GeneratorType `xml:"generator"`

	Document *DocumentType `xml:"document"`

	Questionnaires QuestionnairesType `xml:"questionnaires"`

	TestActions TestActionsType `xml:"test_actions"`

	Questions QuestionsType `xml:"questions"`

	Artifacts *ArtifactsType `xml:"artifacts"`

	Variables *VariablesType `xml:"variables"`

	Results *ResultsType `xml:"results"`

	InnerXml string `xml:",innerxml"`
}

type QuestionnairesType struct {
	Questionnaire []QuestionnaireType `xml:"questionnaire"`

	InnerXml string `xml:",innerxml"`
}

type QuestionnaireType struct {
	Id string `xml:"id,attr"`

	ChildOnly string `xml:"child_only,attr,omitempty"`

	Revision string `xml:"revision,attr"`

	Title *TextType `xml:"title"`

	Description *TextType `xml:"description"`

	References *ReferencesType `xml:"references"`

	Actions OperationType `xml:"actions"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type GeneratorType struct {
	ProductName string `xml:"product_name"`

	ProductVersion string `xml:"product_version"`

	Author []UserType `xml:"author"`

	SchemaVersion float64 `xml:"schema_version"`

	Timestamp string `xml:"timestamp"`

	AdditionalData *ExtensionContainerType `xml:"additional_data"`

	InnerXml string `xml:",innerxml"`
}

type ExtensionContainerType struct {
	InnerXml string `xml:",innerxml"`
}

type DocumentType struct {
	Title string `xml:"title"`

	Description []string `xml:"description"`

	Notice []string `xml:"notice"`

	InnerXml string `xml:",innerxml"`
}

type TestActionsType struct {
	TestAction []TestAction `xml:"test_action"`

	InnerXml string `xml:",innerxml"`
}

type QuestionTestActionType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type BooleanQuestionTestActionType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	WhenTrue TestActionConditionType `xml:"when_true"`

	WhenFalse TestActionConditionType `xml:"when_false"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type ChoiceQuestionTestActionType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	WhenChoice []ChoiceTestActionConditionType `xml:"when_choice"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type NumericQuestionTestActionType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type StringQuestionTestActionType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	WhenPattern []PatternTestActionConditionType `xml:"when_pattern"`

	Title *TextType `xml:"title"`

	WhenUnknown *TestActionConditionType `xml:"when_unknown"`

	WhenNotTested *TestActionConditionType `xml:"when_not_tested"`

	WhenNotApplicable *TestActionConditionType `xml:"when_not_applicable"`

	WhenError *TestActionConditionType `xml:"when_error"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type TestActionRefType struct {
	Negate string `xml:"negate,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type ChoiceTestActionConditionType struct {
	ChoiceRef []string `xml:"choice_ref"`

	ArtifactRefs *ArtifactRefsType `xml:"artifact_refs"`

	Result string `xml:"result"`

	TestActionRef *TestActionRefType `xml:"test_action_ref"`

	InnerXml string `xml:",innerxml"`
}

type EqualsTestActionConditionType struct {
	VarRef string `xml:"var_ref,attr,omitempty"`

	Value []float64 `xml:"value"`

	ArtifactRefs *ArtifactRefsType `xml:"artifact_refs"`

	Result string `xml:"result"`

	TestActionRef *TestActionRefType `xml:"test_action_ref"`

	InnerXml string `xml:",innerxml"`
}

type RangeTestActionConditionType struct {
	Range []RangeType `xml:"range"`

	ArtifactRefs *ArtifactRefsType `xml:"artifact_refs"`

	Result string `xml:"result"`

	TestActionRef *TestActionRefType `xml:"test_action_ref"`

	InnerXml string `xml:",innerxml"`
}

type PatternTestActionConditionType struct {
	Pattern []PatternType `xml:"pattern"`

	ArtifactRefs *ArtifactRefsType `xml:"artifact_refs"`

	Result string `xml:"result"`

	TestActionRef *TestActionRefType `xml:"test_action_ref"`

	InnerXml string `xml:",innerxml"`
}

type PatternType struct {
	VarRef string `xml:"var_ref,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type RangeType struct {
	Min *RangeValueType `xml:"min"`

	Max *RangeValueType `xml:"max"`

	InnerXml string `xml:",innerxml"`
}

type TestActionConditionType struct {
	ArtifactRefs *ArtifactRefsType `xml:"artifact_refs"`

	Result string `xml:"result"`

	TestActionRef *TestActionRefType `xml:"test_action_ref"`

	InnerXml string `xml:",innerxml"`
}

type RangeValueType struct {
	Inclusive string `xml:"inclusive,attr"`

	VarRef string `xml:"var_ref,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type QuestionsType struct {
	Question []Question `xml:"question"`

	ChoiceGroup []ChoiceGroupType `xml:"choice_group"`

	InnerXml string `xml:",innerxml"`
}

type QuestionTextType struct {
	Sub []SubstitutionTextType `xml:"sub"`

	InnerXml string `xml:",innerxml"`
}

type QuestionType struct {
	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type BooleanQuestionType struct {
	DefaultAnswer string `xml:"default_answer,attr,omitempty"`

	Model string `xml:"model,attr,omitempty"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type ChoiceQuestionType struct {
	DefaultAnswerRef string `xml:"default_answer_ref,attr,omitempty"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type NumericQuestionType struct {
	DefaultAnswer string `xml:"default_answer,attr,omitempty"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type StringQuestionType struct {
	DefaultAnswer string `xml:"default_answer,attr,omitempty"`

	Id string `xml:"id,attr"`

	Revision string `xml:"revision,attr"`

	QuestionText []QuestionTextType `xml:"question_text"`

	Instructions *InstructionsType `xml:"instructions"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type ChoiceType struct {
	Id string `xml:"id,attr"`

	VarRef string `xml:"var_ref,attr,omitempty"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type ChoiceGroupType struct {
	Id string `xml:"id,attr"`

	Choice []ChoiceType `xml:"choice"`

	InnerXml string `xml:",innerxml"`
}

type InstructionsType struct {
	Title TextType `xml:"title"`

	Step []StepType `xml:"step"`

	InnerXml string `xml:",innerxml"`
}

type ResultsType struct {
	StartTime string `xml:"start_time,attr,omitempty"`

	EndTime string `xml:"end_time,attr,omitempty"`

	Title *TextType `xml:"title"`

	QuestionnaireResults *QuestionnaireResultsType `xml:"questionnaire_results"`

	TestActionResults *TestActionResultsType `xml:"test_action_results"`

	QuestionResults *QuestionResultsType `xml:"question_results"`

	ArtifactResults *ArtifactResultsType `xml:"artifact_results"`

	Targets *TargetsType `xml:"targets"`

	InnerXml string `xml:",innerxml"`
}

type QuestionnaireResultsType struct {
	QuestionnaireResult []QuestionnaireResultType `xml:"questionnaire_result"`

	InnerXml string `xml:",innerxml"`
}

type TestActionResultsType struct {
	TestActionResult []TestActionResultType `xml:"test_action_result"`

	InnerXml string `xml:",innerxml"`
}

type QuestionResultsType struct {
	QuestionResult []QuestionResult `xml:"question_result"`

	InnerXml string `xml:",innerxml"`
}

type QuestionnaireResultType struct {
	QuestionnaireRef string `xml:"questionnaire_ref,attr"`

	Result string `xml:"result,attr"`

	ArtifactResults *ArtifactResultsType `xml:"artifact_results"`

	InnerXml string `xml:",innerxml"`
}

type TestActionResultType struct {
	TestActionRef string `xml:"test_action_ref,attr"`

	Result string `xml:"result,attr"`

	ArtifactResults *ArtifactResultsType `xml:"artifact_results"`

	InnerXml string `xml:",innerxml"`
}

type QuestionResultType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type BooleanQuestionResultType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr,omitempty"`

	Answer bool `xml:"answer"`

	InnerXml string `xml:",innerxml"`
}

type ChoiceQuestionResultType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr,omitempty"`

	Answer ChoiceAnswerType `xml:"answer"`

	InnerXml string `xml:",innerxml"`
}

type NumericQuestionResultType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr,omitempty"`

	Answer float64 `xml:"answer"`

	InnerXml string `xml:",innerxml"`
}

type StringQuestionResultType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Response string `xml:"response,attr,omitempty"`

	Answer string `xml:"answer"`

	InnerXml string `xml:",innerxml"`
}

type ChoiceAnswerType struct {
	ChoiceRef string `xml:"choice_ref,attr"`

	InnerXml string `xml:",innerxml"`
}

type ArtifactsType struct {
	Artifact []ArtifactType `xml:"artifact"`

	InnerXml string `xml:",innerxml"`
}

type ArtifactType struct {
	Id string `xml:"id,attr"`

	Persistent string `xml:"persistent,attr,omitempty"`

	Revision string `xml:"revision,attr"`

	Title TextType `xml:"title"`

	Description TextType `xml:"description"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type ArtifactRefsType struct {
	ArtifactRef []ArtifactRefType `xml:"artifact_ref"`

	InnerXml string `xml:",innerxml"`
}

type ArtifactRefType struct {
	Idref string `xml:"idref,attr"`

	Required string `xml:"required,attr,omitempty"`

	InnerXml string `xml:",innerxml"`
}

type ArtifactResultsType struct {
	ArtifactResult []ArtifactResultType `xml:"artifact_result"`

	InnerXml string `xml:",innerxml"`
}

type ArtifactValueType struct {
	InnerXml string `xml:",innerxml"`
}

type EmbeddedArtifactValueType struct {
	MimeType string `xml:"mime_type,attr"`

	InnerXml string `xml:",innerxml"`
}

type TextArtifactValueType struct {
	MimeType string `xml:"mime_type,attr"`

	Data string `xml:"data"`

	InnerXml string `xml:",innerxml"`
}

type BinaryArtifactValueType struct {
	MimeType string `xml:"mime_type,attr"`

	Data string `xml:"data"`

	InnerXml string `xml:",innerxml"`
}

type ReferenceArtifactValueType struct {
	Reference Reference `xml:"reference"`

	InnerXml string `xml:",innerxml"`
}

type ArtifactResultType struct {
	ArtifactRef string `xml:"artifact_ref,attr"`

	Timestamp string `xml:"timestamp,attr"`

	ArtifactValue ArtifactValue `xml:"artifact_value"`

	Provider string `xml:"provider"`

	Submitter UserType `xml:"submitter"`

	InnerXml string `xml:",innerxml"`
}

type TargetsType struct {
	Target []Target `xml:"target"`

	InnerXml string `xml:",innerxml"`
}

type UserType struct {
	Revision string `xml:"revision,attr"`

	Organization []string `xml:"organization"`

	Position []string `xml:"position"`

	Email []string `xml:"email"`

	Name string `xml:"name"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type SystemTargetType struct {
	Revision string `xml:"revision,attr"`

	Organization string `xml:"organization"`

	Ipaddress []string `xml:"ipaddress"`

	Description *TextType `xml:"description"`

	Name string `xml:"name"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type VariablesType struct {
	Variable []Variable `xml:"variable"`

	InnerXml string `xml:",innerxml"`
}

type VariableType struct {
	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type ConstantVariableType struct {
	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Value string `xml:"value"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type LocalVariableType struct {
	QuestionRef string `xml:"question_ref,attr"`

	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Set string `xml:"set"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type ExternalVariableType struct {
	Id string `xml:"id,attr"`

	Datatype string `xml:"datatype,attr"`

	Revision string `xml:"revision,attr"`

	Description *TextType `xml:"description"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type SetExpressionBaseType struct {
	Value string `xml:"value"`

	InnerXml string `xml:",innerxml"`
}

type SetExpressionPatternType struct {
	Pattern string `xml:"pattern,attr"`

	Value string `xml:"value"`

	InnerXml string `xml:",innerxml"`
}

type SetExpressionChoiceType struct {
	ChoiceRef string `xml:"choice_ref,attr"`

	Value string `xml:"value"`

	InnerXml string `xml:",innerxml"`
}

type SetExpressionRangeType struct {
	Min string `xml:"min,attr"`

	Max string `xml:"max,attr"`

	Value string `xml:"value"`

	InnerXml string `xml:",innerxml"`
}

type SetExpressionBooleanType struct {
	Value string `xml:"value,attr"`

	ValueElm string `xml:"value"`

	InnerXml string `xml:",innerxml"`
}

type VariableSetType struct {
	Expression []Expression `xml:"expression"`

	InnerXml string `xml:",innerxml"`
}

type SubstitutionTextType struct {
	VarRef string `xml:"var_ref,attr"`

	InnerXml string `xml:",innerxml"`
}

type ReferenceType struct {
	Href string `xml:"href,attr"`

	XmlLang string `xml:"lang,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type StepType struct {
	IsDone string `xml:"is_done,attr,omitempty"`

	IsRequired string `xml:"is_required,attr,omitempty"`

	Description *TextType `xml:"description"`

	Reference []ReferenceType `xml:"reference"`

	Step []StepType `xml:"step"`

	InnerXml string `xml:",innerxml"`
}

type ItemBaseType struct {
	Revision string `xml:"revision,attr"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type NamedItemBaseType struct {
	Revision string `xml:"revision,attr"`

	Name string `xml:"name"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type CompoundTestActionType struct {
	Revision string `xml:"revision,attr"`

	Title *TextType `xml:"title"`

	Description *TextType `xml:"description"`

	References *ReferencesType `xml:"references"`

	Actions OperationType `xml:"actions"`

	Notes []string `xml:"notes"`

	InnerXml string `xml:",innerxml"`
}

type ReferencesType struct {
	Reference []ReferenceType `xml:"reference"`

	InnerXml string `xml:",innerxml"`
}

type OperationType struct {
	Operation string `xml:"operation,attr,omitempty"`

	Negate string `xml:"negate,attr,omitempty"`

	TestActionRef []TestActionRefType `xml:"test_action_ref"`

	InnerXml string `xml:",innerxml"`
}

type TextType struct {
	XmlLang string `xml:"lang,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}
