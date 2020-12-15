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

	pb "github.com/tchebe/abjnet/souscription_service/proto/souscription"
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

func newSouscriptionService(repo repository) *service {
	return &service{repo}
}

//Subscribe -returns the souscription inserted in the DB
func (s *service) Subscribe(ctx context.Context, req *pb.Souscription, res *pb.Response) error {
	log.Println("The request is : %v \n", req)

	resp, err := s.repo.Subscribe(req)
	if err != nil {
		theerror := fmt.Sprintf("%v --from souscription_service", err)
		res.Done = false
		res.Errors = []*pb.Error{{
			Code:        400,
			Description: "souscription echouée.Veuillez contacter l'administrateur système",
		}}
		return errors.New(theerror)
	}

	res.Done = true
	res.Description = "Souscription prise en compte.Un retour vous sera fait d'ici 24h"
	res.Souscription = resp

	//envoi de sms pour confirmation au client

	body := &jsonData{
		Username:  "WEBLOGY",
		Password:  "WEBLOGY",
		Telephone: req.Telephone,
		Expeditor: "Nsia Vie CI",
		TypeEnvoi: "Confirmation de souscription",
		Sms:       fmt.Sprintf("Cher(e) %s, votre souscription a été enregistrée avec succès! Vous pouvez effectuer vos cotisations. Info : 22419800", req.Nom),
	}
	url := "http://154.73.101.139:8080/sendSMS"
	// url := "http://10.11.100.48:8084/sendSMS"
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

//GetAll -returns a slice of souscriptions
func (s *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	subs, err := s.repo.GetAll("")
	if err != nil {
		theerror := fmt.Sprintf("%v --from souscription_service", err)
		return errors.New(theerror)
	}
	res.Souscriptions = subs
	return nil
}

//DeleteAll -deletes all subs from the database
func (s *service) DeleteAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	done, err := s.repo.DeleteAll("")
	if err != nil {
		theerror := fmt.Sprintf("%v --from souscription_service", err)
		return errors.New(theerror)
	}
	res.Done = done
	return nil
}
