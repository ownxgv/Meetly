package participants

import "gorm.io/gorm"

type participantRepository struct {
	db *gorm.DB
}

func NewParticipantRepository(db *gorm.DB) ParticipantRepository {
	return &participantRepository{db: db}
}

func (r *participantRepository) GetAllParticipants() ([]Participant, error) {
	var participants []Participant
	if err := r.db.Find(&participants).Error; err != nil {
		return nil, err
	}
	return participants, nil
}

func (r *participantRepository) GetParticipantsByMeetingID(meetingID uint) ([]Participant, error) {
	var participants []Participant
	if err := r.db.Where("meeting_id = ?", meetingID).Find(&participants).Error; err != nil {
		return nil, err
	}
	return participants, nil
}

func (r *participantRepository) AddParticipant(participant *Participant) error {
	return r.db.Create(participant).Error
}

func (r *participantRepository) UpdateParticipantStatus(meetingID, userID uint, status string) error {
	return r.db.Model(&Participant{}).
		Where("meeting_id = ? AND user_id = ?", meetingID, userID).
		Update("status", status).Error
}

func (r *participantRepository) RemoveParticipant(meetingID, userID uint) error {
	return r.db.Where("meeting_id = ? AND user_id = ?", meetingID, userID).
		Delete(&Participant{}).Error
}
