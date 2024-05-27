package qrcode

import (
	"encoding/json"
	"github.com/ddvalim/go-qr-code-generator/cmd/response"
	"github.com/ddvalim/go-qr-code-generator/core/ports"
	"github.com/ddvalim/go-qr-code-generator/internal/qrcode"
	"io"
	"net/http"
)

type Handler struct {
	qrcodeService qrcode.Service
}

func NewHandler() Handler {
	qrcodeService := qrcode.NewService()

	return Handler{
		qrcodeService: qrcodeService,
	}
}

func (h Handler) Generate(w http.ResponseWriter, r *http.Request) {
	var content ports.Request

	b, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "could not read request body")
	}

	err = json.Unmarshal(b, &content)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "could not unmarshal request body")
	}

	qrCode, err := h.qrcodeService.Generate(content)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "could not generate qr code")
	}

	err = qrCode.Write(content.ImageSize, w)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "could not write into response writer")
	}

	response.Writer(w, http.StatusOK, "image/png", nil)
}
