package services

import (
	"os"
	"testing"
)

var client *Service

func TestMain(m *testing.M) {
	client = NewService()
	exit := m.Run()
	os.Exit(exit)
}

func TestService_BookPortalCategory(t *testing.T) {
	//param := ClassifyBookParam{
	//	BusinessType: 1,
	//	PageNo:       1,
	//	PageSize:     15,
	//}
	resp, err := client.BookModuleContent(400114202, 400133812)
	if err != nil {
		t.Log(err)
	}

	t.Log(resp)
}
