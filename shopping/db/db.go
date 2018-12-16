package db

import (
	"github.com/clvx/go-demos/shopping/models"
)

func LoadItem(id int) *models.Item {
	return &models.Item{
		Price: 9.001,	
	}
}
