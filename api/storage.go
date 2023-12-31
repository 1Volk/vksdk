package api

import (
	"github.com/1Volk/vksdk/object"
)

// StorageGetResponse struct.
type StorageGetResponse []object.BaseRequestParam

// ToMap return map from StorageGetResponse.
func (s StorageGetResponse) ToMap() map[string]string {
	m := make(map[string]string)
	for _, item := range s {
		m[item.Key] = item.Value
	}

	return m
}

// StorageGet returns a value of variable with the name set by key parameter.
//
// StorageGet always return array!
//
// https://vk.com/dev/storage.get
func (vk *VK) StorageGet(params Params) (response StorageGetResponse, err error) {
	// TODO: remove in 5.110
	if _, prs := params["keys"]; !prs {
		params["keys"] = params["key"]
		params["key"] = ""
	}

	err = vk.RequestUnmarshal("storage.get", params, &response)

	return
}

// StorageGetKeysResponse struct.
type StorageGetKeysResponse []string

// StorageGetKeys returns the names of all variables.
//
// https://vk.com/dev/storage.getKeys
func (vk *VK) StorageGetKeys(params Params) (response StorageGetKeysResponse, err error) {
	err = vk.RequestUnmarshal("storage.getKeys", params, &response)
	return
}

// StorageSet saves a value of variable with the name set by key parameter.
//
// https://vk.com/dev/storage.set
func (vk *VK) StorageSet(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("storage.set", params, &response)
	return
}
