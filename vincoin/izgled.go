package vincoin

import (
	"github.com/gioapp/gel"
)

var (
	IzborVrsteRadovaPanelElement = gel.NewPanel()
)

func (v *VinCoin) Nazad() func() {
	return func() {
		btnNazad := v.Tema.Button("NAZAD")
		//btnNazad.Background = tema.HexARGB(v.Tema.Colors["Secondary"])
		for nazadDugme.Clicked(v.Context) {
			v.Strana = "Naslovna"
		}
		btnNazad.Layout(v.Context, nazadDugme)
	}
}
