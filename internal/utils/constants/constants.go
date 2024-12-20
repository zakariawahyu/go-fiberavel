package constants

const (
	KeyConfigs   = "configs"
	KeyCouples   = "couples"
	KeyVenues    = "venues"
	KeyGalleries = "galleries"
	KeyGuests    = "guests"
	KeyWishes    = "wishes"
	KeyGift      = "gifts"

	CSRFContextKey = "__CSRF__"
)

var (
	Configuration = map[string]string{
		"meta":  "Meta",
		"cover": "Cover",
		"event": "Event",
		"story": "Story",
		"venue": "Venue",
		"gift":  "Gift",
		"rsvp":  "RSVP",
		"wish":  "Wish",
		"thank": "Thank",
	}

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
