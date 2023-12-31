package object // import "github.com/1Volk/vksdk/object"

import (
	"encoding/json"
)

// StoriesNarrativeInfo type.
type StoriesNarrativeInfo struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Views  int    `json:"views"`
}

// StoriesPromoData struct.
type StoriesPromoData struct {
	Name        string      `json:"name"`
	Photo50     string      `json:"photo_50"`
	Photo100    string      `json:"photo_100"`
	NotAnimated BaseBoolInt `json:"not_animated"`
}

// StoriesStoryLink struct.
type StoriesStoryLink struct {
	Text string `json:"text"` // Link text
	URL  string `json:"url"`  // Link URL
}

// StoriesReplies struct.
type StoriesReplies struct {
	Count int `json:"count"` // Replies number.
	New   int `json:"new"`   // New replies number.
}

// StoriesQuestions struct.
type StoriesQuestions struct {
	Count int `json:"count"` // Replies number.
	New   int `json:"new"`   // New replies number.
}

// StoriesStoryStats struct.
type StoriesStoryStats struct {
	Answer      StoriesStoryStatsStat `json:"answer"`
	Bans        StoriesStoryStatsStat `json:"bans"`
	OpenLink    StoriesStoryStatsStat `json:"open_link"`
	Replies     StoriesStoryStatsStat `json:"replies"`
	Shares      StoriesStoryStatsStat `json:"shares"`
	Subscribers StoriesStoryStatsStat `json:"subscribers"`
	Views       StoriesStoryStatsStat `json:"views"`
	Likes       StoriesStoryStatsStat `json:"likes"`
}

// StoriesStoryStatsStat struct.
type StoriesStoryStatsStat struct {
	Count int    `json:"count"` // Stat value
	State string `json:"state"`
}

// StoriesStory struct.
type StoriesStory struct {
	AccessKey string      `json:"access_key"` // Access key for private object.
	ExpiresAt int         `json:"expires_at"` // Story expiration time. Unixtime.
	CanHide   BaseBoolInt `json:"can_hide"`
	// Information whether story has question sticker and current user can send question to the author
	CanAsk BaseBoolInt `json:"can_ask"`
	// Information whether story has question sticker and current user can send anonymous question to the author
	CanAskAnonymous      BaseBoolInt              `json:"can_ask_anonymous"`
	CanComment           BaseBoolInt              `json:"can_comment"`   // Information whether current user can comment the story (0 - no, 1 - yes).
	CanReply             BaseBoolInt              `json:"can_reply"`     // Information whether current user can reply to the story (0 - no, 1 - yes).
	CanSee               BaseBoolInt              `json:"can_see"`       // Information whether current user can see the story (0 - no, 1 - yes).
	CanShare             BaseBoolInt              `json:"can_share"`     // Information whether current user can share the story (0 - no, 1 - yes).
	Date                 int                      `json:"date"`          // Date when story has been added in Unixtime.
	ID                   int                      `json:"id"`            // Story ID.
	IsDeleted            BaseBoolInt              `json:"is_deleted"`    // Information whether the story is deleted (false - no, true - yes).
	IsExpired            BaseBoolInt              `json:"is_expired"`    // Information whether the story is expired (false - no, true - yes).
	NoSound              BaseBoolInt              `json:"no_sound"`      // Is video without sound
	IsRestricted         BaseBoolInt              `json:"is_restricted"` // Does author have stories privacy restrictions
	Seen                 BaseBoolInt              `json:"seen"`          // Information whether current user has seen the story or not (0 - no, 1 - yes).
	IsOwnerPinned        BaseBoolInt              `json:"is_owner_pinned"`
	Link                 StoriesStoryLink         `json:"link"`
	OwnerID              int64                      `json:"owner_id"` // Story owner's ID.
	ParentStory          *StoriesStory            `json:"parent_story"`
	ParentStoryAccessKey string                   `json:"parent_story_access_key"` // Access key for private object.
	ParentStoryID        int                      `json:"parent_story_id"`         // Parent story ID.
	ParentStoryOwnerID   int                      `json:"parent_story_owner_id"`   // Parent story owner's ID.
	Photo                PhotosPhoto              `json:"photo"`
	Replies              StoriesReplies           `json:"replies"` // Replies to current story.
	Type                 string                   `json:"type"`
	Video                VideoVideo               `json:"video"`
	Views                int                      `json:"views"` // Views number.
	ClickableStickers    StoriesClickableStickers `json:"clickable_stickers"`
	TrackCode            string                   `json:"track_code"`
	LikesCount           int                      `json:"likes_count"`
	NarrativeID          int                      `json:"narrative_id"`
	NarrativeOwnerID     int                      `json:"narrative_owner_id"`
	NarrativeInfo        StoriesNarrativeInfo     `json:"narrative_info"`
	NarrativesCount      int                      `json:"narratives_count"`
	FirstNarrativeTitle  string                   `json:"first_narrative_title"`
	Questions            StoriesQuestions         `json:"questions"`
}

// StoriesClickableStickers struct.
//
// The field clickable_stickers is available in the history object.
// The sticker object is pasted by the developer on the client himself, only
// coordinates are transmitted to the server.
//
// https://vk.com/dev/objects/clickable_stickers
type StoriesClickableStickers struct {
	OriginalWidth     int                       `json:"original_width"`
	OriginalHeight    int                       `json:"original_height"`
	ClickableStickers []StoriesClickableSticker `json:"clickable_stickers"`
}

// NewClickableStickers return new StoriesClickableStickers.
//
// Requires the width and height of the original photo or video.
func NewClickableStickers(width, height int) *StoriesClickableStickers {
	return &StoriesClickableStickers{
		OriginalWidth:     width,
		OriginalHeight:    height,
		ClickableStickers: []StoriesClickableSticker{},
	}
}

// AddMention add mention sticker.
//
// Mention should be in the format of a VK mentioning, for example: [id1|name] or [club1|name].
func (cs *StoriesClickableStickers) AddMention(mention string, area []StoriesClickablePoint) *StoriesClickableStickers {
	cs.ClickableStickers = append(cs.ClickableStickers, StoriesClickableSticker{
		Type:          ClickableStickerMention,
		ClickableArea: area,
		Mention:       mention,
	})

	return cs
}

// AddHashtag add hashtag sticker.
//
// Hashtag must necessarily begin with the symbol #.
func (cs *StoriesClickableStickers) AddHashtag(hashtag string, area []StoriesClickablePoint) *StoriesClickableStickers {
	cs.ClickableStickers = append(cs.ClickableStickers, StoriesClickableSticker{
		Type:          ClickableStickerHashtag,
		ClickableArea: area,
		Hashtag:       hashtag,
	})

	return cs
}

// TODO: Add more clickable stickers func

// ToJSON returns the JSON encoding of StoriesClickableStickers.
func (cs StoriesClickableStickers) ToJSON() string {
	b, _ := json.Marshal(cs)
	return string(b)
}

// StoriesClickableSticker struct.
type StoriesClickableSticker struct { // nolint: maligned
	ID            int                     `json:"id"`
	Type          string                  `json:"type"`
	ClickableArea []StoriesClickablePoint `json:"clickable_area"`
	Style         string                  `json:"style,omitempty"`

	// type=post
	PostOwnerID int64 `json:"post_owner_id,omitempty"`
	PostID      int `json:"post_id,omitempty"`

	// type=sticker
	StickerID     int `json:"sticker_id,omitempty"`
	StickerPackID int `json:"sticker_pack_id,omitempty"`

	// type=place
	PlaceID int `json:"place_id,omitempty"`

	// type=question
	Question               string      `json:"question,omitempty"`
	QuestionButton         string      `json:"question_button,omitempty"`
	QuestionDefaultPrivate BaseBoolInt `json:"question_default_private,omitempty"`
	Color                  string      `json:"color,omitempty"`

	// type=mention
	Mention string `json:"mention,omitempty"`

	// type=hashtag
	Hashtag string `json:"hashtag,omitempty"`

	// type=link
	LinkObject  BaseLink `json:"link_object,omitempty"`
	TooltipText string   `json:"tooltip_text,omitempty"`

	// type=market_item
	Subtype string `json:"subtype,omitempty"`
	// LinkObject BaseLink         `json:"link_object,omitempty"` // subtype=aliexpress_product
	MarketItem MarketMarketItem `json:"market_item,omitempty"` // subtype=market_item

	// type=story_reply
	OwnerID int64 `json:"owner_id,omitempty"`
	StoryID int `json:"story_id,omitempty"`

	// type=owner
	// OwnerID int `json:"owner_id,omitempty"`

	// type=poll
	Poll PollsPoll `json:"poll,omitempty"`

	// type=music
	Audio          AudioAudioFull `json:"audio,omitempty"`
	AudioStartTime int            `json:"audio_start_time,omitempty"`

	// type=app
	App                      AppsApp     `json:"app"`
	AppContext               string      `json:"app_context"`
	HasNewInteractions       BaseBoolInt `json:"has_new_interactions"`
	IsBroadcastNotifyAllowed BaseBoolInt `json:"is_broadcast_notify_allowed"`
}

// TODO: сделать несколько структур для кликабельного стикера

// Type of clickable sticker.
const (
	ClickableStickerPost       = "post"
	ClickableStickerSticker    = "sticker"
	ClickableStickerPlace      = "place"
	ClickableStickerQuestion   = "question"
	ClickableStickerMention    = "mention"
	ClickableStickerHashtag    = "hashtag"
	ClickableStickerMarketItem = "market_item"
	ClickableStickerLink       = "link"
	ClickableStickerStoryReply = "story_reply"
	ClickableStickerOwner      = "owner"
	ClickableStickerPoll       = "poll"
	ClickableStickerMusic      = "music"
	ClickableStickerApp        = "app"
)

// Subtype of clickable sticker.
const (
	ClickableStickerSubtypeMarketItem        = "market_item"
	ClickableStickerSubtypeAliexpressProduct = "aliexpress_product"
)

// Clickable sticker style.
const (
	ClickableStickerTransparent   = "transparent"
	ClickableStickerBlueGradient  = "blue_gradient"
	ClickableStickerRedGradient   = "red_gradient"
	ClickableStickerUnderline     = "underline"
	ClickableStickerBlue          = "blue"
	ClickableStickerGreen         = "green"
	ClickableStickerWhite         = "white"
	ClickableStickerQuestionReply = "question_reply"
	ClickableStickerLight         = "light"
	ClickableStickerImpressive    = "impressive"
)

// StoriesClickablePoint struct.
type StoriesClickablePoint struct {
	X int `json:"x"`
	Y int `json:"y"`
}
