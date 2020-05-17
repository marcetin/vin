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
	novauplataBtn        = new(gel.Button)
	aktuelnaglasanjaBtn  = new(gel.Button)
	rezultatiglasanjaBtn = new(gel.Button)
	proveraidentitetaBtn = new(gel.Button)
	podesavanjaBtn       = new(gel.Button)

	prijavaBtn      = new(gel.Button)
	registracijaBtn = new(gel.Button)
	prijazalsBtn    = new(gel.Button)
)

func C() {
	vin := vincoin.NewVinCoin()

	go func() {

		vin.Context = layout.NewContext(vin.Window.Queue())
		for e := range vin.Window.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				vin.Context.Reset(e.Config, e.Size)
				tema.DuoUIfill(vin.Context, vin.Tema.Colors["Light"])
				layout.Flex{
					Axis: layout.Vertical,
				}.Layout(vin.Context,
					layout.Rigid(func() {
						layout.Flex{
							Axis: layout.Horizontal,
						}.Layout(vin.Context,
							layout.Rigid(func() {
								vin.Tema.Icons["Grab"].Color = tema.HexARGB("ff444fcf")
								vin.Tema.Icons["Grab"].Layout(vin.Context, unit.Px(float32(32)))
								var logoMeniItem tema.DuoUIbutton
								logoMeniItem = vin.Tema.DuoUIbutton("", "", "", vin.Tema.Colors["Dark"], "", "", "logo", vin.Tema.Colors["Light"], 0, 32, 32, 32, 0, 0, 0, 0)
								for logoBtn.Clicked(vin.Context) {
									vin.Strana = "Novcanik"
								}
								logoMeniItem.IconLayout(vin.Context, logoBtn)
							}),
							layout.Rigid(func() {
								vin.Tema.DuoUIcontainer(0, vin.Tema.Colors["Plava"]).Layout(vin.Context, layout.Center, func() {
									vin.Tema.H2("KRIPTO DRZAVA SRBIJA").Layout(vin.Context)
								})
							}),
						)
					}),
					layout.Flexed(1, strana(vin)),
				)
				e.Frame(vin.Context.Ops)
			}
		}
	}()
	app.Main()
	log.Print("Starting server...")
}

func DugmeZaIzbor(v *vincoin.VinCoin, dugme *gel.Button, sirinaMenija int, tekst string) func() {
	return func() {
		v.Tema.DuoUIcontainer(12, v.Tema.Colors["White"]).Layout(v.Context, layout.Center, func() {
			v.Context.Constraints.Width.Min = sirinaMenija
			for dugme.Clicked(v.Context) {
				fmt.Println("IzbornikroditeL::", tekst)
				v.Strana = tekst
			}
			v.Tema.Button(tekst).Layout(v.Context, dugme)
		})
	}
}

func strana(v *vincoin.VinCoin) func() {
	return func() {
		var pin string
		sirinaMenija := 180
		switch v.Strana {
		case "Naslovna":
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(v.Context,
				layout.Rigid(DugmeZaIzbor(v, prijavaBtn, sirinaMenija, "Prijava za pinom")),
				layout.Rigid(DugmeZaIzbor(v, registracijaBtn, sirinaMenija, "Registracija za potpuno nove clanove")),
				layout.Rigid(DugmeZaIzbor(v, prijazalsBtn, sirinaMenija, "Prijava za clanove koji vec poseduju licnu sifru")),
				layout.Rigid(DugmeZaIzbor(v, podesavanjaBtn, sirinaMenija, "Podesavanja")),
			)
		case "Prijava za pinom":
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(v.Context,
				layout.Rigid(func() {
					v.Tema.DuoUIcontainer(12, v.Tema.Colors["White"]).Layout(v.Context, layout.Center, func() {
						v.Tema.H5(pin).Layout(v.Context)
					})
				}),
				layout.Rigid(DugmeZaIzbor(v, prijavaBtn, sirinaMenija, "Novcanik")),
			)
		case "Novcanik":
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(v.Context,
				layout.Rigid(DugmeZaIzbor(v, licnasifraBtn, sirinaMenija, "Licna sifra (24 reci)")),
				layout.Rigid(DugmeZaIzbor(v, racunBtn, sirinaMenija, "Racun (javni kljuc)")),
				layout.Rigid(DugmeZaIzbor(v, novostiBtn, sirinaMenija, "Novosti")),
				layout.Rigid(DugmeZaIzbor(v, stanjenaracunuBtn, sirinaMenija, "Stanje na racunu")),
				layout.Rigid(DugmeZaIzbor(v, novauplataBtn, sirinaMenija, "Nova uplata")),
				layout.Rigid(DugmeZaIzbor(v, aktuelnaglasanjaBtn, sirinaMenija, "Aktuelna glasanja")),
				layout.Rigid(DugmeZaIzbor(v, rezultatiglasanjaBtn, sirinaMenija, "Rezultati glasanja")),
				layout.Rigid(DugmeZaIzbor(v, proveraidentitetaBtn, sirinaMenija, "Provera identiteta")),
				layout.Rigid(DugmeZaIzbor(v, podesavanjaBtn, sirinaMenija, "Podesavanja")),
			)
		case "Licna sifra (24 reci)":
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(v.Context,
				layout.Rigid(func() {
					v.Tema.DuoUIcontainer(12, v.Tema.Colors["White"]).Layout(v.Context, layout.Center, func() {
						v.Tema.H5("BADEN BADEM MALI BADEM IMA BADEM NEMA BADEM ZEC TU SKOCI BADEM OCI KOME OCI BABEM BODEM").Layout(v.Context)
					})
				}),
			)
		case "Racun (javni kljuc)":
			pub := "0xB9467057a02Ab4d5319e1551E2C29bbc5Affe94f"
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(v.Context,
				layout.Rigid(func() {
					v.Tema.DuoUIcontainer(12, v.Tema.Colors["White"]).Layout(v.Context, layout.Center, func() { v.Tema.H4(pub).Layout(v.Context) })
				}),
				layout.Rigid(func() {
					v.Tema.DuoUIcontainer(12, v.Tema.Colors["White"]).Layout(v.Context, layout.Center, DuoUIqrCode(v.Context, pub, 256))
				}),
			)

		case "Novosti":
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(v.Context,
				layout.Rigid(func() {
					v.Tema.DuoUIcontainer(12, v.Tema.Colors["White"]).Layout(v.Context, layout.Center, func() {
						v.Tema.H5("najnovije INFORMACIJE").Layout(v.Context)
					})
				}),
				layout.Rigid(func() {
					v.Tema.DuoUIcontainer(12, v.Tema.Colors["White"]).Layout(v.Context, layout.Center, func() {
						v.Tema.Body1("U Kripto Drzavi Srbije narod donosi sve vazne odluke. Budite aktivni korisnik sistema lanca blokova i izjasnite se o postavljenim pitanjima. Samo jednom se broji vas glas bez obzira koju opciju izaberete. Sistem lanaca blokova ce Bas za glasanje nagraditi sa 0.1VIN (Vinkoinom). Ako probate glasati bise puta lanac blokova ce vam sveki put obracunati troskove transakcije i odbiti glasanje kao nelegitimno.").Layout(v.Context)
					})
				}),
			)
		case "Stanje na racunu":
		case "Nova uplata":
		case "Aktuelna glasanja":
		case "Rezultati glasanja":
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
