package rota_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sky-uk/support-bot/rota"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestRota(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test suite for team")
}

var _ = Describe("CRUD of team members", func() {

	var myTeam rota.Team = nil
	BeforeEach(func() {
		myTeam = rota.NewTeam(temporaryFilepath())
	})

	Context("Read team members", func() {
		It("List gets data from the members file", func() {
			Expect(myTeam.List()).To(Equal([] string{"person1", "person2", "third person", "some other person"}))
		})
	})

	Context("Adding new team members", func() {
		It("Add new team member adds the member to the list", func() {
			myTeam.Add("new member")
			Expect(myTeam.List()).To(Equal([] string{"person1", "person2", "third person", "some other person", "new member"}))
		})

		It("Add new team member should not fail if the member already exists", func() {
			myTeam.Add("third person")
			Expect(myTeam.List()).To(Equal([] string{"person1", "person2", "third person", "some other person"}))
		})
	})

	Context("Removing team members", func() {
		It("Removing existing team member returns success", func() {
			myTeam.Remove("third person")
			Expect(myTeam.List()).To(Equal([] string{"person1", "person2", "some other person"}))
		})
		It("Removing non-existing team member returns success", func() {
			myTeam.Remove("non-existent person")
			Expect(myTeam.List()).To(Equal([] string{"person1", "person2", "third person", "some other person"}))
		})
	})
})

func temporaryFilepath() string {
	_, filename, _, _ := runtime.Caller(1)
	basePath := filepath.Dir(filename)
	data, err := ioutil.ReadFile(basePath + "/test-team-members.yml")
	if err == nil {
		tempPath, err := ioutil.TempFile("/tmp", "supportrotatest")
		if err == nil {
			ioutil.WriteFile(tempPath.Name(), data, os.ModeTemporary)
			return tempPath.Name()
		} else {
			panic("Unable to create the temp file")
		}
	} else {
		panic(err)
	}
}
