package unit

import (
	"testing"
	"github.com/asaskevich/govalidator"

	"github.com/n6teen/sut-final-lab/entity"
	. "github.com/onsi/gomega"
)

func EmployeesTest(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`Employee Correct`, func(t *testing.T) {
		Emp := entity.Employees{
			Name: "Supachalita",
			URL: "https://github.com/",
			EmployeeID: "EM1234567890",
		}

		ok, err := govalidator.ValidateStruct(Emp)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())



	})

	t.Run(`URl incorrect`, func(t *testing.T) {
		Emp := entity.Employees{
			Name: "Supachalita",
			URL: "httpsgithubcom",
			EmployeeID: "EM1234567890",
		}

		ok, err := govalidator.ValidateStruct(Emp)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("URL is invalid"))



	})
}
