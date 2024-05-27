package ports

type Request struct {
	Content       string `json:"content"`
	ImageSize     int    `json:"image_size"`
	VersionNumber int    `json:"version_number"`
	Level         int    `json:"level"`
}
