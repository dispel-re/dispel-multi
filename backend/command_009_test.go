package backend

import (
	"testing"

	"connectrpc.com/connect"
	v1 "github.com/dispel-re/dispel-multi/gen/multi/v1"
	"github.com/dispel-re/dispel-multi/model"
	"github.com/stretchr/testify/assert"
)

func TestListGamesRequest(t *testing.T) {
	// Arrange
	packet := []byte{
		255, 9,
		4, 0,
	}

	// Act
	req := ListChannelsRequest(packet[4:])

	// Assert
	assert.Empty(t, req)
}

func TestBackend_HandleListGames(t *testing.T) {
	t.Run("no games", func(t *testing.T) {
		b := &Backend{GameClient: &mockGameClient{
			ListGamesResponse: connect.NewResponse(&v1.ListGamesResponse{Games: []*v1.Game{}}),
		}}
		conn := &mockConn{}
		session := &model.Session{ID: "TEST", Conn: conn, UserID: 2137, Username: "JP"}

		assert.NoError(t, b.HandleListGames(session, ListGamesRequest{}))
		assert.Len(t, conn.Written, 8)
		assert.Equal(t, []byte{255, 9, 8, 0}, conn.Written[0:4]) // Header
		assert.Equal(t, []byte{0, 0, 0, 0}, conn.Written[4:8])   // Number of games
	})

	t.Run("with games", func(t *testing.T) {
		b := &Backend{GameClient: &mockGameClient{
			ListGamesResponse: connect.NewResponse(&v1.ListGamesResponse{Games: []*v1.Game{
				{
					GameId:        1,
					Name:          "RoomName",
					Password:      "secret",
					HostIpAddress: "127.0.0.1",
					MapId:         2,
				},
				{
					GameId:        1,
					Name:          "Other",
					Password:      "",
					HostIpAddress: "127.0.0.1",
					MapId:         5,
				},
			}}),
		}}
		conn := &mockConn{}
		session := &model.Session{ID: "TEST", Conn: conn, UserID: 2137, Username: "JP"}

		assert.NoError(t, b.HandleListGames(session, ListGamesRequest{}))
		assert.Len(t, conn.Written, 39)
		assert.Equal(t, []byte{255, 9, 39, 0}, conn.Written[0:4])    // Header
		assert.Equal(t, []byte{2, 0, 0, 0}, conn.Written[4:8])       // Number of games
		assert.Equal(t, []byte{127, 0, 0, 1}, conn.Written[8:12])    // Host IP Address
		assert.Equal(t, []byte("RoomName\x00"), conn.Written[12:21]) // Room name
		assert.Equal(t, []byte("secret\x00"), conn.Written[21:28])   // Password
		assert.Equal(t, []byte{127, 0, 0, 1}, conn.Written[28:32])   // Host IP Address
		assert.Equal(t, []byte("Other\x00"), conn.Written[32:38])    // Room name
		assert.Equal(t, []byte("\x00"), conn.Written[38:39])         // Password
	})
}
