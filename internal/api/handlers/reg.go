package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reforce.pattern/internal/api/entity"
	"reforce.pattern/internal/api/response"
	"reforce.pattern/pkg/mongodb"
)

func Reg(ctx *gin.Context, mdb *mongodb.Collections) {
	body := entity.RegBody{Id: primitive.NewObjectID().Hex()}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.NotAcceptable(ctx, err)
		return
	}
	_, err = mdb.ContactPersons.InsertOne(context.TODO(), body)
	if err != nil {
		response.BadRequest(ctx, err)
		return
	}
	response.Success(ctx)
}
