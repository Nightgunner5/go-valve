package vmf

import (
	"fmt"
	"github.com/Nightgunner5/go-valve/keyvalues"
	"strings"
)

type VisGroups []VisGroup

func (vg *VisGroups) FromKeyValues(kv *keyvalues.KeyValues) error {
	if !strings.EqualFold(kv.Name(), "visgroups") {
		return fmt.Errorf("VisGroups: key name was %q, not %q", kv.Name(), "visgroups")
	}

	*vg = (*vg)[:0]

	for subkey := range kv.Each() {
		var group VisGroup
		if err := group.FromKeyValues(subkey); err != nil {
			return err
		}
		*vg = append(*vg, group)
	}

	return nil
}

func (vg *VisGroups) ToKeyValues() *keyvalues.KeyValues {
	var kv keyvalues.KeyValues

	visgroups := kv.NewSubKey("visgroups")

	for _, group := range *vg {
		visgroups.Append(group.ToKeyValues())
	}

	return visgroups
}

type VisGroup struct {
	Name  string
	ID    int64
	Color RGB
}

func (vg *VisGroup) FromKeyValues(kv *keyvalues.KeyValues) error {
	if !strings.EqualFold(kv.Name(), "visgroup") {
		return fmt.Errorf("VisGroup: key name was %q, not %q", kv.Name(), "visgroup")
	}

	if name := kv.SubKey("name").String(""); name != "" {
		vg.Name = name
	} else {
		return fmt.Errorf("VisGroup: no name")
	}

	if id := kv.SubKey("visgroupid").Int(0); id > 0 {
		vg.ID = id
	} else {
		return fmt.Errorf("VisGroup: %s: ID = %d or no ID", vg.Name, id)
	}

	if err := vg.Color.FromKeyValues(kv.SubKey("color")); err != nil {
		return fmt.Errorf("VisGroup: %s: %v", vg.Name, err)
	}

	return nil
}

func (vg *VisGroup) ToKeyValues() *keyvalues.KeyValues {
	var kv keyvalues.KeyValues

	visgroup := kv.NewSubKey("visgroup")
	visgroup.NewSubKey("name").SetValueString(vg.Name)
	visgroup.NewSubKey("visgroupid").SetValueInt(vg.ID)
	visgroup.Append(vg.Color.ToKeyValues())

	return visgroup
}
