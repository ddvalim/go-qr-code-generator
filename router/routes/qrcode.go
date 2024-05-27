package routes

import (
	"github.com/ddvalim/go-qr-code-generator/cmd/qrcode"
	"github.com/ddvalim/go-qr-code-generator/core/ports"
	"net/http"
)

var QRCode = []ports.Route{
	{
		URI:     "/qrcode",
		Method:  http.MethodGet,
		Handler: qrcode.NewHandler().Generate,
	},
}
