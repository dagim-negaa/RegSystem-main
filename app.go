package main

import (
	"context"
	"math/rand"

	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context
	db  *gorm.DB
}

// NewApp creates a new App application struct
func NewApp(db *gorm.DB) *App {
	return &App{
		db: db,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetAdmins() []Admin {
	var admins []Admin
	a.db.Find(&admins)
	return admins
}

func (a *App) Login(username, password string) (bool, error) {
	var admin Admin
	if err := a.db.Where("username = ? AND password = ?", username, password).First(&admin).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (a *App) Register(username, password string) bool {
	admin := Admin{
		ID:       uint(rand.Uint32()),
		Username: username,
		Password: password,
	}
	if err := a.db.Create(&admin).Error; err != nil {
		return false
	}
	return true
}

func (a *App) GetStudents() []Student {
	var students []Student
	a.db.Find(&students)
	return students
}

func (a *App) AddStudent(
	birthday, address, name, sex string, age uint8,
	birthPlace, nameOfChrist, motherName, kebele, houseNo, phoneNo,
	email, username, priviesSchool, educationLevel, workPosition,
	christFatherName, location, emergencyName, emergencyPhoneNo string,
) bool {

	student := Student{
		ID:               uint(rand.Uint32()),
		Name:             name,
		Sex:              SexType(sex),
		Age:              age,
		Birthday:         birthday,
		Address:          address,
		BirthPlace:       birthPlace,
		NameOfChrist:     nameOfChrist,
		MotherName:       motherName,
		Kebele:           kebele,
		HouseNo:          houseNo,
		PhoneNo:          phoneNo,
		Email:            email,
		Username:         username,
		PriviesSchool:    priviesSchool,
		EducationLevel:   educationLevel,
		WorkPosition:     workPosition,
		ChristFatherName: christFatherName,
		Location:         location,
		EmergencyName:    emergencyName,
		EmergencyPhoneNo: emergencyPhoneNo,
	}

	if err := a.db.Create(&student).Error; err != nil {
		return false
	}
	return true
}

func (a *App) EditStudent(id uint, name, address, sex string, age uint8, birthday string) bool {
	student := Student{
		ID:       id,
		Name:     name,
		Address:  address,
		Sex:      SexType(sex),
		Age:      age,
		Birthday: birthday,
	}
	if err := a.db.Save(&student).Error; err != nil {
		return false
	}
	return true
}

func (a *App) DeleteStudent(id uint) bool {
	if err := a.db.Delete(&Student{}, id).Error; err != nil {
		return false
	}
	return true
}
