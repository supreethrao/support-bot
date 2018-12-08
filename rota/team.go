package rota

import (
	"encoding/binary"
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/sky-uk/support-bot/localdb"
	"gopkg.in/yaml.v2"
	"log"
	"time"
)

const teamKey string = "team_members"

type Team interface {
	List() []string
	Add(newMember string) error
	Remove(existingMember string) error
	SupportHistoryFor(member string) IndividualSupportHistory
	SupportHistoryForTeam() TeamSupportHistory
	SupportPersonFor(date time.Time) string
	SetPersonOnSupport(memberName string) error
	Keys
}

type Keys interface {
	TeamKey() string
	SupportDaysCounterKey(memberName string) string
	SupportPersonOnDayKey(supportDay time.Time) string
	LatestDayOnSupportKey(memberName string) string
}

func (t *team) TeamKey() string {
	return t.name + "::" + teamKey
}

func (t *team) SupportDaysCounterKey(memberName string) string {
	return t.name + "::member::" + memberName
}

func (t *team) SupportPersonOnDayKey(supportDay time.Time) string {
	formattedDay := supportDay.Format("02-01-2006")
	return t.name + "::" + formattedDay
}

func (t *team) LatestDayOnSupportKey(memberName string) string {
	return t.name + "::latest-day::" + memberName
}

// name will be used as the key prefix
type team struct {
	name string
}

type teamMembers struct {
	Members []string `yaml:"members"`
}

func (t *team) List() []string {
	data, err := localdb.Read(t.TeamKey())
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return []string{}
		}
		panic(err)
	}

	members := teamMembers{}

	if err = yaml.Unmarshal(data, &members); err != nil {
		log.Fatalf("Unable to obtain team members: %v", err)
	}

	return members.Members
}

func (t *team) Add(newMember string) error {
	currentMembers := t.List()

	for _, member := range currentMembers {
		if newMember == member {
			log.Printf("%s is already a member", newMember)
			return nil
		}
	}

	updatedTeam := teamMembers{append(currentMembers, newMember)}
	if data, err := yaml.Marshal(updatedTeam); err == nil {
		multiData := map[string][]byte{
			t.TeamKey():                        data,
			t.SupportDaysCounterKey(newMember): uintToBytes(0),
		}
		return localdb.MultiWrite(multiData)
	} else {
		return err
	}
}

func (t *team) Remove(existingMember string) error {
	currentMembers := t.List()
	updatedMembers := make([]string, 0)

	for _, member := range currentMembers {
		if member != existingMember {
			updatedMembers = append(updatedMembers, member)
		}
	}
	updatedTeam := teamMembers{updatedMembers}
	if data, err := yaml.Marshal(updatedTeam); err == nil {
		return localdb.Write(t.TeamKey(), data)
	} else {
		return err
	}
}

func (t *team) SupportHistoryFor(member string) IndividualSupportHistory {
	history := IndividualSupportHistory{member, 0, "UNKNOWN"}
	count, err := localdb.Read(t.SupportDaysCounterKey(member))
	if err == nil {
		history.DaysSupported = bytesToUint(count)
	}

	day, err := localdb.Read(t.LatestDayOnSupportKey(member))
	if err == nil {
		history.LatestSupportDay = string(day)
	}

	return history
}

func (t *team) SupportHistoryForTeam() TeamSupportHistory {
	teamHistory := make([]IndividualSupportHistory, 0)
	for _, member := range t.List() {
		teamHistory = append(teamHistory, t.SupportHistoryFor(member))
	}
	return teamHistory
}

func (t *team) SupportPersonFor(date time.Time) string {
	supportPerson, err := localdb.Read(t.SupportPersonOnDayKey(date))
	if err == nil {
		return string(supportPerson)
	}
	return fmt.Sprintf("Unable to retrieve support person for %v", date)
}

func (t *team) SetPersonOnSupport(memberName string) error {
	supportKeys := make(map[string][]byte)

	// get the next sequence number for the counter
	supportKeys[t.LatestDayOnSupportKey(memberName)] = []byte(today())
	supportKeys[t.SupportPersonOnDayKey(time.Now())] = []byte(memberName)

	return localdb.MultiWrite(supportKeys)
}

func uintToBytes(val uint16) []byte {
	bytesVal := make([]byte, 2)
	binary.BigEndian.PutUint16(bytesVal, val)
	return bytesVal
}

func bytesToUint(val []byte) uint16 {
	return binary.BigEndian.Uint16(val)
}

func NewTeam(name string) Team {
	return &team{name}
}
