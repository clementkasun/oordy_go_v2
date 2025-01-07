package resources

import "fiber_app/models"

// UserResource defines methods to transform user data
type UserResource struct{}

// Transform formats a single user object for API responses
func (ur *UserResource) Transform(user models.User) map[string]interface{} {
	return map[string]interface{}{
		"id":             user.ID,
		"uuid":           user.UUID,
		"first_name":     user.FirstName,
		"last_name":      user.LastName,
		"email":          user.Email,
		"payment_mode":   user.PaymentMode,
		"user_type":      user.UserType,
		"gender":         user.Gender,
		"country_code":   user.CountryCode,
		"mobile":         user.Mobile,
		"wallet_balance": user.WalletBalance,
		"rating":         user.Rating,
		"otp":            user.OTP,
		"language":       user.Language,
		"qrcode_url":     user.QRCodeURL,
		"referral_count": user.ReferralCount,
		"created_at":     user.CreatedAt,
		"updated_at":     user.UpdatedAt,
	}
}

// TransformCollection formats a list of user objects for API responses
func (ur *UserResource) TransformCollection(users []models.User) []map[string]interface{} {
	transformed := make([]map[string]interface{}, len(users))
	for i, user := range users {
		transformed[i] = ur.Transform(user)
	}
	return transformed
}
