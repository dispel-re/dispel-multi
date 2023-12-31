package backend

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateCharacterInventoryRequest(t *testing.T) {
	// Arrange
	packet := []byte{
		255, 44, // Command code
		227, 0, // Packet length
		117, 115, 101, 114, 0, // User name
		99, 104, 97, 114, 97, 99, 116, 101, 114, 0, // Character name
		4, 1, 17, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 97, 11, 101, 97, 11, 101, 97, 11, 101, 97, 11, 101, 97, 11, 101, 97, 0, // Inventory (backpack and belt)
	}

	// Act
	req := UpdateCharacterInventoryRequest(packet[4:])
	data, err := req.Parse()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "user", data.Username)
	assert.Equal(t, "character", data.CharacterName)
	assert.Equal(t,
		[]byte{4, 1, 17, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 121, 11, 101, 97, 11, 101, 97, 11, 101, 97, 11, 101, 97, 11, 101, 97, 11, 101, 97},
		data.Inventory,
	)
}
