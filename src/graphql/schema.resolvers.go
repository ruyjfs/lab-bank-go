package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/ruyjfs/lab-bank-go/src/config"
	"github.com/ruyjfs/lab-bank-go/src/graphql/generated"
	"github.com/ruyjfs/lab-bank-go/src/graphql/model"
)

func (r *mutationResolver) CreateAccount(ctx context.Context, input model.NewAccount) (*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Accounts(ctx context.Context) ([]*model.Account, error) {
	db := config.Db()
	var accounts []*model.Account
	db.Find(&accounts)
	return accounts, nil
}

func (r *queryResolver) Account(ctx context.Context, id int) (*model.Account, error) {
	log.Println(id)
	db := config.Db()
	var account *model.Account
	db.First(&account, id)
	return account, nil
}

func (r *queryResolver) Transactions(ctx context.Context) ([]*model.Transaction, error) {
	db := config.Db()
	var transactions []*model.Transaction
	db.Joins("Account").Joins("OperationsType").Find(&transactions)
	return transactions, nil
}

func (r *queryResolver) OperationsTypes(ctx context.Context) ([]*model.OperationsType, error) {
	db := config.Db()
	var operationsTypes []*model.OperationsType
	db.Find(&operationsTypes)
	return operationsTypes, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
