package plugin

import "testing"

func TestDatoNotificationURL(t *testing.T) {
	got := DatoNotificationURL("123")
	want := "https://webhooks.datocms.com/123/deploy-results"
	if got != want {
		t.Errorf("Expected %s got %s", want, got)
	}
}

func Test(t *testing.T) {

}
