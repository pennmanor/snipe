package snipe

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	defaultLimit = "100"
)

type Snipe struct {
	Client *http.Client
	Key    string
	Server string
}

type AssetQueryParams map[string]string

func NewDefaultClient(server string, key string) (*Snipe, error) {
	return &Snipe{Key: key, Server: server, Client: &http.Client{Transport: snipeAgentTransport{
		apikey: key,
		base:   http.DefaultTransport,
	}}}, nil
}

func makeQueryFromMap(params map[string]string) string {
	r := ""
	first := true

	for k, v := range params {

		v = makeQueryFromString(v)

		if !first {
			r += "&"
		} else {
			first = false
		}

		r += fmt.Sprintf("%+v=%+v", url.QueryEscape(k), v)
	}

	return r
}

func makeQueryFromString(text string) string {

	textWords := strings.Fields(text)

	for i, word := range textWords {
		textWords[i] = url.QueryEscape(word)
	}

	r := strings.Join(textWords[:], "+")
	return r
}

func (s *Snipe) GetAssets(params AssetQueryParams) (*Assets, error) {

	var r = new(Assets)

	qp := makeQueryFromMap(params)

	_, limitSet := params["limit"]
	if !limitSet {
		params["limit"] = defaultLimit
		qp = makeQueryFromMap(params)
	}

	_, offsetSet := params["offset"]
	if !offsetSet {
		params["offset"] = "0"
		qp = makeQueryFromMap(params)
	}

	url := fmt.Sprintf("%+v/api/v1/hardware?%+v", s.Server, qp)

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := s.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	// We need to iterate the api to fetch all the results
	if !limitSet && len(r.Rows) != r.Total && !offsetSet {

		for len(r.Rows) < r.Total {
			var r2 = new(Assets)

			os, _ := params["offset"]

			offset, err := strconv.Atoi(os)
			if err != nil {
				log.Fatal(err)
			}

			limits, _ := params["limit"]
			limit, err := strconv.Atoi(limits)
			if err != nil {
				log.Fatal(err)
			}

			offset += limit

			params["offset"] = fmt.Sprintf("%+v", offset)
			qp = makeQueryFromMap(params)

			url := fmt.Sprintf("%+v/api/v1/hardware?%+v", s.Server, qp)

			req, _ := http.NewRequest("GET", url, nil)
			resp, err := s.Client.Do(req)
			if err != nil {
				log.Fatal(err)
			}

			err = json.NewDecoder(resp.Body).Decode(r2)
			if err != nil {
				return nil, err
			}

			r.Rows = append(r.Rows, r2.Rows...)

		}

	}

	return r, nil

}

func (s *Snipe) GetAssetByTag(assetTag string) (*Asset, error) {
	var r = new(Asset)

	url := fmt.Sprintf("%+v/api/v1/hardware/bytag/%+v", s.Server, assetTag)

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := s.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Snipe) GetUserHistory(itemID int) (*Activities, error) {
	var r = new(Activities)

	url := fmt.Sprintf("%+v/api/v1/reports/activity?target_id=%+v&target_type=user", s.Server, itemID)

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := s.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Snipe) GetAssetHistory(itemID int) (*Activities, error) {
	var r = new(Activities)

	url := fmt.Sprintf("%+v/api/v1/reports/activity?item_id=%+v&item_type=asset", s.Server, itemID)

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := s.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Snipe) GetAllActivity(params AssetQueryParams) (*Activities, error) {
	qp := makeQueryFromMap(params)

	_, limitSet := params["limit"]
	if !limitSet {
		params["limit"] = defaultLimit
		qp = makeQueryFromMap(params)
	}

	_, offsetSet := params["offset"]
	if !offsetSet {
		params["offset"] = "0"
		qp = makeQueryFromMap(params)
	}

	var r = new(Activities)

	url := fmt.Sprintf("%+v/api/v1/reports/activity?%v", s.Server, qp)

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := s.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Snipe) GetUserCurrentAssets(userID int) (*Assets, error) {
	var r = new(Assets)

	url := fmt.Sprintf("%+v/api/v1/users/%v/assets", s.Server, userID)

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := s.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Snipe) CheckoutAsset(assetID int, userID int) (*Checkout, error) {
	var r = new(Checkout)

	url := fmt.Sprintf("%+v/api/v1/hardware/%+v/checkout?checkout_to_type=user&assigned_user=%+v", s.Server, assetID, userID)

	req, _ := http.NewRequest("POST", url, nil)
	resp, err := s.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Snipe) CheckinAsset(assetID int) (*Checkin, error) {
	var r = new(Checkin)

	url := fmt.Sprintf("%+v/api/v1/hardware/%+v/checkin", s.Server, assetID)

	req, _ := http.NewRequest("POST", url, nil)
	resp, err := s.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// read the json response
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (s *Snipe) GetUsers(params AssetQueryParams) (*Users, error) {

	var r = new(Users)

	qp := makeQueryFromMap(params)

	_, limitSet := params["limit"]
	if !limitSet {
		params["limit"] = defaultLimit
		qp = makeQueryFromMap(params)
	}

	_, offsetSet := params["offset"]
	if !offsetSet {
		params["offset"] = "0"
		qp = makeQueryFromMap(params)
	}

	url := fmt.Sprintf("%+v/api/v1/users?%+v", s.Server, qp)

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := s.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	// We need to iterate the api to fetch all the results
	if !limitSet && len(r.Rows) != r.Total && !offsetSet {

		for len(r.Rows) < r.Total {
			var r2 = new(Users)

			os, _ := params["offset"]

			offset, err := strconv.Atoi(os)
			if err != nil {
				log.Fatal(err)
			}

			limits, _ := params["limit"]
			limit, err := strconv.Atoi(limits)
			if err != nil {
				log.Fatal(err)
			}

			offset += limit

			params["offset"] = fmt.Sprintf("%+v", offset)
			qp = makeQueryFromMap(params)

			url := fmt.Sprintf("%+v/api/v1/users?%+v", s.Server, qp)

			req, _ := http.NewRequest("GET", url, nil)
			resp, err := s.Client.Do(req)
			if err != nil {
				log.Fatal(err)
			}

			err = json.NewDecoder(resp.Body).Decode(r2)
			if err != nil {
				return nil, err
			}

			r.Rows = append(r.Rows, r2.Rows...)

		}

	}

	return r, nil

}

func (s *Snipe) GetUser(userID int) (*User, error) {

	var r = new(User)

	url := fmt.Sprintf("%+v/api/v1/users/%v", s.Server, userID)

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := s.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Snipe) GetLocations() (*Locations, error) {

	var r = new(Locations)

	url := fmt.Sprintf("%+v/api/v1/locations", s.Server)

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := s.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Snipe) UpdateCustomField(assetID int, fieldName string, value string) (*Update, error) {

	var r = new(Update)

	// get the field name
	var a = new(Asset)

	url := fmt.Sprintf("%v/api/v1/hardware/%v", s.Server, assetID)

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := s.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(a)
	if err != nil {
		return nil, err
	}

	cfinternalmap := a.CustomFields.(map[string]interface{})
	fieldMap := cfinternalmap[fieldName].(map[string]interface{})
	fieldDBName, ok := fieldMap["field"].(string)
	if !ok { // ok is false if there is no value for field
		return nil, fmt.Errorf("Could not get the field name")
	}

	// create the json for the request
	jsonString := []byte(fmt.Sprintf(`{"%v":"%v"}`, fieldDBName, value))

	url = fmt.Sprintf("%v/api/v1/hardware/%v", s.Server, assetID)

	req, _ = http.NewRequest("PUT", url, bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")
	resp, err = s.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Snipe) PatchUser(userID int, fields map[string]interface{}) (*User, error) {

	r := UserPatch{}

	url := fmt.Sprintf("%+v/api/v1/users/%v", s.Server, userID)

	body, err := json.Marshal(fields)
	if err != nil {
		return nil, err
	}

	br := bytes.NewBuffer(body)

	req, err := http.NewRequest("PATCH", url, br)
	req.Header.Add("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	if r.Status != "success" {
		return nil, errors.New(r.Status)
	}

	return &r.User, nil
}

type snipeAgentTransport struct {
	apikey string
	base   http.RoundTripper
}

func (t snipeAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %+v", t.apikey))
	return t.base.RoundTrip(req)
}
