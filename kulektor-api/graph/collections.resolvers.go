package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"
	"fmt"

	"jirku.sk/kulektor/auth"
	"jirku.sk/kulektor/db"
	"jirku.sk/kulektor/graph/model"
	"jirku.sk/kulektor/ksuid"
)

// CreateCollection is the resolver for the createCollection field.
func (r *mutationResolver) CreateCollection(ctx context.Context, input model.CreateCollectionInput) (*model.CreateCollectionResp, error) {
	user := auth.GetUser(ctx)
	collection := db.CreateCollectionParams{
		ID:          ksuid.NewKSUID(),
		UserID:      user.UserID,
		Description: input.Description,
		Title:       input.Title,
		Type:        input.Type,
	}

	output, err := r.Queries.CreateCollection(ctx, collection)
	if err != nil {
		return &model.CreateCollectionResp{
			Success: false,
		}, fmt.Errorf("creating collection: %s", err)
	}
	c := model.CollectionFromDb(output)
	return &model.CreateCollectionResp{
		Success: true,
		Data:    &c,
	}, nil
}

// MyCollectionsList is the resolver for the myCollectionsList field.
func (r *queryResolver) MyCollectionsList(ctx context.Context, input model.CollectionsListInput) (*model.CollectionsListResp, error) {
	user := auth.GetUser(ctx)
	params := db.GetUsersCollectionsListParams{
		UserID: user.UserID,
		Limit:  int32(input.Paging.GetLimit()),
		Offset: int32(input.Paging.GetOffset()),
	}

	items, err := r.Queries.GetUsersCollectionsList(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("requesting collection list with offset %d, limit %d", params.Offset, params.Limit)
	}
	total, err := r.Queries.GetUsersCollectionListTotal(ctx, user.UserID)
	if err != nil {
		return nil, fmt.Errorf("requesting total list with offset %d, limit %d", params.Offset, params.Limit)
	}
	collections := make([]model.Collection, 0, len(items))
	for _, i := range items {
		collections = append(collections, model.CollectionFromDb(i))
	}

	meta := model.Meta{
		Total:       int(total),
		PrevPage:    input.Paging.PrevPage(),
		CurrentPage: input.Paging.CurrentPage(),
		NextPage:    input.Paging.NextPage(total),
	}

	return &model.CollectionsListResp{
		Items: collections,
		Meta:  &meta,
	}, nil
}

// MyCollectionDetail is the resolver for the myCollectionDetail field.
func (r *queryResolver) MyCollectionDetail(ctx context.Context, collectionID ksuid.KSUID) (*model.Collection, error) {
	user := auth.GetUser(ctx)
	params := db.GetUserCollectionByIDParams{
		UserID: user.UserID,
		ID:     collectionID,
	}
	collection, err := r.Queries.GetUserCollectionByID(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("getting collection by ID %s for user %s: %s", params.ID, params.UserID, err)
	}
	resp := model.CollectionFromDb(collection)
	return &resp, nil
}

// CollectionsList is the resolver for the collectionsList field.
func (r *queryResolver) CollectionsList(ctx context.Context, input model.CollectionsListInput) (*model.CollectionsListResp, error) {
	panic(fmt.Errorf("not implemented: CollectionsList - collectionsList"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
