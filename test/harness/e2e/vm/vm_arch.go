package vm

import (
	_ "embed"
	"runtime"
)

// Separate libvirt domain template for aarch64 hosts. The default
// domain-template.xml is x86_64-only (q35 machine, Intel CPU model, isa-serial,
// tpm-tis), none of which are valid on the aarch64 'virt' machine.
//
//go:embed domain-template-aarch64.xml
var domainTemplateAarch64 string

// selectDomainTemplate returns the libvirt domain XML template matching the
// host architecture the binary is running on.
func selectDomainTemplate() string {
	if runtime.GOARCH == "arm64" {
		return domainTemplateAarch64
	}
	return domainTemplate
}
