package api 

import (
	"github.com/1Volk/vksdk/object"
)

// SecureAddAppEventResponse struct.
type SecureAddAppEventResponse int // FIXME: not found documentation. https://github.com/VKCOM/vk-api-schema/issues/98

// SecureAddAppEvent adds user activity information to an application.
//
// https://vk.com/dev/secure.addAppEvent
func (vk *VK) SecureAddAppEvent(params Params) (response SecureAddAppEventResponse, err error) {
	err = vk.RequestUnmarshal("secure.addAppEvent", params, &response)
	return
}

// SecureCheckTokenResponse struct.
type SecureCheckTokenResponse object.SecureTokenChecked

// SecureCheckToken checks the user authentication in IFrame and Flash apps using the access_token parameter.
//
// https://vk.com/dev/secure.checkToken
func (vk *VK) SecureCheckToken(params Params) (response SecureCheckTokenResponse, err error) {
	err = vk.RequestUnmarshal("secure.checkToken", params, &response)
	return
}

// SecureGetAppBalance returns payment balance of the application in hundredth of a vote.
//
// https://vk.com/dev/secure.getAppBalance
func (vk *VK) SecureGetAppBalance(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("secure.getAppBalance", params, &response)
	return
}

// SecureGetSMSHistoryResponse struct.
type SecureGetSMSHistoryResponse []object.SecureSmsNotification

// SecureGetSMSHistory shows a list of SMS notifications sent by the application using secure.sendSMSNotification method.
//
// https://vk.com/dev/secure.getSMSHistory
func (vk *VK) SecureGetSMSHistory(params Params) (response SecureGetSMSHistoryResponse, err error) {
	err = vk.RequestUnmarshal("secure.getSMSHistory", params, &response)
	return
}

// SecureGetTransactionsHistoryResponse struct.
type SecureGetTransactionsHistoryResponse []object.SecureTransaction

// SecureGetTransactionsHistory shows history of votes transaction between users and the application.
//
// https://vk.com/dev/secure.getTransactionsHistory
func (vk *VK) SecureGetTransactionsHistory(params Params) (response SecureGetTransactionsHistoryResponse, err error) {
	err = vk.RequestUnmarshal("secure.getTransactionsHistory", params, &response)
	return
}

// SecureGetUserLevelResponse struct.
type SecureGetUserLevelResponse []object.SecureLevel

// SecureGetUserLevel returns one of the previously set game levels of one or more users in the application.
//
// https://vk.com/dev/secure.getUserLevel
func (vk *VK) SecureGetUserLevel(params Params) (response SecureGetUserLevelResponse, err error) {
	err = vk.RequestUnmarshal("secure.getUserLevel", params, &response)
	return
}

// SecureGiveEventStickerResponse struct.
type SecureGiveEventStickerResponse []struct {
	UserID int64    `json:"user_id"`
	Status string `json:"status"`
}

// SecureGiveEventSticker method.
//
// https://vk.com/dev/secure.giveEventSticker
func (vk *VK) SecureGiveEventSticker(params Params) (response SecureGiveEventStickerResponse, err error) {
	err = vk.RequestUnmarshal("secure.giveEventSticker", params, &response)
	return
}

// SecureSendNotificationResponse struct.
type SecureSendNotificationResponse []int //User ID

// SecureSendNotification sends notification to the user.
//
// https://vk.com/dev/secure.sendNotification
func (vk *VK) SecureSendNotification(params Params) (response SecureSendNotificationResponse, err error) {
	err = vk.RequestUnmarshal("secure.sendNotification", params, &response)
	return
}

// SecureSendSMSNotification sends SMS notification to a user's mobile device.
//
// https://vk.com/dev/secure.sendSMSNotification
func (vk *VK) SecureSendSMSNotification(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("secure.sendSMSNotification", params, &response)
	return
}

// SecureSetCounter sets a counter which is shown to the user in bold in the left menu.
//
// https://vk.com/dev/secure.setCounter
func (vk *VK) SecureSetCounter(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("secure.setCounter", params, &response)
	return
}
