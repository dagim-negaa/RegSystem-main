package main

import (
	"embed"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

type SexType string

const (
	Male   SexType = "Male"
	Female SexType = "Female"
)

type Admin struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Student struct {
	ID               uint    `gorm:"primaryKey"`
	Name             string  `gorm:"not null"`
	Sex              SexType `sql:"type:ENUM('Male', 'Female')" gorm:"column:SexType;not null"`
	Age              uint8   `gorm:"not null;check:age >= 0"`
	Birthday         string  `gorm:"not null"`
	Address          string  `gorm:"not null"`
	BirthPlace       string  `gorm:"not null"`
	NameOfChrist     string  `gorm:"not null"`
	MotherName       string  `gorm:"not null"`
	Kebele           string  `gorm:"not null"`
	HouseNo          string  `gorm:"not null"`
	PhoneNo          string  `gorm:"not null"`
	Email            string  `gorm:"not null;unique"`
	Username         string  `gorm:"not null;unique"`
	PriviesSchool    string  `gorm:"not null"`
	EducationLevel   string  `gorm:"not null"`
	WorkPosition     string  `gorm:"not null"`
	ChristFatherName string  `gorm:"not null"`
	Location         string  `gorm:"not null"`
	EmergencyName    string  `gorm:"not null"`
	EmergencyPhoneNo string  `gorm:"not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func main() {
	// Get the user's AppData directory
	appDataDir, err := os.UserConfigDir()
	if err != nil {
		panic("failed to get user config directory")
	}

	// Create a subdirectory for your app
	appDir := filepath.Join(appDataDir, "RegSystem")
	err = os.MkdirAll(appDir, os.ModePerm)
	if err != nil {
		panic("failed to create app directory")
	}

	// Set the database path
	dbPath := filepath.Join(appDir, "RegSystem.db")

	// Open the database
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Admin{}, &Student{}) // Automatically migrate the schema

	// Create an instance of the app structure
	app := NewApp(db)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "RegSystem",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []any{
			app,
		},
		EnumBind: []any{
			Admin{},
			Student{},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
