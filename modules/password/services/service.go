package services

import (
	"errors"
	"log"
	"strings"
	"users/modules/password/models"
	"users/modules/password/repositories"
)

type IService interface {
	InitPasswordService(initPassword models.InitPassword) error
	ChangePasswordService(changePassword models.ChangePassword) error
}

type service struct {
	r repositories.IRepositorie
}

func NewService(r repositories.IRepositorie) IService {
	return &service{r: r}
}

func (s *service) InitPasswordService(initPassword models.InitPassword) error {
	// ตรวจสอบว่า orgmbEmail เป็นอีเมลที่ถูกต้อง
	if !strings.Contains(initPassword.OrgmbEmail, "@") || strings.HasPrefix(initPassword.OrgmbEmail, "@") || strings.HasSuffix(initPassword.OrgmbEmail, "@") || strings.Count(initPassword.OrgmbEmail, "@") != 1 {
		return errors.New("orgmbEmail must be a valid email address (init)")
	}

	// ตรวจสอบว่า password มีความยาวอย่างน้อย 8 ตัวอักษร
	if initPassword.NewPassword == "" || len(initPassword.NewPassword) < 8 {
		return errors.New("password must not be empty and must be at least 8 characters long (init)")
	}

	// เรียกใช้ repository ในการตั้งค่ารหัสผ่านใหม่
	err := s.r.InitPasswordRepository(initPassword)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *service) ChangePasswordService(changePassword models.ChangePassword) error {
	// ตรวจสอบว่า orgmbEmail เป็นอีเมลที่ถูกต้อง
	if !strings.Contains(changePassword.OrgmbEmail, "@") || strings.HasPrefix(changePassword.OrgmbEmail, "@") || strings.HasSuffix(changePassword.OrgmbEmail, "@") || strings.Count(changePassword.OrgmbEmail, "@") != 1 {
		return errors.New("orgmbEmail must be a valid email address")
	}
	// ตรวจสอบว่า password เก่ามีความยาวอย่างน้อย 8 ตัวอักษร
	if changePassword.Oldpassword == "" || len(changePassword.Oldpassword) < 8 {
		return errors.New("old password must not be empty and must be at least 8 characters long")
	}
	// ตรวจสอบว่า password ใหม่มีความยาวอย่างน้อย 8 ตัวอักษร
	if changePassword.Newpassword == "" || len(changePassword.Newpassword) < 8 {
		return errors.New("new password must not be empty and must be at least 8 characters long")
	}
	// เรียกใช้ repository ในการเปลี่ยนรหัสผ่าน
	err := s.r.ChangePasswordRepository(changePassword)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
