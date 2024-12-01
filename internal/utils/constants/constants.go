package constants

const (
	KeyConfigs   = "configs"
	KeyCouples   = "couples"
	KeyVenues    = "venues"
	KeyGalleries = "galleries"
	KeyGuestList = "guest_lists"
	KeyWishes    = "wishes"
	KeyGift      = "gifts"

	CSRFContextKey = "__CSRF__"
)

var (
	CoupleTypes = map[string]string{
		"cpw": "CPW",
		"cpp": "CPP",
	}

	Banks = map[string]string{
		"Mandiri": "Bank Mandiri",
		"BCA":     "Bank Central Asia (BCA)",
		"BRI":     "Bank Rakyat Indonesia (BRI)",
		"BNI":     "Bank Negara Indonesia (BNI)",
		"Jago":    "Bank Artos Indonesia (Jago)",
		"SeaBank": "PT Bank Seabank Indonesia",
		"Krom":    "Krom Bank Indonesia",
	}
)
