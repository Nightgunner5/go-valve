package vmf

import (
	"fmt"
	"github.com/Nightgunner5/go-valve/keyvalues"
	"strings"
)

type ViewSettings struct {
	SnapToGrid      bool
	ShowGrid        bool
	ShowLogicalGrid bool
	GridSpacing     int64
	Show3DGrid      bool
}

func (vs *ViewSettings) FromKeyValues(kv *keyvalues.KeyValues) error {
	if !strings.EqualFold(kv.Name(), "viewsettings") {
		return fmt.Errorf("ViewSettings: key name was %q, not %q", kv.Name(), "viewsettings")
	}

	vs.SnapToGrid = kv.SubKey("bSnapToGrid").Bool(true)
	vs.ShowGrid = kv.SubKey("bShowGrid").Bool(true)
	vs.ShowLogicalGrid = kv.SubKey("bShowLogicalGrid").Bool(false)
	vs.GridSpacing = kv.SubKey("nGridSpacing").Int(64)
	vs.Show3DGrid = kv.SubKey("bShow3DGrid").Bool(false)

	return nil
}

func (vs *ViewSettings) ToKeyValues() *keyvalues.KeyValues {
	var kv keyvalues.KeyValues

	viewsettings := kv.NewSubKey("viewsettings")
	viewsettings.NewSubKey("bSnapToGrid").SetValueBool(vs.SnapToGrid)
	viewsettings.NewSubKey("bShowGrid").SetValueBool(vs.ShowGrid)
	viewsettings.NewSubKey("bShowLogicalGrid").SetValueBool(vs.ShowLogicalGrid)
	viewsettings.NewSubKey("nGridSpacing").SetValueInt(vs.GridSpacing)
	viewsettings.NewSubKey("bShow3DGrid").SetValueBool(vs.Show3DGrid)

	return viewsettings
}
