package mapgetservice

import (
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/mapmodel"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model/mapdb"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

func GetMaps() (*[]mapmodel.Map, db.DatabaseError) {
	return mapdb.GetMaps()
}
