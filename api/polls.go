package api 

import "github.com/1Volk/vksdk/object"

// PollsAddVote adds the current user's vote to the selected answer in the poll.
//
// https://vk.com/dev/polls.addVote
func (vk *VK) PollsAddVote(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("polls.addVote", params, &response)
	return
}

// PollsCreateResponse struct.
type PollsCreateResponse object.PollsPoll

// PollsCreate creates polls that can be attached to the users' or communities' posts.
//
// https://vk.com/dev/polls.create
func (vk *VK) PollsCreate(params Params) (response PollsCreateResponse, err error) {
	err = vk.RequestUnmarshal("polls.create", params, &response)
	return
}

// PollsDeleteVote deletes the current user's vote from the selected answer in the poll.
//
// https://vk.com/dev/polls.deleteVote
func (vk *VK) PollsDeleteVote(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("polls.deleteVote", params, &response)
	return
}

// PollsEdit edits created polls.
//
// https://vk.com/dev/polls.edit
func (vk *VK) PollsEdit(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("polls.edit", params, &response)
	return
}

// PollsGetBackgroundsResponse struct.
type PollsGetBackgroundsResponse []object.PollsBackground

// PollsGetBackgrounds return default backgrounds for polls.
//
// https://vk.com/dev/polls.getBackgrounds
func (vk *VK) PollsGetBackgrounds(params Params) (response PollsGetBackgroundsResponse, err error) {
	err = vk.RequestUnmarshal("polls.getBackgrounds", params, &response)
	return
}

// PollsGetByIDResponse struct.
type PollsGetByIDResponse object.PollsPoll

// PollsGetByID returns detailed information about a poll by its ID.
//
// https://vk.com/dev/polls.getById
func (vk *VK) PollsGetByID(params Params) (response PollsGetByIDResponse, err error) {
	err = vk.RequestUnmarshal("polls.getById", params, &response)
	return
}

// PollsGetPhotoUploadServerResponse struct.
type PollsGetPhotoUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// PollsGetPhotoUploadServer returns a URL for uploading a photo to a poll.
//
// https://vk.com/dev/polls.getPhotoUploadServer
func (vk *VK) PollsGetPhotoUploadServer(params Params) (response PollsGetPhotoUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("polls.getPhotoUploadServer", params, &response)
	return
}

// PollsGetVotersResponse struct.
type PollsGetVotersResponse []object.PollsVoters

// PollsGetVoters returns a list of IDs of users who selected specific answers in the poll.
//
// https://vk.com/dev/polls.getVoters
func (vk *VK) PollsGetVoters(params Params) (response PollsGetVotersResponse, err error) {
	err = vk.RequestUnmarshal("polls.getVoters", params, &response)
	return
}

// PollsGetVotersFieldsResponse struct.
type PollsGetVotersFieldsResponse []object.PollsVotersFields

// PollsGetVotersFields returns a list of IDs of users who selected specific answers in the poll.
//
// https://vk.com/dev/polls.getVoters
func (vk *VK) PollsGetVotersFields(params Params) (response PollsGetVotersFieldsResponse, err error) {
	err = vk.RequestUnmarshal("polls.getVoters", params, &response)
	return
}

// PollsSavePhotoResponse struct.
type PollsSavePhotoResponse object.PollsPhoto

// PollsSavePhoto allows to save poll's uploaded photo.
//
// https://vk.com/dev/polls.savePhoto
func (vk *VK) PollsSavePhoto(params Params) (response PollsSavePhotoResponse, err error) {
	err = vk.RequestUnmarshal("polls.savePhoto", params, &response)
	return
}
