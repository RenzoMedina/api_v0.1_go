package model

import (
	"time"
)

/*
? Change of name and type
*/
type Product struct {
	ID        uint
	Title     string
	Body      string
	Create_At time.Time
	Update_At time.Time
}

type Products []*Product

type Storage interface {
	Migrate() error
	Create(*Product) error
	Update(*Product) error
	//GetAll()(Products, error)
	//GetId(uint)(*Product, error)
	Delete(uint) error
}

type Service struct {
	storage Storage
}

func NewServices(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

func (s *Service) Create(p *Product) error {
	return s.storage.Create(p)
}

func (s *Service) Update(p *Product) error {
	return s.storage.Update(p)
}

func (s *Service) Delete(id uint) error {
	return s.storage.Delete(id)
}
