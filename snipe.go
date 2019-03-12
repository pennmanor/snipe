package snipe

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const (
	defaultLimit = "2000"
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

		if !first {
			r += "&"
		} else {
			first = false
		}

		r += fmt.Sprintf("%+v=%+v", url.QueryEscape(k), url.QueryEscape(v))
	}

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

func (s *Snipe) CheckoutAsset(assetID string, userID string) (*Checkout, error) {
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

func (s *Snipe) CheckinAsset(assetID string) (*Checkin, error) {
	var r = new(Checkin)

	url := fmt.Sprintf("%+v/api/v1/hardware/%+v/checkin", s.Server, assetID)

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

func (s *Snipe) GetUserByUsername(username string) (*User, error) {
	var r = new(Users)

	url := fmt.Sprintf("%+v/api/v1/users?search=%+v", s.Server, username)

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := s.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return &r.Rows[0], nil
}

type snipeAgentTransport struct {
	apikey string
	base   http.RoundTripper
}

func (t snipeAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %+v", t.apikey))
	return t.base.RoundTrip(req)
}
