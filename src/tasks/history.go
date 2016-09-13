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
	events []Event
}

// LogEvent writes new event to history
func (history *History) LogEvent(status string) error {
	if history.isNotEmpty() && history.isInStatus(status) {
		return errors.New("This task already has " + messages[status] + " status")
	}

	event := Event{Status: status, Time: getCurrentDateTime()}

	history.events = append(history.events, event)

	return nil
}

// GetLoggedTime returns logged time in seconds
func (history History) GetLoggedTime() (int, error) {
	history = history.fillHistoryGap()

	if len(history.events)%2 != 0 {
		return 0, errors.New("Wrong history records number")
	}

	loggedTime := 0

	for i := 0; i < len(history.events); i += 2 {
		startTime := history.events[i].Time
		endTime := history.events[i+1].Time
		timeDiff := endTime.Sub(startTime)

		loggedTime += int(timeDiff.Seconds())
	}

	return loggedTime, nil
}

func (history History) isNotEmpty() bool {
	return len(history.events) != 0
}

func (history History) isInStatus(status string) bool {
	length := len(history.events)

	if length == 0 {
		return false
	}

	return history.events[length-1].Status == status
}

func (history History) fillHistoryGap() History {
	if history.isInStatus("start") {
		event := Event{Status: "tmp", Time: getCurrentDateTime()}

		history.events = append(history.events, event)
	}

	return history
}
