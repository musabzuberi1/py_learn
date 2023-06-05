package service

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminService struct {
	repo *repository.AdminRepository
}

func (as *AdminService) GetUser(ctx context.Context, id string) (model.User, error) {
	return as.repo.GetUser(ctx, id)
}

func (as *AdminService) GetInstructor(ctx context.Context, id string) (model.Instructor, error) {
	return as.repo.GetInstructor(ctx, id)
}

func (as *AdminService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return as.repo.GetAllUsers(ctx)
}

func (as *AdminService) GetAllInstructors(ctx context.Context) ([]model.Instructor, error) {
	return as.repo.GetAllInstructors(ctx)
}

func (as *AdminService) UpdateUser(ctx context.Context, id string, u model.User) error {
	return as.repo.UpdateUser(ctx, id, u)
}

func (as *AdminService) UpdateInstructor(ctx context.Context, id string, i model.Instructor) error {
	return as.repo.UpdateInstructor(ctx, id, i)
}

func (as *AdminService) DeleteUser(ctx context.Context, id string) error {
	return as.repo.DeleteUser(ctx, id)
}

func (as *AdminService) DeleteInstructor(ctx context.Context, id string) error {
	return as.repo.DeleteInstructor(ctx, id)
}
