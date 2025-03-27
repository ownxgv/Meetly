package meetings

import appctx "meetly/internal/context"

type MeetingRepository interface {
	GetAllMeetings() ([]Meeting, error)
	GetMeetingByID(id uint) (*Meeting, error)
	CreateMeeting(meeting *Meeting) error
	UpdateMeeting(meeting *Meeting) error
	DeleteMeeting(id uint) error
}

type MeetingService interface {
	GetAllMeetings() ([]Meeting, error)
	GetMeetingByID(id uint) (*Meeting, error)
	CreateMeeting(meeting *Meeting) error
	UpdateMeeting(meeting *Meeting) error
	DeleteMeeting(id uint) error
}

type MeetingHandler interface {
	GetAllMeetings(c appctx.Context)
	CreateMeeting(c appctx.Context)
	GetMeetingByID(c appctx.Context)
	UpdateMeeting(c appctx.Context)
	DeleteMeeting(c appctx.Context)
}
