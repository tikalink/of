// Code generated by entc, DO NOT EDIT.

package kernel

import (
	"github.com/tikafog/of/dbc/kernel/instruct"
	"github.com/tikafog/of/dbc/kernel/pin"
	"github.com/tikafog/of/dbc/kernel/schema"
	"github.com/tikafog/of/dbc/kernel/version"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	instructFields := schema.Instruct{}.Fields()
	_ = instructFields
	// instructDescCurrentUnix is the schema descriptor for current_unix field.
	instructDescCurrentUnix := instructFields[0].Descriptor()
	// instruct.DefaultCurrentUnix holds the default value on creation for the current_unix field.
	instruct.DefaultCurrentUnix = instructDescCurrentUnix.Default.(int64)
	// instructDescUpdatedUnix is the schema descriptor for updated_unix field.
	instructDescUpdatedUnix := instructFields[1].Descriptor()
	// instruct.DefaultUpdatedUnix holds the default value on creation for the updated_unix field.
	instruct.DefaultUpdatedUnix = instructDescUpdatedUnix.Default.(int64)
	pinFields := schema.Pin{}.Fields()
	_ = pinFields
	// pinDescStatus is the schema descriptor for status field.
	pinDescStatus := pinFields[1].Descriptor()
	// pin.DefaultStatus holds the default value on creation for the status field.
	pin.DefaultStatus = pinDescStatus.Default.(string)
	// pinDescStep is the schema descriptor for step field.
	pinDescStep := pinFields[2].Descriptor()
	// pin.DefaultStep holds the default value on creation for the step field.
	pin.DefaultStep = pinDescStep.Default.(string)
	// pinDescPriority is the schema descriptor for priority field.
	pinDescPriority := pinFields[3].Descriptor()
	// pin.DefaultPriority holds the default value on creation for the priority field.
	pin.DefaultPriority = pinDescPriority.Default.(int)
	// pinDescRelate is the schema descriptor for relate field.
	pinDescRelate := pinFields[4].Descriptor()
	// pin.DefaultRelate holds the default value on creation for the relate field.
	pin.DefaultRelate = pinDescRelate.Default.(string)
	// pinDescUpdatedUnix is the schema descriptor for updated_unix field.
	pinDescUpdatedUnix := pinFields[5].Descriptor()
	// pin.DefaultUpdatedUnix holds the default value on creation for the updated_unix field.
	pin.DefaultUpdatedUnix = pinDescUpdatedUnix.Default.(int64)
	versionFields := schema.Version{}.Fields()
	_ = versionFields
	// versionDescCurrent is the schema descriptor for Current field.
	versionDescCurrent := versionFields[0].Descriptor()
	// version.DefaultCurrent holds the default value on creation for the Current field.
	version.DefaultCurrent = versionDescCurrent.Default.(int)
	// version.CurrentValidator is a validator for the "Current" field. It is called by the builders before save.
	version.CurrentValidator = versionDescCurrent.Validators[0].(func(int) error)
	// versionDescLast is the schema descriptor for Last field.
	versionDescLast := versionFields[1].Descriptor()
	// version.DefaultLast holds the default value on creation for the Last field.
	version.DefaultLast = versionDescLast.Default.(int)
	// version.LastValidator is a validator for the "Last" field. It is called by the builders before save.
	version.LastValidator = versionDescLast.Validators[0].(func(int) error)
}
