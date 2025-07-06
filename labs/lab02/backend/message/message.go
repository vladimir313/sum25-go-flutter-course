package message

import (
	"sync"
)

// Message represents a chat message
// TODO: Add more fields if needed

type Message struct {
	Sender    string
	Content   string
	Timestamp int64
}

// MessageStore stores chat messages
// Contains a slice of messages and a mutex for concurrency

type MessageStore struct {
	messages []Message
	mutex    sync.RWMutex
	// TODO: Add more fields if needed
}

// NewMessageStore creates a new MessageStore
func NewMessageStore() *MessageStore {
	// TODO: Initialize MessageStore fields
	return &MessageStore{
		messages: make([]Message, 0, 100),
	}
}

// AddMessage stores a new message
func (s *MessageStore) AddMessage(msg Message) error {
	// TODO: Add message to storage (concurrent safe)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.messages = append(s.messages, msg)
	return nil
}

// GetMessages retrieves messages (optionally by user)
func (s *MessageStore) GetMessages(user string) ([]Message, error) {
	// TODO: Retrieve messages (all or by user)
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if user == "" {
		// Return all messages
		result := make([]Message, len(s.messages))
		copy(result, s.messages)
		return result, nil
	}

	// Return messages by specific user
	var result []Message
	for _, msg := range s.messages {
		if msg.Sender == user {
			result = append(result, msg)
		}
	}
	return result, nil
}
