package vmf

import "github.com/Nightgunner5/go-valve/keyvalues"

var _ keyvalues.Marshaler = new(VersionInfo)
var _ keyvalues.Marshaler = new(VisGroups)
var _ keyvalues.Marshaler = new(VisGroup)
//var _ keyvalues.Marshaler = new(ViewSettings)
var _ keyvalues.Marshaler = new(RGB)
