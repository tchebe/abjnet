package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	pb "github.com/zjjt/abjnet/souscription_service/proto/souscription"
)

type repository interface {
	Subscribe(sub *pb.Souscription) error
	GetAll() ([]*pb.Souscription, error)
	DeleteAll() (bool, error)
}

type SubRepository struct {
	db *gorm.DB
}

func newSubRepository(db *gorm.DB) *SubRepository {
	return &SubRepository{db}
}

//GetSub gets a subscription
func (repo *SubRepository) GetSub(sub *pb.Souscription) (*pb.Souscription, error) {
	sup := new(pb.Souscription)
	if err := repo.db.First(&sup, sub).Error; err != nil {
		return nil, err
	}
	return sup, nil
}

//Subscribe creates a subscription in the DB
func (repo *SubRepository) Subscribe(sub *pb.Souscription) error {
	subTime := time.Now().Format("02-01-2006 15:04:05")
	sub.CreatedAt = subTime
	//check if the subscription doesnt exist already
	subexist := new(pb.Souscription)
	repo.db.First(subexist, sub)
	if subexist != nil {
		return errors.New("Cette souscription existe déjà")
	}

	if err := repo.db.Create(sub).Error; err != nil {
		return err
	}
	return nil
}

//GetAll gets all the subscription in db based on the etattraitement if it is set
func (repo *SubRepository) GetAll(etat string) ([]*pb.Souscription, error) {
	var subs []*pb.Souscription
	if etat != "CREE" || etat != "TRAITEMENT" || etat != "TRAITEE" {
		if err := repo.db.Find(&subs).Error; err != nil {
			return nil, err
		}
	} else {
		sub := &pb.Souscription{Etattraitement: etat}
		if err := repo.db.Find(&subs, sub).Error; err != nil {
			return nil, err
		}
	}

	return subs, nil

}

//DeleteAll deletes all the subscriptions in db based on the etattraitement=TRAITEE if it is set or just removes everything
func (repo *SubRepository) DeleteAll(etat string) (bool, error) {
	if etat != "TRAITEE" {
		if err := repo.db.Exec("TRUNCATE TABLE souscriptions RESTART IDENTITY;").Error; err != nil {
			return false, err
		}
	} else {
		if err := repo.db.Where("etattraitement LIKE ?", fmt.Sprintf("%%v%", etat)).Delete(pb.Souscription{}).Error; err != nil {
			return false, err
		}
	}

	return true, nil

}
