package messagebird

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Recording struct {
	ID         string    `json:"id"`
	Format     string    `json:"format"`
	Duration   float64   `json:"duration"`
	Status     string    `json:"status"`
	CreateAt   time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	StartedAt  time.Time `json:"startedAt"`
	FinishedAt time.Time `json:"finishedAt"`
}

type RecordingList struct {
	Pagination struct {
		PerPage     int `json:"perPage"`
		CurrentPage int `json:"currentPage"`
		PageCount   int `json:"pageCount"`
		TotalCount  int `json:"totalCount"`
	} `json:"pagination"`
	Data []Recording `json:"data"`
}

type Call struct {
	ID          string    `json:"id"`
	CallID      string    `json:"callId"`
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	Hold        bool      `json:"hold"`
	Direction   string    `json:"direction"`
	Cost        float64   `json:"cost"`
	Currency    string    `json:"currency"`
	Duration    float64   `json:"duration"`
	Status      string    `json:"status"`
	CreateAt    time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	AnsweredAt  time.Time `json:"answeredAt"`
	EndedAt     time.Time `json:"endedAt"`
}

type CallList struct {
	Pagination struct {
		PerPage     int `json:"perPage"`
		CurrentPage int `json:"currentPage"`
		PageCount   int `json:"pageCount"`
		TotalCount  int `json:"totalCount"`
	} `json:"pagination"`
	Data []Call `json:"data"`
}

type CallItem struct {
	Data []Call `json:"data"`
}

func (messageBird *messageBird[MessageContent, NewMessageCreated]) GetAllCalls() (calls *CallList, err error) {
	responseBody, err := messageBird.apiRequest("GET", "voice", "/calls/", nil, nil)
	if err != nil {
		return calls, err
	}
	err = json.Unmarshal(responseBody, &calls)
	if err != nil {
		return calls, err
	}
	return calls, err
}

func (messageBird *messageBird[MessageContent, NewMessageCreated]) GetCall(id string) (call *CallItem, err error) {
	path := fmt.Sprintf("/calls/%s/legs", id)
	responseBody, err := messageBird.apiRequest("GET", "voice", path, nil, nil)
	if err != nil {
		return call, err
	}
	err = json.Unmarshal(responseBody, &call)
	if err != nil {
		return call, err
	}
	return call, err
}

func (messageBird *messageBird[MessageContent, NewMessageCreated]) GetCallRecording(id string) (recording *RecordingList, err error) {
	leg, err := messageBird.GetCall(id)
	if err != nil {
		return recording, err
	}
	if len(leg.Data) == 0 {
		return recording, errors.New("none leg found")
	}
	legID := leg.Data[0].ID
	path := fmt.Sprintf("/calls/%s/legs/%s/recordings", id, legID)
	responseBody, err := messageBird.apiRequest("GET", "voice", path, nil, nil)
	if err != nil {
		return recording, err
	}
	err = json.Unmarshal(responseBody, &recording)
	if err != nil {
		return recording, err
	}
	return recording, err
}

func (messageBird *messageBird[MessageContent, NewMessageCreated]) GetCallRecordingFile(id string, format string) (recordingFile *[]byte, err error) {
	recording, err := messageBird.GetCallRecording(id)
	if err != nil {
		return recordingFile, err
	}
	if len(recording.Data) == 0 {
		return recordingFile, errors.New("none recording found")
	}
	recordingID := recording.Data[0].ID
	path := fmt.Sprintf("/recordings/%s.%s", recordingID, format)
	responseBody, err := messageBird.apiRequest("GET", "voice", path, nil, nil)
	if err != nil {
		return recordingFile, err
	}
	return &responseBody, err
}

func (messageBird *messageBird[MessageContent, NewMessageCreated]) GetCallRecordingTranscription(id string) (recordingTranscription *string, err error) {
	recording, err := messageBird.GetCallRecording(id)
	if err != nil {
		return recordingTranscription, err
	}
	if len(recording.Data) == 0 {
		return recordingTranscription, errors.New("none recording found")
	}
	recordingID := recording.Data[0].ID
	path := fmt.Sprintf("/recordings/%s/transcriptions", recordingID)
	responseBody, err := messageBird.apiRequest("GET", "voice", path, nil, nil)
	if err != nil {
		return recordingTranscription, err
	}
	err = json.Unmarshal(responseBody, &recording)
	if err != nil {
		return recordingTranscription, err
	}
	recordingID = recording.Data[0].ID
	path = fmt.Sprintf("/transcriptions/%s.txt", recordingID)
	responseBody, err = messageBird.apiRequest("GET", "voice", path, nil, nil)
	if err != nil {
		return recordingTranscription, err
	}
	transcription := string(responseBody)
	return &transcription, err
}
