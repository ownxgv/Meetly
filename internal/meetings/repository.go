package meetings

import "gorm.io/gorm"

type meetingRepository struct {
	db *gorm.DB
}

func NewMeetingRepository(db *gorm.DB) MeetingRepository {
	return &meetingRepository{db: db}
}

func (r *meetingRepository) GetAllMeetings() ([]Meeting, error) {
	var meetings []Meeting
	if err := r.db.Preload("Creator").Find(&meetings).Error; err != nil {
		return nil, err
	}
	return meetings, nil
}

func (r *meetingRepository) GetMeetingByID(id uint) (*Meeting, error) {
	var meeting Meeting
	if err := r.db.Preload("Creator").First(&meeting, id).Error; err != nil {
		return nil, err
	}
	return &meeting, nil
}

func (r *meetingRepository) CreateMeeting(meeting *Meeting) error {
	return r.db.Create(meeting).Error
}

func (r *meetingRepository) UpdateMeeting(meeting *Meeting) error {
	return r.db.Save(meeting).Error
}

func (r *meetingRepository) DeleteMeeting(id uint) error {
	return r.db.Delete(&Meeting{}, id).Error
}
