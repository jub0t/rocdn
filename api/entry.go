package rblx

import (
	"encoding/json"
	"fmt"
  "io"
	"io/ioutil"
	"net/http"
	"rblx/structs"
)

func GetHeadshot(user_id int, size int, format string, circle bool) (structs.RobloxImage, error) {
	if req, err := http.NewRequest("GET", fmt.Sprintf("https://thumbnails.roblox.com/v1/users/avatar-headshot?userIds=%v&size=%vx%v&format=%v&isCircular=%v", user_id, size, size, format, circle), nil); err != nil {
		return structs.RobloxImage{}, err
	} else {
		if res, err := http.DefaultClient.Do(req); err != nil {
			return structs.RobloxImage{}, err
		} else {
			if body, err := ioutil.ReadAll(res.Body); err != nil {
				return structs.RobloxImage{}, err
			} else {
				var i structs.RobloxResponse

				if err := json.Unmarshal(body, &i); err != nil {
					return structs.RobloxImage{}, err
				} else {
					if len(i.Data) > 0 {
						return i.Data[0], nil
					} else {
						return structs.RobloxImage{}, err
					}
				}
			}
		}
	}
}

func GetAvatar(user_id int, size int, format string, circle bool) (structs.RobloxImage, error) {
	if req, err := http.NewRequest("GET", fmt.Sprintf("https://thumbnails.roblox.com/v1/users/avatar?userIds=%v&size=%vx%v&format=%v&isCircular=%v", user_id, size, size, format, circle), nil); err != nil {
		return structs.RobloxImage{}, err
	} else {
		if res, err := http.DefaultClient.Do(req); err != nil {
			return structs.RobloxImage{}, err
		} else {
			if body, err := ioutil.ReadAll(res.Body); err != nil {
				return structs.RobloxImage{}, err
			} else {
				var i structs.RobloxResponse

				if err := json.Unmarshal(body, &i); err != nil {
					return structs.RobloxImage{}, err
				} else {
					if len(i.Data) > 0 {
						return i.Data[0], nil
					} else {
						return structs.RobloxImage{}, err
					}
				}
			}
		}
	}
}

func GetImageDataFromURL(url string) ([]byte, error) {
    // Make a GET request to the image URL
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Check if the response status code is OK
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to fetch image: %s", resp.Status)
    }

    // Read the response body into a byte slice
    imageData, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return imageData, nil
}
