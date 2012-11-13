package vmf

import (
	"fmt"
	"github.com/Nightgunner5/go-valve/keyvalues"
	"strings"
)

type VersionInfo struct {
	EditorVersion int64
	EditorBuild   int64
	MapVersion    int64
	FormatVersion int64
	Prefab        bool
}

func (info *VersionInfo) ToKeyValues() *keyvalues.KeyValues {
	var kv keyvalues.KeyValues
	versioninfo := kv.NewSubKey("versioninfo")
	versioninfo.NewSubKey("editorversion").SetValueInt(info.EditorVersion)
	versioninfo.NewSubKey("editorbuild").SetValueInt(info.EditorBuild)
	versioninfo.NewSubKey("mapversion").SetValueInt(info.MapVersion)
	versioninfo.NewSubKey("formatversion").SetValueInt(info.FormatVersion)
	versioninfo.NewSubKey("prefab").SetValueBool(info.Prefab)

	return versioninfo
}

func (info *VersionInfo) FromKeyValues(kv *keyvalues.KeyValues) error {
	if !strings.EqualFold(kv.Name(), "versioninfo") {
		return fmt.Errorf("VersionInfo: key name was %q, not %q", kv.Name(), "versioninfo")
	}

	info.EditorVersion = kv.SubKey("editorversion").Int(0)
	info.EditorBuild = kv.SubKey("editorbuild").Int(0)
	info.MapVersion = kv.SubKey("mapversion").Int(0)
	info.FormatVersion = kv.SubKey("formatversion").Int(0)
	info.Prefab = kv.SubKey("prefab").Bool(false)

	return nil
}
