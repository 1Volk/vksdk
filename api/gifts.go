package api 

import "github.com/1Volk/vksdk/object"

// GiftsGetResponse struct.
type GiftsGetResponse struct {
	Count int                `json:"count"`
	Items []object.GiftsGift `json:"items"`
}

// GiftsGet returns a list of user gifts.
//
// https://vk.com/dev/gifts.get
func (vk *VK) GiftsGet(params Params) (response GiftsGetResponse, err error) {
	err = vk.RequestUnmarshal("gifts.get", params, &response)
	return
}

// GiftsGetCatalogResponse struct.
type GiftsGetCatalogResponse []struct {
	Name  string             `json:"name"`
	Title string             `json:"title"`
	Items []object.GiftsGift `json:"items"`
}

// GiftsGetCatalog returns catalog.
//
// https://vk.com/dev/gifts.get
func (vk *VK) GiftsGetCatalog(params Params) (response GiftsGetCatalogResponse, err error) {
	err = vk.RequestUnmarshal("gifts.getCatalog", params, &response)
	return
}
