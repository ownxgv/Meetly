package meetings

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

type HTTPContext interface {
	JSON(code int, obj interface{}) error
	BindJSON(obj interface{}) error
	Param(s string) string
}

type MeetingHandler interface {
	GetAllMeetings(c HTTPContext)
	CreateMeeting(c HTTPContext)
	GetMeetingByID(c HTTPContext)
	UpdateMeeting(c HTTPContext)
	DeleteMeeting(c HTTPContext)
}
