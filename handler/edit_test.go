package handler

import (
	"fmt"
	"testing"
	"time"

	"github.com/class/orb/model"
	"github.com/class/orb/types"
)

func assertModifiedUpdated(t *testing.T, record types.TimeStamper, previousRecord types.TimeStamper) {
	if !record.GetModified().After(previousRecord.GetModified()) {
		t.Errorf("Modified date %v should be after NOW: %v", record.GetModified(), previousRecord.GetModified())
	}
}

func testUpdateUser(t *testing.T) {
	// t.SkipNow()
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.Input = "test_edit_user.json"
	rt.User = TestAdminUser
	rt.H = UserEdit
	rt.Method = "PUT"
	rt.Pattern = "/users"
	rt.URL = "/users"
	rt.RequestHeader = JSONHeader()

	rt.Run(t)
}

func testUpdateOrg(t *testing.T) {
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.Input = "test_update_org.json"
	rt.User = TestAdminUser
	rt.H = EditOrg
	rt.Method = "PUT"
	rt.Pattern = "/orgs/:id"
	rt.URL = "/orgs/2"
	rt.RequestHeader = JSONHeader()

	rt.Run(t)

	_, err := model.NewOrganizationTable().GetByName("Updated Test Org 2")
	if err != nil {
		t.Error("Did not retrieve updated record from database - got error: ", err)
	}
}

func testUpdateEvent(t *testing.T) {
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.Input = "test_update_event.json"
	rt.User = TestAdminUser
	rt.H = EditEvent
	rt.Method = "PUT"
	rt.Pattern = "/events/:id"
	rt.URL = "/events/15"
	rt.RequestHeader = JSONHeader()

	rt.Run(t)
	ev, err := model.NewEventTable().GetByName("Updated Event 1")
	if err != nil {
		t.Fatal("Could not retrieve updated event record: ", err)
	}

	var expected time.Time
	if expected, _ = time.Parse(time.RFC1123, "Wed, 02 Nov 2016 19:10:19 GMT"); !ev.GetStart().Equal(expected) {
		t.Error("Event start was not updated properly. Expected %v but got %v", expected, ev.GetStart())
	}
	if expected, _ = time.Parse(time.RFC1123, "Fri, 04 Nov 2016 19:10:19 GMT"); !ev.GetEnd().Equal(expected) {
		t.Errorf("Event end was not updated properly. Expected %v but got %v", expected, ev.GetEnd())
	}
}

func testUpdateZone(t *testing.T) {
	oldID := 6
	oldZ, err := model.NewZoneTable().Get(oldID)
	if err != nil {
		t.Fatal("Unable to read original zone with ID : ", oldID)
	}
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.ExpectedResponseFragment = "This description has been updated."
	rt.Input = "test_update_zone.json"
	rt.User = TestAdminUser
	rt.H = EditZone
	rt.Method = "PUT"
	rt.Pattern = "/zones/:id"
	rt.URL = fmt.Sprintf("/zones/%v", oldID)
	rt.RequestHeader = JSONHeader()

	rt.Run(t)

	z, err := model.NewZoneTable().GetByName("Updated Zone 1")
	if err != nil {
		t.Fatal("Could not retrieve updated zone record: ", err)
	}

	// if !z.Modified.After(now) {
	// 	t.Errorf("Modified date %v should be after NOW: %v", z.Modified, now)
	// }
	assertModifiedUpdated(t, z, oldZ)
}

func testUpdateHotel(t *testing.T) {
	oldID := 14
	oldH, err := model.NewHotelTable().Get(oldID)
	if err != nil {
		t.Fatal("Could not retrieve original hotel record: ", oldID)
	}
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.ExpectedResponseFragment = "Now With Contact Info"
	rt.Input = "test_update_hotel.json"
	rt.User = TestAdminUser
	rt.H = EditHotel
	rt.Method = "PUT"
	rt.Pattern = "/hotels/:id"
	rt.URL = fmt.Sprintf("/hotels/%v", oldID)
	rt.RequestHeader = JSONHeader()

	rt.Run(t)

	h, err := model.NewHotelTable().GetByName("Updated Hotel 1")
	if err != nil {
		t.Fatal("Could not retreive updated record from the database.")
	}

	assertModifiedUpdated(t, h, oldH)
	if h.ModifierID != TestAdminUser.ID {
		t.Error("Modifier was not updated correctly.")
	}
}

func testUpdateRoomType(t *testing.T) {
	oldID := 34
	oldRoom, err := model.NewRoomTypeTable().Get(oldID)
	if err != nil {
		t.Fatal("Unable to retrieve record to edit.")
	}
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.ExpectedResponseFragment = "Updated Room Type 1"
	rt.Input = "test_update_room_type.json"
	rt.User = TestAdminUser
	rt.H = EditRoomType
	rt.Method = "PUT"
	rt.Pattern = "/room_types/:id"
	rt.URL = fmt.Sprintf("/room_types/%v", oldID)
	rt.RequestHeader = JSONHeader()

	rt.Run(t)

	room, err := model.NewRoomTypeTable().GetByName("Updated Room Type 1")
	if err != nil {
		t.Fatal("Could not retrieve updated record  by name: ", err)
	}
	if room.Rate != 150 {
		t.Fatal("Room rate did not update properly")
	}
	assertModifiedUpdated(t, room, oldRoom)
	if room.ModifierID != TestAdminUser.ID {
		t.Error("Modifier was not updated correctly.")
	}
}

func testUpdateAttendeeCategory(t *testing.T) {
	testID := 5
	originalCategory, err := model.NewAttendeeCategoryTable().Get(testID)
	if err != nil {
		t.Fatal("Unable to retrieve record to edit.")
	}
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.Input = "test_update_attendee_category.json"
	rt.ExpectedResponseFragment = "Updated Test Category 5"
	rt.User = TestAdminUser
	rt.H = EditAttendeeCategory
	rt.Method = "PUT"
	rt.Pattern = "/attendee_categories/:id"
	rt.URL = fmt.Sprintf("/attendee_categories/%v", testID)
	rt.RequestHeader = JSONHeader()

	rt.Run(t)

	ac, err := model.NewAttendeeCategoryTable().Get(testID)
	if err != nil {
		t.Fatal("Could not retreive update record.")
	}
	if ac.Name != rt.ExpectedResponseFragment {
		t.Fatal("Record did not update.")
	}
	assertModifiedUpdated(t, ac, originalCategory)
}
