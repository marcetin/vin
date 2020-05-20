package vincoin

import (
	"github.com/gioapp/gel"
)

var (
	IzborVrsteRadovaPanelElement = gel.NewPanel()
)

func (в *ВинКоин) Назад() func() {
	return func() {
		btnNazad := в.Тема.Button("NAZAD")
		//btnNazad.Background = tema.HexARGB(в.Тема.Colors["Secondary"])
		for nazadDugme.Clicked(в.Context) {
			в.Страна = "Naslovna"
		}
		btnNazad.Layout(в.Context, nazadDugme)
	}
}
