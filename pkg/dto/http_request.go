package dto

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/storeapp/pkg/constants"
	"github.com/storeapp/pkg/err"
	e "github.com/storeapp/pkg/err"
)

type HttpRequestDetails struct {
	QueryParams map[string]string `json:"query_params"`
	Headers     map[string]string `json:"headers"`
	Body        []byte            `json:"body"`
}
type HttpParamsStruct struct {
	HttpClient *http.Client
	Method     string
	Url        string
	Details    HttpRequestDetails
}

// func NewHttpRequest(h HttpParamsStruct) ([]byte, error) {
// 	req, err := http.NewRequest(h.Method, h.Url, bytes.NewBuffer(h.Details.Body))
// 	req.Header.Add("Accept-Encoding", "application/json")
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("X-Shopify-Access-Token", constants.XShopifyAccessToken)
// 	fmt.Print(req.URL)
// 	if h.Details.Headers != nil {
// 		for k, v := range h.Details.Headers {
// 			req.Header.Add(k, v)
// 		}
// 	}

// 	if h.Details.QueryParams != nil {
// 		q := req.URL.Query()
// 		for k, v := range h.Details.QueryParams {
// 			q.Add(k, v)
// 			q.Add(k, v)
// 			req.URL.RawQuery = q.Encode()
// 		}
// 	}

// 	resp, err := h.HttpClient.Do(req)

// 	if err != nil {
// 		fmt.Printf("http request failed with error : " + err.Error())
// 		return nil, err
// 	}

// 	defer func(Body io.ReadCloser) {
// 		err := Body.Close()
// 		if err != nil {
// 			return
// 		}
// 	}(resp.Body)

// 	data, err := ioutil.ReadAll(resp.Body)
// 	//logger.Info("Response Body: " + string(data))
// 	if err != nil {
// 		//logger.Error("reading response body failed with error: " + err.Error())
// 		return nil, err

// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		if resp.StatusCode != http.StatusCreated {
// 			//logger.Error("request failed with status: " + resp.Status + " having response body:" + string(data))
// 			return nil, err
// 		}
// 	}

// 	//logger.Info("ending new http request from util")
// 	return data, nil
// }
func NewHttpRequest(httpClient *http.Client, method, url string, details HttpRequestDetails) ([]byte, *err.AppError) {
	fmt.Print("calling new http request from util")
	fmt.Print("HTTP Request Body: " + string(details.Body))
	req, err := http.NewRequest(method, url, bytes.NewBuffer(details.Body))
	req.Header.Add("Accept-Encoding", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Shopify-Access-Token", constants.XShopifyAccessToken)
	if details.Headers != nil {
		for k, v := range details.Headers {
			req.Header.Add(k, v)
		}
	}

	if details.QueryParams != nil {
		q := req.URL.Query()
		for k, v := range details.QueryParams {
			q.Add(k, v)
			q.Add(k, v)
			req.URL.RawQuery = q.Encode()
		}
	}

	fmt.Print("Request url: " + req.URL.String())
	fmt.Print("Http Request: " + fmt.Sprintln(req))

	resp, err := httpClient.Do(req)
	fmt.Print("Raw Response: " + fmt.Sprintln(resp))
	fmt.Print("Response Body: " + fmt.Sprintln(resp.Body))
	if err != nil {
		fmt.Print("http request failed with error : " + err.Error())
		return nil, e.NewUnexpectedError("http request failed with error : " + err.Error())
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	fmt.Print("Response Body: " + string(data))
	if err != nil {
		fmt.Print("reading response body failed with error: " + err.Error())
		return nil, e.NewUnexpectedError("reading response body failed with error: " + err.Error())

	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode != http.StatusCreated {
			fmt.Print("request failed with status: " + resp.Status + " having response body:" + string(data))
			return nil, e.NewCustomErrorCodeFromThirdPartyAPI(resp.StatusCode, "request failed with status: "+resp.Status+" having response body:"+string(data))
		}
	}

	fmt.Print("ending new http request from util")
	return data, nil
}
