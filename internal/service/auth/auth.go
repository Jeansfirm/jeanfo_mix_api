package auth_service

import (
	"errors"
	"jeanfo_mix/util"
	session_util "jeanfo_mix/util/session"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

type ClientToken = string // jwt token string
type SessionID = string

type ClientData struct {
	UserID    int
	UserName  string
	SessionID SessionID

	dataLoaded  bool
	clientToken ClientToken
	sessionData *SessionData
}

type SessionData = session_util.SessionData

const (
	ClientTokenValidSeconds = 60 * 60 * 24 * 3
)

func (cd *ClientData) Load(token ClientToken) error {
	data, err := util.JwtParseToken(token)
	if err != nil {
		return err
	}

	err = mapstructure.Decode(data, cd)
	if err != nil {
		return err
	}
	cd.dataLoaded = true

	return nil
}

func (cd *ClientData) GetToken() (ClientToken, error) {
	if cd.clientToken != "" {
		return cd.clientToken, nil
	}
	if !cd.dataLoaded {
		return "", errors.New("clientData must gen token after data loaded")
	}

	data := structs.Map(cd)
	jwt, err := util.JwtGenerateToken(data, ClientTokenValidSeconds)
	if err != nil {
		return "", errors.New("clientData to jwt token fail: " + err.Error())
	}
	cd.clientToken = jwt

	return cd.clientToken, nil
}

func (cd *ClientData) GetSessionData() (*SessionData, error) {
	if cd.sessionData != nil {
		return cd.sessionData, nil
	}

	if !cd.dataLoaded {
		return nil, errors.New("clientData must load data before gen session data")
	}

	sessionData := &SessionData{UserID: cd.UserID, UserName: cd.UserName}
	sessionData.SessionID = cd.SessionID
	err := sessionData.Load()
	if err != nil {
		return nil, errors.New("get session data from client data fail: " + err.Error())
	}
	cd.sessionData = sessionData

	return sessionData, nil
}

func LoginUser(sessionData *SessionData) (ClientToken, error) {
	err := sessionData.Save()
	if err != nil {
		return "", err
	}

	clientData := ClientData{
		UserID: sessionData.UserID, UserName: sessionData.UserName,
		SessionID: sessionData.SessionID,
	}
	clientData.dataLoaded = true
	token, err := clientData.GetToken()

	return token, err
}

func LogoutUser(jwt ClientToken) error {
	var cdata ClientData
	err := cdata.Load(jwt)
	if err != nil {
		return err
	}

	sdata, err := cdata.GetSessionData()
	if err != nil {
		return err
	}
	err = sdata.Delete()

	return err
}
