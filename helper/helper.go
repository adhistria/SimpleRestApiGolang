package helper

import (
	"errors"
	"net/http"
	"omahihrom/respond"
)

func CheckMimeType(mimeType string) error {
	switch mimeType {
	case "image/jpeg":
		return nil
	case "image/png":
		return nil
	default:
		return errors.New("The format file is not valid.")
	}
}

func ConvertToString(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

func CheckError(w http.ResponseWriter, err error) {
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
}
