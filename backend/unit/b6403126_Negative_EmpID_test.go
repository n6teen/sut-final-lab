package unit

import (
	"testing"
	"github.com/asaskevich/govalidator"

	"github.com/n6teen/sut-final-lab/entity"
	. "github.com/onsi/gomega"
)

func EmpIdTest(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`URl incorrect`, func(t *testing.T) {
		Emp := entity.Employees{
			Name: "Supachalita",
			URL: "https://github.com/",
			EmployeeID: "ME123456789",
		}

		ok, err := govalidator.ValidateStruct(Emp)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("EmployeeID is invalid"))



	})
}
