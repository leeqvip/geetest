package geetestbot

import (
	"errors"
	"fmt"
	"github.com/GeeTeam/gt3-golang-sdk/geetest"
	"time"
)

type Api interface {
	OneLogin
	SenseBot
}

func NewApi(appID string, appKey string) Api {
	gt := &api{
		appID:  appID,
		appKey: appKey,
	}

	gt.geetestlib = geetest.NewGeetestLib(gt.appID, gt.appKey, 5*time.Second)

	return gt
}

func NewApiFromConfig(c *ConfigApi) (Api, error) {
	if c == nil {
		return nil, errors.New("config can not be nil")
	}
	gt := &api{
		appID:  c.AppId,
		appKey: c.AppKey,
	}

	timeout := 5 * time.Second
	if c.Timeout != 0 {
		timeout = timeout * time.Second
	}

	gt.geetestlib = geetest.NewGeetestLib(gt.appID, gt.appKey, timeout)

	return gt, nil
}

type api struct {
	appID  string
	appKey string

	geetestlib geetest.GeetestLib
}

func (o *api) GetSign(timestamp int64) string {
	return HmacSha256([]byte(fmt.Sprintf("%s&&%d", o.appID, timestamp)), []byte(o.appKey))
}
