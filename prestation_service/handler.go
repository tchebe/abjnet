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

type jsonData struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Telephone string `json:"telephone"`
	Expeditor string `json:"expeditor"`
	TypeEnvoi string `json:"typeEnvoi"`
	Sms       string `json:"sms"`
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

	res.Done = true
	res.Description = "Prestation prise en compte.Un retour vous sera fait d'ici 24h"
	res.Prestation = resp

	//envoi de sms pour confirmation au client
	body := &jsonData{
		Username:  "WEBLOGY",
		Password:  "WEBLOGY",
		Telephone: req.Telephone,
		Expeditor: "Nsia Vie CI",
		TypeEnvoi: "Confirmation de prestation",
		Sms:       fmt.Sprintf("Cher(e) %s, Votre demande de rachat partiel de %s est en cours de traitement, Votre carte cristal sera crédité dans 24H00, Info : 22419800", req.Nomclient, req.Montantdemande),
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
