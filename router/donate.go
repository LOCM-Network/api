package router

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"sync"
	"time"

	"github.com/locm-team/api/util"
)

var callbackResponse *ResponseCard = nil

const (
	PENDING = "99"
)

type CardData struct {
	Telecom       string `json:"telecom"`
	Pin           string `json:"pin"`
	Serial        string `json:"serial"`
	Amount        string `json:"amount"`
	TransactionID string `json:"transaction_id"`
	DateTime      string `json:"date_time"`
	Command       string `json:"command"`
	Key           *Key   `json:"key"`
}

type ResponseCard struct {
	Status         string `json:"status"`
	Telco          string `json:"telco"`
	Code           string `json:"code"`
	Serial         string `json:"serial"`
	Declared_value string `json:"declared_value"`
	Value          string `json:"value"`
	Trans_id       string `json:"trans_id"`
	Request_id     string `json:"request_id"`
	Callback_sign  string `json:"callback_sign"`
}

type Key struct {
	Partner_id  string `json:"partner_id"`
	Partner_key string `json:"partner_key"`
}

func (c *CardData) PostCard(driver string, r *http.Request) ([]byte, bool) {
	url := "http://" + driver + "/chargingws/v2"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("telco", c.Telecom)
	writer.WriteField("code", c.Pin)
	writer.WriteField("serial", c.Serial)
	writer.WriteField("amount", c.Amount)
	writer.WriteField("partner_id", c.Key.Partner_id)
	writer.WriteField("request_id", c.TransactionID)
	writer.WriteField("command", c.Command)
	writer.WriteField("sign", string(c.MD5()))
	err := writer.Close()
	if err != nil {
		return nil, false
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, false
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return nil, false
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, false
	}

	return body, true
}

func (c *CardData) MD5() []byte {
	return util.MD5Byte([]byte(c.Key.Partner_key + c.Pin + c.Command + c.Key.Partner_id + c.TransactionID + c.Serial + c.Telecom))
}

func postCardHandler(w http.ResponseWriter, r *http.Request) {
	driver := r.FormValue("driver")
	card := &CardData{
		Telecom:  r.FormValue("telecom"),
		Pin:      r.FormValue("pin"),
		Serial:   r.FormValue("serial"),
		Amount:   r.FormValue("amount"),
		DateTime: r.FormValue("datetime"),
		Command:  r.FormValue("command"),
		Key: &Key{
			Partner_id:  r.FormValue("partner_id"),
			Partner_key: r.FormValue("partner_key"),
		},
		TransactionID: r.FormValue("transaction_id"),
	}
	response, ok := card.PostCard(driver, r)
	if !ok {
		json.NewEncoder(w).Encode(Response{Status: ResponseStatusInternalServerError, Message: ResponseInternalServerErrorMessage, Data: nil})
		return
	}
	respCard, ok := parseResponse(response)
	if !ok {
		json.NewEncoder(w).Encode(Response{Status: ResponseStatusInternalServerError, Message: ResponseInternalServerErrorMessage, Data: nil})
		return
	}
	if respCard.Status == PENDING {
		log.Printf("Pending: %s", respCard.Request_id)
		wait := &sync.WaitGroup{}
		wait.Add(1)
		go func() {
			waitCallback(5, wait)
		}()
		wait.Wait()
	}
	json.NewEncoder(w).Encode(respCard)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := parseRequest(r)
	if !ok {
		json.NewEncoder(w).Encode(Response{Status: ResponseStatusInternalServerError, Message: ResponseInternalServerErrorMessage, Data: nil})
		return
	}
	if req.Status == PENDING {
		log.Printf("Pending: %s", req.Request_id)
		return
	}
	callbackResponse = req
}

func parseRequest(r *http.Request) (*ResponseCard, bool) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, false
	}
	var card ResponseCard
	err = json.Unmarshal(body, &card)
	if err != nil {
		return nil, false
	}
	return &card, true
}

func parseResponse(response []byte) (*ResponseCard, bool) {
	var card ResponseCard
	err := json.Unmarshal(response, &card)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	return &card, true
}

func waitCallback(second int, wait *sync.WaitGroup) {
	time.Sleep(time.Second * time.Duration(second))
	if callbackResponse == nil {
		waitCallback(5, wait)
	}
	defer wait.Done()
}
