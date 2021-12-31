package main

import (
	"context"

	"github.com/GavukaAlexandr/ent-sandbox/db"
	"go.uber.org/zap"
)

func main() {
	zap.L().Info("STARTED CLOUD AUDIT APP")
	err := db.OpenConnection()
	if err != nil {
		zap.S().Fatal("connection error: ", err)
	}
	defer db.Client.Close()
	ctx := context.TODO()
	db.AutoMigration(&ctx)

	oliver, err := db.Client.User.
		Create().
		SetName("Oliver").
		SetAge(35).
		Save(ctx)
	if err != nil {
		zap.S().Fatal("create user err: ", err)
	}
	zap.S().Debug(oliver)

	max, err := db.Client.Pet.
		Create().
		SetName("Max").
		SetAge(1).
		SetOwner(oliver).
		Save(ctx)
	if err != nil {
		zap.S().Fatal("create pet err: ", err)
	}
	zap.S().Debug(max)

	oliverID, err := db.Client.User.
		Create().
		SetName("Oliver").
		SetAge(30). // for example update age
		OnConflict().
		DoNothing(). // or UpdateNewValues()
		ID(ctx)
	if err != nil {
		zap.S().Warn("upsert user err: ", err)
	} else {
		zap.S().Info(oliverID)
	}

}
