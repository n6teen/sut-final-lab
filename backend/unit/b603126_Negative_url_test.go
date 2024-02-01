package unit

import (
	"testing"
	"github.com/asaskevich/govalidator"

	"github.com/n6teen/sut-final-lab/entity"
	. "github.com/onsi/gomega"
)

func UrlTest(t *testing.T) {
	g := NewGomegaWithT(t)
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
