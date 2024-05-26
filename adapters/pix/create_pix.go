package pix

func (pix *pix) CreatePix(amountValue *float64, description *string, transactionID *string) (copyPaste string, qrCode string, err error) {
	copyPaste, err = pix.CreatePixCopyPasteCode(amountValue, description, transactionID)
	if err != nil {
		return copyPaste, qrCode, err
	}
	qrCode, err = pix.CreatePixQRCode(copyPaste)
	if err != nil {
		return copyPaste, qrCode, err
	}
	return copyPaste, qrCode, err
}
