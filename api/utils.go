package api 

import (
	"encoding/json"

	"github.com/1Volk/vksdk/object"
)

// UtilsCheckLinkResponse struct.
type UtilsCheckLinkResponse object.UtilsLinkChecked

// UtilsCheckLink checks whether a link is blocked in VK.
//
// https://vk.com/dev/utils.checkLink
func (vk *VK) UtilsCheckLink(params Params) (response UtilsCheckLinkResponse, err error) {
	err = vk.RequestUnmarshal("utils.checkLink", params, &response)
	return
}

// UtilsDeleteFromLastShortened deletes shortened link from user's list.
//
// https://vk.com/dev/utils.deleteFromLastShortened
func (vk *VK) UtilsDeleteFromLastShortened(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("utils.deleteFromLastShortened", params, &response)
	return
}

// UtilsGetLastShortenedLinksResponse struct.
type UtilsGetLastShortenedLinksResponse struct {
	Count int                             `json:"count"`
	Items []object.UtilsLastShortenedLink `json:"items"`
}

// UtilsGetLastShortenedLinks returns a list of user's shortened links.
//
// https://vk.com/dev/utils.getLastShortenedLinks
func (vk *VK) UtilsGetLastShortenedLinks(params Params) (response UtilsGetLastShortenedLinksResponse, err error) {
	err = vk.RequestUnmarshal("utils.getLastShortenedLinks", params, &response)
	return
}

// UtilsGetLinkStatsResponse struct.
type UtilsGetLinkStatsResponse object.UtilsLinkStats

// UtilsGetLinkStats returns stats data for shortened link.
//
// extended=0
//
// https://vk.com/dev/utils.getLinkStats
func (vk *VK) UtilsGetLinkStats(params Params) (response UtilsGetLinkStatsResponse, err error) {
	params["extended"] = false
	err = vk.RequestUnmarshal("utils.getLinkStats", params, &response)

	return
}

// UtilsGetLinkStatsExtendedResponse struct.
type UtilsGetLinkStatsExtendedResponse object.UtilsLinkStatsExtended

// UtilsGetLinkStatsExtended returns stats data for shortened link.
//
// extended=1
//
// https://vk.com/dev/utils.getLinkStats
func (vk *VK) UtilsGetLinkStatsExtended(params Params) (response UtilsGetLinkStatsExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("utils.getLinkStats", params, &response)

	return
}

// UtilsGetServerTime returns the current time of the VK server.
//
// https://vk.com/dev/utils.getServerTime
func (vk *VK) UtilsGetServerTime(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("utils.getServerTime", params, &response)
	return
}

// UtilsGetShortLinkResponse struct.
type UtilsGetShortLinkResponse object.UtilsShortLink

// UtilsGetShortLink allows to receive a link shortened via vk.cc.
//
// https://vk.com/dev/utils.getShortLink
func (vk *VK) UtilsGetShortLink(params Params) (response UtilsGetShortLinkResponse, err error) {
	err = vk.RequestUnmarshal("utils.getShortLink", params, &response)
	return
}

// UtilsResolveScreenNameResponse struct.
type UtilsResolveScreenNameResponse object.UtilsDomainResolved

// UtilsResolveScreenName detects a type of object (e.g., user, community, application) and its ID by screen name.
//
// https://vk.com/dev/utils.resolveScreenName
func (vk *VK) UtilsResolveScreenName(params Params) (response UtilsResolveScreenNameResponse, err error) {
	rawResponse, err := vk.Request("utils.resolveScreenName", params)
	// Если короткое имя screen_name не занято, то будет возвращён пустой объект.
	if err != nil || string(rawResponse) == "[]" {
		return
	}

	err = json.Unmarshal(rawResponse, &response)

	return
}
