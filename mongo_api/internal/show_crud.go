package internal

import (
	"context"
	"log"

	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/config"
	requestmodels "github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/request_models"
	responsemodels "github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/response_models"
)

//simply create crud for the show model
//get all shows
//get a single show
//create one show
//create a single show
// create a multiple show
//update a single show
//update multiple show
//delete single show
//delete multiple show

func insertSingleShow(ctx context.Context, showbd requestmodels.ShowReq)(responsemodels.ShowRes, error){
	select{
	case <- ctx.Done():
		return ctx.Err()
	default:
		//do nothing
	}

	res, err :=config.ShowClcn.InsertOne(ctx, showbd)
	if err != nil{
		log.Println("error occured while inserting the docs", err)
		return  nil, err
	}

	return responsemodels.ShowRes{
		
	} , nil
}