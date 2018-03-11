package x11colors

import (
	"image/color"
	"math/rand"
	"time"
)

// X11Color defines single color
type X11Color struct {
	// Name is the color name as defined in <X11root>/lib/X11/rgb.txt
	Name string
	// RGBA is the color RGB with A component always set to 0xFF (255)
	RGBA color.RGBA
	// CaptionBlack defines the text color that can be used on the current color background:
	// true for black and false for white - as it is in
	// https://ru.wikipedia.org/wiki/%D0%A1%D0%BF%D0%B8%D1%81%D0%BE%D0%BA_%D1%86%D0%B2%D0%B5%D1%82%D0%BE%D0%B2_%D0%B2_X11
	CaptionBlack bool
}

var (
	AliceBlue         = X11Color{"Alice Blue", color.RGBA{0xF0, 0xF8, 0xFF, 0xFF}, true}
	AntiqueWhite      = X11Color{"Antique White", color.RGBA{0xFA, 0xEB, 0xD7, 0xFF}, true}
	Aqua              = X11Color{"Aqua", color.RGBA{0x00, 0xFF, 0xFF, 0xFF}, false}
	Aquamarine        = X11Color{"Aquamarine", color.RGBA{0x7F, 0xFF, 0xD4, 0xFF}, false}
	Azure             = X11Color{"Azure", color.RGBA{0xF0, 0xFF, 0xFF, 0xFF}, true}
	Beige             = X11Color{"Beige", color.RGBA{0xF5, 0xF5, 0xDC, 0xFF}, true}
	Bisque            = X11Color{"Bisque", color.RGBA{0xFF, 0xE4, 0xC4, 0xFF}, false}
	Black             = X11Color{"Black", color.RGBA{0x00, 0x00, 0x00, 0xFF}, false}
	BlanchedAlmond    = X11Color{"Blanched Almond", color.RGBA{0xFF, 0xEB, 0xCD, 0xFF}, true}
	Blue              = X11Color{"Blue", color.RGBA{0x00, 0x00, 0xFF, 0xFF}, false}
	BlueViolet        = X11Color{"Blue Violet", color.RGBA{0x8A, 0x2B, 0xE2, 0xFF}, false}
	Brown             = X11Color{"Brown", color.RGBA{0xA5, 0x2A, 0x2A, 0xFF}, false}
	Burlywood         = X11Color{"Burlywood", color.RGBA{0xDE, 0xB8, 0x87, 0xFF}, false}
	CadetBlue         = X11Color{"Cadet Blue", color.RGBA{0x5F, 0x9E, 0xA0, 0xFF}, false}
	Chartreuse        = X11Color{"Chartreuse", color.RGBA{0x7F, 0xFF, 0x00, 0xFF}, false}
	Chocolate         = X11Color{"Chocolate", color.RGBA{0xD2, 0x69, 0x1E, 0xFF}, false}
	Coral             = X11Color{"Coral", color.RGBA{0xFF, 0x7F, 0x50, 0xFF}, false}
	Cornflower        = X11Color{"Cornflower", color.RGBA{0x64, 0x95, 0xED, 0xFF}, false}
	Cornsilk          = X11Color{"Cornsilk", color.RGBA{0xFF, 0xF8, 0xDC, 0xFF}, true}
	Crimson           = X11Color{"Crimson", color.RGBA{0xDC, 0x14, 0x3C, 0xFF}, false}
	Cyan              = X11Color{"Cyan", color.RGBA{0x00, 0xFF, 0xFF, 0xFF}, false}
	DarkBlue          = X11Color{"Dark Blue", color.RGBA{0x00, 0x00, 0x8B, 0xFF}, false}
	DarkCyan          = X11Color{"Dark Cyan", color.RGBA{0x00, 0x8B, 0x8B, 0xFF}, false}
	DarkGoldenrod     = X11Color{"Dark Goldenrod", color.RGBA{0xB8, 0x86, 0x0B, 0xFF}, false}
	DarkGray          = X11Color{"Dark Gray", color.RGBA{0xA9, 0xA9, 0xA9, 0xFF}, false}
	DarkGreen         = X11Color{"Dark Green", color.RGBA{0x00, 0x64, 0x00, 0xFF}, false}
	DarkKhaki         = X11Color{"Dark Khaki", color.RGBA{0xBD, 0xB7, 0x6B, 0xFF}, false}
	DarkMagenta       = X11Color{"Dark Magenta", color.RGBA{0x8B, 0x00, 0x8B, 0xFF}, false}
	DarkOliveGreen    = X11Color{"Dark Olive Green", color.RGBA{0x55, 0x6B, 0x2F, 0xFF}, false}
	DarkOrange        = X11Color{"Dark Orange", color.RGBA{0xFF, 0x8C, 0x00, 0xFF}, false}
	DarkOrchid        = X11Color{"Dark Orchid", color.RGBA{0x99, 0x32, 0xCC, 0xFF}, false}
	DarkRed           = X11Color{"Dark Red", color.RGBA{0x8B, 0x00, 0x00, 0xFF}, false}
	DarkSalmon        = X11Color{"Dark Salmon", color.RGBA{0xE9, 0x96, 0x7A, 0xFF}, false}
	DarkSeaGreen      = X11Color{"Dark Sea Green", color.RGBA{0x8F, 0xBC, 0x8F, 0xFF}, false}
	DarkSlateBlue     = X11Color{"Dark Slate Blue", color.RGBA{0x48, 0x3D, 0x8B, 0xFF}, false}
	DarkSlateGray     = X11Color{"Dark Slate Gray", color.RGBA{0x2F, 0x4F, 0x4F, 0xFF}, false}
	DarkTurquoise     = X11Color{"Dark Turquoise", color.RGBA{0x00, 0xCE, 0xD1, 0xFF}, false}
	DarkViolet        = X11Color{"Dark Violet", color.RGBA{0x94, 0x00, 0xD3, 0xFF}, false}
	DeepPink          = X11Color{"Deep Pink", color.RGBA{0xFF, 0x14, 0x93, 0xFF}, false}
	DeepSkyBlue       = X11Color{"Deep Sky Blue", color.RGBA{0x00, 0xBF, 0xFF, 0xFF}, false}
	DimGray           = X11Color{"Dim Gray", color.RGBA{0x69, 0x69, 0x69, 0xFF}, false}
	DodgerBlue        = X11Color{"Dodger Blue", color.RGBA{0x1E, 0x90, 0xFF, 0xFF}, false}
	Firebrick         = X11Color{"Firebrick", color.RGBA{0xB2, 0x22, 0x22, 0xFF}, false}
	FloralWhite       = X11Color{"Floral White", color.RGBA{0xFF, 0xFA, 0xF0, 0xFF}, true}
	ForestGreen       = X11Color{"Forest Green", color.RGBA{0x22, 0x8B, 0x22, 0xFF}, false}
	Fuchsia           = X11Color{"Fuchsia", color.RGBA{0xFF, 0x00, 0xFF, 0xFF}, false}
	Gainsboro         = X11Color{"Gainsboro", color.RGBA{0xDC, 0xDC, 0xDC, 0xFF}, true}
	GhostWhite        = X11Color{"Ghost White", color.RGBA{0xF8, 0xF8, 0xFF, 0xFF}, true}
	Gold              = X11Color{"Gold", color.RGBA{0xFF, 0xD7, 0x00, 0xFF}, false}
	Goldenrod         = X11Color{"Goldenrod", color.RGBA{0xDA, 0xA5, 0x20, 0xFF}, false}
	GrayX11           = X11Color{"Gray (X11)", color.RGBA{0xBE, 0xBE, 0xBE, 0xFF}, false}
	GrayW3C           = X11Color{"Gray (W3C)", color.RGBA{0x7F, 0x7F, 0x7F, 0xFF}, false}
	GreenX11          = X11Color{"Green (X11)", color.RGBA{0x00, 0xFF, 0x00, 0xFF}, false}
	GreenW3C          = X11Color{"Green (W3C)", color.RGBA{0x00, 0x7F, 0x00, 0xFF}, false}
	GreenYellow       = X11Color{"Green Yellow", color.RGBA{0xAD, 0xFF, 0x2F, 0xFF}, false}
	Honeydew          = X11Color{"Honeydew", color.RGBA{0xF0, 0xFF, 0xF0, 0xFF}, true}
	HotPink           = X11Color{"Hot Pink", color.RGBA{0xFF, 0x69, 0xB4, 0xFF}, false}
	IndianRed         = X11Color{"Indian Red", color.RGBA{0xCD, 0x5C, 0x5C, 0xFF}, false}
	Indigo            = X11Color{"Indigo", color.RGBA{0x4B, 0x00, 0x82, 0xFF}, false}
	Ivory             = X11Color{"Ivory", color.RGBA{0xFF, 0xFF, 0xF0, 0xFF}, true}
	Khaki             = X11Color{"Khaki", color.RGBA{0xF0, 0xE6, 0x8C, 0xFF}, false}
	Lavender          = X11Color{"Lavender", color.RGBA{0xE6, 0xE6, 0xFA, 0xFF}, false}
	LavenderBlush     = X11Color{"Lavender Blush", color.RGBA{0xFF, 0xF0, 0xF5, 0xFF}, true}
	LawnGreen         = X11Color{"Lawn Green", color.RGBA{0x7C, 0xFC, 0x00, 0xFF}, false}
	LemonChiffon      = X11Color{"Lemon Chiffon", color.RGBA{0xFF, 0xFA, 0xCD, 0xFF}, true}
	LightBlue         = X11Color{"Light Blue", color.RGBA{0xAD, 0xD8, 0xE6, 0xFF}, false}
	LightCoral        = X11Color{"Light Coral", color.RGBA{0xF0, 0x80, 0x80, 0xFF}, false}
	LightCyan         = X11Color{"Light Cyan", color.RGBA{0xE0, 0xFF, 0xFF, 0xFF}, true}
	LightGoldenrod    = X11Color{"Light Goldenrod", color.RGBA{0xFA, 0xFA, 0xD2, 0xFF}, true}
	LightGray         = X11Color{"Light Gray", color.RGBA{0xD3, 0xD3, 0xD3, 0xFF}, false}
	LightGreen        = X11Color{"Light Green", color.RGBA{0x90, 0xEE, 0x90, 0xFF}, false}
	LightPink         = X11Color{"Light Pink", color.RGBA{0xFF, 0xB6, 0xC1, 0xFF}, false}
	LightSalmon       = X11Color{"Light Salmon", color.RGBA{0xFF, 0xA0, 0x7A, 0xFF}, false}
	LightSeaGreen     = X11Color{"Light Sea Green", color.RGBA{0x20, 0xB2, 0xAA, 0xFF}, false}
	LightSkyBlue      = X11Color{"Light Sky Blue", color.RGBA{0x87, 0xCE, 0xFA, 0xFF}, false}
	LightSlateGray    = X11Color{"Light Slate Gray", color.RGBA{0x77, 0x88, 0x99, 0xFF}, false}
	LightSteelBlue    = X11Color{"Light Steel Blue", color.RGBA{0xB0, 0xC4, 0xDE, 0xFF}, false}
	LightYellow       = X11Color{"Light Yellow", color.RGBA{0xFF, 0xFF, 0xE0, 0xFF}, true}
	Lime              = X11Color{"Lime", color.RGBA{0x00, 0xFF, 0x00, 0xFF}, false}
	LimeGreen         = X11Color{"Lime Green", color.RGBA{0x32, 0xCD, 0x32, 0xFF}, false}
	Linen             = X11Color{"Linen", color.RGBA{0xFA, 0xF0, 0xE6, 0xFF}, true}
	Magenta           = X11Color{"Magenta", color.RGBA{0xFF, 0x00, 0xFF, 0xFF}, false}
	MaroonX11         = X11Color{"Maroon (X11)", color.RGBA{0xB0, 0x30, 0x60, 0xFF}, false}
	MaroonW3C         = X11Color{"Maroon (W3C)", color.RGBA{0x7F, 0x00, 0x00, 0xFF}, false}
	MediumAquamarine  = X11Color{"Medium Aquamarine", color.RGBA{0x66, 0xCD, 0xAA, 0xFF}, false}
	MediumBlue        = X11Color{"Medium Blue", color.RGBA{0x00, 0x00, 0xCD, 0xFF}, false}
	MediumOrchid      = X11Color{"Medium Orchid", color.RGBA{0xBA, 0x55, 0xD3, 0xFF}, false}
	MediumPurple      = X11Color{"Medium Purple", color.RGBA{0x93, 0x70, 0xDB, 0xFF}, false}
	MediumSeaGreen    = X11Color{"Medium Sea Green", color.RGBA{0x3C, 0xB3, 0x71, 0xFF}, false}
	MediumSlateBlue   = X11Color{"Medium Slate Blue", color.RGBA{0x7B, 0x68, 0xEE, 0xFF}, false}
	MediumSpringGreen = X11Color{"Medium Spring Green", color.RGBA{0x00, 0xFA, 0x9A, 0xFF}, false}
	MediumTurquoise   = X11Color{"Medium Turquoise", color.RGBA{0x48, 0xD1, 0xCC, 0xFF}, false}
	MediumVioletRed   = X11Color{"Medium Violet Red", color.RGBA{0xC7, 0x15, 0x85, 0xFF}, false}
	MidnightBlue      = X11Color{"Midnight Blue", color.RGBA{0x19, 0x19, 0x70, 0xFF}, false}
	MintCream         = X11Color{"Mint Cream", color.RGBA{0xF5, 0xFF, 0xFA, 0xFF}, true}
	MistyRose         = X11Color{"Misty Rose", color.RGBA{0xFF, 0xE4, 0xE1, 0xFF}, true}
	Moccasin          = X11Color{"Moccasin", color.RGBA{0xFF, 0xE4, 0xB5, 0xFF}, false}
	NavajoWhite       = X11Color{"Navajo White", color.RGBA{0xFF, 0xDE, 0xAD, 0xFF}, false}
	Navy              = X11Color{"Navy", color.RGBA{0x00, 0x00, 0x80, 0xFF}, false}
	OldLace           = X11Color{"Old Lace", color.RGBA{0xFD, 0xF5, 0xE6, 0xFF}, true}
	Olive             = X11Color{"Olive", color.RGBA{0x80, 0x80, 0x00, 0xFF}, false}
	OliveDrab         = X11Color{"Olive Drab", color.RGBA{0x6B, 0x8E, 0x23, 0xFF}, false}
	Orange            = X11Color{"Orange", color.RGBA{0xFF, 0xA5, 0x00, 0xFF}, false}
	OrangeRed         = X11Color{"Orange Red", color.RGBA{0xFF, 0x45, 0x00, 0xFF}, false}
	Orchid            = X11Color{"Orchid", color.RGBA{0xDA, 0x70, 0xD6, 0xFF}, false}
	PaleGoldenrod     = X11Color{"Pale Goldenrod", color.RGBA{0xEE, 0xE8, 0xAA, 0xFF}, false}
	PaleGreen         = X11Color{"Pale Green", color.RGBA{0x98, 0xFB, 0x98, 0xFF}, false}
	PaleTurquoise     = X11Color{"Pale Turquoise", color.RGBA{0xAF, 0xEE, 0xEE, 0xFF}, false}
	PaleVioletRed     = X11Color{"Pale Violet Red", color.RGBA{0xDB, 0x70, 0x93, 0xFF}, false}
	PapayaWhip        = X11Color{"Papaya Whip", color.RGBA{0xFF, 0xEF, 0xD5, 0xFF}, true}
	PeachPuff         = X11Color{"Peach Puff", color.RGBA{0xFF, 0xDA, 0xB9, 0xFF}, false}
	Peru              = X11Color{"Peru", color.RGBA{0xCD, 0x85, 0x3F, 0xFF}, false}
	Pink              = X11Color{"Pink", color.RGBA{0xFF, 0xC0, 0xCB, 0xFF}, false}
	Plum              = X11Color{"Plum", color.RGBA{0xDD, 0xA0, 0xDD, 0xFF}, false}
	PowderBlue        = X11Color{"Powder Blue", color.RGBA{0xB0, 0xE0, 0xE6, 0xFF}, false}
	PurpleX11         = X11Color{"Purple (X11)", color.RGBA{0xA0, 0x20, 0xF0, 0xFF}, false}
	PurpleW3C         = X11Color{"Purple (W3C)", color.RGBA{0x7F, 0x00, 0x7F, 0xFF}, false}
	Red               = X11Color{"Red", color.RGBA{0xFF, 0x00, 0x00, 0xFF}, false}
	RosyBrown         = X11Color{"Rosy Brown", color.RGBA{0xBC, 0x8F, 0x8F, 0xFF}, false}
	RoyalBlue         = X11Color{"Royal Blue", color.RGBA{0x41, 0x69, 0xE1, 0xFF}, false}
	SaddleBrown       = X11Color{"Saddle Brown", color.RGBA{0x8B, 0x45, 0x13, 0xFF}, false}
	Salmon            = X11Color{"Salmon", color.RGBA{0xFA, 0x80, 0x72, 0xFF}, false}
	SandyBrown        = X11Color{"Sandy Brown", color.RGBA{0xF4, 0xA4, 0x60, 0xFF}, false}
	SeaGreen          = X11Color{"Sea Green", color.RGBA{0x2E, 0x8B, 0x57, 0xFF}, false}
	Seashell          = X11Color{"Seashell", color.RGBA{0xFF, 0xF5, 0xEE, 0xFF}, true}
	Sienna            = X11Color{"Sienna", color.RGBA{0xA0, 0x52, 0x2D, 0xFF}, false}
	Silver            = X11Color{"Silver", color.RGBA{0xC0, 0xC0, 0xC0, 0xFF}, false}
	SkyBlue           = X11Color{"Sky Blue", color.RGBA{0x87, 0xCE, 0xEB, 0xFF}, false}
	SlateBlue         = X11Color{"Slate Blue", color.RGBA{0x6A, 0x5A, 0xCD, 0xFF}, false}
	SlateGray         = X11Color{"Slate Gray", color.RGBA{0x70, 0x80, 0x90, 0xFF}, false}
	Snow              = X11Color{"Snow", color.RGBA{0xFF, 0xFA, 0xFA, 0xFF}, true}
	SpringGreen       = X11Color{"Spring Green", color.RGBA{0x00, 0xFF, 0x7F, 0xFF}, false}
	SteelBlue         = X11Color{"Steel Blue", color.RGBA{0x46, 0x82, 0xB4, 0xFF}, false}
	Tan               = X11Color{"Tan", color.RGBA{0xD2, 0xB4, 0x8C, 0xFF}, false}
	Teal              = X11Color{"Teal", color.RGBA{0x00, 0x80, 0x80, 0xFF}, false}
	Thistle           = X11Color{"Thistle", color.RGBA{0xD8, 0xBF, 0xD8, 0xFF}, false}
	Tomato            = X11Color{"Tomato", color.RGBA{0xFF, 0x63, 0x47, 0xFF}, false}
	Turquoise         = X11Color{"Turquoise", color.RGBA{0x40, 0xE0, 0xD0, 0xFF}, false}
	Violet            = X11Color{"Violet", color.RGBA{0xEE, 0x82, 0xEE, 0xFF}, false}
	Wheat             = X11Color{"Wheat", color.RGBA{0xF5, 0xDE, 0xB3, 0xFF}, false}
	White             = X11Color{"White", color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, true}
	WhiteSmoke        = X11Color{"White Smoke", color.RGBA{0xF5, 0xF5, 0xF5, 0xFF}, true}
	Yellow            = X11Color{"Yellow", color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, true}
	YellowGreen       = X11Color{"Yellow Green", color.RGBA{0x9A, 0xCD, 0x32, 0xFF}, false}

	colors = []X11Color{
		AliceBlue,
		AntiqueWhite,
		Aqua,
		Aquamarine,
		Azure,
		Beige,
		Bisque,
		Black,
		BlanchedAlmond,
		Blue,
		BlueViolet,
		Brown,
		Burlywood,
		CadetBlue,
		Chartreuse,
		Chocolate,
		Coral,
		Cornflower,
		Cornsilk,
		Crimson,
		Cyan,
		DarkBlue,
		DarkCyan,
		DarkGoldenrod,
		DarkGray,
		DarkGreen,
		DarkKhaki,
		DarkMagenta,
		DarkOliveGreen,
		DarkOrange,
		DarkOrchid,
		DarkRed,
		DarkSalmon,
		DarkSeaGreen,
		DarkSlateBlue,
		DarkSlateGray,
		DarkTurquoise,
		DarkViolet,
		DeepPink,
		DeepSkyBlue,
		DimGray,
		DodgerBlue,
		Firebrick,
		FloralWhite,
		ForestGreen,
		Fuchsia,
		Gainsboro,
		GhostWhite,
		Gold,
		Goldenrod,
		GrayX11,
		GrayW3C,
		GreenX11,
		GreenW3C,
		GreenYellow,
		Honeydew,
		HotPink,
		IndianRed,
		Indigo,
		Ivory,
		Khaki,
		Lavender,
		LavenderBlush,
		LawnGreen,
		LemonChiffon,
		LightBlue,
		LightCoral,
		LightCyan,
		LightGoldenrod,
		LightGray,
		LightGreen,
		LightPink,
		LightSalmon,
		LightSeaGreen,
		LightSkyBlue,
		LightSlateGray,
		LightSteelBlue,
		LightYellow,
		Lime,
		LimeGreen,
		Linen,
		Magenta,
		MaroonX11,
		MaroonW3C,
		MediumAquamarine,
		MediumBlue,
		MediumOrchid,
		MediumPurple,
		MediumSeaGreen,
		MediumSlateBlue,
		MediumSpringGreen,
		MediumTurquoise,
		MediumVioletRed,
		MidnightBlue,
		MintCream,
		MistyRose,
		Moccasin,
		NavajoWhite,
		Navy,
		OldLace,
		Olive,
		OliveDrab,
		Orange,
		OrangeRed,
		Orchid,
		PaleGoldenrod,
		PaleGreen,
		PaleTurquoise,
		PaleVioletRed,
		PapayaWhip,
		PeachPuff,
		Peru,
		Pink,
		Plum,
		PowderBlue,
		PurpleX11,
		PurpleW3C,
		Red,
		RosyBrown,
		RoyalBlue,
		SaddleBrown,
		Salmon,
		SandyBrown,
		SeaGreen,
		Seashell,
		Sienna,
		Silver,
		SkyBlue,
		SlateBlue,
		SlateGray,
		Snow,
		SpringGreen,
		SteelBlue,
		Tan,
		Teal,
		Thistle,
		Tomato,
		Turquoise,
		Violet,
		Wheat,
		White,
		WhiteSmoke,
		Yellow,
		YellowGreen,
	}

	names = map[string]X11Color{
		"Alice Blue":          AliceBlue,
		"Antique White":       AntiqueWhite,
		"Aqua":                Aqua,
		"Aquamarine":          Aquamarine,
		"Azure":               Azure,
		"Beige":               Beige,
		"Bisque":              Bisque,
		"Black":               Black,
		"Blanched Almond":     BlanchedAlmond,
		"Blue":                Blue,
		"Blue Violet":         BlueViolet,
		"Brown":               Brown,
		"Burlywood":           Burlywood,
		"Cadet Blue":          CadetBlue,
		"Chartreuse":          Chartreuse,
		"Chocolate":           Chocolate,
		"Coral":               Coral,
		"Cornflower":          Cornflower,
		"Cornsilk":            Cornsilk,
		"Crimson":             Crimson,
		"Cyan":                Cyan,
		"Dark Blue":           DarkBlue,
		"Dark Cyan":           DarkCyan,
		"Dark Goldenrod":      DarkGoldenrod,
		"Dark Gray":           DarkGray,
		"Dark Green":          DarkGreen,
		"Dark Khaki":          DarkKhaki,
		"Dark Magenta":        DarkMagenta,
		"Dark Olive Green":    DarkOliveGreen,
		"Dark Orange":         DarkOrange,
		"Dark Orchid":         DarkOrchid,
		"Dark Red":            DarkRed,
		"Dark Salmon":         DarkSalmon,
		"Dark Sea Green":      DarkSeaGreen,
		"Dark Slate Blue":     DarkSlateBlue,
		"Dark Slate Gray":     DarkSlateGray,
		"Dark Turquoise":      DarkTurquoise,
		"Dark Violet":         DarkViolet,
		"Deep Pink":           DeepPink,
		"Deep Sky Blue":       DeepSkyBlue,
		"Dim Gray":            DimGray,
		"Dodger Blue":         DodgerBlue,
		"Firebrick":           Firebrick,
		"Floral White":        FloralWhite,
		"Forest Green":        ForestGreen,
		"Fuchsia":             Fuchsia,
		"Gainsboro":           Gainsboro,
		"Ghost White":         GhostWhite,
		"Gold":                Gold,
		"Goldenrod":           Goldenrod,
		"Gray (X11)":          GrayX11,
		"Gray (W3C)":          GrayW3C,
		"Green (X11)":         GreenX11,
		"Green (W3C)":         GreenW3C,
		"Green Yellow":        GreenYellow,
		"Honeydew":            Honeydew,
		"Hot Pink":            HotPink,
		"Indian Red":          IndianRed,
		"Indigo":              Indigo,
		"Ivory":               Ivory,
		"Khaki":               Khaki,
		"Lavender":            Lavender,
		"Lavender Blush":      LavenderBlush,
		"Lawn Green":          LawnGreen,
		"Lemon Chiffon":       LemonChiffon,
		"Light Blue":          LightBlue,
		"Light Coral":         LightCoral,
		"Light Cyan":          LightCyan,
		"Light Goldenrod":     LightGoldenrod,
		"Light Gray":          LightGray,
		"Light Green":         LightGreen,
		"Light Pink":          LightPink,
		"Light Salmon":        LightSalmon,
		"Light Sea Green":     LightSeaGreen,
		"Light Sky Blue":      LightSkyBlue,
		"Light Slate Gray":    LightSlateGray,
		"Light Steel Blue":    LightSteelBlue,
		"Light Yellow":        LightYellow,
		"Lime":                Lime,
		"Lime Green":          LimeGreen,
		"Linen":               Linen,
		"Magenta":             Magenta,
		"Maroon (X11)":        MaroonX11,
		"Maroon (W3C)":        MaroonW3C,
		"Medium Aquamarine":   MediumAquamarine,
		"Medium Blue":         MediumBlue,
		"Medium Orchid":       MediumOrchid,
		"Medium Purple":       MediumPurple,
		"Medium Sea Green":    MediumSeaGreen,
		"Medium Slate Blue":   MediumSlateBlue,
		"Medium Spring Green": MediumSpringGreen,
		"Medium Turquoise":    MediumTurquoise,
		"Medium Violet Red":   MediumVioletRed,
		"Midnight Blue":       MidnightBlue,
		"Mint Cream":          MintCream,
		"Misty Rose":          MistyRose,
		"Moccasin":            Moccasin,
		"Navajo White":        NavajoWhite,
		"Navy":                Navy,
		"Old Lace":            OldLace,
		"Olive":               Olive,
		"Olive Drab":          OliveDrab,
		"Orange":              Orange,
		"Orange Red":          OrangeRed,
		"Orchid":              Orchid,
		"Pale Goldenrod":      PaleGoldenrod,
		"Pale Green":          PaleGreen,
		"Pale Turquoise":      PaleTurquoise,
		"Pale Violet Red":     PaleVioletRed,
		"Papaya Whip":         PapayaWhip,
		"Peach Puff":          PeachPuff,
		"Peru":                Peru,
		"Pink":                Pink,
		"Plum":                Plum,
		"Powder Blue":         PowderBlue,
		"Purple (X11)":        PurpleX11,
		"Purple (W3C)":        PurpleW3C,
		"Red":                 Red,
		"Rosy Brown":          RosyBrown,
		"Royal Blue":          RoyalBlue,
		"Saddle Brown":        SaddleBrown,
		"Salmon":              Salmon,
		"Sandy Brown":         SandyBrown,
		"Sea Green":           SeaGreen,
		"Seashell":            Seashell,
		"Sienna":              Sienna,
		"Silver":              Silver,
		"Sky Blue":            SkyBlue,
		"Slate Blue":          SlateBlue,
		"Slate Gray":          SlateGray,
		"Snow":                Snow,
		"Spring Green":        SpringGreen,
		"Steel Blue":          SteelBlue,
		"Tan":                 Tan,
		"Teal":                Teal,
		"Thistle":             Thistle,
		"Tomato":              Tomato,
		"Turquoise":           Turquoise,
		"Violet":              Violet,
		"Wheat":               Wheat,
		"White":               White,
		"White Smoke":         WhiteSmoke,
		"Yellow":              Yellow,
		"Yellow Green":        YellowGreen,
	}
)

// Random returns random color
func Random() X11Color {
	return colors[rand.Intn(len(colors))]
}

// RandomSeeded initialises generator with time-based seed and returns random color
func RandomSeeded() X11Color {
	rand.Seed(time.Now().Unix())
	return Random()
}

// GetByName returns X11Color by its name if found
func GetByName(name string) (x11color X11Color, found bool) {
	x11color, found = names[name]
	return
}
