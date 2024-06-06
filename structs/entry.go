package structs

type Storage struct {
	Data []Image
}

type RobloxResponse struct {
	Data []RobloxImage
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Image struct {
	Size      int    `json:"size"`
	TargetId  int    `json:"targetId"`
	ImageUrl  string `json:"imageUrl"`
	Timestamp int64  `json:"timestamp"`
}

type RobloxImage struct {
	TargetId int    `json:"targetId"`
	State    string `json:"state"`
	ImageUrl string `json:"imageUrl"`
}

type DatabaseInfo struct {
	Avatars   int `json:"avatars"`
	Headshots int `json:"headshots"`
}
