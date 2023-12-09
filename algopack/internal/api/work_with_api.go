package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"algopack/internal/model"
)

func GetPredictByTicket(ctx context.Context, latestTicketData []interface{}) (*model.TicketPredict, error) {
	ticketData := model.TicketData{
		Ticket: latestTicketData,
	}

	jsonData, err := json.Marshal(ticketData)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(jsonData)
	apiUrl := fmt.Sprintf("http://api:5000/api-data")
	res, err := http.Post(apiUrl, "application/json", buffer)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(res.Body)
	bodyData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	ticketPredict := new(model.TicketPredict)
	err = json.Unmarshal(bodyData, ticketPredict)
	if err != nil {
		return nil, err
	}

	return ticketPredict, nil
}
