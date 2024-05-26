package pix

import (
	"encoding/base64"
	"strings"

	"github.com/skip2/go-qrcode"
)

func (*pix) CreatePixQRCode(qrCodeContent string) (qrCode string, err error) {
	base64String := &strings.Builder{}
	bytes, err := qrcode.Encode(qrCodeContent, qrcode.Medium, 256)
	if err != nil {
		return qrCode, err
	}
	_, err = base64String.WriteString("data:image/png;base64,")
	if err != nil {
		return qrCode, err
	}
	_, err = base64String.WriteString(base64.StdEncoding.EncodeToString(bytes))
	if err != nil {
		return qrCode, err
	}
	qrCode = base64String.String()
	return qrCode, err
}
