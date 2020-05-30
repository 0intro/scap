package ds_test

import (
	"testing"

	"github.com/gocomply/scap/pkg/scap/scap_document"
	"github.com/stretchr/testify/assert"
)

func TestSanityDsParsing(t *testing.T) {
	doc, err := scap_document.ReadDocumentFromFile("../../../../examples/scap/ds/1.3/sds.ds.xml")
	if err != nil {
		t.Fatal(err)
	}
	ds := doc.DataStreamCollection
	assert.NotNil(t, ds)
	assert.Equal(t, ds.XMLName.Space, "http://scap.nist.gov/schema/scap/source/1.2")
	assert.Equal(t, ds.Id, "scap_org.open-scap_collection_from_xccdf_test_deriving_xccdf_result_from_oval_pass.oval.xml")
	assert.Equal(t, ds.SchematronVersion, "1.3")
	assert.Equal(t, len(ds.DataStream), 1)
	assert.Empty(t, ds.ExtendedComponent)
	assert.Empty(t, ds.Signature)

	dataStream := ds.DataStream[0]
	assert.Equal(t, dataStream.Id, "scap_org.open-scap_datastream_from_xccdf_test_deriving_xccdf_result_from_oval_pass.oval.xml")
	assert.Equal(t, dataStream.ScapVersion, "1.3")
	assert.Equal(t, dataStream.UseCase, "OTHER")

	assert.Nil(t, dataStream.Dictionaries)
	assert.Nil(t, dataStream.ExtendedComponents)
	checklists := dataStream.Checklists.ComponentRef
	assert.Equal(t, len(checklists), 2)
	assert.Equal(t, checklists[0].Id, "scap_org.open-scap_cref_xccdf.xml")
	assert.Equal(t, checklists[1].Id, "scap_org.open-scap_cref_xccdf-file.xml")

	checks := dataStream.Checks.ComponentRef
	assert.Equal(t, len(checks), 1)
	assert.Equal(t, checks[0].Id, "scap_org.open-scap_cref_oval.xml")

	assert.Equal(t, len(ds.Component), 1)
	component := ds.Component[0]
	assert.Equal(t, component.Id, "scap_1_comp_pass.oval.xml")
	assert.Equal(t, component.Timestamp, "2016-02-02T14:06:14")
}
