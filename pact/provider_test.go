package pact

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"testing"
	"web-based-todo-list-backend/controller"
	"web-based-todo-list-backend/mock"
	"web-based-todo-list-backend/models"
	"web-based-todo-list-backend/server"
)

func TestProvider(t *testing.T) {
	port, _ := utils.GetFreePort()
	svr := server.NewServer()

	args := "buy some egg"
	serviceReturnGetTodoList := &models.DataResponse{
		"buy some milk": models.Todo{
			ID:          1,
			Description: "buy some milk",
		},
	}
	serviceReturnAddTodoList := &models.Todo{
		ID:          2,
		Description: "buy some egg",
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockSvc := mock.NewMockIService(ctl)
	mockSvc.
		EXPECT().
		GetTodoList().
		Return(serviceReturnGetTodoList, nil)

	mockSvc.
		EXPECT().
		AddTodoList(args).
		Return(serviceReturnAddTodoList, nil)

	handler := controller.NewTodoListController(mockSvc)

	go svr.StartServer(port, handler)

	pact := dsl.Pact{
		Host:                     "127.0.0.1",
		Provider:                 "Backend",
		Consumer:                 "Frontend",
		DisableToolValidityCheck: true,
	}

	request := types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost:%d", port),
		BrokerURL:                  "https://unsalsenturkk.pactflow.io",
		BrokerToken:                "mCUautWrurN9Z4mTW4WMdA",
		ProviderVersion:            "1.0.0",
		PublishVerificationResults: true,
	}

	verifyResponses, err := pact.VerifyProvider(t, request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(verifyResponses), "pact tests run")
}
