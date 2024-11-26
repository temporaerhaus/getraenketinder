package api

type Getraenk struct {
	Name         string   `json:"name"`
	Hash         string   `json:"hash,omitempty"`
	Presentation string   `json:"presentation,omitempty"`
	Images       []string `json:"images,omitempty"`
	Likes        int64    `json:"no_likes,omitempty"`
	Dislikes     int64    `json:"no_dislikes,omitempty"`
	Superlikes   int64    `json:"no_superlikes,omitempty"`
}

type GetraenkeList struct {
	GetraenkeList []Getraenk `json:"GetraenkeList"`
}

type Bicture struct {
	Url            string `json:"url"`
	Hash           string `json:"hash,omitempty"`
	Height         int64  `json:"height,omitempty"`
	Width          int64  `json:"width,omitempty"`
	BackgroundHash string `json:"weird_background_hash,omitempty"`
	MimeType       string `json:"mime_type,omitempty"`
}
