# Overview

This package is intended to provide a mechanism for interacting with the snipe api. The snipe api documentation is
available [here](https://snipe-it.readme.io/reference#api-overview) with examples
in php, python, javascript, etc.

## Getting Started

To use this package, the following code may prove helpful. 

    var {
        apiKey = "<your Snipe-IT api key>"
        snipeServer = "<your snipe server address>"
    }

    type snipeAgentTransport struct {
        apiKey string
        base   http.RoundTripper
    }

    // within your function:
    s := &snipe.Snipe{Key: snipeAPIKey, Server: snipeServer, Client: &http.Client{Transport: snipeAgentTransport{
        apiKey: snipeAPIKey,
        base:   http.DefaultTransport,
    }}}

    assets, err := s.GetAssets(snipe.AssetQueryParams{"search": query})
    if err != nil {
        log.Fatal(err)
    }

    // you should now have a variable "assets" that is a struct with Rows[] of individual snipe.Asset structs.



## Methods

The following methods are available.

* snipe.GetAssets		- get a list of assets
* snipe.GetAssetByTag		- get a single asset using an asset tag
* snipe.CheckoutAsset		- checkout an asset to a user
* snipe.CheckinAsset		- checkin (return) an asset
* snipe.GetAssetHistory		- get the checkin-checkout history of an asset
* snipe.GetUsers		- get a list of users
* snipe.GetUser			- get one and only one user
* snipe.GetUserHistory		- get the checkin-checkout history of a user
