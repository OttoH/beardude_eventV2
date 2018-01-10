package models

type Events struct {
  *Event
}

type Event struct {
    ManagerIds int    `json:"ManagerIds"`
    UniqueName string `json:"UniqueName"`
    StartTime  string `json:"StartTime"`
    EndTime    string `json:"EndTime"`}
