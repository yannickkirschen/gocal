package ical

import (
	"github.com/google/uuid"
	"gocal/geo"
	"testing"
	"time"
)

func buildEvent(t *testing.T) *Event {
	creation, err := time.Parse(time.RFC3339, "2023-11-27T10:00:00Z")
	if err != nil {
		t.Errorf("Error creating time: %s", err)
	}

	begin, err := time.Parse(time.RFC3339, "2023-11-28T12:05:00Z")
	if err != nil {
		t.Errorf("Error creating time: %s", err)
	}

	end, err := time.Parse(time.RFC3339, "2023-11-28T12:35:00Z")
	if err != nil {
		t.Errorf("Error creating time: %s", err)
	}

	return NewEventBuilder().
		UUID(uuid.MustParse("7bb770da-8d04-11ee-b41a-f202af5473b8")).
		Creation(creation).
		Begin(begin).
		End(end).
		Organizer(NewOrganizerBuilder().Name("John Doe").Email("doe@example.com").Build()).
		Summary("Test").
		Geo(geo.NewCoordinates(49.474193, 8.534854)).
		Build()
}

func TestEvent(t *testing.T) {
	event := buildEvent(t)
	eventString := event.String()
	expected := "BEGIN:VEVENT\n" +
		"UID:7bb770da-8d04-11ee-b41a-f202af5473b8\n" +
		"DTSTAMP:20231127T100000Z\n" +
		"DTSTART:20231128T120500Z\n" +
		"DTEND:20231128T123500Z\n" +
		"ORGANIZER;CN=John Doe:MAILTO:doe@example.com\n" +
		"SUMMARY:Test\n" +
		"GEO:49.474193;8.534854\n" +
		"END:VEVENT"

	if eventString != expected {
		t.Errorf("Event string is not as expected:\n%s\n%s", eventString, expected)
	}
}

func TestCalendar(t *testing.T) {
	event := buildEvent(t)
	calendar := NewCalendarBuilder().LatestVersion().ProdID([]string{"Test", "local", "EN"}).AddEvent(event).Build()

	calendarString := calendar.String()
	expected := "BEGIN:VCALENDAR\n" +
		"VERSION:2.0\n" +
		"PRODID://Test//local//EN\n" +
		"BEGIN:VEVENT\n" +
		"UID:7bb770da-8d04-11ee-b41a-f202af5473b8\n" +
		"DTSTAMP:20231127T100000Z\n" +
		"DTSTART:20231128T120500Z\n" +
		"DTEND:20231128T123500Z\n" +
		"ORGANIZER;CN=John Doe:MAILTO:doe@example.com\n" +
		"SUMMARY:Test\n" +
		"GEO:49.474193;8.534854\n" +
		"END:VEVENT\n" +
		"END:VCALENDAR"

	if calendarString != expected {
		t.Errorf("Calendar string is not as expected:\n%s\n%s", calendarString, expected)
	}
}
