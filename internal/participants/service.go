package participants

type participantService struct {
	repo ParticipantRepository
}

func NewParticipantService(repo ParticipantRepository) ParticipantService {
	return &participantService{repo: repo}
}

func (s *participantService) GetAllParticipants() ([]Participant, error) {
	return s.repo.GetAllParticipants()
}

func (s *participantService) GetParticipantsByMeetingID(meetingID uint) ([]Participant, error) {
	return s.repo.GetParticipantsByMeetingID(meetingID)
}

func (s *participantService) AddParticipant(participant *Participant) error {
	return s.repo.AddParticipant(participant)
}

func (s *participantService) UpdateParticipantStatus(meetingID, userID uint, status string) error {
	return s.repo.UpdateParticipantStatus(meetingID, userID, status)
}

func (s *participantService) RemoveParticipant(meetingID, userID uint) error {
	return s.repo.RemoveParticipant(meetingID, userID)
}
