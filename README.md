# SCAP ![Build CI](https://github.com/0intro/scap/workflows/Build%20CI/badge.svg) [![Gitter](https://badges.gitter.im/GoComply/community.svg)](https://gitter.im/GoComply/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge) [![PkgGoDev](https://pkg.go.dev/badge/github.com/gocomply/scap)](https://pkg.go.dev/github.com/gocomply/scap)
A GO module of the Security Content Automation Protocol (SCAP) Specification

This project represents complete parsers of SCAP standard. Parsers are generated automatically by [xsd2go](https://github.com/GoComply/xsd2go) project.

Users are advised to use [OpenSCAP](https://www.open-scap.org/) for SCAP evaluation and for any advanced use of SCAP standard. This project is beneficial for quick manipulation of SCAP data within golang code.

## Examplary usage
```
	import "github.com/0intro/scap/pkg/scap/scap_document"
	doc, err := scap_document.ReadDocumentFromFile("DOCUMENT.XML")
```

## Developer information
How to re-generate SCAP parsers with the latest [xsd2go](https://github.com/GoComply/xsd2go).
```
make all
```
