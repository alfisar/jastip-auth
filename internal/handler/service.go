package handler

import (
	"fmt"
	"jastip/domain"
	"jastip/internal/errorhandler"
)

func PanicError() (err domain.ErrorData) {

	if r := recover(); r != nil {
		err = errorhandler.ErrInternal(errorhandler.ErrCodePanic, fmt.Errorf(fmt.Sprintf("%s", r)))
	}

	return
}
