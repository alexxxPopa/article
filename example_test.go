package article

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock is the workhorse used to track activity on another object.
type MockedConnection struct {
	mock.Mock
}

func (mock *MockedConnection) isTitleUnique(title string) bool {
	args := mock.Called(title)
	return args.Bool(0)
}


func TestApi(t *testing.T) {
	conn := new (MockedConnection)
	api :=  API{
		conn: conn,
	}
	conn.On("isTitleUnique","Moby-Dick").Return(true)

	assert.Equal(t,api.isDuplicateTitle("Moby-Dick"), "No good")
}