package hcl

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclparse"
	log "github.com/sirupsen/logrus"
)

func ParseHCLFile(filename string) error {
	p := hclparse.NewParser()
	file, diags := p.ParseHCLFile(filename)
	if diags.HasErrors() {
		return fmt.Errorf("%#v", diags.Error())
	}
	log.Debugf("%#v\n", file.Body)

	// log.Debugf("%s\n", file.Bytes)

	// attributes, diags := file.Body.JustAttributes()
	// if diags.HasErrors() {
	// 	return fmt.Errorf("%#v", diags.Error())
	// }
	// log.Debugf("%#v\n", attributes)
	return nil
}
