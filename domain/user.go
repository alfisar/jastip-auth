package domain

type User struct {
	Id       int    `gorm:"primaryKey; column:id" json:"id"`
	FullName string `gorm:"column:full_name" json:"full_name"`
	Username string `gorm:"column:username" json:"username"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	NoHP     string `gorm:"column:nohp" json:"nohp"`
	Role     int    `gorm:"column:role" json:"role"`
	Status   int    `gorm:"column:status" json:"status"`
}

type UserResponse struct {
	Id       int    `gorm:"primaryKey; column:id" json:"id"`
	FullName string `gorm:"column:full_name" json:"full_name"`
	Username string `gorm:"column:username" json:"username"`
	Status   int    `gorm:"column:status" json:"status"`
}

type UserResendOtpRequest struct {
	FullName string `gorm:"column:full_name" json:"full_name"`
	Email    string `gorm:"column:email" json:"email"`
	NoHP     string `gorm:"column:nohp" json:"nohp"`
}

type UserVerifyOtpRequest struct {
	Otp   string `json:"otp"`
	Email string `json:"email"`
	NoHP  string `json:"nohp"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ProfileResponse struct {
	Id       int    `gorm:"primaryKey; column:id" json:"id"`
	FullName string `gorm:"column:full_name" json:"full_name"`
	Username string `gorm:"column:username" json:"username"`
	Email    string `gorm:"column:email" json:"email"`
	NoHP     string `gorm:"column:nohp" json:"nohp"`
	Role     int    `gorm:"column:role" json:"role"`
	Status   int    `gorm:"column:status" json:"status"`
}

type AddressResponse struct {
	Id          int    `gorm:"primaryKey; column:id" json:"id"`
	Street      string `gorm:"column:street" json:"street"`
	City        string `gorm:"column:city" json:"city"`
	District    string `gorm:"column:district" json:"district"`
	SUbDistrict string `gorm:"column:subdistrict" json:"subdistrict"`
	PostalCode  string `gorm:"column:postalcode" json:"postal_code"`
}

type AddressRequest struct {
	Id          int    `gorm:"primaryKey; column:id" json:"id"`
	UserID      int    `gorm:"column:user_id" json:"user_id"`
	Street      string `gorm:"column:street" json:"street"`
	City        string `gorm:"column:city" json:"city"`
	District    string `gorm:"column:district" json:"district"`
	SUbDistrict string `gorm:"column:subdistrict" json:"subdistrict"`
	PostalCode  string `gorm:"column:postalcode" json:"postalcode"`
}
