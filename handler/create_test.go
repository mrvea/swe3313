package handler

import (
	"testing"
	"time"

	"github.com/class/pizza/model"
)

func testCreateUser(t *testing.T) {
	// t.SkipNow()
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.Input = "test_create_user.json"
	rt.ExpectedResponseFragment = `"ID":`
	rt.User = TestAdminUser
	rt.H = UserAdd
	rt.Method = "POST"
	rt.Pattern = "/users"
	rt.URL = "/users"
	rt.RequestHeader = JSONHeader()

	rt.Run(t)
}

func testCreateOrg(t *testing.T) {
	// TODO add users, auth
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.ExpectedResponseFragments = []string{
		"Test Insert Organization 1",
		`"ID":`,
	}
	rt.Input = "test_create_org.json"
	rt.User = TestAdminUser
	rt.H = CreateOrg
	rt.Method = "POST"
	rt.Pattern = "/orgs"
	rt.URL = "/orgs"
	rt.RequestHeader = JSONHeader()

	rt.Run(t)
	_, err := model.NewOrganizationTable().GetByName("Test Insert Organization 1")
	if err != nil {
		t.Error("Did not retrieve new record from database - got error: ", err)
	}
}

func testCreateEvent(t *testing.T) {
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.ExpectedResponseFragments = []string{
		"Test Inserted Event",
		`"ID":`,
	}
	rt.Input = "test_create_event.json"
	rt.User = TestAdminUser
	rt.H = CreateEvent
	rt.Method = "POST"
	rt.Pattern = "/events"
	rt.URL = "/events"
	rt.RequestHeader = JSONHeader()

	rt.Run(t)
	_, err := model.NewEventTable().GetByName("Test Inserted Event")
	if err != nil {
		t.Error("Did not retrieve record from database - got error: ", err)
	}
}

func testCreateZone(t *testing.T) {
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.ExpectedResponseFragments = []string{
		"This should describe the new zone.",
		`"ID":`,
	}
	rt.Input = "test_create_zone.json"
	rt.User = TestAdminUser
	rt.H = CreateZone
	rt.Method = "POST"
	rt.Pattern = "/zones"
	rt.URL = "/zones"
	rt.RequestHeader = JSONHeader()

	rt.Run(t)

	_, err := model.NewZoneTable().GetByName("Test Inserted Zone 1")
	if err != nil {
		t.Error("Did not retrieve new zone record from database - got error: ", err)
	}
}

func testCreateHotel(t *testing.T) {
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.Input = "test_create_hotel.json"
	rt.ExpectedResponseFragments = []string{
		"Newly Created Hotel 1",
		`"ID":`,
	}
	rt.User = TestAdminUser
	rt.H = CreateHotel
	rt.Method = "POST"
	rt.Pattern = "/hotels"
	rt.URL = "/hotels"
	rt.RequestHeader = JSONHeader()

	rt.Run(t)

	_, err := model.NewHotelTable().GetByName("Newly Created Hotel 1")
	if err != nil {
		t.Error("Did not retrieve new hotel record from database - got error: ", err)
	}
}

func testCreateRoomType(t *testing.T) {
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.Input = "test_create_room_type.json"
	rt.ExpectedResponseFragments = []string{
		"Newly Created Room Type 1",
		`"ID":`,
	}
	rt.User = TestAdminUser
	rt.H = CreateRoomType
	rt.Method = "POST"
	rt.Pattern = "/room_types"
	rt.URL = "/room_types"
	rt.RequestHeader = JSONHeader()

	rt.Run(t)

	_, err := model.NewRoomTypeTable().GetByName("Newly Created Room Type 1")
	if err != nil {
		t.Error("Did not retrieve new room type record from database - got error: ", err)
	}
}

func testCreateRoomTypeWithInventory(t *testing.T) {
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.Input = "test_create_room_type_with_inventory.json"
	rt.ExpectedResponseFragments = []string{
		"Newly Created Room Type With Inventory",
		`"ID":`,
	}
	rt.User = TestAdminUser
	rt.H = CreateRoomType
	rt.Method = "POST"
	rt.Pattern = "/room_types"
	rt.URL = "/room_types"
	rt.RequestHeader = JSONHeader()

	rt.Run(t)

	newRoomType, err := model.NewRoomTypeTable().GetByName("Newly Created Room Type With Inventory")
	if err != nil {
		t.Fatal("Did not retrieve new room type record from database - got error: ", err)
	}

	dateFormat := "2006-01-02"
	day, err := time.Parse(dateFormat, "2016-11-03")
	if err != nil {
		t.Error("Error parsing date: ", err)
	}
	availability, err := model.NewRoomTypeTable().GetAvailability(newRoomType.ID, day)
	if availability != 55 {
		t.Errorf("Expected availability of %v, for new room, but got %v", 55, availability)
	}

	day, err = time.Parse(dateFormat, "2016-11-02")
	if err != nil {
		t.Error("Error parsing date: ", err)
	}
	availability, err = model.NewRoomTypeTable().GetAvailability(newRoomType.ID, day)
	if availability != 125 {
		t.Errorf("Expected availability of %v for new room, but got %v", 125, availability)
	}
}

func testCreateSubblock(t *testing.T) {
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	// rt.ExpectedResponseFragment = "" TODO
	rt.Input = "test_create_subblock.json"
	rt.User = TestAdminUser
	rt.H = CreateSubblock
	rt.Method = "POST"
	rt.Pattern = "/subblocks"
	rt.URL = "/subblocks"
	rt.RequestHeader = JSONHeader()

	rt.Run(t)

	day, err := time.Parse("2006-01-02", "2016-11-02")
	if err != nil {
		t.Error("Error parsing date: ", err)
		t.FailNow()
	}
	availability, err := model.NewRoomTypeTable().GetAvailability(25, day)
	if availability != -15 {
		t.Errorf("Availability should be %v, but got %v.  The value of error is: %v",
			5, availability, err)
	}
}
