package arm

import (
	"context"
	"io"
	"io/fs"

	"github.com/aquasecurity/defsec/pkg/scan"

	"github.com/aquasecurity/defsec/pkg/framework"
	"github.com/aquasecurity/defsec/pkg/scanners"
	"github.com/aquasecurity/defsec/pkg/scanners/options"
)

var _ scanners.FSScanner = (*Scanner)(nil)
var _ options.ConfigurableScanner = (*Scanner)(nil)

type Scanner struct {
	scannerOptions []options.ScannerOption
	parserOptions  []options.ParserOption
}

func (s Scanner) Name() string {
	//TODO implement me
	panic("implement me")
}

func (s Scanner) ScanFS(ctx context.Context, fs fs.FS, dir string) (scan.Results, error) {
	//TODO implement me
	panic("implement me")
}

func (s Scanner) SetDebugWriter(writer io.Writer) {
	//TODO implement me
	panic("implement me")
}

func (s Scanner) SetTraceWriter(writer io.Writer) {
	//TODO implement me
	panic("implement me")
}

func (s Scanner) SetPerResultTracingEnabled(b bool) {
	//TODO implement me
	panic("implement me")
}

func (s Scanner) SetPolicyDirs(s2 ...string) {
	//TODO implement me
	panic("implement me")
}

func (s Scanner) SetDataDirs(s2 ...string) {
	//TODO implement me
	panic("implement me")
}

func (s Scanner) SetPolicyNamespaces(s2 ...string) {
	//TODO implement me
	panic("implement me")
}

func (s Scanner) SetSkipRequiredCheck(b bool) {
	//TODO implement me
	panic("implement me")
}

func (s Scanner) SetPolicyReaders(readers []io.Reader) {
	//TODO implement me
	panic("implement me")
}

func (s Scanner) SetPolicyFilesystem(fs fs.FS) {
	//TODO implement me
	panic("implement me")
}

func (s Scanner) SetUseEmbeddedPolicies(b bool) {
	//TODO implement me
	panic("implement me")
}

func (s Scanner) SetFrameworks(frameworks []framework.Framework) {
	//TODO implement me
	panic("implement me")
}