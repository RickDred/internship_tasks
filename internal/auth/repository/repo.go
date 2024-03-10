package repository

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/RickDred/internship_tasks/tree/fifth_task/config"
	"github.com/RickDred/internship_tasks/tree/fifth_task/internal/auth"
	"github.com/RickDred/internship_tasks/tree/fifth_task/internal/models"
	"github.com/google/uuid"
)

// i decide to save data in json file, instead of connecting database like postgres, mongo or else.

type Repo struct {
	cfg *config.Config
}

func NewAuthService(cfg *config.Config) auth.Repo {
	return &Repo{
		cfg: cfg,
	}
}

func (r *Repo) Insert(user *models.User) error {
	file, err := os.OpenFile(r.cfg.JsonDB.Filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println(1, err)
		return err
	}
	defer file.Close()

	var users []*models.User
	if err := json.NewDecoder(file).Decode(&users); err != nil && err != io.EOF {
		log.Println(2, err)
		return err
	}

	user.CreatedAt = time.Now()
	users = append(users, user)

	data, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		log.Println(3, err)
		return err
	}

	if _, err := file.Seek(0, 0); err != nil {
		log.Println(4, err)
		return err
	}
	if err := file.Truncate(0); err != nil {
		log.Println(5, err)
		return err
	}
	if _, err := file.Write(data); err != nil {
		log.Println(6, err)
		return err
	}
	return nil
}

func (r *Repo) GetAll() ([]models.User, error) {
	var users []models.User

	data, err := ioutil.ReadFile(r.cfg.JsonDB.Filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repo) GetByEmail(email string) (*models.User, error) {
	users, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, errors.New("user not found by email")
}

func (r *Repo) GetByID(id uuid.UUID) (*models.User, error) {
	users, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, errors.New("user not found by ID")
}
