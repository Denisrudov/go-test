package api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func NewService() *service {
	return &service{}
}

type service struct {
}

func (s *service) GetName() (Name, error) {
	data, err := makeCall(RANDOM_NAME)
	if err != nil {
		return Name{}, err
	}
	var NameData Name
	err = json.Unmarshal(data, &NameData)
	if err != nil {
		return Name{}, err
	}

	return NameData, nil
}

func (s *service) GetNorris() (ChuckNorris, error) {
	data, err := makeCall(CHUCKNORRIS_JOKE)
	if err != nil {
		return ChuckNorris{}, err
	}

	var ChukJoke ChuckNorris
	err = json.Unmarshal(data, &ChukJoke)
	if err != nil {
		return ChuckNorris{}, err
	}
	return ChukJoke, err
}

func makeCall(url string) (responseBytes []byte, err error) {
	client := http.DefaultClient
	resp, err := client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	responseBytes, err = ioutil.ReadAll(resp.Body)
	return
}
