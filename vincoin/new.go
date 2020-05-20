package vincoin

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/unit"
	"github.com/marcetin/vin/db"
	"github.com/marcetin/vin/model"
	"github.com/marcetin/vin/pkg/tema"
)

func НовиВинКоин() *ВинКоин {
	gofont.Register()
	vin := &ВинКоин{
		Назив: "Vincoin Novcanik",
		Window: app.NewWindow(
			app.Size(unit.Dp(999), unit.Dp(999)),
			app.Title("Vincoin Novcanik"),
		),
		Тема:             tema.NewDuoUItheme(),
		Страна:           "Naslovna",
		Db:               db.DuoUIdbInit("./BAZA"),
		PrikazaniElement: &model.WingVrstaRadova{},
		Suma: &model.WingIzabraniElementi{
			UkupanNeophodanMaterijal: map[int]model.WingNeophodanMaterijal{},
		},
	}

	return vin
}
