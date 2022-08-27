package worker

import (
	"gorm.io/gorm"
)

type Worker interface {
	Run(data string, db gorm.DB)
	//AvroRun(data string, db gorm.DB)
}

func Run(data string, db gorm.DB) {

}
