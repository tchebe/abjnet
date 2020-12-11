package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	pb "github.com/tchebe/abjnet/payment_service/proto/payment"
)

type service struct {
	repo repository
}

type jsonData struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Telephone string `json:"telephone"`
	Expeditor string `json:"expeditor"`
	TypeEnvoi string `json:"typeEnvoi"`
	Sms       string `json:"sms"`
}

func newPaymentService(repo repository) *service {
	return &service{repo}
}

func (s *service) Pay(ctx context.Context, req *pb.Payment, res *pb.Response) error {
	log.Println(req)
	resp, err := s.repo.MakePayment(req)
	if err != nil {
		theerror := fmt.Sprintf("%v --from payment_service", err)
		res.Done = false
		res.Errors = []*pb.Error{{
			Code:        400,
			Description: "paiement echouée.Veuillez contacter l'administrateur système",
		}}
		return errors.New(theerror)
	}

	res.Done = true
	res.Description = "Paiement pris en compte.Un retour vous sera fait d'ici 24h"
	res.Payment = resp

	//envoi de sms pour confirmation au client
	body := &jsonData{
		Username:  "WEBLOGY",
		Password:  "WEBLOGY",
		Telephone: req.Telephone,
		Expeditor: "Nsia Vie CI",
		TypeEnvoi: "Confirmation de souscription",
		Sms:       fmt.Sprintf("Cher(e) %s, le paiement de %d FCFA  a été prise en compte, Merci d’avoir utilisé le service de Weblogy Info : 22419800 ", req.Nomclient, req.Montant),
	}

	url := "http://10.11.100.48:8084/sendSMS"
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)
	request, _ := http.NewRequest("POST", url, payloadBuf)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, er := client.Do(request)

	if er != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	return nil
}
