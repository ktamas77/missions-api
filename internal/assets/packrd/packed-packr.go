// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "8cbb14e86e369f5d22b4176e28308530"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"64775b443fc934611808a6e1794828cb": "1f8b08000000000000ff8cd2514fc2301000e0f7fe8a7b64511234e20b4f03165dc481cd66c25353db0a8dd036ed91857f6fe6181313abf776c9d7bb5cef8643b8daeb8de7a8a07284cc6896961994e97491c15e87a0ad210302005dc6b46c3298e60f79514255e42f5576fd25a40ac26b87da9a267d4de9ec31a583dbf1386981e17b057d74e066344aa0589650548b452bdf2cf7529b0d43dd3c39356be352eef8c188edc945a532f2cce2521e3cefa688cbee53f0e8545c7a55732fcfb34764705ca8b0d58e091bf03fdd9dad958fd6448b7cc79aa2e18feee61d99b0063d17d82ce2888ac3af12ed8732ece075bfcdfbbbe4a75cd1fc39a56b78cad630e8ef2821c98490ef2738b7b521644e97abcb139c90cf000000ffffc7f374f8a9020000",
		"a48aaf22f5ae0ea54776b2d1b1d1b972": "1f8b08000000000000ff84d04f4bc3401005f0fb7c8a776cd17e829e52bb4830a6356cc09e968d1974b53b1b76d63f1f5fb41e5a84768ec38f07ef2d16b88ae139fbc2e827a29bce54d6c056abc6e0350571fcc15268460050b217f54f25247161fcf96055dfd6ad45dfd60fbdb9fe55fc35ed53e6ecfc3866560560cda305d06e2cdabe690e2e06d5900447f79776eae43d0e9c9dbe8449cf392dfe8d5dd9c70b790737c870c16dbbfabeea76b8333bcc4eabcf09a0f992e878bf75fa14a275b7d9fedb6f49df010000ffffc887d1b969010000",
		"e5bb19fd7b62848f30fef472880fbbd6": "1f8b08000000000000ff74cfcf6a843010c7f1fb3cc5efa8b43e81a75843094da34d13a82789188ad47fc440f7f1177641dc5d766ec37c0ef3cd32bc4cc36f70d1c3ae446f9a33c361582139fc691d97e0032504605fdba1c765beb9164cbede1e5ddf07bf6d30fcc700509581b252c22af165f915c725bab1dda2fbf36d1c27a010ef42991d3faa6eee9ea85a8b4fa61b7cf006c9e1c594d29ce8d8572eff3351a9abfaae2f3f070000ffffc6cbbd0706010000",
		"f33a8d342c06f4e323d5cd7c8280c6a0": "1f8b08000000000000ff8c924b6e83301086f7738a5906b53941563c26112a35917116595920dcc60a60845dd1e3575541316d1ef5d2f3e9f7e79959aff1a9d5ef43e9141e7a809853280845186584eab36fcca006d96a6bb5e9600588f8e75aea1a0be269983d2feae89d28dda54c20b25c203b64133905e06372d4ee540fe5e8b391318d2abb5f64f7d1566a90f6a47b7b3fd3195736d2baf2aca46bda7f925557dd21f73c7d0df9115fe888ab2b8d0a7eb0386785e0e177c0db59ce1cccbedb9c53ba63cb9400fc3671da1227165371e9f7e5415d2fe99c61421909c2827cdda5c73ce6ab1a53f1a6c53ccb95f7d9070e106c00fc0d4cccd801243cdfdfd8c0cd57000000ffff6619491fb0020000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("migrations", "./migrations")
		b.SetResolver("001_mission.sql", packr.Pointer{ForwardBox: gk, ForwardPath: "64775b443fc934611808a6e1794828cb"})
		b.SetResolver("002_explorer.sql", packr.Pointer{ForwardBox: gk, ForwardPath: "e5bb19fd7b62848f30fef472880fbbd6"})
		b.SetResolver("003_explorer_mission.sql", packr.Pointer{ForwardBox: gk, ForwardPath: "f33a8d342c06f4e323d5cd7c8280c6a0"})
		b.SetResolver("004_joinevent.sql", packr.Pointer{ForwardBox: gk, ForwardPath: "a48aaf22f5ae0ea54776b2d1b1d1b972"})
	}()
	return nil
}()
