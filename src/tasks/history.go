package tasks

import (
	"errors"
	"time"
)

// Event describes change of status
type Event struct {
	Status string
	Time   time.Time
}

// History contains list of events
type History struct {
	Events []Event
}

// LogEvent writes new event to history
func (history *History) LogEvent(status string) error {
	if history.isNotEmpty() && history.isInStatus(status) {
		return errors.New("This task already has " + messages[status] + " status")
	}

	event := Event{Status: status, Time: getCurrentDateTime()}

	history.Events = append(history.Events, event)

	return nil
}

// GetLoggedTime returns logged time in seconds
func (history History) GetLoggedTime() int {
	history = history.fillHistoryGap()

	loggedTime := 0

	for i := 0; i < len(history.Events); i += 2 {
		startTime := history.Events[i].Time
		endTime := history.Events[i+1].Time
		timeDiff := endTime.Sub(startTime)

		loggedTime += int(timeDiff.Seconds())
	}

	return loggedTime
}

func (history History) isNotEmpty() bool {
	return len(history.Events) != 0
}

func (history History) isInStatus(status string) bool {
	length := len(history.Events)

	if length == 0 {
		return false
	}

	return history.Events[length-1].Status == status
}

func (history History) fillHistoryGap() History {
	if history.isInStatus("start") {
		event := Event{Status: "tmp", Time: getCurrentDateTime()}

		history.Events = append(history.Events, event)
	}

	return history
}
