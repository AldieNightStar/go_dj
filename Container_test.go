package go_dj

import (
	"testing"
)

func TestContainer(t *testing.T) {
	c := NewContainer()
	s := &tService{}

	c.Register("controller", func(args ... interface{}) interface{} {
		return &tController{args[0].(*tService)}
	}, "service")

	c.Register("service", func(args ... interface{}) interface{} {
		return s
	})

	ctrl1, _ := c.Provide("controller")
	ctrl := ctrl1.(*tController)
	ctrl.Send("Test!")

	if s.LatestText != "Test!" {
		t.Fail()
	}

	_, err := c.Provide("NonExisten")
	if err == nil {
		t.Fail()
	}
}

// ======

type tController struct {
	Service *tService
}

func (t *tController) Send(text string) {
	t.Service.Serve(text)
}

// ======

type tService struct {
	LatestText string
}

func (t *tService) Serve(text string) {
	t.LatestText = text
}
