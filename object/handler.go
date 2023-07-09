package object // import "github.com/1Volk/vksdk/object"

import (
	"bytes"
	"encoding/json"
)

// MessageNewFunc func.
type MessageNewFunc func(MessageNewObject, int64)

// MessageNewObject struct.
type MessageNewObject struct {
	Message    MessagesMessage `json:"message"`
	ClientInfo ClientInfo      `json:"client_info"`
}

type EventNewObject struct {
	UserID       int64            `json:"user_id"`
	PeerID       int64            `json:"peer_id"`
	EventID      string         `json:"event_id"`
	Payload      map[string]int `json:"payload"`
	Conversation int            `json:"conversation_message_id"`
}

// UnmarshalJSON need for support api version < 5.103.
//
// To unmarshal JSON into a value implementing the Unmarshaler interface,
// Unmarshal calls that value's UnmarshalJSON method.
// See more https://golang.org/pkg/encoding/json/#Unmarshal
func (obj *MessageNewObject) UnmarshalJSON(data []byte) (err error) {
	// The renamed type is necessary to avoid loops
	type renamedMessageNewObject MessageNewObject

	var renamedObj renamedMessageNewObject

	if bytes.Contains(data, []byte(`"message":`)) {
		// v >= 5.103
		err = json.Unmarshal(data, &renamedObj)
	} else {
		// Support v < 5.103
		err = json.Unmarshal(data, &renamedObj.Message)
	}

	obj.Message = renamedObj.Message
	obj.ClientInfo = renamedObj.ClientInfo

	return
}

// Message Event func.
type MessageEventFunc func(EventNewObject, int64)

// MessageReplyFunc func.
type MessageReplyFunc func(MessageReplyObject, int64)

// MessageReplyObject struct.
type MessageReplyObject MessagesMessage

// MessageEditFunc func.
type MessageEditFunc func(MessageEditObject, int64)

// MessageEditObject struct.
type MessageEditObject MessagesMessage

// MessageAllowFunc func.
type MessageAllowFunc func(MessageAllowObject, int64)

// MessageAllowObject struct.
type MessageAllowObject struct {
	UserID int64    `json:"user_id"`
	Key    string `json:"key"`
}

// MessageDenyFunc func.
type MessageDenyFunc func(MessageDenyObject, int64)

// MessageDenyObject struct.
type MessageDenyObject struct {
	UserID int64 `json:"user_id"`
}

// MessageTypingStateFunc func.
type MessageTypingStateFunc func(MessageTypingStateObject, int64)

// MessageTypingStateObject struct.
type MessageTypingStateObject struct {
	State  string `json:"state"`
	FromID int64    `json:"from_id"`
	ToID   int    `json:"to_id"`
}

// PhotoNewFunc func.
type PhotoNewFunc func(PhotoNewObject, int64)

// PhotoNewObject struct.
type PhotoNewObject PhotosPhoto

// PhotoCommentNewFunc func.
type PhotoCommentNewFunc func(PhotoCommentNewObject, int64)

// PhotoCommentNewObject struct.
type PhotoCommentNewObject WallWallComment

// PhotoCommentEditFunc func.
type PhotoCommentEditFunc func(PhotoCommentEditObject, int64)

// PhotoCommentEditObject struct.
type PhotoCommentEditObject WallWallComment

// PhotoCommentRestoreFunc func.
type PhotoCommentRestoreFunc func(PhotoCommentRestoreObject, int64)

// PhotoCommentRestoreObject struct.
type PhotoCommentRestoreObject WallWallComment

// PhotoCommentDeleteFunc func.
type PhotoCommentDeleteFunc func(PhotoCommentDeleteObject, int64)

// PhotoCommentDeleteObject struct.
type PhotoCommentDeleteObject struct {
	OwnerID   int64 `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int64 `json:"user_id"`
	DeleterID int64 `json:"deleter_id"`
	PhotoID   int `json:"photo_id"`
}

// AudioNewFunc func.
type AudioNewFunc func(AudioNewObject, int64)

// AudioNewObject struct.
type AudioNewObject AudioAudioFull

// VideoNewFunc func.
type VideoNewFunc func(VideoNewObject, int64)

// VideoNewObject struct.
type VideoNewObject VideoVideo

// VideoCommentNewFunc func.
type VideoCommentNewFunc func(VideoCommentNewObject, int64)

// VideoCommentNewObject struct.
type VideoCommentNewObject WallWallComment

// VideoCommentEditFunc func.
type VideoCommentEditFunc func(VideoCommentEditObject, int64)

// VideoCommentEditObject struct.
type VideoCommentEditObject WallWallComment

// VideoCommentRestoreFunc func.
type VideoCommentRestoreFunc func(VideoCommentRestoreObject, int64)

// VideoCommentRestoreObject struct.
type VideoCommentRestoreObject WallWallComment

// VideoCommentDeleteFunc func.
type VideoCommentDeleteFunc func(VideoCommentDeleteObject, int64)

// VideoCommentDeleteObject struct.
type VideoCommentDeleteObject struct {
	OwnerID   int64 `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int64 `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	VideoID   int `json:"video_id"`
}

// WallPostNewFunc func.
type WallPostNewFunc func(WallPostNewObject, int64)

// WallPostNewObject struct.
type WallPostNewObject WallWallpost

// WallRepostFunc func.
type WallRepostFunc func(WallRepostObject, int64)

// WallRepostObject struct.
type WallRepostObject WallWallpost

// WallReplyNewFunc func.
type WallReplyNewFunc func(WallReplyNewObject, int64)

// WallReplyNewObject struct.
type WallReplyNewObject WallWallComment

// WallReplyEditFunc func.
type WallReplyEditFunc func(WallReplyEditObject, int64)

// WallReplyEditObject struct.
type WallReplyEditObject WallWallComment

// WallReplyRestoreFunc func.
type WallReplyRestoreFunc func(WallReplyRestoreObject, int64)

// WallReplyRestoreObject struct.
type WallReplyRestoreObject WallWallComment

// WallReplyDeleteFunc func.
type WallReplyDeleteFunc func(WallReplyDeleteObject, int64)

// WallReplyDeleteObject struct.
type WallReplyDeleteObject struct {
	OwnerID   int64 `json:"owner_id"`
	ID        int `json:"id"`
	DeleterID int `json:"deleter_id"`
	PostID    int `json:"post_id"`
}

// BoardPostNewFunc func.
type BoardPostNewFunc func(BoardPostNewObject, int64)

// BoardPostNewObject struct.
type BoardPostNewObject BoardTopicComment

// BoardPostEditFunc func.
type BoardPostEditFunc func(BoardPostEditObject, int64)

// BoardPostEditObject struct.
type BoardPostEditObject BoardTopicComment

// BoardPostRestoreFunc func.
type BoardPostRestoreFunc func(BoardPostRestoreObject, int64)

// BoardPostRestoreObject struct.
type BoardPostRestoreObject BoardTopicComment

// BoardPostDeleteFunc func.
type BoardPostDeleteFunc func(BoardPostDeleteObject, int64)

// BoardPostDeleteObject struct.
type BoardPostDeleteObject struct {
	TopicOwnerID int64 `json:"topic_owner_id"`
	TopicID      int `json:"topic_id"`
	ID           int `json:"id"`
}

// MarketCommentNewFunc func.
type MarketCommentNewFunc func(MarketCommentNewObject, int64)

// MarketCommentNewObject struct.
type MarketCommentNewObject WallWallComment

// MarketCommentEditFunc func.
type MarketCommentEditFunc func(MarketCommentEditObject, int64)

// MarketCommentEditObject struct.
type MarketCommentEditObject WallWallComment

// MarketCommentRestoreFunc func.
type MarketCommentRestoreFunc func(MarketCommentRestoreObject, int64)

// MarketCommentRestoreObject struct.
type MarketCommentRestoreObject WallWallComment

// MarketCommentDeleteFunc func.
type MarketCommentDeleteFunc func(MarketCommentDeleteObject, int64)

// MarketCommentDeleteObject struct.
type MarketCommentDeleteObject struct {
	OwnerID   int64 `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int64 `json:"user_id"`
	DeleterID int64 `json:"deleter_id"`
	ItemID    int `json:"item_id"`
}

// MarketOrderNewFunc func.
type MarketOrderNewFunc func(MarketOrderNewObject, int64)

// MarketOrderNewObject struct.
type MarketOrderNewObject MarketOrder

// MarketOrderEditFunc func.
type MarketOrderEditFunc func(MarketOrderEditObject, int64)

// MarketOrderEditObject struct.
type MarketOrderEditObject MarketOrder

// GroupLeaveFunc func.
type GroupLeaveFunc func(GroupLeaveObject, int64)

// GroupLeaveObject struct.
type GroupLeaveObject struct {
	UserID int64         `json:"user_id"`
	Self   BaseBoolInt `json:"self"`
}

// GroupJoinFunc func.
type GroupJoinFunc func(GroupJoinObject, int64)

// GroupJoinObject struct.
type GroupJoinObject struct {
	UserID   int64    `json:"user_id"`
	JoinType string `json:"join_type"`
}

// UserBlockFunc func.
type UserBlockFunc func(UserBlockObject, int64)

// UserBlockObject struct.
type UserBlockObject struct {
	AdminID     int64    `json:"admin_id"`
	UserID      int64    `json:"user_id"`
	UnblockDate int    `json:"unblock_date"`
	Reason      int    `json:"reason"`
	Comment     string `json:"comment"`
}

// UserUnblockFunc func.
type UserUnblockFunc func(UserUnblockObject, int64)

// UserUnblockObject struct.
type UserUnblockObject struct {
	AdminID   int64 `json:"admin_id"`
	UserID    int64 `json:"user_id"`
	ByEndDate int `json:"by_end_date"`
}

// PollVoteNewFunc func.
type PollVoteNewFunc func(PollVoteNewObject, int64)

// PollVoteNewObject struct.
//
// BUG(VK): при голосовании за несколько вариантов, возвращается только один.
type PollVoteNewObject struct {
	OwnerID  int64 `json:"owner_id"`
	PollID   int `json:"poll_id"`
	OptionID int `json:"option_id"`
	UserID   int64 `json:"user_id"`
}

// GroupOfficersEditFunc func.
type GroupOfficersEditFunc func(GroupOfficersEditObject, int64)

// GroupOfficersEditObject struct.
type GroupOfficersEditObject struct {
	AdminID  int64 `json:"admin_id"`
	UserID   int64 `json:"user_id"`
	LevelOld int `json:"level_old"`
	LevelNew int `json:"level_new"`
}

// Changes struct.
type Changes struct {
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
}

// ChangesInt struct.
type ChangesInt struct {
	OldValue int `json:"old_value"`
	NewValue int `json:"new_value"`
}

// GroupChangeSettingsFunc func.
type GroupChangeSettingsFunc func(GroupChangeSettingsObject, int64)

// GroupChangeSettingsObject struct
// спасибо vk.com/eee
// BUG(VK): Phone https://vk.com/bugtracker?act=show&id=64240
// BUG(VK): Email https://vk.com/bugtracker?act=show&id=86650
type GroupChangeSettingsObject struct {
	UserID  int64 `json:"user_id"`
	Changes struct {
		Title                 Changes    `json:"title"`
		Description           Changes    `json:"description"`
		Access                ChangesInt `json:"access"`
		ScreenName            Changes    `json:"screen_name"`
		PublicCategory        ChangesInt `json:"public_category"`
		PublicSubcategory     ChangesInt `json:"public_subcategory"`
		AgeLimits             ChangesInt `json:"age_limits"`
		Website               Changes    `json:"website"`
		StatusDefault         Changes    `json:"status_default"`
		Wall                  ChangesInt `json:"wall"`                    // на основе ответа
		Replies               ChangesInt `json:"replies"`                 // на основе ответа
		Topics                ChangesInt `json:"topics"`                  // на основе ответа
		Audio                 ChangesInt `json:"audio"`                   // на основе ответа
		Photos                ChangesInt `json:"photos"`                  // на основе ответа
		Video                 ChangesInt `json:"video"`                   // на основе ответа
		Market                ChangesInt `json:"market"`                  // на основе ответа
		Docs                  ChangesInt `json:"docs"`                    // на основе ответа
		Messages              ChangesInt `json:"messages"`                // на основе ответа
		EventGroupID          ChangesInt `json:"event_group_id"`          // на основе ответа
		Links                 Changes    `json:"links"`                   // на основе ответа
		Email                 Changes    `json:"email"`                   // на основе ответа
		EventStartDate        ChangesInt `json:"event_start_date::"`      // на основе ответа
		EventFinishDate       ChangesInt `json:"event_finish_date:"`      // на основе ответа
		Subject               Changes    `json:"subject"`                 // на основе ответа
		MarketWiki            Changes    `json:"market_wiki"`             // на основе ответа
		DisableMarketComments ChangesInt `json:"disable_market_comments"` // на основе ответа
		Phone                 ChangesInt `json:"phone"`                   // на основе ответа
		CountryID             ChangesInt `json:"country_id"`              // на основе ответа
		CityID                ChangesInt `json:"city_id"`                 // на основе ответа
	} `json:"Changes"`
}

// GroupChangePhotoFunc func.
type GroupChangePhotoFunc func(GroupChangePhotoObject, int64)

// GroupChangePhotoObject struct.
type GroupChangePhotoObject struct {
	UserID int64         `json:"user_id"`
	Photo  PhotosPhoto `json:"photo"`
}

// VkpayTransactionFunc func.
type VkpayTransactionFunc func(VkpayTransactionObject, int64)

// VkpayTransactionObject struct.
type VkpayTransactionObject struct {
	FromID      int64    `json:"from_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Date        int    `json:"date"`
}

// LeadFormsNewFunc func.
type LeadFormsNewFunc func(LeadFormsNewObject, int64)

// LeadFormsNewObject struct.
type LeadFormsNewObject struct {
	LeadID   int    `json:"lead_id"`
	GroupID  int64    `json:"group_id"`
	UserID   int64    `json:"user_id"`
	FormID   int    `json:"form_id"`
	FormName string `json:"form_name"`
	AdID     int    `json:"ad_id"`
	Answers  []struct {
		Key      string `json:"key"`
		Question string `json:"question"`
		Answer   string `json:"answer"`
	} `json:"answers"`
}

// AppPayloadFunc func.
type AppPayloadFunc func(AppPayloadObject, int64)

// AppPayloadObject struct.
type AppPayloadObject struct {
	UserID  int64    `json:"user_id"`
	AppID   int    `json:"app_id"`
	Payload string `json:"payload"`
}

// MessageReadFunc func.
type MessageReadFunc func(MessageReadObject, int64)

// MessageReadObject struct.
type MessageReadObject struct {
	FromID        int64 `json:"from_id"`
	PeerID        int64 `json:"peer_id"`
	ReadMessageID int `json:"read_message_id"`
}

// LikeAddFunc func.
type LikeAddFunc func(LikeAddObject, int64)

// LikeAddObject struct.
type LikeAddObject struct {
	LikerID       int    `json:"liker_id"`
	ObjectType    string `json:"object_type"`
	ObjectOwnerID int64    `json:"object_owner_id"`
	ObjectID      int    `json:"object_id"`
	ThreadReplyID int    `json:"thread_reply_id"`
	PostID        int    `json:"post_id"` // for comment
}

// LikeRemoveFunc func.
type LikeRemoveFunc func(LikeRemoveObject, int64)

// LikeRemoveObject struct.
type LikeRemoveObject struct {
	LikerID       int    `json:"liker_id"`
	ObjectType    string `json:"object_type"`
	ObjectOwnerID int64    `json:"object_owner_id"`
	ObjectID      int    `json:"object_id"`
	ThreadReplyID int    `json:"thread_reply_id"`
	PostID        int    `json:"post_id"` // for comment
}
