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

// CreateMyCollection is the resolver for the createMyCollection field.
func (r *mutationResolver) CreateMyCollection(ctx context.Context, input model.CollectionInput) (*model.CreateCollectionResp, error) {
	if input.Title == nil {
		return &model.CreateCollectionResp{}, fmt.Errorf("parsing collection fields: missing title")
	}
	user := auth.GetUser(ctx)
	collection := db.CreateCollectionParams{
		ID:          ksuid.NewKSUID(),
		UserID:      user.UserID,
		Description: input.Description,
		Title:       *input.Title,
		Type:        input.Type,
	}

	output, err := r.Queries.CreateCollection(ctx, collection)
	if err != nil {
		return &model.CreateCollectionResp{}, fmt.Errorf("creating collection: %s", err)
	}
	c := model.CollectionFromDb(output)
	return &model.CreateCollectionResp{
		Success: true,
		Data:    &c,
	}, nil
}

// UpdateMyCollection is the resolver for the updateMyCollection field.
func (r *mutationResolver) UpdateMyCollection(ctx context.Context, input model.UpdateCollectionInput) (*model.UpdateCollectionResp, error) {
	user := auth.GetUser(ctx)
	fields := make([]string, 0, len(input.FieldsToUpdate))
	for _, f := range input.FieldsToUpdate {
		fields = append(fields, string(f))
	}
	collection := db.Collection{
		ID:          input.ID,
		UserID:      user.UserID,
		Description: input.Collection.Description,
		Type:        input.Collection.Type,
	}
	if input.Collection.Title != nil {
		collection.Title = *input.Collection.Title
	}
	output, err := r.Queries.UpdateMyCollection(ctx, collection, fields)
	if err != nil {
		return &model.UpdateCollectionResp{}, fmt.Errorf("updating collection: %s", err)
	}
	c := model.CollectionFromDb(output)
	return &model.UpdateCollectionResp{
		Success: true,
		Data:    &c,
	}, nil
}

// DeleteMyCollection is the resolver for the deleteMyCollection field.
func (r *mutationResolver) DeleteMyCollection(ctx context.Context, collectionID ksuid.KSUID) (*model.DeleteMyCollectionResp, error) {
	user := auth.GetUser(ctx)
	err := r.Queries.DeleteUserCollectionByID(ctx, db.DeleteUserCollectionByIDParams{ID: collectionID, UserID: user.UserID})
	if err != nil {
		return &model.DeleteMyCollectionResp{
			Success: false,
		}, fmt.Errorf("deleting my collection: %s", err)
	}
	return &model.DeleteMyCollectionResp{
		Success: true,
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
		return nil, fmt.Errorf("requesting collection list with offset %d, limit %d: %s", params.Offset, params.Limit, err)
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
