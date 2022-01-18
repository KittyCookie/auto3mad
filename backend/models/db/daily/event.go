package daily

import (
	"backend/models/db/base"
	"backend/models/util"
	"fmt"
)

type Event struct {
	ID            int `orm:"column(id)"`
	StartTime     int64
	EndTime       int64
	SpecificEvent string
	RoutineId     int
}

func (o *Event) TableName() string {
	return "daily_time_use"
}

func (o *Event) GetID() int {
	return o.ID
}

func (o *Event) NewObjectOnlyID(id int) interface{} {
	ooid := new(Event)
	ooid.ID = id
	return ooid
}

type EventModel struct {
	base.BaseModel
}

func NewEventModel() *EventModel {
	m := new(EventModel)
	m.BaseModel = *base.NewBaseModel(&Event{})
	return m
}

func (m *EventModel) GetEventByDate(date string) (events []Event, err error) {
	firstSecond, lastSecond := util.GetDateTimestamp(date)
	sql := fmt.Sprintf("SELECT start_time, end_time, specific_event, routine_id FROM %s WHERE start_time BETWEEN %d AND %d ORDER BY start_time DESC", m.TableName, firstSecond, lastSecond)
	_, err = m.ORM.Raw(sql).QueryRows(&events)
	return
}

type TimeUse struct {
	RoutineId int
	Use       int
}

func (m *EventModel) GetTimeUseGroupByRoutine(date string) (uses []TimeUse, err error) {
	firstSecond, lastSecond := util.GetDateTimestamp(date)
	sql := fmt.Sprintf("SELECT routine_id, ROUND(SUM(end_time-start_time)/60) as `use` FROM %s WHERE start_time BETWEEN %d AND %d GROUP BY routine_id", m.TableName, firstSecond, lastSecond)
	_, err = m.ORM.Raw(sql).QueryRows(&uses)
	return
}

func (m *EventModel) GetAllTimeUseGroupByRoutine() (uses []TimeUse, err error) {
	sql := fmt.Sprintf("SELECT routine_id, ROUND(SUM(end_time-start_time)/3600) as `use` FROM %s GROUP BY routine_id", m.TableName)
	_, err = m.ORM.Raw(sql).QueryRows(&uses)
	return
}
