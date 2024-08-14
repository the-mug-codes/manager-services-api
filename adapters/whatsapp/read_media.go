package whatsapp

import (
	"encoding/json"
)

func (whatsapp *whatsapp) ReadMedia(mediaID string) (body []byte, mimeType string, err error) {
	var media MediaDownload
	body, err = whatsapp.apiRequest("GET", mediaID, nil)
	if err != nil {
		return body, mimeType, err
	}
	err = json.Unmarshal(body, &media)
	if err != nil {
		return body, media.MimeType, err
	}
	body, err = whatsapp.apiMediaDownloadRequest(media.URL, media.MimeType)
	if err != nil {
		return body, media.MimeType, err
	}
	return body, media.MimeType, err
}
