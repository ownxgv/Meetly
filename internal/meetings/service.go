package meetings

type meetingService struct {
	repo MeetingRepository
}

func NewMeetingService(repo MeetingRepository) MeetingService {
	return &meetingService{repo: repo}
}

func (s *meetingService) GetAllMeetings() ([]Meeting, error) {
	return s.repo.GetAllMeetings()
}

func (s *meetingService) GetMeetingByID(id uint) (*Meeting, error) {
	return s.repo.GetMeetingByID(id)
}

func (s *meetingService) CreateMeeting(meeting *Meeting) error {
	return s.repo.CreateMeeting(meeting)
}

func (s *meetingService) UpdateMeeting(meeting *Meeting) error {
	return s.repo.UpdateMeeting(meeting)
}

func (s *meetingService) DeleteMeeting(id uint) error {

	return s.repo.DeleteMeeting(id)
}
