package httpx

import (
	"context"

	"github.com/stretchr/testify/suite"
)

type TestRequestSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *TestRequestSuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *TestRequestSuite) Test_SetHeader_ShouldRunSuccesfully() {
	// Arrange
	req := NewReq("http://localhost")

	// Act
	req.SetHeader("Content-Type", "application/json")

	// Assert
	s.Equal("application/json", req.Headers["Content-Type"][0])
}
