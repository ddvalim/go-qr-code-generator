package qrcode

import (
	"github.com/ddvalim/go-qr-code-generator/core/ports"
	"github.com/skip2/go-qrcode"
)

type Service interface {
	Generate(content ports.Request) (*qrcode.QRCode, error)
}

type ServiceImpl struct {
}

func NewService() ServiceImpl {
	return ServiceImpl{}
}

func (s ServiceImpl) Generate(request ports.Request) (*qrcode.QRCode, error) {
	qrCode, err := qrcode.New(request.Content, qrcode.RecoveryLevel(request.Level))
	if err != nil {
		return nil, err
	}

	qrCode.VersionNumber = request.VersionNumber

	return qrCode, nil
}
