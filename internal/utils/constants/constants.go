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
)
