package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/MahmudulTushar/graphql/graph/generated"
	"github.com/MahmudulTushar/graphql/graph/model"
	"github.com/MahmudulTushar/graphql/repository"
)

func (r *mutationResolver) CreateNewCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	course := &model.Course{
		ID:          strconv.Itoa(rand.Int()),
		Name:        input.Name,
		Description: input.Description,
		User:        &model.User{ID: input.UserID, Name: "User " + input.UserID},
	}
	courseRepository.Save(course)
	return course, nil
}

func (r *mutationResolver) UpdateCourse(ctx context.Context, id string, input *model.UpdateCourse) (string, error) {
	course := &model.Course{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
	}
	ret := courseRepository.UpdateById(id, course)
	return ret, nil
}

func (r *mutationResolver) DeleteCourse(ctx context.Context, id string) (string, error) {
	ret := courseRepository.Delete(id)
	return ret, nil
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return courseRepository.FindAll(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var courseRepository repository.CourseRepository = repository.NewDatabaseInstance()
