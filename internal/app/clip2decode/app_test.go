package clip2decode_test

import (
	"github/tr3mor/clip2decode/internal/app/clip2decode"
	"github/tr3mor/clip2decode/internal/app/clip2decode/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type AppTestSuite struct {
	suite.Suite

	controller *gomock.Controller
	clipboard  *mocks.MockClipboardInterface
	app        *clip2decode.Application
}

func (s *AppTestSuite) SetupTest() {
	s.controller = gomock.NewController(s.T())
	s.clipboard = mocks.NewMockClipboardInterface(s.controller)
	s.app = clip2decode.NewApplication(s.clipboard)
}

func (s *AppTestSuite) TestCheckDecode() {
	s.clipboard.EXPECT().GetData().Return("aGVsbG8=", nil).Times(1)
	s.clipboard.EXPECT().WriteData("hello").Return(nil).Times(1)
	data, err := s.app.DecodeData()
	s.Equal("hello", data)
	s.NoError(err)
}

func (s *AppTestSuite) TestIncorrectBase64() {
	s.clipboard.EXPECT().GetData().Return("notBase64", nil).Times(1)
	data, err := s.app.DecodeData()
	s.Error(err)
	s.Equal("Failed to decode base64", data)
}
func TestC2DSuite(t *testing.T) {
	suite.Run(t, new(AppTestSuite))
}
