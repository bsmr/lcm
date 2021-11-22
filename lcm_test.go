package lcm

import (
	"testing"

	"github.com/bsmr/lcm/page"
	"github.com/bsmr/lcm/widgets/label"
)

func TestExampleApp(t *testing.T) {

	app, err := application.New(
		page.New(
			"login",
			"Test Application",
			table.New(
				label.New("title", "Test",
					table.RowSpan(2),
				),
				labelText.New("Login:"),
				labelPassword.New("Password:"),
				spacer.New(),
				button.New(
					"submit",
					"Submit",
					action.Clicked(application.Page("hello")),
				),
				button.New(
					"cancel",
					"Cancel",
					action.Clicked(application.Quit),
				),
			),
		),
	)

	if err != nil {
		t.Fatalf("application.New() failed with: %s")
	}

	if err := app.Run(); err != nil {
		t.Errorf("app.Run() failed with: %s", err)
	}
}
