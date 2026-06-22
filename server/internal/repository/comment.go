package repository

import (
	"github.com/VinceZCL/FinalYearProject/internal/client"
	"github.com/VinceZCL/FinalYearProject/tools"
	"github.com/VinceZCL/FinalYearProject/types/model"
)

type CommentRepository interface {
	NewComment(input model.Comment) error
	GetTeamComments(teamID uint, date string) ([]model.Comment, error)
}

type commentRepository struct {
	client *client.PostgresClient
}

func NewCommentRepository(dbclient *client.PostgresClient) CommentRepository {
	return &commentRepository{client: dbclient}
}

func (r *commentRepository) NewComment(input model.Comment) error {
	return r.client.DB.Create(&input).Error
}

func (r *commentRepository) GetTeamComments(teamID uint, date string) ([]model.Comment, error) {
	var comments []model.Comment
	start, end, err := tools.GetTimes(date)
	if err != nil {
		return nil, err
	}
	err = r.client.DB.Model(&model.Comment{}).
		Where(`fyp_scrum_comments.created_at >= ? AND fyp_scrum_comments.created_at < ?`, start, end).
		Where(`fyp_scrum_comments.team_id = ?`, teamID).
		Preload("User").
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
