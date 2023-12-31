package api 

import "github.com/1Volk/vksdk/object"

// PrettyCardsCreateResponse struct.
type PrettyCardsCreateResponse struct {
	OwnerID int64    `json:"owner_id"` // Owner ID of created pretty card
	CardID  string `json:"card_id"`  // Card ID of created pretty card
}

// PrettyCardsCreate method.
//
// https://vk.com/dev/prettyCards.create
func (vk *VK) PrettyCardsCreate(params Params) (response PrettyCardsCreateResponse, err error) {
	err = vk.RequestUnmarshal("prettyCards.create", params, &response)
	return
}

// PrettyCardsDeleteResponse struct.
type PrettyCardsDeleteResponse struct {
	OwnerID int64    `json:"owner_id"` // Owner ID of created pretty card
	CardID  string `json:"card_id"`  // Card ID of created pretty card
	Error   string `json:"error"`    // Error reason if error happened
}

// PrettyCardsDelete method.
//
// https://vk.com/dev/prettyCards.delete
func (vk *VK) PrettyCardsDelete(params Params) (response PrettyCardsDeleteResponse, err error) {
	err = vk.RequestUnmarshal("prettyCards.delete", params, &response)
	return
}

// PrettyCardsEditResponse struct.
type PrettyCardsEditResponse struct {
	OwnerID int64    `json:"owner_id"` // Owner ID of created pretty card
	CardID  string `json:"card_id"`  // Card ID of created pretty card
}

// PrettyCardsEdit method.
//
// https://vk.com/dev/prettyCards.edit
func (vk *VK) PrettyCardsEdit(params Params) (response PrettyCardsEditResponse, err error) {
	err = vk.RequestUnmarshal("prettyCards.edit", params, &response)
	return
}

// PrettyCardsGetResponse struct.
type PrettyCardsGetResponse struct {
	Count int                            `json:"count"` // Total number
	Items []object.PrettyCardsPrettyCard `json:"items"`
}

// PrettyCardsGet method.
//
// https://vk.com/dev/prettyCards.get
func (vk *VK) PrettyCardsGet(params Params) (response PrettyCardsGetResponse, err error) {
	err = vk.RequestUnmarshal("prettyCards.get", params, &response)
	return
}

// PrettyCardsGetByIDResponse struct.
type PrettyCardsGetByIDResponse []object.PrettyCardsPrettyCard

// PrettyCardsGetByID method.
//
// https://vk.com/dev/prettyCards.getById
func (vk *VK) PrettyCardsGetByID(params Params) (response PrettyCardsGetByIDResponse, err error) {
	err = vk.RequestUnmarshal("prettyCards.getById", params, &response)
	return
}

// PrettyCardsGetUploadURL method.
//
// https://vk.com/dev/prettyCards.getUploadURL
func (vk *VK) PrettyCardsGetUploadURL(params Params) (response string, err error) {
	err = vk.RequestUnmarshal("prettyCards.getUploadURL", params, &response)
	return
}
