package vincoin

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/text"
	"github.com/gioapp/gel"
	"github.com/marcetin/vin/db"
	"github.com/marcetin/vin/model"
	"github.com/marcetin/vin/pkg/tema"
)

var (
	izbornikRadova = &layout.List{
		Axis: layout.Vertical,
	}
	sumList = &layout.List{
		Axis: layout.Vertical,
	}
	neophodanMaterijalList = &layout.List{
		Axis: layout.Vertical,
	}
	ukupanNeophodanMaterijalList = &layout.List{
		Axis: layout.Vertical,
	}
	nazadDugme = new(gel.Button)
	dodajDugme = new(gel.Button)
	kolicina   = &gel.DuoUIcounter{
		Value:        1,
		OperateValue: 1,
		From:         1,
		To:           100,
		CounterInput: &gel.Editor{
			Alignment:  text.Middle,
			SingleLine: true,
		},
		CounterIncrease: new(gel.Button),
		CounterDecrease: new(gel.Button),
		CounterReset:    new(gel.Button),
	}
)

type VinCoin struct {
	Naziv                    string
	Window                   *app.Window
	Context                  *layout.Context
	Tema                     *tema.DuoUItheme
	Strana                   string
	Edit                     bool
	LinkoviIzboraVrsteRadova map[int]*gel.Button
	EditPolja                *model.EditabilnaPoljaVrsteRadova
	Materijal                map[int]*model.WingMaterijal
	Radovi                   model.WingVrstaRadova
	IzbornikRadova           *model.WingVrstaRadova
	Transfered               model.VinCoinGrupaRadova
	Db                       *db.DuoUIdb
	Client                   *model.Client
	PrikazaniElement         *model.WingVrstaRadova
	Suma                     *model.WingIzabraniElementi
}
