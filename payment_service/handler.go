package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	pb "github.com/tchebe/abjnet/payment_service/proto/payment"
)

type service struct {
	repo repository
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
	
	//envoi de sms pour confirmation au client
	jsonData := map[string]string{
		"username": "WEBLOGY",
		"password": "WEBLOGY", 
		"telephone": req.telephone,
		"expeditor": "Nsia Vie CI",
		"typeEnvoi": "Confirmation de paiement",
		"sms": fmt.Sprintf("Cher(e) %s, votre paiement a bien été enregistré. Infoline 22419800",req.nomclient)
	}
    jsonValue, _ := json.Marshal(jsonData)
    response, err = http.Post("http://10.11.100.48:8084/sendSMS", "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }


	res.Done = true
	res.Description = "Paiement pris en compte.Un retour vous sera fait d'ici 24h"
	res.Payment = resp
	return nil
}
