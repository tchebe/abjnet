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

	pb "github.com/tchebe/abjnet/prestation_service/proto/prestation"
)

type service struct {
	repo repository
}

func newPrestationService(repo repository) *service {
	return &service{repo}
}

func (s *service) Rachat(ctx context.Context, req *pb.Prestation, res *pb.Response) error {
	resp, err := s.repo.MakeRachat(req)
	if err != nil {
		theerror := fmt.Sprintf("%v --from prestation_service", err)
		res.Done = false
		res.Errors = []*pb.Error{{
			Code:        400,
			Description: "prestation echouée.Veuillez contacter l'administrateur système",
		}}
		return errors.New(theerror)
	}

	//envoi de sms pour confirmation au client
	jsonData := map[string]string{
		"username":  "WEBLOGY",
		"password":  "WEBLOGY",
		"telephone": req.Telephone,
		"expeditor": "Nsia Vie CI",
		"typeEnvoi": "Confirmation de souscription",
		"sms":       fmt.Sprintf("Cher(e) %s, votre souscription a bien été enregistré. Infoline 22419800", req.Nomclient),
	}
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("http://10.11.100.48:8084/sendSMS", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	res.Done = true
	res.Description = "Prestation prise en compte.Un retour vous sera fait d'ici 24h"
	res.Prestation = resp
	return nil
}
func (s *service) ValeurRachat(ctx context.Context, req *pb.Request, res *pb.Response) error {
	log.Println("police for GetVR in handler is ", req)

	vr, err := s.repo.GetVR(&pb.Prestation{Policeno: req.Police})
	if err != nil {
		res.Done = false
		res.Maximumrachetable = "nothing"
		return err
	}
	res.Done = true
	res.Maximumrachetable = vr
	return nil
}
