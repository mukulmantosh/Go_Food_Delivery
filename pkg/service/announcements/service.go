package announcements

import "Go_Food_Delivery/pkg/database"

type AnnouncementService struct {
	db  database.Database
	env string
}

func NewAnnouncementService(db database.Database, env string) *AnnouncementService {
	return &AnnouncementService{db, env}
}
