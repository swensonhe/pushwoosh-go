package pushwoosh_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/labstack/gommon/random"
	"github.com/swensonhe/actually-backend/src/svc/pushwoosh"
)

var _ = Describe("Pushwoosh", func() {

	Describe("RegisterDevice", func() {
		It("is successful", func() {
			token := random.String(6)
			hwid := random.String(6)
			err := pushwoosh.SharedClient().RegisterDevice(token, hwid, pushwoosh.DeviceTypeiOS)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("UnregisterDevice", func() {
		It("is successful", func() {
			hwid := random.String(6)
			err := pushwoosh.SharedClient().UnregisterDevice(hwid)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("CreateMessage", func() {
		It("is successful", func() {
			token := random.String(6)
			err := pushwoosh.SharedClient().CreateMessage(pushwoosh.CreateMessageOpts{
				Content: "",
				ContentAvailable: pushwoosh.ContentAvailable,
				Tokens: []string{token},
			})
			Expect(err).ToNot(HaveOccurred())
		})
	})

})
