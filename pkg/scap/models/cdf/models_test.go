package cdf_test

import (
	"testing"

	"github.com/gocomply/scap/pkg/scap/scap_document"
	"github.com/stretchr/testify/assert"
)

func TestRFCFeedParsing(t *testing.T) {
	doc, err := scap_document.ReadDocumentFromFile("../../../../examples/scap/xccdf/1.2/test_xccdf_complex_check_nand.xccdf.xml")
	if err != nil {
		t.Fatal(err)
	}
	bench := doc.Benchmark
	assert.NotNil(t, bench)
	assert.Equal(t, bench.XMLName.Space, "http://checklists.nist.gov/xccdf/1.2")
	assert.Equal(t, bench.Id, "testing-xcccdf")
	assert.Equal(t, len(bench.Status), 1)
	assert.Empty(t, bench.Title)
	assert.Empty(t, bench.Description)
}
