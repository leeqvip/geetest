package geetestbot

import (
	"encoding/json"
	"github.com/GeeTeam/gt3-golang-sdk/geetest"
)

// SenseBot
type SenseBot interface {
	Register(userID, userIP string) (*RegisterResult, error)
	Validate(status int8, challenge, validate, secCode string, userID, userIP string) bool
}

type RegisterResult struct {
	geetest.FailbackRegisterRespnse
	Status int8 `json:"status"`
}

func (s *api) Register(userID, userIP string) (*RegisterResult, error) {
	status, response := s.geetestlib.PreProcess(userID, userIP)
	var result RegisterResult
	var err error

	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	result.Status = status

	return &result, nil
}

func (s *api) Validate(status int8, challenge, validate, secCode string, userID, userIP string) bool {
	var geetestRes bool
	if status == 1 {
		geetestRes = s.geetestlib.SuccessValidate(challenge, validate, secCode, userID, userIP)
	} else {
		geetestRes = s.geetestlib.FailbackValidate(challenge, validate, secCode)
	}

	return geetestRes
}
