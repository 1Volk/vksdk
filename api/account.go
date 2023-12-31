package api 

import (
	"github.com/1Volk/vksdk/object"
)

// AccountBan account.ban.
//
// https://vk.com/dev/account.ban
func (vk *VK) AccountBan(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.ban", params, &response)
	return
}

// AccountChangePasswordResponse struct.
type AccountChangePasswordResponse struct {
	Token string `json:"token"`
}

// AccountChangePassword changes a user password after access is successfully restored with the auth.restore method.
//
// https://vk.com/dev/account.changePassword
func (vk *VK) AccountChangePassword(params Params) (response AccountChangePasswordResponse, err error) {
	err = vk.RequestUnmarshal("account.changePassword", params, &response)
	return
}

// AccountGetActiveOffersResponse struct.
type AccountGetActiveOffersResponse struct {
	Count int                   `json:"count"`
	Items []object.AccountOffer `json:"items"`
}

// AccountGetActiveOffers returns a list of active ads (offers).
// If the user fulfill their conditions, he will be able to get
// the appropriate number of votes to his balance.
//
// https://vk.com/dev/account.getActiveOffers
func (vk *VK) AccountGetActiveOffers(params Params) (response AccountGetActiveOffersResponse, err error) {
	err = vk.RequestUnmarshal("account.getActiveOffers", params, &response)
	return
}

// AccountGetAppPermissions gets settings of the user in this application.
//
// https://vk.com/dev/account.getAppPermissions
func (vk *VK) AccountGetAppPermissions(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.getAppPermissions", params, &response)
	return
}

// AccountGetBannedResponse struct.
type AccountGetBannedResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
	object.ExtendedResponse
}

// AccountGetBanned returns a user's blacklist.
//
// https://vk.com/dev/account.getBanned
func (vk *VK) AccountGetBanned(params Params) (response AccountGetBannedResponse, err error) {
	err = vk.RequestUnmarshal("account.getBanned", params, &response)
	return
}

// AccountGetCountersResponse struct.
type AccountGetCountersResponse object.AccountAccountCounters

// AccountGetCounters returns non-null values of user counters.
//
// https://vk.com/dev/account.getCounters
func (vk *VK) AccountGetCounters(params Params) (response AccountGetCountersResponse, err error) {
	err = vk.RequestUnmarshal("account.getCounters", params, &response)
	return
}

// AccountGetInfoResponse struct.
type AccountGetInfoResponse object.AccountInfo

// AccountGetInfo returns current account info.
//
// https://vk.com/dev/account.getInfo
func (vk *VK) AccountGetInfo(params Params) (response AccountGetInfoResponse, err error) {
	err = vk.RequestUnmarshal("account.getInfo", params, &response)
	return
}

// AccountGetProfileInfoResponse struct.
type AccountGetProfileInfoResponse object.AccountUserSettings

// AccountGetProfileInfo returns the current account info.
//
// https://vk.com/dev/account.getProfileInfo
func (vk *VK) AccountGetProfileInfo(params Params) (response AccountGetProfileInfoResponse, err error) {
	err = vk.RequestUnmarshal("account.getProfileInfo", params, &response)
	return
}

// AccountGetPushSettingsResponse struct.
type AccountGetPushSettingsResponse object.AccountPushSettings

// AccountGetPushSettings account.getPushSettings Gets settings of push notifications.
//
// https://vk.com/dev/account.getPushSettings
func (vk *VK) AccountGetPushSettings(params Params) (response AccountGetPushSettingsResponse, err error) {
	err = vk.RequestUnmarshal("account.getPushSettings", params, &response)
	return
}

// AccountRegisterDevice subscribes an iOS/Android/Windows/Mac based device to receive push notifications.
//
// https://vk.com/dev/account.registerDevice
func (vk *VK) AccountRegisterDevice(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.registerDevice", params, &response)
	return
}

// AccountSaveProfileInfoResponse struct.
type AccountSaveProfileInfoResponse struct {
	Changed     object.BaseBoolInt        `json:"changed"`
	NameRequest object.AccountNameRequest `json:"name_request"`
}

// AccountSaveProfileInfo edits current profile info.
//
// https://vk.com/dev/account.saveProfileInfo
func (vk *VK) AccountSaveProfileInfo(params Params) (response AccountSaveProfileInfoResponse, err error) {
	err = vk.RequestUnmarshal("account.saveProfileInfo", params, &response)
	return
}

// AccountSetInfo allows to edit the current account info.
//
// https://vk.com/dev/account.setInfo
func (vk *VK) AccountSetInfo(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.setInfo", params, &response)
	return
}

// AccountSetNameInMenu sets an application screen name
// (up to 17 characters), that is shown to the user in the left menu.
//
// https://vk.com/dev/account.setNameInMenu
func (vk *VK) AccountSetNameInMenu(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.setNameInMenu", params, &response)
	return
}

// AccountSetOffline marks a current user as offline.
//
// https://vk.com/dev/account.setOffline
func (vk *VK) AccountSetOffline(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.setOffline", params, &response)
	return
}

// AccountSetOnline marks the current user as online for 5 minutes.
//
// https://vk.com/dev/account.setOnline
func (vk *VK) AccountSetOnline(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.setOnline", params, &response)
	return
}

// AccountSetPushSettings change push settings.
//
// https://vk.com/dev/account.setPushSettings
func (vk *VK) AccountSetPushSettings(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.setPushSettings", params, &response)
	return
}

// AccountSetSilenceMode mutes push notifications for the set period of time.
//
// https://vk.com/dev/account.setSilenceMode
func (vk *VK) AccountSetSilenceMode(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.setSilenceMode", params, &response)
	return
}

// AccountUnban account.unban.
//
// https://vk.com/dev/account.unban
func (vk *VK) AccountUnban(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.unban", params, &response)
	return
}

// AccountUnregisterDevice unsubscribes a device from push notifications.
//
// https://vk.com/dev/account.unregisterDevice
func (vk *VK) AccountUnregisterDevice(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.unregisterDevice", params, &response)
	return
}
