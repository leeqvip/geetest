package geetest

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/forgoer/openssl"
	"io/ioutil"
	"net/http"
	"time"
)

const OneLoginURL = "https://onelogin.geetest.com/check_phone"

var iv = []byte("0000000000000000")

type CheckPhoneRequest struct {
	ProcessID     string `json:"process_id"`
	Sign          string `json:"sign"`
	Token         string `json:"token"`
	IsPhoneEncode bool   `json:"is_phone_encode"`
	Timestamp     string `json:"timestamp"`
	AuthCode      string `json:"authcode"`
}

type CheckPhoneResponse struct {
	Status      int    `json:"status"`
	ErrorMsg    string `json:"error_msg"`
	Result      string `json:"result"`
	Charge      bool   `json:"charge"`
	RiskLevel   int    `json:"risk_level"`
	FingerPrint string `json:"finger_print"`
	DeviceName  string `json:"device_name"`
}

type OneLogin interface {
	CheckPhone(processID, token string, authCode string) (*CheckPhoneResponse, error)
	DecodePhone(phone string) (string, error)
	GetSign(timestamp int64) string
}

type oneLogin struct {
	appID  string
	appKey string
}

func NewOneLogin(appID string, appKey string) OneLogin {
	return &oneLogin{
		appID:  appID,
		appKey: appKey,
	}
}

// CheckPhone 一键登录APP
func (o *oneLogin) CheckPhone(processID, token string, authCode string) (*CheckPhoneResponse, error) {
	timestamp := time.Now().UnixNano() / 1e6
	sign := o.GetSign(timestamp)
	r := CheckPhoneRequest{
		ProcessID:     processID,
		Sign:          sign,
		Token:         token,
		IsPhoneEncode: false,
		Timestamp:     fmt.Sprintf("%d", timestamp),
		AuthCode:      authCode,
	}

	b, err := json.Marshal(&r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", OneLoginURL, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response CheckPhoneResponse
	err = json.Unmarshal(body, &response)
	return &response, err
}

func (o *oneLogin) GetSign(timestamp int64) string {
	return HmacSha256([]byte(fmt.Sprintf("%s&&%d", o.appID, timestamp)), []byte(o.appKey))
}

func (o *oneLogin) DecodePhone(phone string) (string, error) {
	p, err := hex.DecodeString(phone)
	if err != nil {
		return "", err
	}

	key := o.appKey

	b, err := openssl.AesCBCDecrypt(p, []byte(key), iv, openssl.PKCS7_PADDING)
	if err != nil {
		return "", nil
	}

	return string(b), nil
}

func (o *oneLogin) EncodePhone(phone string) (string, error) {
	b, err := openssl.AesCBCEncrypt([]byte(phone), []byte(o.appKey), iv, openssl.PKCS7_PADDING)
	if err != nil {
		return "", nil
	}

	d := hex.EncodeToString(b)
	return d, nil
}
