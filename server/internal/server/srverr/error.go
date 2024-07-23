package srverr

import (
	"net/http"

	"github.com/haguirrear/coffeeassistant/server/internal/server/srvwrite"
	"github.com/haguirrear/coffeeassistant/server/pkg/logger"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h HTTPError) Error() string {
	return h.Message
}

func HandleErr(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	if err, ok := err.(HTTPError); ok {
		if err := srvwrite.JSON(w, err.Code, err); err != nil {
			logger.Logger.Error().Err(err).Msgf("defaultErrorHandler()")
		}

		return
	}

	errMsg := err.Error()
	if errMsg == "" {
		return
	}

	if err := srvwrite.JSON(w, http.StatusInternalServerError, HTTPError{Code: http.StatusInternalServerError, Message: errMsg}); err != nil {
		logger.Logger.Error().Err(err).Msgf("defaultErrorHandler()")
	}
}
