package tests

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/Weeping-Willow/bdd-go-testing/src"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

const (
	userDeleteRoute  = "/public/v1/users/%d"
	userGetRoute     = "/public/v1/users/%d"
	userPutRoute     = "/public/v1/users/%d"
	userPostRoute    = "/public/v1/users"
	userListGetRoute = "/public/v1/users"
	urlToTest        = "https://gorest.co.in"
)

func TestGoRestAPISuite(t *testing.T) {
	c, err := src.NewConfig("../.env")
	assert.NoError(t, err)
	testConfig = c
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoRest API Suite")
}

func TestRunningRegularTestBesidesGinkgo(t *testing.T) {
	assert.Equal(t, "Hello World", "Hello World")
}

var _ = Describe("Sample test run", func() {
	Describe("the strings package", func() {
		Context("strings.Contains()", func() {
			When("the string contains the substring in the middle", func() {
				It("returns `true`", func() {
					Expect(strings.Contains("Ginkgo is awesome", "is")).To(BeTrue())
				})
			})
		})
	})
})

var _ = Describe("GoRest API User create : task 1", func() {
	BeforeEach(func() {
		cleanUpTestUser("")
	})
	AfterEach(func() {
		cleanUpTestUser("")
	})

	Describe("Create a user: Positive scenario", func() {
		Context(contextName(http.MethodPost, userPostRoute), func() {
			When("user is created successfully", func() {
				It("returns 201 status code with user object inside the body under the data field", func() {
					u, body, status := createUser(getDefaultUser())
					if status != http.StatusCreated {
						fmt.Println(body)
					}

					Expect(u).To(Equal(getDefaultUser().userWithId(u.Id)))
					Expect(status).To(Equal(http.StatusCreated))
				})
			})
		})
	})
})

var _ = Describe("GoRest API Update user name : task 2", func() {
	userTest := getDefaultUser().userWithEmail("randomemail2@kmail.com")
	cleanUpTestUser(userTest.Email)
	defer cleanUpTestUser(userTest.Email)

	userId := 0
	Describe("Update user name. Positive scenario", func() {
		cleanUpTestUser("")
		Context(contextName(http.MethodPost, userPostRoute), func() {
			When("user that we will updated is created successfully", func() {
				It("returns 201 status code with user object inside the body under the data field", func() {
					u, body, status := createUser(userTest)
					if status != http.StatusCreated {
						fmt.Println(body)
					}
					userId = u.Id
					Expect(u).To(Equal(userTest.userWithId(u.Id)))
					Expect(status).To(Equal(http.StatusCreated))
				})
			})
		})
		Context(contextName(http.MethodPatch, userPutRoute), func() {
			When("user name is updated successfully", func() {
				It("returns 200 status code with user object inside the body under the data field, this user object should contain the new name", func() {
					newName := "Joe Exotic"
					u, _, status := updateUser(userId, &user{
						Name: newName,
					})

					Expect(u.Id).To(Equal(userId))
					Expect(u.Name).To(Equal(newName))
					Expect(status).To(Equal(http.StatusOK))
				})
			})
		})
	})
})

var _ = Describe("GoRest API User create : task 3", func() {
	userTest := getDefaultUser().userWithEmail("testuser21@kmail.coa")
	cleanUpTestUser(userTest.Email)
	defer cleanUpTestUser(userTest.Email)

	Describe("Create a user with email already in use. Negative scenario.", func() {
		Context(contextName(http.MethodPost, userPostRoute), func() {
			When("first user is created successfully", func() {
				It("returns 201 status code with user object inside the body under the data field", func() {
					u, body, status := createUser(userTest)
					if status != http.StatusCreated {
						fmt.Println(body)
					}
					Expect(u).To(Equal(userTest.userWithId(u.Id)))
					Expect(status).To(Equal(http.StatusCreated))
				})
			})
			When("the second user creation attempt is made with using the exact same details from first user", func() {
				It("should return 422 error and have error about email already being taken in body", func() {
					u, body, status := createUser(userTest)
					Expect(status).To(Equal(http.StatusUnprocessableEntity))
					Expect(u).To(Equal(&user{}))
					Expect(body).To(ContainSubstring(`{"field":"email","message":"has already been taken"}`))
				})
			})
		})
	})
})

var _ = Describe("GoRest API User delete : task 4", func() {
	userTest := getDefaultUser().userWithEmail("testuser54@kmail.coa")
	cleanUpTestUser(userTest.Email)
	defer cleanUpTestUser(userTest.Email)

	userId := 0
	Describe("Remove the already removed user. Negative scenario.", func() {
		Context(contextName(http.MethodPost, userPostRoute), func() {
			When("first user is created successfully", func() {
				It("returns 201 status code with user object inside the body under the data field", func() {
					u, body, status := createUser(userTest)
					if status != http.StatusCreated {
						fmt.Println(body)
					}

					userId = u.Id
					Expect(u).To(Equal(userTest.userWithId(u.Id)))
					Expect(status).To(Equal(http.StatusCreated))
				})
			})
		})
		Context(contextName(http.MethodDelete, userDeleteRoute), func() {
			When("deleting recently created user", func() {
				It("should return 204 status code and have no body", func() {
					body, status := deleteUser(userId)
					Expect(status).To(Equal(http.StatusNoContent))
					Expect(body).To(Equal(""))
				})
			})
			When("trying to delete the same user as before", func() {
				It("should return 404 error and have error about resource not found", func() {
					body, status := deleteUser(userId)
					Expect(status).To(Equal(http.StatusNotFound))
					Expect(body).To(ContainSubstring(`{"message":"Resource not found"}`))
				})
			})
		})
	})
})
