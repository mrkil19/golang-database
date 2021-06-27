package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	golang_database "golang-database"
	"golang-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email: "alfi.nd@gmail.com",
		Comment: "Semangat abang, keep u spirit",

	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	CommentRepository := NewCommentRepository(golang_database.GetConnection())

	comment, err := CommentRepository.FindByID(context.Background(), 32)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(golang_database.GetConnection())

	comments, err := CommentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments{
		fmt.Println(comment)
	}
}
