package object // import "github.com/1Volk/vksdk/object"

// WidgetsCommentMedia struct.
type WidgetsCommentMedia struct {
	ItemID   int    `json:"item_id"`   // Media item ID
	OwnerID  int64    `json:"owner_id"`  // Media owner's ID
	ThumbSrc string `json:"thumb_src"` // URL of the preview image (type=photo only)
	Type     string `json:"type"`
}

// WidgetsCommentReplies struct.
type WidgetsCommentReplies struct {
	CanPost       BaseBoolInt                 `json:"can_post"` // Information whether current user can comment the post
	GroupsCanPost BaseBoolInt                 `json:"groups_can_post"`
	Count         int                         `json:"count"` // Comments number
	Replies       []WidgetsCommentRepliesItem `json:"replies"`
}

// WidgetsCommentRepliesItem struct.
type WidgetsCommentRepliesItem struct {
	Cid   int                `json:"cid"`  // Comment ID
	Date  int                `json:"date"` // Date when the comment has been added in Unixtime
	Likes WidgetsWidgetLikes `json:"likes"`
	Text  string             `json:"text"` // Comment text
	UID   int64                `json:"uid"`  // User ID
	User  UsersUser          `json:"user"`
}

// WidgetsWidgetComment struct.
type WidgetsWidgetComment struct {
	Attachments []WallCommentAttachment `json:"attachments"`
	CanDelete   BaseBoolInt             `json:"can_delete"` // Information whether current user can delete the comment
	IsFavorite  BaseBoolInt             `json:"is_favorite"`
	Comments    WidgetsCommentReplies   `json:"comments"`
	Date        int                     `json:"date"`    // Date when the comment has been added in Unixtime
	FromID      int64                     `json:"from_id"` // Comment author ID
	ID          int                     `json:"id"`      // Comment ID
	Likes       BaseLikesInfo           `json:"likes"`
	Media       WidgetsCommentMedia     `json:"media"`
	PostType    string                  `json:"post_type"` // Post type
	Reposts     BaseRepostsInfo         `json:"reposts"`
	Text        string                  `json:"text"`  // Comment text
	ToID        int                     `json:"to_id"` // Wall owner
	PostSource  WallPostSource          `json:"post_source"`
	Views       struct {
		Count int `json:"count"`
	} `json:"views"`
}

// WidgetsWidgetLikes struct.
type WidgetsWidgetLikes struct {
	Count int `json:"count"` // Likes number
}

// WidgetsWidgetPage struct.
type WidgetsWidgetPage struct {
	Comments    WidgetsWidgetLikes `json:"comments,omitempty"`
	Date        int                `json:"date,omitempty"`        // Date when Widgets on the page has been initialized firstly in Unixtime
	Description string             `json:"description,omitempty"` // Page description
	ID          int                `json:"id,omitempty"`          // Page ID
	Likes       WidgetsWidgetLikes `json:"likes,omitempty"`
	PageID      string             `json:"page_id,omitempty"` // page_id parameter value
	Photo       string             `json:"photo,omitempty"`   // URL of the preview image
	Title       string             `json:"title,omitempty"`   // Page title
	URL         string             `json:"url,omitempty"`     // Page absolute URL
}
