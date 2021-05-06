package message_test

import (
	"testing"

	"github.com/mythrnr/go-message"
	"github.com/stretchr/testify/assert"
)

func Test_Localizer_Get_and_GetNum(t *testing.T) {
	t.Parallel()

	t.Run("Get", func(t *testing.T) {
		t.Parallel()

		l := message.NewLocalizer(&message.M{
			"0-0": []string{},
			"0-1": "apple",
			"0-2": []string{"no apples", "one apple", "{0} apples"},
		}).Selector(func(n uint) uint { return n })

		tests := []struct{ key, want string }{
			{key: "0-0", want: "0-0"},
			{key: "0-1", want: "apple"},
			{key: "0-2", want: "no apples"},
		}

		for _, tt := range tests {
			t.Log(tt.key, tt.want)
			assert.Equal(t, tt.want, l.Get(tt.key))
		}
	})

	t.Run("GetNum", func(t *testing.T) {
		t.Parallel()

		l := message.NewLocalizer(&message.M{
			"0-0": []string{},
			"0-1": "apple",
			"0-2": []string{"no apples", "one apple", "{0} apples"},
		}).Selector(func(n uint) uint {
			if 2 < n {
				return 2
			}

			return n
		})

		tests := []struct {
			key  string
			num  uint
			want string
		}{
			{key: "0-0", num: 0, want: "0-0"},
			{key: "0-0", num: 1, want: "0-0"},
			{key: "0-1", num: 0, want: "apple"},
			{key: "0-1", num: 1, want: "apple"},
			{key: "0-2", num: 0, want: "no apples"},
			{key: "0-2", num: 1, want: "one apple"},
			{key: "0-2", num: 2, want: "{0} apples"},
			{key: "0-2", num: 3, want: "{0} apples"},
		}

		for _, tt := range tests {
			t.Log(tt.key, tt.want)
			assert.Equal(t, tt.want, l.GetNum(tt.key, tt.num))
		}
	})
}

func Test_Localizer_Getf_and_GetNumf(t *testing.T) {
	t.Parallel()

	t.Run("Getf", func(t *testing.T) {
		t.Parallel()

		l := message.NewLocalizer(&message.M{
			"0-0": []string{},
			"0-1": "apple",
			"0-2": []string{"no apples", "one apple", "{0} apples"},
		}).Selector(func(n uint) uint { return n })

		tests := []struct {
			key  string
			args []interface{}
			want string
		}{
			{key: "0-0", args: []interface{}{1}, want: "0-0"},
			{key: "0-1", args: []interface{}{2}, want: "apple"},
			{key: "0-2", args: []interface{}{3}, want: "no apples"},
		}

		for _, tt := range tests {
			t.Log(tt.key, tt.want)
			assert.Equal(t, tt.want, l.Getf(tt.key, tt.args...))
		}
	})

	t.Run("GetNumf", func(t *testing.T) {
		t.Parallel()

		l := message.NewLocalizer(&message.M{
			"0-0": []string{},
			"0-1": "apple",
			"0-2": []string{"no apples", "one apple", "{0} apples"},
		}).Selector(func(n uint) uint {
			if 2 < n {
				return 2
			}

			return n
		})

		tests := []struct {
			key  string
			num  uint
			args []interface{}
			want string
		}{
			{key: "0-0", num: 0, args: []interface{}{1}, want: "0-0"},
			{key: "0-0", num: 1, args: []interface{}{2}, want: "0-0"},
			{key: "0-1", num: 0, args: []interface{}{1}, want: "apple"},
			{key: "0-1", num: 1, args: []interface{}{2}, want: "apple"},
			{key: "0-2", num: 0, args: []interface{}{1}, want: "no apples"},
			{key: "0-2", num: 1, args: []interface{}{2}, want: "one apple"},
			{key: "0-2", num: 2, args: []interface{}{20}, want: "20 apples"},
			{key: "0-2", num: 3, args: []interface{}{30}, want: "30 apples"},
		}

		for _, tt := range tests {
			t.Log(tt.key, tt.want)
			assert.Equal(t, tt.want, l.GetNumf(tt.key, tt.num, tt.args...))
		}
	})
}

func Test_Localizer_NamedGetf_and_NamedGetNumf(t *testing.T) {
	t.Parallel()

	t.Run("NamedGetf", func(t *testing.T) {
		t.Parallel()

		l := message.NewLocalizer(&message.M{
			"0-0": []string{},
			"0-1": "apple",
			"0-2": []string{"no apples", "one apple", "{count} apples"},
		}).Selector(func(n uint) uint { return n })

		tests := []struct {
			key  string
			rep  message.Replacement
			want string
		}{{
			key:  "0-0",
			rep:  message.Replacement{"count": 1},
			want: "0-0",
		}, {
			key:  "0-1",
			rep:  message.Replacement{"count": 1},
			want: "apple",
		}, {
			key:  "0-2",
			rep:  message.Replacement{"count": 1},
			want: "no apples",
		}}

		for _, tt := range tests {
			t.Log(tt.key, tt.want)
			assert.Equal(t, tt.want, l.NamedGetf(tt.key, tt.rep))
		}
	})

	t.Run("NamedGetNumf", func(t *testing.T) {
		t.Parallel()

		l := message.NewLocalizer(&message.M{
			"0-0": []string{},
			"0-1": "apple",
			"0-2": []string{"no apples", "one apple", "{count} apples"},
		}).Selector(func(n uint) uint {
			if 2 < n {
				return 2
			}

			return n
		})

		tests := []struct {
			key  string
			num  uint
			rep  message.Replacement
			want string
		}{{
			key:  "0-0",
			num:  0,
			rep:  message.Replacement{"count": 1},
			want: "0-0",
		}, {
			key:  "0-0",
			num:  1,
			rep:  message.Replacement{"count": 2},
			want: "0-0",
		}, {
			key:  "0-1",
			num:  0,
			rep:  message.Replacement{"count": 1},
			want: "apple",
		}, {
			key:  "0-1",
			num:  1,
			rep:  message.Replacement{"count": 2},
			want: "apple",
		}, {
			key:  "0-2",
			num:  0,
			rep:  message.Replacement{"count": 1},
			want: "no apples",
		}, {
			key:  "0-2",
			num:  1,
			rep:  message.Replacement{"count": 2},
			want: "one apple",
		}, {
			key:  "0-2",
			num:  2,
			rep:  message.Replacement{"count": 20},
			want: "20 apples",
		}, {
			key:  "0-2",
			num:  3,
			rep:  message.Replacement{"count": 30},
			want: "30 apples",
		}}

		for _, tt := range tests {
			t.Log(tt.key, tt.want)
			assert.Equal(t, tt.want, l.NamedGetNumf(tt.key, tt.num, tt.rep))
		}
	})
}
