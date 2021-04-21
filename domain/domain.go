package domain

import (
	"context"
	"github.com/JIeeiroSst/go-app/proto"
)

type User struct {
	ID int 				`gorm:"primaryKey,autoIncrement"`
	Name string 
	Address string 
	Profile Profile     `gorm:"foreignkey:UserId;references:ID"`

}

type Profile struct {
	ID int 				`gorm:"primaryKey,autoIncrement"`
	Name string 
	Email string 
	UserId int 
}

type Service interface {
	UpdateEmail(context.Context,*proto.UpdateEmailRequest) (*proto.Response,error)
	CreateEmail(context.Context,*proto.CreateEmailRequest) (*proto.Response,error)
	DeleteEmail(context.Context,*proto.DeleteEmailRequest) (*proto.Response,error)
}

type Repository interface{
	UserAll() (users []User, err error)
	UserById(int) (user User,err error)
	CreateUser(User) (err error)
	UpdateUser(int,User) (err error)
	DeleteUser(int) (err error)

	ProfileAll() (profiles []Profile,err error)
	ProfileById(int) (profile Profile,err error)
	CreateProfile(Profile) (err error)
	UpdateProfile(int,Profile) (err error)
	DeleteProfile(int) (err error)
}