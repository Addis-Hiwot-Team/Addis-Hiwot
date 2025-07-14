package usecases

import (
	"addis-hiwot/internal/domain/interfaces"
	"addis-hiwot/internal/domain/models"
	"addis-hiwot/internal/repository"
	"addis-hiwot/internal/service"
	"addis-hiwot/utils"
	"errors"
	"os"
	"time"
)

type UserUsecase struct {
	repo         interfaces.UserRepository
	jwtService   interfaces.JWTService
	emailService service.EmailService
	otpRepo      repository.OtpRepo
}

func NewUserUsecase(
	r interfaces.UserRepository,
	j interfaces.JWTService,
	em service.EmailService,
	or repository.OtpRepo,
) *UserUsecase {
	return &UserUsecase{
		repo:         r,
		jwtService:   j,
		emailService: em,
		otpRepo:      or,
	}
}

func (uc *UserUsecase) GetAll() ([]*models.UserResponse, error) {
	return uc.repo.GetAll()
}

func (uc *UserUsecase) GetByID(id int) (*models.UserResponse, error) {
	user, err := uc.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return user.ToResponse(), nil
}

func (uc *UserUsecase) ChangePassword(userID int, password, newPassword string) error {
	user, err := uc.repo.Get(userID)
	if err != nil {
		return err
	}
	err = utils.CheckPassword(user.PasswordHash, password)
	if err != nil {
		return errors.New("current password is incorrect")
	}
	hashedPwd, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}
	return uc.repo.ChangePassword(userID, hashedPwd)
}

func (uc *UserUsecase) ForgotPassword(email string) error {
	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	code := utils.GenerateOTP() // 6-digit or similar
	otp := &models.Otp{
		UserID:    uint(user.ID),
		Code:      code,
		Type:      "password_reset",
		Exp:       time.Now().Add(15 * time.Minute),
		CreatedAt: time.Now(),
	}

	if _, err := uc.otpRepo.Create(otp); err != nil {
		return err
	}

	// Send the email (abstract email service)
	return uc.emailService.SendEmail(user.Email, "Forgot Password", "forgotpassword.html", map[string]any{
		"UserName":          user.Name,
		"ResetLink":         os.Getenv("FRONTEND_URL") + "/reset-password?token=" + code,
		"ExpirationMinutes": 15,
		"CurrentYear":       time.Now().Year(),
		"WebsiteLink":       "https://www.addishiwt.com",
		"SupportLink":       "mailto:support@addishiwt.com",
	})
}

func (uc *UserUsecase) ResetPassword(code, newPassword string) error {
	otp, err := uc.otpRepo.Get(code, "password_reset")
	if err != nil {
		return errors.New("invalid or expired code")
	}
	if otp.Exp.Before(time.Now()) {
		return errors.New("code has expired")
	}

	hashed, _ := utils.HashPassword(newPassword)
	if err := uc.repo.ChangePassword(int(otp.UserID), string(hashed)); err != nil {
		return err
	}

	// Invalidate the OTP
	utils.LogIfError("PWD:Reset", uc.otpRepo.Delete(otp.ID))
	return nil
}

func (uc *UserUsecase) ActivateAccount(code string) error {
	otp, err := uc.otpRepo.Get(code, "account_verification")
	if err != nil {
		return errors.New("invalid or expired code")
	}
	if otp.Exp.Before(time.Now()) {
		return errors.New("code has expired")
	}

	user, err := uc.repo.Get(int(otp.UserID))
	if err != nil {
		return errors.New("user not found")
	}

	if err := uc.repo.Activate(user.ID); err != nil {
		return err
	}

	// Invalidate the OTP
	utils.LogIfError("ACC:Verify", uc.otpRepo.Delete(otp.ID))
	return nil
}
