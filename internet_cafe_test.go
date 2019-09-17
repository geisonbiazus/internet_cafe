package internet_cafe

import (
	"reflect"
	"testing"
)

func TestInternetCafe(t *testing.T) {
	t.Run("When there is 1 user for 15 minutes", func(t *testing.T) {
		randomizer := &FakeRandomizer{15}

		cafe := NewInternetCafe(1, randomizer)
		events := cafe.Start(1)

		if !reflect.DeepEqual(events, []string{
			"Tourist 1 is online",
			"Tourist 1 is done, having spent 15 minutes online.",
			"The place is empty, let's close up and go to the beach!",
		}) {
			t.Errorf("It didn't work: %v", events)
		}
	})
	t.Run("When there is 1 user for 20 minutes", func(t *testing.T) {
		randomizer := &FakeRandomizer{20}

		cafe := NewInternetCafe(1, randomizer)
		events := cafe.Start(1)

		if !reflect.DeepEqual(events, []string{
			"Tourist 1 is online",
			"Tourist 1 is done, having spent 20 minutes online.",
			"The place is empty, let's close up and go to the beach!",
		}) {
			t.Errorf("It didn't work: %v", events)
		}
	})

	t.Run("When there is 2 user an 1 computer", func(t *testing.T) {
		randomizer := &FakeRandomizer{20}

		cafe := NewInternetCafe(1, randomizer)
		events := cafe.Start(2)

		if !reflect.DeepEqual(events, []string{
			"Tourist 1 is online",
			"Tourist 2 waiting for turn.",
			"Tourist 1 is done, having spent 20 minutes online.",
			"Tourist 2 is online",
			"Tourist 2 is done, having spent 20 minutes online.",
			"The place is empty, let's close up and go to the beach!",
		}) {
			t.Errorf("It didn't work: %v", events)
		}
	})
}

type FakeRandomizer struct {
	Result int
}

func (f *FakeRandomizer) CalculateTime() int {
	return f.Result
}
