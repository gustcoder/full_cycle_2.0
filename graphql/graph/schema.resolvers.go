package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/gustcoder/fcl2-graphql/graph/generated"
	"github.com/gustcoder/fcl2-graphql/graph/model"
)

func generateID() string {
	return fmt.Sprintf("T%d", rand.Int())
}

func findCategoryByID(r *mutationResolver, categoryId string) *model.Category {
	var category *model.Category

	for _, cat := range r.Categories {
		if cat.ID == categoryId {
			category = cat
		}
	}

	return category
}

func findCourseByID(r *mutationResolver, courseId string) *model.Course {
	var myCourse *model.Course

	for _, course := range r.Courses {
		if course.ID == courseId {
			myCourse = course
		}
	}

	return myCourse
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	// panic(fmt.Errorf("not implemented"))
	category := model.Category{
		ID:          generateID(),
		Name:        input.Name,
		Description: input.Description,
	}
	r.Categories = append(r.Categories, &category)

	return &category, nil
}

func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	// panic(fmt.Errorf("not implemented"))

	category := findCategoryByID(r, input.CategoryID)

	course := model.Course{
		ID:          generateID(),
		Name:        input.Name,
		Description: input.Description,
		Category:    category,
	}
	r.Courses = append(r.Courses, &course)

	return &course, nil
}

func (r *mutationResolver) CreateChapter(ctx context.Context, input model.NewChapter) (*model.Chapter, error) {
	// panic(fmt.Errorf("not implemented"))

	course := findCourseByID(r, input.CourseID)

	chapter := model.Chapter{
		ID:     generateID(),
		Name:   input.Name,
		Course: course,
	}
	r.Chapters = append(r.Chapters, &chapter)

	return &chapter, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.Resolver.Categories, nil
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return r.Resolver.Courses, nil
}

func (r *queryResolver) Chapters(ctx context.Context) ([]*model.Chapter, error) {
	return r.Resolver.Chapters, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
