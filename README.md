# whiteboxrisque go client

This is an API that provides whitebox risk scoring

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation

To install the package, run the following command:

```bash
go get github.com/januusio/cryptowallet_risk_scoring_go
```

Then import the package in your Go code:

```golang
import "github.com/januusio/cryptowallet_risk_scoring_go"
```

## Getting started

```golang
package main

import (
	"context"
	"fmt"

	"github.com/januusio/cryptowallet_risk_scoring_go"
)

func main() {
	cfg := risk_scoring.NewConfiguration()
	apiClient := risk_scoring.NewAPIClient(cfg)

	address := "0xbb0ea877a85df253ccc312b80c644da31443abfd"

	riskReport, _, err := apiClient.DefaultApi.ScoreBtcAddressGet(context.Background(), address)
	if err != nil {
		fmt.Printf("Error retrieving risk report: %s", err.Error())
		return
	}

	fmt.Printf("Risk report: %+v", riskReport)
}
```

## Documentation for API Endpoints

All URIs are relative to *https://risk.charybdis.januus.io/*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**ScoreBtcAddressGet**](docs/DefaultApi.md#scorebtcaddressget) | **Get** /score/btc/{address} | 
*DefaultApi* | [**ScoreEthAddressGet**](docs/DefaultApi.md#scoreethaddressget) | **Get** /score/eth/{address} | 
*DefaultApi* | [**ScoreOtherAddressGet**](docs/DefaultApi.md#scoreotheraddressget) | **Get** /score/other/{address} | 

## Documentation For Models

 - [Offsets](docs/Offsets.md)
 - [Reason](docs/Reason.md)
 - [RiskElaboration](docs/RiskElaboration.md)
 - [RiskReport](docs/RiskReport.md)
 - [RiskScores](docs/RiskScores.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author


