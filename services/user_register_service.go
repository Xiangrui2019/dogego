package services

type UserRegisterService struct {
	PhoneNumber     string `form:"phone_number" json:"phone_number" binding:"required,min=11,max=30"`
	NickName        string `form:"nick_name" json:"nick_name" binding:"required,min=2,max=100"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

func (service *UserRegisterService) Register() {

}
