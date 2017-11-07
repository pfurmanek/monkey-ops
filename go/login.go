package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Login(loginInput *LoginInput) *LoginOutput {

	var token string

	urlTarget := loginInput.Url + "/oauth/authorize?client_id=openshift-challenging-client&response_type=token"

	var redirectAttemptedError = errors.New("redirect")

	// Set up the HTTP request
	req, err := http.NewRequest("GET", urlTarget, nil)
	req.SetBasicAuth(loginInput.User, loginInput.Password)
	if err != nil {
		panic(err)
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	cli := &http.Client{
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return redirectAttemptedError
		}}

	resp, err := cli.Do(req)

	if urlError, ok := err.(*url.Error); ok && urlError.Err == redirectAttemptedError {
		location := resp.Header["Location"]

		token = StrExtract(location[0], "access_token=", "&expires_in")

		err = nil
	}

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Set up the HTTP request to get projects
	req, err = http.NewRequest("GET", loginInput.Url+"/oapi/v1/projects", nil)
	req.Header.Add("Authorization", "Bearer "+token)

	if err != nil {
		panic(err)
	}

	//reconfiguring the http client
	cli = &http.Client{
		Transport: transport}

	//getting project names from ose for an unique user
	resp, err = cli.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	projects, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	projectsCustom := map[string]interface{}{}
	json.Unmarshal(projects, &projectsCustom)

	items := projectsCustom["items"].([]interface{})
	loginOutput := &LoginOutput{token, make([]string, 0, 0)}

	for _, item := range items {
		itemObject := item.(map[string]interface{})
		metadataMap := itemObject["metadata"].(map[string]interface{})
		loginOutput.Projects = append(loginOutput.Projects, metadataMap["name"].(string))
	}

	return loginOutput
}
