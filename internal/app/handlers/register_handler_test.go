package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"patreon/internal/models"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type LoginTestSuite struct {
	SuiteTestStore
}

func TestRegisterHandler(t *testing.T) {
	suite.Run(t, new(SuiteTestStore))
}

func (s *SuiteTestStore) TestRegisterHandler_ServeHTTP_EmptyBody() {
	s.test = TestTable{
		name:              "Empty body from request",
		data:              &models.RequestRegistration{},
		expectedMockTimes: 0,
		expectedCode:      http.StatusUnprocessableEntity,
	}
	recorder := httptest.NewRecorder()
	handler := NewRegisterHandler()
	logger := logrus.New()
	str := bytes.Buffer{}
	logger.SetOutput(&str)

	handler.SetLogger(logger)

	b := bytes.Buffer{}
	err := json.NewEncoder(&b).Encode(s.test.data)

	assert.NoError(s.T(), err)

	reader, _ := http.NewRequest(http.MethodPost, "/register", &b)
	handler.ServeHTTP(recorder, reader)
	assert.Equal(s.T(), s.test.expectedCode, recorder.Code)
}

func (s *SuiteTestStore) TestRegisterHandler_ServeHTTP_InvalidBody() {
	s.test = TestTable{
		name:              "Invalid body",
		expectedMockTimes: 0,
		expectedCode:      http.StatusUnprocessableEntity,
	}
	data := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{
		Login:    "nickname",
		Password: "password",
	}
	recorder := httptest.NewRecorder()
	handler := NewRegisterHandler()
	logger := logrus.New()
	str := bytes.Buffer{}
	logger.SetOutput(&str)

	handler.SetLogger(logger)

	b := bytes.Buffer{}
	err := json.NewEncoder(&b).Encode(data)

	assert.NoError(s.T(), err)

	reader, _ := http.NewRequest(http.MethodPost, "/register", &b)
	handler.ServeHTTP(recorder, reader)
	assert.Equal(s.T(), s.test.expectedCode, recorder.Code)
}
func (s *SuiteTestStore) TestRegisterHandler_ServeHTTP_UserAlreadyExist() {
	s.test = TestTable{
		name: "User exist in database",
		data: models.RequestRegistration{
			Login:    "dmitriy",
			Nickname: "linux1998",
			Password: "mail.ru",
		},
		expectedMockTimes: 1,
		expectedCode:      http.StatusConflict,
	}
	recorder := httptest.NewRecorder()
	handler := NewRegisterHandler()
	logger := logrus.New()
	str := bytes.Buffer{}
	logger.SetOutput(&str)

	handler.SetLogger(logger)

	handler.SetStore(s.store)

	req := s.test.data.(models.RequestRegistration)
	user := &models.User{
		Login:    req.Login,
		Nickname: req.Nickname,
		Password: req.Password,
	}

	s.mockUserRepository.EXPECT().
		FindByLogin(user.Login).
		Times(s.test.expectedMockTimes).
		Return(user, nil)

	b := bytes.Buffer{}
	err := json.NewEncoder(&b).Encode(s.test.data)

	assert.NoError(s.T(), err)

	reader, _ := http.NewRequest(http.MethodPost, "/register", &b)
	handler.ServeHTTP(recorder, reader)
	assert.Equal(s.T(), s.test.expectedCode, recorder.Code)
}
func (s *SuiteTestStore) TestRegisterHandler_ServeHTTP_SmallPassword() {
	s.test = TestTable{
		name: "Small password in request",
		data: models.RequestRegistration{
			Login:    "dmitriy",
			Nickname: "linux1998",
			Password: "mail",
		},
		expectedMockTimes: 1,
		expectedCode:      http.StatusBadRequest,
	}
	recorder := httptest.NewRecorder()
	handler := NewRegisterHandler()
	logger := logrus.New()
	str := bytes.Buffer{}
	logger.SetOutput(&str)

	handler.SetLogger(logger)

	handler.SetStore(s.store)

	req := s.test.data.(models.RequestRegistration)
	user := &models.User{
		Login:    req.Login,
		Nickname: req.Nickname,
		Password: req.Password,
	}

	s.mockUserRepository.EXPECT().
		FindByLogin(user.Login).
		Times(s.test.expectedMockTimes).
		Return(nil, nil)

	b := bytes.Buffer{}
	err := json.NewEncoder(&b).Encode(s.test.data)

	assert.NoError(s.T(), err)

	reader, _ := http.NewRequest(http.MethodPost, "/register", &b)
	handler.ServeHTTP(recorder, reader)
	assert.Equal(s.T(), s.test.expectedCode, recorder.Code)
}
func (s *SuiteTestStore) TestRegisterHandler_ServeHTTP_CreateSuccess() {
	s.test = TestTable{
		name: "Success create user",
		data: models.RequestRegistration{
			Login:    "dmitriy",
			Password: "mail.ru",
			Nickname: "linux1998",
		},
		expectedMockTimes: 1,
		expectedCode:      http.StatusOK,
	}
	recorder := httptest.NewRecorder()
	handler := NewRegisterHandler()
	logger := logrus.New()
	str := bytes.Buffer{}
	logger.SetOutput(&str)

	handler.SetLogger(logger)

	req := s.test.data.(models.RequestRegistration)
	user := &models.User{
		Login:    req.Login,
		Password: req.Password,
		Nickname: req.Nickname,
	}

	b := bytes.Buffer{}
	err := json.NewEncoder(&b).Encode(s.test.data)
	assert.NoError(s.T(), err)

	s.mockUserRepository.EXPECT().
		FindByLogin(user.Login).
		Times(s.test.expectedMockTimes).
		Return(nil, nil)

	assert.NoError(s.T(), user.BeforeCreate())

	s.mockUserRepository.EXPECT().Create(user).Return(nil).Times(1)

	handler.SetStore(s.store)
	reader, _ := http.NewRequest(http.MethodPost, "/register", &b)
	handler.ServeHTTP(recorder, reader)
	//assert.Equal(s.T(), s.test.expectedCode, recorder.Code)
}
