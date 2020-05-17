package novcanik

import (
	"gioui.org/layout"
	"github.com/marcetin/vin/vincoin"
)

func glavniDeo(w *vincoin.VinCoin) func() {
	return func() {
		if w.Edit {
			layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(w.Context)
		} else {
			layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(w.Context)
		}
	}
}
