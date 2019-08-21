package models

func MigrationModels() error {
	DB.AutoMigrate(&User{})
	return nil
}
