package object // import "github.com/1Volk/vksdk/object"

// AppsApp type application type.
const (
	AppTypeApp          = "app"
	AppTypeGame         = "game"
	AppTypeSite         = "site"
	AppTypeStandalone   = "standalone"
	AppTypeVkApp        = "vk_app"
	AppTypeCommunityApp = "community_app"
	AppTypeHTML5Game    = "html5_game"
)

// AppsLeaderboardType leaderboardType type.
const (
	AppsLeaderboardTypeNotSupported = iota
	AppsLeaderboardTypeLevels
	AppsLeaderboardTypePoints
)

// AppsScreenOrientation supported screen orientation.
type AppsScreenOrientation int

// Possible values.
const (
	AppsScreenOrientationBoth AppsScreenOrientation = iota
	AppsScreenOrientationLandscape
	AppsScreenOrientationPortrait
)

// AppsCatalogBanner struct.
type AppsCatalogBanner struct {
	BackgroundColor  string `json:"background_color"`
	DescriptionColor string `json:"description_color"`
	Description      string `json:"description"`
	TitleColor       string `json:"title_color"`
}

// AppsApp struct.
type AppsApp struct {
	AuthorGroup       int                   `json:"author_group"`     // Official community's ID
	AuthorID          int64                   `json:"author_id"`        // Application author's ID
	AuthorURL         string                `json:"author_url"`       // Application author's URL
	Banner1120        string                `json:"banner_1120"`      // URL of the app banner with 1120 px in width
	Banner560         string                `json:"banner_560"`       // URL of the app banner with 560 px in width
	CatalogPosition   int                   `json:"catalog_position"` // Catalog position
	Description       string                `json:"description"`      // Application description
	Friends           []int                 `json:"friends"`
	Genre             string                `json:"genre"`         // Genre name
	GenreID           int                   `json:"genre_id"`      // Genre ID
	Icon139           string                `json:"icon_139"`      // URL of the app icon with 139 px in width
	Icon150           string                `json:"icon_150"`      // URL of the app icon with 150 px in width
	Icon278           string                `json:"icon_278"`      // URL of the app icon with 279 px in width
	Icon75            string                `json:"icon_75"`       // URL of the app icon with 75 px in width
	ID                int                   `json:"id"`            // Application ID
	International     BaseBoolInt           `json:"international"` // Information whether the application is multi language
	IsInCatalog       BaseBoolInt           `json:"is_in_catalog"` // Information whether application is in mobile catalog
	Installed         BaseBoolInt           `json:"installed"`
	PushEnabled       BaseBoolInt           `json:"push_enabled"`
	HideTabbar        BaseBoolInt           `json:"hide_tabbar"`
	IsNew             BaseBoolInt           `json:"is_new"`
	New               BaseBoolInt           `json:"new"`
	LeaderboardType   int                   `json:"leaderboard_type"`
	MembersCount      int                   `json:"members_count"`  // Members number
	PlatformID        int                   `json:"platform_id"`    // Application ID in store
	PublishedDate     int                   `json:"published_date"` // Date when the application has been published in Unixtime
	ScreenName        string                `json:"screen_name"`    // Screen name
	Screenshots       []PhotosPhoto         `json:"screenshots"`
	Section           string                `json:"section"` // Application section name
	Title             string                `json:"title"`   // Application title
	Type              string                `json:"type"`
	Icon16            string                `json:"icon_16"`
	Icon576           string                `json:"icon_576"`
	ScreenOrientation AppsScreenOrientation `json:"screen_orientation"`
	CatalogBanner     AppsCatalogBanner     `json:"catalog_banner"`
}

// AppsLeaderboard struct.
type AppsLeaderboard struct {
	Level  int `json:"level"`   // Level
	Points int `json:"points"`  // Points number
	Score  int `json:"score"`   // Score number
	UserID int64 `json:"user_id"` // User ID
}

// AppsScope Scope description.
type AppsScope struct {
	Name  string `json:"name"`  // Scope name
	Title string `json:"title"` // Scope title
}
