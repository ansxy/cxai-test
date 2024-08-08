package utils

import (
	"context"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/ansxy/golang-boilerplate-gin/pkg/constant"
	custom_error "github.com/ansxy/golang-boilerplate-gin/pkg/error"
)

func GetUserIDFromCtx(ctx context.Context) (token string, err error) {
	userID := ctx.Value(constant.USERIDKEY).(string)

	if userID == "" {
		err := custom_error.SetCustomeError(&custom_error.ErrorContext{
			Code:    constant.DefaultUnauthorizedErrorCode,
			Message: constant.ErrorMessageMap[constant.DefaultUnauthorizedErrorCode],
		})

		return "", err
	}

	return userID, nil
}

func GetTokenFromHeader(r *http.Request) (token string, err error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		err := custom_error.SetCustomeError(&custom_error.ErrorContext{
			Code:    constant.DefaultUnauthorizedErrorCode,
			Message: constant.ErrorMessageMap[constant.DefaultUnauthorizedErrorCode],
		})

		return "", err
	}

	lenToken := 2

	splitToken := strings.Split(authHeader, " ")

	if len(splitToken) != lenToken {
		err := custom_error.SetCustomeError(&custom_error.ErrorContext{
			Code:    constant.DefaultUnauthorizedErrorCode,
			Message: constant.ErrorMessageMap[constant.DefaultUnauthorizedErrorCode],
		})

		return "", err
	}

	token = splitToken[1]
	return token, err
}

func CalculateOffset(page, perPage int) int {
	return (page - 1) * perPage
}

func getFrame(skipFrames int) runtime.Frame {
	targetFrameIndex := skipFrames + 2

	pc := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, pc)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(pc[:n])
		for more := true; more; {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if targetFrameIndex == 0 {
				frame = frameCandidate
				break
			}
			targetFrameIndex--
		}
	}

	return frame

}

func MyCaller() string {
	return getFrame(2).Function
}

func DateTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
