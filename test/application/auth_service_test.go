package application

import (
	"testing"

	"github.com/ViniciusMartinsS/manager/internal/application/service"
	constant "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/test/mock/infrastructure"
	. "github.com/onsi/gomega"
)

const (
	EMAIL_INVALID    = "viniciussimonemartins@yahoo.com"
	PASSWORD_INVALID = "123456"

	EMAIL_VALID    = "visimonemartins@hotmail.com"
	PASSWORD_VALID = "Sw@rd2021"
)

func TestProcessorApplicationSuite(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Expect login successfully", func(it *testing.T) {
		var accessTokenType string

		UserRepository := infrastructure.NewuserRepositoryMock(false, false)
		auth := service.NewAuthService(UserRepository).
			Login(EMAIL_VALID, PASSWORD_VALID)

		g.Expect(auth.AccessToken).To(BeAssignableToTypeOf(accessTokenType))
		g.Expect(auth.Code).To(Equal(constant.SUCCESS_CODE))
		g.Expect(auth.Message).To(BeEmpty())
	})

	t.Run("Expect to unauthorized login by email", func(it *testing.T) {
		UserRepository := infrastructure.NewuserRepositoryMock(true, false)
		auth := service.NewAuthService(UserRepository).
			Login(EMAIL_INVALID, PASSWORD_INVALID)

		g.Expect(auth.AccessToken).To(BeEmpty())
		g.Expect(auth.Code).To(Equal(constant.NOT_AUTHORIZED_ERROR_CODE))
		g.Expect(auth.Message).To(Equal(constant.NOT_AUTHORIZED_ERROR_MESSAGE))
	})

	t.Run("Expect to unauthorized login by password", func(it *testing.T) {
		UserRepository := infrastructure.NewuserRepositoryMock(false, false)
		auth := service.NewAuthService(UserRepository).
			Login(EMAIL_INVALID, PASSWORD_INVALID)

		g.Expect(auth.AccessToken).To(BeEmpty())
		g.Expect(auth.Code).To(Equal(constant.NOT_AUTHORIZED_ERROR_CODE))
		g.Expect(auth.Message).To(Equal(constant.NOT_AUTHORIZED_ERROR_MESSAGE))
	})

	t.Run("Expect internal server error", func(it *testing.T) {
		UserRepository := infrastructure.NewuserRepositoryMock(true, true)
		auth := service.NewAuthService(UserRepository).
			Login(EMAIL_INVALID, PASSWORD_INVALID)

		g.Expect(auth.AccessToken).To(BeEmpty())
		g.Expect(auth.Code).To(Equal(constant.INTERNAL_SERVER_ERROR_CODE))
		g.Expect(auth.Message).To(Equal(constant.INTERNAL_SERVER_ERROR_MESSAGE))
	})
}
