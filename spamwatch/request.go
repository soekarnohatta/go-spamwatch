// Package spamwatch provides a HTTP wrapper for the SpamWatch API.
package spamwatch

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"net/http"
	"time"
)

// ApiUrl contains default URL for SpamwatchAPI
const ApiUrl = "https://api.spamwat.ch"

// DefaultApiReq contains minimal configuration
// to make requests to the API.
var DefaultApiReq = BaseRequester{
	client: http.Client{
		Timeout: time.Millisecond * 1500,
	},
	apiUrl: ApiUrl,
}

// BaseRequester contains basic data to request to
// the API.
type BaseRequester struct {
	// client is the underlying HTTP client used to run the requests
	client http.Client

	// ApiUrl is the API endpoint
	apiUrl string

	// Token holds the user token
	token string
}

// Requester is an interface that implemements
// MakeRequest method.
type Requester interface{
	MakeRequest(l *zap.SugaredLogger, method string, param string, input interface{}) ([]byte, error)
}

// MakeRequest creates and reads a new HTTP request
// from or to the API.
func (b *BaseRequester) MakeRequest(l *zap.SugaredLogger, method string,
	param string, input interface{}) ([]byte, error) {
	token := b.token
	if token == "" {
		l.Debug("Token invalid")
		return nil, errors.Errorf("Token Is Invalid")
	}

	var req *http.Request
	var err error
	if input != nil {
		marshalStruct, _ := json.Marshal(input)
		newBytesBuffer := bytes.NewBuffer(marshalStruct)
		req, err = http.NewRequest(method, b.apiUrl+"/"+param, newBytesBuffer)
		l.Debugf(method + " request with body: %+v", input)
	} else {
		req, err = http.NewRequest(method, b.apiUrl+"/"+param, nil)
	}

	if err != nil {
		l.Debugw("failed to create" + method + "request",
			zapcore.Field{
				Key:    "param",
				Type:   zapcore.StringType,
				String: param,
			},
			zap.Error(err))
		return nil, errors.Wrapf(err, "client error executing" + method + "request to %v\n", param)
	}
	req.Header.Set("Authorization", "Bearer " + token)

	l.Debugf("executing " + method + ": %+v", req)
	resp, err := b.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	l.Debugf("successful " + method + " request: %+v", resp)

	switch resp.StatusCode {
	case http.StatusBadRequest:
		return nil, errors.New("API Error: 400 BadRequest, Please Check Your Configuration")
	case http.StatusForbidden:
		return nil, errors.New("API Error: 403 Forbidden, Please Check Your Permissions")
	case http.StatusUnauthorized:
		return nil, errors.New("API Error: 401 Unauthorized, Please Check Your Token")
	case http.StatusTooManyRequests:
		return nil, errors.New("API Error: 429 TooManyRequests, Please Retry In A Few Moments")
	case http.StatusNoContent:
		return nil, nil
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	l.Debugw("obtained " +method+ " result",
		zapcore.Field{
			Key:    "method",
			Type:   zapcore.StringType,
			String: method,
		},
		zapcore.Field{
			Key:    "result",
			Type:   zapcore.StringType,
			String: string(bodyBytes),
		},
	)

	return bodyBytes, nil
}