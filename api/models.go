package api

type UploadGetraenk struct {
	Name         string   `form:"name"`
	UUID         string   `form:"hash"`
	Presentation string   `form:"presentation"`
	Images       []string `form:"images"`
	Likes        int64    `form:"no_likes"`
	Dislikes     int64    `form:"no_dislikes"`
	Superlikes   int64    `form:"no_superlikes"`
}

type Image struct {
	Url            string `json:"url"`
	Hash           string `json:"hash,omitempty"`
	Height         int64  `json:"height,omitempty"`
	Width          int64  `json:"width,omitempty"`
	BackgroundHash string `json:"weird_background_hash,omitempty"`
	MimeType       string `json:"mime_type,omitempty"`
}

type Getraenk struct {
	Name         string   `json:"name"`
	UUID         string   `json:"hash,omitempty"`
	Presentation string   `json:"presentation,omitempty"`
	Images       []Image  `json:"images,omitempty"`
	Likes        int64    `json:"no_likes,omitempty"`
	Dislikes     int64    `json:"no_dislikes,omitempty"`
	Superlikes   int64    `json:"no_superlikes,omitempty"`
}

type GetraenkeList struct {
	GetraenkeList []Getraenk `json:"GetraenkeList"`
}
