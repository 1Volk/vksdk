package api 

import (
	"github.com/1Volk/vksdk/object"
)

// LeadsCheckUserResponse struct.
type LeadsCheckUserResponse object.LeadsChecked

// LeadsCheckUser checks if the user can start the lead.
//
// https://vk.com/dev/leads.checkUser
func (vk *VK) LeadsCheckUser(params Params) (response LeadsCheckUserResponse, err error) {
	err = vk.RequestUnmarshal("leads.checkUser", params, &response)
	return
}

// LeadsCompleteResponse struct.
type LeadsCompleteResponse object.LeadsComplete

// LeadsComplete completes the lead started by user.
//
// https://vk.com/dev/leads.complete
func (vk *VK) LeadsComplete(params Params) (response LeadsCompleteResponse, err error) {
	err = vk.RequestUnmarshal("leads.complete", params, &response)
	return
}

// LeadsGetStatsResponse struct.
type LeadsGetStatsResponse object.LeadsLead

// LeadsGetStats returns lead stats data.
//
// https://vk.com/dev/leads.getStats
func (vk *VK) LeadsGetStats(params Params) (response LeadsGetStatsResponse, err error) {
	err = vk.RequestUnmarshal("leads.getStats", params, &response)
	return
}

// LeadsGetUsersResponse struct.
type LeadsGetUsersResponse object.LeadsEntry

// LeadsGetUsers returns a list of last user actions for the offer.
//
// https://vk.com/dev/leads.getUsers
func (vk *VK) LeadsGetUsers(params Params) (response LeadsGetUsersResponse, err error) {
	err = vk.RequestUnmarshal("leads.getUsers", params, &response)
	return
}

// LeadsMetricHitResponse struct.
type LeadsMetricHitResponse struct {
	Result       object.BaseBoolInt `json:"result"`        // Information whether request has been processed successfully
	RedirectLink string             `json:"redirect_link"` // Redirect link
}

// LeadsMetricHit counts the metric event.
//
// https://vk.com/dev/leads.metricHit
func (vk *VK) LeadsMetricHit(params Params) (response LeadsMetricHitResponse, err error) {
	err = vk.RequestUnmarshal("leads.metricHit", params, &response)
	return
}

// LeadsStartResponse struct.
type LeadsStartResponse object.LeadsStart

// LeadsStart creates new session for the user passing the offer.
//
// https://vk.com/dev/leads.start
func (vk *VK) LeadsStart(params Params) (response LeadsStartResponse, err error) {
	err = vk.RequestUnmarshal("leads.start", params, &response)
	return
}
