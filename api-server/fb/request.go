package fb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const facebookApiEndpoint = "https://graph.facebook.com/v2.5"
const Album2ID = "615849938471305"
const Album3ID = "855587011164262"

type Photo struct {
	Images  []Image
	Picture string
}

type Image struct {
	Height int
	Source string
	Width  int
}

type PhotoRequest struct {
	Data   []*Photo
	Paging Pagination
}

type Pagination struct {
	Cursors  interface{}
	Previous *string
	Next     *string
}

func RequestFBAccessToken(secret string) (string, error) {
	resp, err := http.Get(facebookApiEndpoint + "/oauth/access_token?" +
		"client_id=162401547444127" +
		"&client_secret=" + secret +
		"&grant_type=client_credentials")

	if err != nil {
		return "", fmt.Errorf("Unable to access facebook oauth api\n%v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Can't read facebook oath api response body\n%v", err)
	}

	var responseMap interface{}
	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		return "", fmt.Errorf("Unable to marshal from JSON\n%v", err)
	}

	m := responseMap.(map[string]interface{})
	return m["access_token"].(string), nil
}

func RequestPhotosForAlbum(accessToken string, albumID string) ([]*Photo, error) {

	photosResponse, err := requestPhotos(facebookApiEndpoint + "/" + albumID + "/photos?" +
		"access_token=" + accessToken +
		"&fields=images,picture" +
		"&limit=1024")

	if err != nil {
		return nil, err
	}

	photos := photosResponse.Data
	for photosResponse.Paging.Next != nil {
		photosResponse, err = requestPhotos(*photosResponse.Paging.Next)

		if err != nil {
			return nil, err
		}

		photos = append(photos, photosResponse.Data...)
	}

	return photos, nil
}

func requestPhotos(url string) (*PhotoRequest, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to access the Facebook photos api\n%v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var f PhotoRequest
	err = json.Unmarshal(body, &f)

	if err != nil {
		return nil, fmt.Errorf("Unable to marshal from JSON\n%v", err)
	}

	return &f, nil
}
