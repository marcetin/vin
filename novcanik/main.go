package novcanik

import "C"
import (
	"fmt"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/gioapp/gel"
	"github.com/marcetin/vin/pkg/qrcode"
	"github.com/marcetin/vin/pkg/tema"
	"github.com/marcetin/vin/vincoin"
	"github.com/nfnt/resize"
	"image"
	"log"
)

const (
	Address = ":9999"

	addr = "localhost:4242"
)

var (
	logoBtn = new(gel.Button)

	btn = new(gel.Button)

	licnasifraBtn        = new(gel.Button)
	racunBtn             = new(gel.Button)
	novostiBtn           = new(gel.Button)
	stanjenaracunuBtn    = new(gel.Button)
	novaуплатаBtn        = new(gel.Button)
	aktuelnaгласањаBtn   = new(gel.Button)
	rezultatiгласањаBtn  = new(gel.Button)
	proveraidentitetaBtn = new(gel.Button)
	podesavanjaBtn       = new(gel.Button)

	prijavaBtn      = new(gel.Button)
	registracijaBtn = new(gel.Button)
	prijazalsBtn    = new(gel.Button)
)

func ВИН() {
	vin := vincoin.НовиВинКоин()

	go func() {

		vin.Context = layout.NewContext(vin.Window.Queue())
		for e := range vin.Window.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				vin.Context.Reset(e.Config, e.Size)
				tema.DuoUIfill(vin.Context, vin.Тема.Colors["Light"])
				layout.Flex{
					Axis: layout.Vertical,
				}.Layout(vin.Context,
					layout.Rigid(func() {
						layout.Flex{
							Axis: layout.Horizontal,
						}.Layout(vin.Context,
							layout.Rigid(func() {
								vin.Тема.Icons["Grab"].Color = tema.HexARGB("ff444fcf")
								vin.Тема.Icons["Grab"].Layout(vin.Context, unit.Px(float32(32)))
								var logoMeniItem tema.DuoUIbutton
								logoMeniItem = vin.Тема.DuoUIbutton("", "", "", vin.Тема.Colors["Dark"], "", "", "logo", vin.Тема.Colors["Light"], 0, 32, 32, 32, 0, 0, 0, 0)
								for logoBtn.Clicked(vin.Context) {
									vin.Страна = "Novcanik"
								}
								logoMeniItem.IconLayout(vin.Context, logoBtn)
							}),
							layout.Rigid(func() {
								vin.Тема.DuoUIcontainer(0, vin.Тема.Colors["Plava"]).Layout(vin.Context, layout.Center, func() {
									vin.Тема.H2("Крипто Држава Србија").Layout(vin.Context)
								})
							}),
						)
					}),
					layout.Flexed(1, страна(vin)),
				)
				e.Frame(vin.Context.Ops)
			}
		}
	}()
	app.Main()
	log.Print("Starting server...")
}

func ДугмеЗаИзбор(v *vincoin.ВинКоин, dugme *gel.Button, ширинаМенија int, tekst string) func() {
	return func() {
		v.Тема.DuoUIcontainer(12, v.Тема.Colors["White"]).Layout(v.Context, layout.Center, func() {
			v.Context.Constraints.Width.Min = ширинаМенија
			for dugme.Clicked(v.Context) {
				fmt.Println("IzbornikroditeL::", tekst)
				v.Страна = tekst
			}
			v.Тема.Button(tekst).Layout(v.Context, dugme)
		})
	}
}

func страна(v *vincoin.ВинКоин) func() {
	return func() {
		var pin string
		ширинаМенија := 180
		switch v.Страна {
		case "Naslovna":
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(v.Context,
				layout.Rigid(ДугмеЗаИзбор(v, prijavaBtn, ширинаМенија, "Пријава са пином")),
				layout.Rigid(ДугмеЗаИзбор(v, registracijaBtn, ширинаМенија, "Регистрација за потпуно нове чланове")),
				layout.Rigid(ДугмеЗаИзбор(v, prijazalsBtn, ширинаМенија, "Пријава за чланове који већ поседују личну шифру")),
				layout.Rigid(ДугмеЗаИзбор(v, podesavanjaBtn, ширинаМенија, "Подешавања")),
			)
		case "Пријава са пином":
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(v.Context,
				layout.Rigid(func() {
					v.Тема.DuoUIcontainer(12, v.Тема.Colors["White"]).Layout(v.Context, layout.Center, func() {
						v.Тема.H5(pin).Layout(v.Context)
					})
				}),
				layout.Rigid(ДугмеЗаИзбор(v, prijavaBtn, ширинаМенија, "Novcanik")),
			)
		case "Novcanik":
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(v.Context,
				layout.Rigid(ДугмеЗаИзбор(v, licnasifraBtn, ширинаМенија, "Licna sifra (24 reci)")),
				layout.Rigid(ДугмеЗаИзбор(v, racunBtn, ширинаМенија, "Racun (javni kljuc)")),
				layout.Rigid(ДугмеЗаИзбор(v, novostiBtn, ширинаМенија, "Novosti")),
				layout.Rigid(ДугмеЗаИзбор(v, stanjenaracunuBtn, ширинаМенија, "Stanje na racunu")),
				layout.Rigid(ДугмеЗаИзбор(v, novaуплатаBtn, ширинаМенија, "Nova уплата")),
				layout.Rigid(ДугмеЗаИзбор(v, aktuelnaгласањаBtn, ширинаМенија, "Aktuelna гласања")),
				layout.Rigid(ДугмеЗаИзбор(v, rezultatiгласањаBtn, ширинаМенија, "Rezultati гласања")),
				layout.Rigid(ДугмеЗаИзбор(v, proveraidentitetaBtn, ширинаМенија, "Provera identiteta")),
				layout.Rigid(ДугмеЗаИзбор(v, podesavanjaBtn, ширинаМенија, "Podesavanja")),
			)
		case "Licna sifra (24 reci)":
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(v.Context,
				layout.Rigid(func() {
					v.Тема.DuoUIcontainer(12, v.Тема.Colors["White"]).Layout(v.Context, layout.Center, func() {
						v.Тема.H5("BADEN BADEM MALI BADEM IMA BADEM NEMA BADEM ZEC TU SKOCI BADEM OCI KOME OCI BABEM BODEM").Layout(v.Context)
					})
				}),
			)
		case "Racun (javni kljuc)":
			pub := "0xB9467057a02Ab4d5319e1551E2C29bbc5Affe94f"
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(v.Context,
				layout.Rigid(func() {
					v.Тема.DuoUIcontainer(12, v.Тема.Colors["White"]).Layout(v.Context, layout.Center, func() { v.Тема.H4(pub).Layout(v.Context) })
				}),
				layout.Rigid(func() {
					v.Тема.DuoUIcontainer(12, v.Тема.Colors["White"]).Layout(v.Context, layout.Center, DuoUIqrCode(v.Context, pub, 256))
				}),
			)

		case "Novosti":
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(v.Context,
				layout.Rigid(func() {
					v.Тема.DuoUIcontainer(12, v.Тема.Colors["White"]).Layout(v.Context, layout.Center, func() {
						v.Тема.H5("najnovije INFORMACIJE").Layout(v.Context)
					})
				}),
				layout.Rigid(func() {
					v.Тема.DuoUIcontainer(12, v.Тема.Colors["White"]).Layout(v.Context, layout.Center, func() {
						v.Тема.Body1("U Kripto Drzavi Srbije narod donosi sve vazne odluke. Budite aktivni korisnik sistema lanca blokova i izjasnite se o postavljenim pitanjima. Samo jednom se broji vas glas bez obzira koju opciju izaberete. Sistem lanaca blokova ce Bas za glasanje nagraditi sa 0.1VIN (Vinkoinom). Ako probate glasati bise puta lanac blokova ce vam sveki put obracunati troskove transakcije i odbiti glasanje kao nelegitimno.").Layout(v.Context)
					})
				}),
			)
		case "Stanje na racunu":
		case "Nova уплата":
		case "Aktuelna гласања":
		case "Rezultati гласања":
		case "Provera identiteta":
		case "Podesavanja":

		}
	}
}

func DuoUIqrCode(gtx *layout.Context, hash string, size uint) func() {
	return func() {
		qr, err := qrcode.Encode(hash, 3, qrcode.ECLevelM)
		if err != nil {
		}
		qrResize := resize.Resize(size, 0, qr, resize.NearestNeighbor)
		addrQR := paint.NewImageOp(qrResize)
		sz := gtx.Constraints.Width.Constrain(gtx.Px(unit.Dp(float32(size))))
		addrQR.Add(gtx.Ops)
		paint.PaintOp{
			Rect: f32.Rectangle{
				Max: f32.Point{
					X: float32(sz), Y: float32(sz),
				},
			},
		}.Add(gtx.Ops)
		gtx.Dimensions.Size = image.Point{X: sz, Y: sz}
	}
}
