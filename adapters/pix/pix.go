package pix

type PixInterface interface {
	CreatePixCopyPasteCode(amountValue *float64, description *string, transactionID *string) (copyPaste string, err error)
	CreatePixQRCode(qrCodeContent string) (qrCode string, err error)
	CreatePix(amountValue *float64, description *string, transactionID *string) (copyPaste string, qrCode string, err error)
}

type pix struct {
	Name string
	Key  string
	City string
}

func Start(key string, name string, city string) PixInterface {
	return &pix{
		Key:  key,
		Name: name,
		City: city,
	}
}
