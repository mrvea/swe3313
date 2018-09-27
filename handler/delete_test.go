package handler

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/class/log"
	"github.com/class/pizza/model"
)

func testDeleteOrg(t *testing.T) {
	toDelete := 5

	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.User = TestAdminUser
	rt.H = DeleteOrg
	rt.Method = "DELETE"
	rt.Pattern = "/orgs/:id"
	rt.URL = fmt.Sprintf("/orgs/%v", toDelete)

	rt.Run(t)

	_, err := model.NewOrganizationTable().Get(toDelete)
	if err != sql.ErrNoRows {
		log.Debug("Error retrieving record: ", err)
		t.Error("Organization was not deleted as expected.")
	}
}

func testDeleteEvent(t *testing.T) {
	toDelete := 16

	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.User = TestAdminUser
	rt.H = DeleteEvent
	rt.Method = "DELETE"
	rt.Pattern = "/events/:id"
	rt.URL = fmt.Sprintf("/events/%v", toDelete)

	rt.Run(t)
	if _, err := model.NewEventTable().Get(toDelete); err != sql.ErrNoRows {
		log.Debug("Error retrieving record: ", err)
		t.Error("Event was not deleted properly.")
	}
}

func testDeleteZone(t *testing.T) {
	toDelete := 7
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.ExpectedResponse = ""
	rt.User = TestAdminUser
	rt.H = DeleteZone
	rt.Method = "DELETE"
	rt.Pattern = "/zones/:id"
	rt.URL = fmt.Sprintf("/zones/%v", toDelete)

	rt.Run(t)
	if _, err := model.NewZoneTable().GetByName("Delete Me!"); err != sql.ErrNoRows {
		log.Debug("Error retrieving record: ", err)
		t.Error("Event was not deleted properly.")
	}
}

func testDeleteHotel(t *testing.T) {
	toDelete := 15
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.User = TestAdminUser
	rt.H = DeleteHotel
	rt.Method = "DELETE"
	rt.Pattern = "/hotels/:id"
	rt.URL = fmt.Sprintf("/hotels/%v", toDelete)

	rt.Run(t)
	if _, err := model.NewHotelTable().GetByName("Delete Me!"); err != sql.ErrNoRows {
		log.Debug("Error retrieving record: ", err)
		t.Error("Hotel was not deleted properly.")
	}
}

func testDeleteRoomType(t *testing.T) {
	toDelete := 35
	rt := NewRouteTest()
	rt.ExpectedCode = 200
	rt.User = TestAdminUser
	rt.H = DeleteRoomType
	rt.Method = "DELETE"
	rt.Pattern = "/room_types/:id"
	rt.URL = fmt.Sprintf("/room_types/%v", toDelete)

	rt.Run(t)
	if _, err := model.NewRoomTypeTable().GetByName("Delete Me!"); err != sql.ErrNoRows {
		log.Debug("Error retrieving record: ", err)
		t.Error("Room type was not deleted properly.")
	}
}
