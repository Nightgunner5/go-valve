package vmf

import (
	"fmt"
	"github.com/Nightgunner5/go-valve/keyvalues"
	"strconv"
	"strings"
)

type RGB struct {
	Red, Green, Blue uint8
}

func (c *RGB) FromKeyValues(kv *keyvalues.KeyValues) error {
	parts := strings.Fields(kv.String(""))

	if len(parts) != 3 {
		return fmt.Errorf("RGB: color has %d components, but expected %d", len(parts), 3)
	}

	components := [...]*uint8{&c.Red, &c.Green, &c.Blue}
	componentNames := [...]string{"red", "green", "blue"}
	for i := range components {
		if v, err := strconv.ParseUint(parts[i], 10, 8); err == nil {
			*components[i] = uint8(v)
		} else {
			return fmt.Errorf("RGB: error parsing %s component: %v", componentNames[i], err)
		}
	}

	return nil
}

func (c *RGB) ToKeyValues() *keyvalues.KeyValues {
	var kv keyvalues.KeyValues

	color := kv.NewSubKey("color")
	color.SetValueString(fmt.Sprintf("%d %d %d", c.Red, c.Green, c.Blue))

	return color
}
