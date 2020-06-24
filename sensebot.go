package geetest

import (
	"encoding/json"
	geetestlib "github.com/GeeTeam/gt3-golang-sdk/geetest"
	"time"
)

type SenseBot struct {
	geeTestId  string
	geeTestKey string
	geetestlib geetestlib.GeetestLib
}

type RegisterResult struct {
	geetestlib.FailbackRegisterRespnse
	Status int8 `json:"status"`
}

func NewSenseBot(geeTestId string, geeTestKey string) *SenseBot {
	senseBot := &SenseBot{
		geeTestId:  geeTestId,
		geeTestKey: geeTestKey,
	}

	senseBot.geetestlib = geetestlib.NewGeetestLib(senseBot.geeTestId, senseBot.geeTestKey, 5*time.Second)

	return senseBot
}

func (s *SenseBot) Register(userID, userIP string) (*RegisterResult, error) {
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

func (s *SenseBot) Validate(status int8, challenge, validate, secCode string, userID, userIP string) bool {
	var geetestRes bool
	if status == 1 {
		geetestRes = s.geetestlib.SuccessValidate(challenge, validate, secCode, userID, userIP)
	} else {
		geetestRes = s.geetestlib.FailbackValidate(challenge, validate, secCode)
	}

	return geetestRes
}
