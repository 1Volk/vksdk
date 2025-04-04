package api

import (
	"github.com/1Volk/vksdk/object"
)

// MessagesAddChatUser adds a new user to a chat.
//
// https://vk.com/dev/messages.addChatUser
func (vk *VK) MessagesAddChatUser(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.addChatUser", params, &response)
	return
}

// MessagesAllowMessagesFromGroup allows sending messages from community to the current user.
//
// https://vk.com/dev/messages.allowMessagesFromGroup
func (vk *VK) MessagesAllowMessagesFromGroup(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.allowMessagesFromGroup", params, &response)
	return
}

// MessagesCreateChat creates a chat with several participants.
//
// https://vk.com/dev/messages.createChat
func (vk *VK) MessagesCreateChat(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.createChat", params, &response)
	return
}

// MessagesDeleteResponse struct.
type MessagesDeleteResponse map[string]int

// MessagesDelete deletes one or more messages.
//
// https://vk.com/dev/messages.delete
func (vk *VK) MessagesDelete(params Params) (response MessagesDeleteResponse, err error) {
	err = vk.RequestUnmarshal("messages.delete", params, &response)
	return
}

// MessagesDeleteChatPhotoResponse struct.
type MessagesDeleteChatPhotoResponse struct {
	MessageID int                 `json:"message_id"`
	Chat      object.MessagesChat `json:"chat"`
}

// MessagesDeleteChatPhoto deletes a chat's cover picture.
//
// https://vk.com/dev/messages.deleteChatPhoto
func (vk *VK) MessagesDeleteChatPhoto(params Params) (response MessagesDeleteChatPhotoResponse, err error) {
	err = vk.RequestUnmarshal("messages.deleteChatPhoto", params, &response)
	return
}

// MessagesDeleteConversationResponse struct.
type MessagesDeleteConversationResponse struct {
	LastDeletedID int `json:"last_deleted_id"` // Id of the last message, that was deleted
}

// MessagesDeleteConversation deletes private messages in a conversation.
//
// https://vk.com/dev/messages.deleteConversation
func (vk *VK) MessagesDeleteConversation(params Params) (response MessagesDeleteConversationResponse, err error) {
	err = vk.RequestUnmarshal("messages.deleteConversation", params, &response)
	return
}

// MessagesDenyMessagesFromGroup denies sending message from community to the current user.
//
// https://vk.com/dev/messages.denyMessagesFromGroup
func (vk *VK) MessagesDenyMessagesFromGroup(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.denyMessagesFromGroup", params, &response)
	return
}

// MessageEventAnswer edits the message.
//
// https://vk.com/dev/sendMessageEventAnswer
func (vk *VK) MessageEventAnswer(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.sendMessageEventAnswer", params, &response)
	return
}

//	edits the message.
//
// https://vk.com/dev/messages.edit
func (vk *VK) MessagesEdit(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.edit", params, &response)
	return
}

// MessagesEditChat edits the title of a chat.
//
// https://vk.com/dev/messages.editChat
func (vk *VK) MessagesEditChat(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.editChat", params, &response)
	return
}

// MessagesGetByConversationMessageIDResponse struct.
type MessagesGetByConversationMessageIDResponse struct {
	Count int                      `json:"count"`
	Items []object.MessagesMessage `json:"items"`
	object.ExtendedResponse
}

// MessagesGetByConversationMessageID messages.getByConversationMessageId.
//
// https://vk.com/dev/messages.getByConversationMessageId
func (vk *VK) MessagesGetByConversationMessageID(params Params) (response MessagesGetByConversationMessageIDResponse, err error) {
	err = vk.RequestUnmarshal("messages.getByConversationMessageId", params, &response)
	return
}

// MessagesGetByIDResponse struct.
type MessagesGetByIDResponse struct {
	Count int                      `json:"count"`
	Items []object.MessagesMessage `json:"items"`
}

// MessagesGetByID returns messages by their IDs.
//
// extended=0
//
// https://vk.com/dev/messages.getById
func (vk *VK) MessagesGetByID(params Params) (response MessagesGetByIDResponse, err error) {
	params["extended"] = false
	err = vk.RequestUnmarshal("messages.getById", params, &response)

	return
}

// MessagesGetByIDExtendedResponse struct.
type MessagesGetByIDExtendedResponse struct {
	Count int                      `json:"count"`
	Items []object.MessagesMessage `json:"items"`
	object.ExtendedResponse
}

// MessagesGetByIDExtended returns messages by their IDs.
//
// extended=1
//
// https://vk.com/dev/messages.getById
func (vk *VK) MessagesGetByIDExtended(params Params) (response MessagesGetByIDExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("messages.getById", params, &response)

	return
}

// MessagesGetChatResponse struct.
type MessagesGetChatResponse object.MessagesChat

// MessagesGetChat returns information about a chat.
//
// https://vk.com/dev/messages.getChat
func (vk *VK) MessagesGetChat(params Params) (response MessagesGetChatResponse, err error) {
	err = vk.RequestUnmarshal("messages.getChat", params, &response)
	return
}

// MessagesGetChatChatIDsResponse struct.
type MessagesGetChatChatIDsResponse []object.MessagesChat

// MessagesGetChatChatIDs returns information about a chat.
//
// https://vk.com/dev/messages.getChat
func (vk *VK) MessagesGetChatChatIDs(params Params) (response MessagesGetChatChatIDsResponse, err error) {
	err = vk.RequestUnmarshal("messages.getChat", params, &response)
	return
}

// MessagesGetChatPreviewResponse struct.
type MessagesGetChatPreviewResponse struct {
	Preview object.MessagesChat `json:"preview"`
	object.ExtendedResponse
}

// MessagesGetChatPreview allows to receive chat preview by the invitation link.
//
// https://vk.com/dev/messages.getChatPreview
func (vk *VK) MessagesGetChatPreview(params Params) (response MessagesGetChatPreviewResponse, err error) {
	err = vk.RequestUnmarshal("messages.getChatPreview", params, &response)
	return
}

// MessagesGetConversationMembersResponse struct.
type MessagesGetConversationMembersResponse struct {
	Items []struct {
		MemberID     int64              `json:"member_id"`
		JoinDate     int                `json:"join_date"`
		InvitedBy    int                `json:"invited_by"`
		IsRestricted object.BaseBoolInt `json:"is_restricted_to_write"`
		IsOwner      object.BaseBoolInt `json:"is_owner,omitempty"`
		IsAdmin      object.BaseBoolInt `json:"is_admin,omitempty"`
		CanKick      object.BaseBoolInt `json:"can_kick,omitempty"`
	} `json:"items"`
	Count            int `json:"count"`
	ChatRestrictions struct {
		OnlyAdminsInvite   object.BaseBoolInt `json:"only_admins_invite"`
		OnlyAdminsEditPin  object.BaseBoolInt `json:"only_admins_edit_pin"`
		OnlyAdminsEditInfo object.BaseBoolInt `json:"only_admins_edit_info"`
		AdminsPromoteUsers object.BaseBoolInt `json:"admins_promote_users"`
	} `json:"chat_restrictions"`
	object.ExtendedResponse
}

// MessagesGetConversationMembers returns a list of IDs of users participating in a conversation.
//
// https://vk.com/dev/messages.getConversationMembers
func (vk *VK) MessagesGetConversationMembers(params Params) (response MessagesGetConversationMembersResponse, err error) {
	err = vk.RequestUnmarshal("messages.getConversationMembers", params, &response)
	return
}

// MessagesGetConversationsResponse struct.
type MessagesGetConversationsResponse struct {
	Count       int                                      `json:"count"`
	Items       []object.MessagesConversationWithMessage `json:"items"`
	UnreadCount int                                      `json:"unread_count"`
	object.ExtendedResponse
}

// MessagesGetConversations returns a list of conversations.
//
// https://vk.com/dev/messages.getConversations
func (vk *VK) MessagesGetConversations(params Params) (response MessagesGetConversationsResponse, err error) {
	err = vk.RequestUnmarshal("messages.getConversations", params, &response)
	return
}

// MessagesGetConversationsByIDResponse struct.
type MessagesGetConversationsByIDResponse struct {
	Count int                           `json:"count"`
	Items []object.MessagesConversation `json:"items"`
}

// MessagesGetConversationsByID returns conversations by their IDs.
//
// extended=0
//
// https://vk.com/dev/messages.getConversationsById
func (vk *VK) MessagesGetConversationsByID(params Params) (response MessagesGetConversationsByIDResponse, err error) {
	params["extended"] = false
	err = vk.RequestUnmarshal("messages.getConversationsById", params, &response)

	return
}

// MessagesGetConversationsByIDExtendedResponse struct.
type MessagesGetConversationsByIDExtendedResponse struct {
	Count int                           `json:"count"`
	Items []object.MessagesConversation `json:"items"`
	object.ExtendedResponse
}

// MessagesGetConversationsByIDExtended returns conversations by their IDs.
//
// extended=1
//
// https://vk.com/dev/messages.getConversationsById
func (vk *VK) MessagesGetConversationsByIDExtended(params Params) (response MessagesGetConversationsByIDExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("messages.getConversationsById", params, &response)

	return
}

// MessagesGetHistoryResponse struct.
type MessagesGetHistoryResponse struct {
	Count   int                      `json:"count"`
	Items   []object.MessagesMessage `json:"items"`
	InRead  int                      `json:"in_read"`
	OutRead int                      `json:"out_read"`
	object.ExtendedResponse
}

// MessagesGetHistory returns message history for the specified user or group chat.
//
// https://vk.com/dev/messages.getHistory
func (vk *VK) MessagesGetHistory(params Params) (response MessagesGetHistoryResponse, err error) {
	err = vk.RequestUnmarshal("messages.getHistory", params, &response)
	return
}

// MessagesGetHistoryAttachmentsResponse struct.
type MessagesGetHistoryAttachmentsResponse struct {
	Items    []object.MessagesHistoryAttachment `json:"items"`
	NextFrom string                             `json:"next_from"`
	object.ExtendedResponse
}

// MessagesGetHistoryAttachments returns media files from the dialog or group chat.
//
// https://vk.com/dev/messages.getHistoryAttachments
func (vk *VK) MessagesGetHistoryAttachments(params Params) (response MessagesGetHistoryAttachmentsResponse, err error) {
	err = vk.RequestUnmarshal("messages.getHistoryAttachments", params, &response)
	return
}

// MessagesGetImportantMessagesResponse struct.
type MessagesGetImportantMessagesResponse struct {
	Messages struct {
		Count int                      `json:"count"`
		Items []object.MessagesMessage `json:"items"`
	} `json:"messages"`
	Conversations []object.MessagesConversation `json:"conversations"`
	object.ExtendedResponse
}

// MessagesGetImportantMessages messages.getImportantMessages.
//
// https://vk.com/dev/messages.getImportantMessages
func (vk *VK) MessagesGetImportantMessages(params Params) (response MessagesGetImportantMessagesResponse, err error) {
	err = vk.RequestUnmarshal("messages.getImportantMessages", params, &response)
	return
}

// MessagesGetInviteLinkResponse struct.
type MessagesGetInviteLinkResponse struct {
	Link string `json:"link"`
}

// MessagesGetInviteLink receives a link to invite a user to the chat.
//
// https://vk.com/dev/messages.getInviteLink
func (vk *VK) MessagesGetInviteLink(params Params) (response MessagesGetInviteLinkResponse, err error) {
	err = vk.RequestUnmarshal("messages.getInviteLink", params, &response)
	return
}

// MessagesGetLastActivityResponse struct.
type MessagesGetLastActivityResponse object.MessagesLastActivity

// MessagesGetLastActivity returns a user's current status and date of last activity.
//
// https://vk.com/dev/messages.getLastActivity
func (vk *VK) MessagesGetLastActivity(params Params) (response MessagesGetLastActivityResponse, err error) {
	err = vk.RequestUnmarshal("messages.getLastActivity", params, &response)
	return
}

// MessagesGetLongPollHistoryResponse struct.
type MessagesGetLongPollHistoryResponse struct {
	History  [][]int              `json:"history"`
	Groups   []object.GroupsGroup `json:"groups"`
	Messages struct {
		Count int                      `json:"count"`
		Items []object.MessagesMessage `json:"items"`
	} `json:"messages"`
	Profiles []object.UsersUser `json:"profiles"`
	// Chats struct {} `json:"chats"`
	NewPTS        int                           `json:"new_pts"`
	FromPTS       int                           `json:"from_pts"`
	More          object.BaseBoolInt            `json:"chats"`
	Conversations []object.MessagesConversation `json:"conversations"`
}

// MessagesGetLongPollHistory returns updates in user's private messages.
//
// https://vk.com/dev/messages.getLongPollHistory
func (vk *VK) MessagesGetLongPollHistory(params Params) (response MessagesGetLongPollHistoryResponse, err error) {
	err = vk.RequestUnmarshal("messages.getLongPollHistory", params, &response)
	return
}

// MessagesGetLongPollServerResponse struct.
type MessagesGetLongPollServerResponse object.MessagesLongpollParams

// MessagesGetLongPollServer returns data required for connection to a Long Poll server.
//
// https://vk.com/dev/messages.getLongPollServer
func (vk *VK) MessagesGetLongPollServer(params Params) (response MessagesGetLongPollServerResponse, err error) {
	err = vk.RequestUnmarshal("messages.getLongPollServer", params, &response)
	return
}

// MessagesIsMessagesFromGroupAllowedResponse struct.
type MessagesIsMessagesFromGroupAllowedResponse struct {
	IsAllowed object.BaseBoolInt `json:"is_allowed"`
}

// MessagesIsMessagesFromGroupAllowed returns information whether
// sending messages from the community to current user is allowed.
//
// https://vk.com/dev/messages.isMessagesFromGroupAllowed
func (vk *VK) MessagesIsMessagesFromGroupAllowed(params Params) (response MessagesIsMessagesFromGroupAllowedResponse, err error) {
	err = vk.RequestUnmarshal("messages.isMessagesFromGroupAllowed", params, &response)
	return
}

// MessagesJoinChatByInviteLinkResponse struct.
type MessagesJoinChatByInviteLinkResponse struct {
	ChatID int `json:"chat_id"`
}

// MessagesJoinChatByInviteLink allows to enter the chat by the invitation link.
//
// https://vk.com/dev/messages.joinChatByInviteLink
func (vk *VK) MessagesJoinChatByInviteLink(params Params) (response MessagesJoinChatByInviteLinkResponse, err error) {
	err = vk.RequestUnmarshal("messages.joinChatByInviteLink", params, &response)
	return
}

// MessagesMarkAsAnsweredConversation messages.markAsAnsweredConversation.
//
// https://vk.com/dev/messages.markAsAnsweredConversation
func (vk *VK) MessagesMarkAsAnsweredConversation(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.markAsAnsweredConversation", params, &response)
	return
}

// MessagesMarkAsImportantResponse struct.
type MessagesMarkAsImportantResponse []int

// MessagesMarkAsImportant marks and un marks messages as important (starred).
//
// https://vk.com/dev/messages.markAsImportant
func (vk *VK) MessagesMarkAsImportant(params Params) (response MessagesMarkAsImportantResponse, err error) {
	err = vk.RequestUnmarshal("messages.markAsImportant", params, &response)
	return
}

// MessagesMarkAsImportantConversation messages.markAsImportantConversation.
//
// https://vk.com/dev/messages.markAsImportantConversation
func (vk *VK) MessagesMarkAsImportantConversation(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.markAsImportantConversation", params, &response)
	return
}

// MessagesMarkAsRead marks messages as read.
//
// https://vk.com/dev/messages.markAsRead
func (vk *VK) MessagesMarkAsRead(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.markAsRead", params, &response)
	return
}

// MessagesPinResponse struct.
type MessagesPinResponse object.MessagesMessage

// MessagesPin messages.pin.
//
// https://vk.com/dev/messages.pin
func (vk *VK) MessagesPin(params Params) (response MessagesPinResponse, err error) {
	err = vk.RequestUnmarshal("messages.pin", params, &response)
	return
}

// MessagesRemoveChatUser allows the current user to leave a chat or, if the current user started the chat, allows the user to remove another user from the chat.
//
// https://vk.com/dev/messages.removeChatUser
func (vk *VK) MessagesRemoveChatUser(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.removeChatUser", params, &response)
	return
}

// MessagesRestore restores a deleted message.
//
// https://vk.com/dev/messages.restore
func (vk *VK) MessagesRestore(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.restore", params, &response)
	return
}

// MessagesSearchResponse struct.
type MessagesSearchResponse struct {
	Count int                      `json:"count"`
	Items []object.MessagesMessage `json:"items"`
	object.ExtendedResponse
	Conversations []object.MessagesConversation `json:"conversations,omitempty"`
}

// MessagesSearch returns a list of the current user's private messages that match search criteria.
//
// https://vk.com/dev/messages.search
func (vk *VK) MessagesSearch(params Params) (response MessagesSearchResponse, err error) {
	err = vk.RequestUnmarshal("messages.search", params, &response)
	return
}

// MessagesSearchConversationsResponse struct.
type MessagesSearchConversationsResponse struct {
	Count int                           `json:"count"`
	Items []object.MessagesConversation `json:"items"`
	object.ExtendedResponse
}

// MessagesSearchConversations returns a list of conversations that match search criteria.
//
// https://vk.com/dev/messages.searchConversations
func (vk *VK) MessagesSearchConversations(params Params) (response MessagesSearchConversationsResponse, err error) {
	err = vk.RequestUnmarshal("messages.searchConversations", params, &response)
	return
}

// MessagesSend sends a message.
//
// For user_ids or peer_ids parameters, use MessagesSendUserIDs.
//
// https://vk.com/dev/messages.send
func (vk *VK) MessagesSend(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.send", params, &response)

	return
}

// MessagesSendUserIDsResponse struct.
type MessagesSendUserIDsResponse []struct {
	PeerID                int64 `json:"peer_id"`
	MessageID             int   `json:"message_id"`
	ConversationMessageID int   `json:"conversation_message_id"`
	Error                 struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	} `json:"error"`
}

// MessagesSendUserIDs sends a message.
//
// need user_ids or peer_ids;
//
// https://vk.com/dev/messages.send
func (vk *VK) MessagesSendUserIDs(params Params) (response MessagesSendUserIDsResponse, err error) {
	err = vk.RequestUnmarshal("messages.send", params, &response)
	return
}

// MessagesSendSticker sends a message.
//
// https://vk.com/dev/messages.sendSticker
func (vk *VK) MessagesSendSticker(params Params) (response int, err error) {
	params["user_ids"] = ""
	err = vk.RequestUnmarshal("messages.sendSticker", params, &response)

	return
}

// MessagesSetActivity changes the status of a user as typing in a conversation.
//
// https://vk.com/dev/messages.setActivity
func (vk *VK) MessagesSetActivity(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.setActivity", params, &response)
	return
}

// MessagesSetChatPhotoResponse struct.
type MessagesSetChatPhotoResponse struct {
	MessageID int                 `json:"message_id"`
	Chat      object.MessagesChat `json:"chat"`
}

// MessagesSetChatPhoto sets a previously-uploaded picture as the cover picture of a chat.
//
// https://vk.com/dev/messages.setChatPhoto
func (vk *VK) MessagesSetChatPhoto(params Params) (response MessagesSetChatPhotoResponse, err error) {
	err = vk.RequestUnmarshal("messages.setChatPhoto", params, &response)
	return
}

// MessagesUnpin messages.unpin.
//
// https://vk.com/dev/messages.unpin
func (vk *VK) MessagesUnpin(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.unpin", params, &response)
	return
}

// MessagesReaction messages.sendReaction?.
//
// https://vk.com/dev/messages.sendReaction
func (vk *VK) MessagesReaction(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("messages.sendReaction", params, &response)
	return
}

// MessagesResctrionsResponse struct.
type MessagesResctrionsResponse struct {
	FailedMemberIds []int64 `json:"failed_member_ids"`
}

// MessagesRestrictions messages.changeConversationMemberRestrictions
//
// https://vk.com/dev/messages.changeConversationMemberRestrictions
func (vk *VK) MessagesRestrictions(params Params) (response MessagesResctrionsResponse, err error) {
	err = vk.RequestUnmarshal("messages.changeConversationMemberRestrictions", params, &response)
	return
}
