package models

import "time"

type InputJsonStructJob struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Position    string    `json:"position"`
	Url         string    `json:"url"`
	Location    string    `json:"location"`
	Salary      string    `json:"salary"`
	Type        uint      `json:"type"`
	Apply       bool      `json:"apply"`
	CreatedAt   time.Time `json:"created_at"`
	JobReaction bool      `json:"job_reaction"`
	JobReject   bool      `json:"job_reject"`
}
