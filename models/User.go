package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID             string  `json:"uuid"`
	FirstName        string  `json:"first_name"`
	LastName         string  `json:"last_name"`
	Email            string  `json:"email" gorm:"unique"`
	Password         string  `json:"password"`
	PaymentMode      string  `json:"payment_mode"`
	UserType         string  `json:"user_type" gorm:"default:NORMAL"`
	Gender           string  `json:"gender" gorm:"default:MALE"`
	CountryCode      string  `json:"country_code"`
	Mobile           string  `json:"mobile"`
	Picture          string  `json:"picture"`
	DeviceToken      string  `json:"device_token"`
	DeviceID         string  `json:"device_id"`
	DeviceType       string  `json:"device_type"`
	LoginBy          string  `json:"login_by"`
	SocialUniqueID   string  `json:"social_unique_id"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	StripeCustID     string  `json:"stripe_cust_id"`
	WalletBalance    float64 `json:"wallet_balance" gorm:"default:0.00"`
	Rating           float64 `json:"rating" gorm:"default:5.00"`
	OTP              int     `json:"otp" gorm:"default:0"`
	Language         string  `json:"language"`
	QRCodeURL        string  `json:"qrcode_url"`
	ReferralUniqueID string  `json:"referral_unique_id"`
	ReferralCount    int     `json:"referral_count" gorm:"default:0"`
	RememberToken    string  `json:"remember_token"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	DeletedAt        *string `json:"deleted_at"`
}
