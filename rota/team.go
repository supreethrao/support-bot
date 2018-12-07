package rota

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Team interface {
	List() []string
	Add(newMember string)
	Remove(existingMember string)
}

type team struct {
	membersFilePath string
}

type teamMembers struct {
	Members []string `yaml:"members"`
}

func (t *team) List() []string {
	val, err := ioutil.ReadFile(t.membersFilePath)
	if err != nil {
		log.Fatalf("Unable to read the team members file: %v", err)
	}

	members := teamMembers{}

	if err = yaml.Unmarshal(val, &members); err != nil {
		log.Fatalf("Unable to obtain team members: %v", err)
	}

	return members.Members
}

func (t *team) Add(newMember string) {
	currentMembers := t.List()

	for _, member := range currentMembers {
		if newMember == member {
			log.Printf("%s is already a member", newMember)
			return
		}
	}

	updatedMembers := append(currentMembers, newMember)

	if err := t.updateMembersFile(updatedMembers); err != nil {
		log.Fatalf("Unable to remove team member %s", newMember)
	}
}

func (t *team) Remove(existingMember string) {
	currentMembers := t.List()
	updatedMembers := make([]string, 0)

	for _, member := range currentMembers {
		if member != existingMember {
			updatedMembers = append(updatedMembers, member)
		}
	}

	 if err := t.updateMembersFile(updatedMembers); err != nil {
		log.Fatalf("Unable to remove team member %s", existingMember)
	}
}

func (t *team) updateMembersFile(updatedMembers []string) error {
	updatedTeam := teamMembers{updatedMembers}
	if data, err := yaml.Marshal(updatedTeam); err == nil {
		if err := ioutil.WriteFile(t.membersFilePath, data, os.ModePerm); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

func NewTeam(membersFilePath string) Team {
	return &team{membersFilePath}
}
