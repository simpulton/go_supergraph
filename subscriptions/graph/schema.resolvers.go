package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"fmt"
	"go_supergraph/graph/model"
)

// Subscription is the resolver for the subscription field.
func (r *accountResolver) Subscription(ctx context.Context, obj *model.Account) (*model.Plan, error) {
	return &model.Plan{
		ID: fmt.Sprintf("plan-%s", obj.ID),
		Type: model.PlanTypeFree,
	}, nil
}

// Account returns AccountResolver implementation.
func (r *Resolver) Account() AccountResolver { return &accountResolver{r} }

type accountResolver struct{ *Resolver }
