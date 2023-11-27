package ical

import (
	"fmt"
	"github.com/google/uuid"
	"gocal/geo"
	"strings"
	"time"
)

type Calendar struct {
	Version string
	ProdID  []string
	Events  []*Event
}

type CalendarBuilder struct {
	calendar *Calendar
}

func NewCalendarBuilder() *CalendarBuilder {
	return &CalendarBuilder{
		calendar: &Calendar{},
	}
}

func (b *CalendarBuilder) LatestVersion() *CalendarBuilder {
	b.calendar.Version = "2.0"
	return b
}

func (b *CalendarBuilder) Version(version string) *CalendarBuilder {
	b.calendar.Version = version
	return b
}

func (b *CalendarBuilder) ProdID(prodID []string) *CalendarBuilder {
	b.calendar.ProdID = prodID
	return b
}

func (b *CalendarBuilder) Events(events []*Event) *CalendarBuilder {
	b.calendar.Events = events
	return b
}

func (b *CalendarBuilder) AddEvent(event *Event) *CalendarBuilder {
	b.calendar.Events = append(b.calendar.Events, event)
	return b
}

func (b *CalendarBuilder) Build() *Calendar {
	return b.calendar
}

func (c Calendar) String() string {
	var events string
	for _, event := range c.Events {
		events += event.String() + "\n"
	}

	return fmt.Sprintf(""+
		"BEGIN:VCALENDAR\n"+
		"VERSION:%s\n"+
		"PRODID://%s\n"+
		"%s"+
		"END:VCALENDAR",
		c.Version,
		strings.Join(c.ProdID, "//"),
		events,
	)
}

type Event struct {
	UUID      uuid.UUID
	Creation  time.Time
	Begin     time.Time
	End       time.Time
	Organizer *Organizer
	Summary   string
	Geo       *geo.Coordinates
}

type EventBuilder struct {
	event *Event
}

func NewEventBuilder() *EventBuilder {
	return &EventBuilder{
		event: &Event{},
	}
}

func (b *EventBuilder) UUID(uuid uuid.UUID) *EventBuilder {
	b.event.UUID = uuid
	return b
}

func (b *EventBuilder) Creation(creation time.Time) *EventBuilder {
	b.event.Creation = creation
	return b
}

func (b *EventBuilder) Begin(begin time.Time) *EventBuilder {
	b.event.Begin = begin
	return b
}

func (b *EventBuilder) End(end time.Time) *EventBuilder {
	b.event.End = end
	return b
}

func (b *EventBuilder) Organizer(organizer *Organizer) *EventBuilder {
	b.event.Organizer = organizer
	return b
}

func (b *EventBuilder) Summary(summary string) *EventBuilder {
	b.event.Summary = summary
	return b
}

func (b *EventBuilder) Geo(geo *geo.Coordinates) *EventBuilder {
	b.event.Geo = geo
	return b
}

func (b *EventBuilder) Build() *Event {
	return b.event
}

func (e Event) String() string {
	return fmt.Sprintf(""+
		"BEGIN:VEVENT\n"+
		"UID:%s\n"+
		"DTSTAMP:%s\n"+
		"DTSTART:%s\n"+
		"DTEND:%s\n"+
		"ORGANIZER;%s\n"+
		"SUMMARY:%s\n"+
		"GEO:%s\n"+
		"END:VEVENT",
		e.UUID.String(),
		e.Creation.Format("20060102T150405Z"),
		e.Begin.Format("20060102T150405Z"),
		e.End.Format("20060102T150405Z"),
		e.Organizer.String(),
		e.Summary,
		e.Geo.StringSep(";"),
	)
}

type Organizer struct {
	Name  string
	Email string
	LDAP  string
}

type OrganizerBuilder struct {
	organizer *Organizer
}

func NewOrganizerBuilder() *OrganizerBuilder {
	return &OrganizerBuilder{
		organizer: &Organizer{},
	}
}

func (b *OrganizerBuilder) Name(name string) *OrganizerBuilder {
	b.organizer.Name = name
	return b
}

func (b *OrganizerBuilder) Email(email string) *OrganizerBuilder {
	b.organizer.Email = email
	return b
}

func (b *OrganizerBuilder) LDAP(ldap string) *OrganizerBuilder {
	b.organizer.LDAP = ldap
	return b
}

func (b *OrganizerBuilder) Build() *Organizer {
	return b.organizer
}

func (o Organizer) String() string {
	if o.LDAP == "" {
		return fmt.Sprintf("CN=%s:MAILTO:%s", o.Name, o.Email)
	}

	return fmt.Sprintf("CN=%s;DIR=\"%s\":MAILTO:%s", o.Name, o.LDAP, o.Email)
}
