package novcanik

import "C"
import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"github.com/gioapp/gel"
	"github.com/gioapp/gelook"
	"github.com/marcetin/vin/calc"
	"github.com/marcetin/vin/pkg/tema"
	"log"
)

const (
	Address = ":9999"

	addr = "localhost:4242"
)

var (
	btn = new(gel.Button)

	licnasifraBtn        = new(gel.Button)
	racunBtn             = new(gel.Button)
	novostiBtn           = new(gel.Button)
	stanjenaracunuBtn    = new(gel.Button)
	novauplataBtn        = new(gel.Button)
	aktuelnaglasanjaBtn  = new(gel.Button)
	rezultatiglasanjaBtn = new(gel.Button)
	proveraidentitetaBtn = new(gel.Button)
	podesavanjaBtn       = new(gel.Button)
)

func C() {
	vin := calc.NewWingCal()

	vin.LinkoviIzboraVrsteRadova = map[int]*gel.Button{}

	vin.GenerisanjeLinkova(vin.Radovi.PodvrsteRadova)

	go func() {

		vin.Context = layout.NewContext(vin.Window.Queue())
		for e := range vin.Window.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				vin.Context.Reset(e.Config, e.Size)
				gelook.DuoUIfill(vin.Context, vin.Tema.Colors["Light"])
				layout.Flex{
					Axis: layout.Vertical,
				}.Layout(vin.Context,
					layout.Rigid(func() {
						vin.Tema.DuoUIcontainer(0, vin.Tema.Colors["Plava"]).Layout(vin.Context, layout.Center, func() {
							vin.Tema.H2("KRIPTO DRZAVA SRBIJA").Layout(vin.Context)
						})
					}),
					layout.Flexed(1, func() {
						layout.Flex{
							Axis: layout.Vertical,
						}.Layout(vin.Context,
							layout.Rigid(func() {
								vin.Tema.DuoUIcontainer(0, vin.Tema.Colors["White"]).Layout(vin.Context, layout.Center, func() {
									vin.Tema.Button("Licna sifra (24 reci)").Layout(vin.Context, licnasifraBtn)
								})

							}),
						)
					}),
				)
				e.Frame(vin.Context.Ops)
			}
		}
	}()
	app.Main()
	log.Print("Starting server...")
}
