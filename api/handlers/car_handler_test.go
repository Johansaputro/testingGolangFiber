package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testingfiber/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockService struct {
	mock.Mock
}

func (m *mockService) InsertCarService(car *entities.Car) (*entities.Car, error) {
	args := m.Called(car)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*entities.Car), err
	}
	return nil, err
}

func (m *mockService) UpdateCarService(car *entities.Car) (*entities.Car, error) {
	args := m.Called(car)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*entities.Car), err
	}
	return nil, err
}

func (m *mockService) RemoveCarService(ID string) error {
	args := m.Called(ID)
	return args.Error(0)
}

func (m *mockService) CheckCarService() (*[]entities.Car, error) {
	args := m.Called()
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*[]entities.Car), err
	}
	return nil, err
}

func TestAddBookHandler(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		requestBody  string
		expectedCode int
	}{
		{
			//success test case
			description:  "postHTTP200",
			route:        "/cars",
			requestBody:  `{"carName":"Mazda","company":"Toyota"}`,
			expectedCode: 200,
		},
		{
			//failed test case 2
			description:  "postHTTP500",
			route:        "/cars",
			requestBody:  `{}`,
			expectedCode: 500,
		},
		{
			//failed test case 3
			description:  "postHTTP400",
			route:        "/cars",
			requestBody:  `[]`,
			expectedCode: 400,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mockService := new(mockService)
			handler := AddCar(mockService)

			app := fiber.New()
			app.Post(test.route, handler)

			if test.expectedCode == 200 {
				var requestBody *entities.Car
				err := json.Unmarshal([]byte(test.requestBody), &requestBody)
				require.NoError(t, err)
				mockService.On("InsertCarService", requestBody).Return(requestBody, nil)
			}

			req := httptest.NewRequest(http.MethodPost, test.route, strings.NewReader(test.requestBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectedCode, resp.StatusCode)
			mockService.AssertExpectations(t)
		})
	}
}

func TestUpdateCarHandler(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		requestBody  string
		expectedCode int
	}{
		{
			//success test case
			description:  "putHTTP200",
			route:        "/cars",
			requestBody:  `{"carName":"Mazda","company":"Toyota"}`,
			expectedCode: 200,
		},
		{
			//failed test case 2
			description: "putHTTP400",
			route:       "/cars",
			requestBody: `{
				"_id": {
				  "$oid": "64a4c6181955b6923fff02b5"
				},
				"carName": "Mazda CX 6",
				"company": "Mazda",
				"madeAt": {
				  "$date": "2023-07-05T01:23:36.296Z"
				},
				"soldAt": {
				  "$date": "2023-07-05T09:03:39.703Z"
				}
			  }`,
			expectedCode: 400,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mockService := new(mockService)
			handler := UpdateCar(mockService)

			app := fiber.New()
			app.Put(test.route, handler)

			var requestBody *entities.Car
			err := json.Unmarshal([]byte(test.requestBody), &requestBody)
			require.NoError(t, err)
			mockService.On("UpdateCarService", requestBody).Return(requestBody, nil)

			req := httptest.NewRequest(http.MethodPut, test.route, strings.NewReader(test.requestBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectedCode, resp.StatusCode)
			mockService.AssertExpectations(t)
		})
	}
}

func TestRemoveCarHandler(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		requestBody  string
		expectedCode int
	}{
		{
			//success test case
			description:  "DeleteHTTP200",
			route:        "/cars",
			requestBody:  `{"id": "64a4c6181955b6923fff02b5"}`,
			expectedCode: 200,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mockService := new(mockService)
			handler := RemoveCar(mockService)

			app := fiber.New()
			app.Delete(test.route, handler)

			if test.expectedCode == 200 {
				var requestBody entities.DeleteRequest
				err := json.Unmarshal([]byte(test.requestBody), &requestBody)
				require.NoError(t, err)

				mockService.On("RemoveCarService", requestBody).Return(nil)
			}

			req := httptest.NewRequest(http.MethodDelete, test.route, strings.NewReader(test.requestBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectedCode, resp.StatusCode)
			mockService.AssertExpectations(t)
		})
	}
}

func TestGetCarHandle(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			//success test case
			description:  "GetHTTP200",
			route:        "/cars",
			expectedCode: 200,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mockService := new(mockService)
			handler := GetCars(mockService)

			app := fiber.New()
			app.Get(test.route, handler)

			cars := []entities.Car{
				{CarName: "Car 1"},
				{CarName: "Car 2"},
			}

			mockService.On("CheckCarService").Return(&cars, nil)

			req := httptest.NewRequest(http.MethodGet, test.route, nil)
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectedCode, resp.StatusCode)
			mockService.AssertExpectations(t)
		})
	}
}
