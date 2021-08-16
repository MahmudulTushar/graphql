package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/MahmudulTushar/graphql/repository"
	"math/rand"
	"strconv"

	"github.com/MahmudulTushar/graphql/graph/generated"
	"github.com/MahmudulTushar/graphql/graph/model"
)

var courseRepository repository.CourseRepository = repository.NewDatabaseInstance()

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

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return courseRepository.FindAll(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
