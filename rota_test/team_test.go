package rota_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sky-uk/support-bot/localdb"
	"github.com/sky-uk/support-bot/rota"
	"github.com/sky-uk/support-bot/rota_test/helper"
)

var _ = Describe("CRUD of team members", func() {

	var myTeam rota.Team = nil
	BeforeEach(func() {
		myTeam = rota.NewTeam("test_team")
		Expect(localdb.Remove(myTeam.TeamKey())).To(Succeed())
		Expect(localdb.Write(myTeam.TeamKey(), helper.TestTeamMembers()))
	})

	Context("Read team members", func() {
		It("List gets data from the members file", func() {
			Expect(myTeam.List()).To(Equal([] string{"person1", "person2", "third person"}))
		})
	})

	Context("Adding new team members", func() {
		It("Add new team member adds the member to the list", func() {
			Expect(myTeam.Add("new member")).To(Succeed())
			Expect(myTeam.List()).To(Equal([] string{"person1", "person2", "third person", "new member"}))
		})

		It("Add new team member should not fail if the member already exists", func() {
			personToAdd := "third person"
			Expect(myTeam.Add(personToAdd)).To(Succeed())
			Expect(myTeam.List()).To(Equal([] string{"person1", "person2", "third person"}))
		})

		It("Add new team member initialise their support counter key to 0", func() {
			newTeamMember := "fourth person"
			Expect(myTeam.Add(newTeamMember)).To(Succeed())
			Expect(myTeam.SupportHistoryFor(newTeamMember).DaysSupported).To(Equal(uint16(0)))
		})

		It("Adding an existing team member again should not reset the supported days ", func() {
			existingTeamMember := "person1"
			Expect(localdb.Write(myTeam.SupportDaysCounterKey(existingTeamMember), helper.Uint16ToBytes(7))).To(Succeed())

			Expect(myTeam.Add(existingTeamMember)).To(Succeed())
			Expect(myTeam.SupportHistoryFor(existingTeamMember).DaysSupported).To(Equal(uint16(7)))
		})
	})

	Context("Removing team members", func() {
		It("Removing existing team member returns success", func() {
			Expect(myTeam.Remove("third person")).To(Succeed())
			Expect(myTeam.List()).To(Equal([] string{"person1", "person2"}))
		})
		It("Removing non-existing team member returns success", func() {
			Expect(myTeam.Remove("non-existent person")).To(Succeed())
			Expect(myTeam.List()).To(Equal([] string{"person1", "person2", "third person"}))
		})
	})

	Context("Batch add of team members or initialise the whole team", func() {
		It("Creates the entire team members from scratch", func() {

		})

		It("Adds multiple team members retaining old team members", func() {

		})

		It("Adding multiple team members don't add duplicates", func() {

		})
	})
})

