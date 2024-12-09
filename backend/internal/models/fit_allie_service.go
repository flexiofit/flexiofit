package models

type FitAllieService struct {
    FitAllieID int `gorm:"primaryKey"`
    ServiceID  int `gorm:"primaryKey"`
}
