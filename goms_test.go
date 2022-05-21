package goms

import (
	"errors"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		/* year */
		{
			name:    "year - y",
			args:    args{val: "1y"},
			want:    float64(31536000000),
			wantErr: false,
		},
		{
			name:    "year - year",
			args:    args{val: "1year"},
			want:    float64(31536000000),
			wantErr: false,
		},
		{
			name:    "year - years",
			args:    args{val: "2years"},
			want:    float64(2 * 31536000000),
			wantErr: false,
		},
		{
			name:    "year - yrs",
			args:    args{val: "2yrs"},
			want:    float64(2 * 31536000000),
			wantErr: false,
		},

		/* day */
		{
			name:    "day - d",
			args:    args{val: "1d"},
			want:    float64(86400000),
			wantErr: false,
		},
		{
			name:    "day - day",
			args:    args{val: "1day"},
			want:    float64(86400000),
			wantErr: false,
		},
		{
			name:    "days - days",
			args:    args{val: "2days"},
			want:    float64(2 * 86400000),
			wantErr: false,
		},

		/* week */
		{
			name:    "week - week",
			args:    args{val: "1week"},
			want:    float64(604800000),
			wantErr: false,
		},
		{
			name:    "week - w",
			args:    args{val: "1w"},
			want:    float64(604800000),
			wantErr: false,
		},

		/* hour */
		{
			name:    "hour - h",
			args:    args{val: "1h"},
			want:    float64(3600000),
			wantErr: false,
		},
		{
			name:    "hour - hour",
			args:    args{val: "1hour"},
			want:    float64(3600000),
			wantErr: false,
		},
		{
			name:    "hours - hours",
			args:    args{val: "2hour"},
			want:    float64(2 * 3600000),
			wantErr: false,
		},
		{
			name:    "hours - hrs",
			args:    args{val: "2hrs"},
			want:    float64(2 * 3600000),
			wantErr: false,
		},

		/* minute */
		{
			name:    "minute - m",
			args:    args{val: "1m"},
			want:    float64(60000),
			wantErr: false,
		},
		{
			name:    "minute - minute",
			args:    args{val: "1minute"},
			want:    float64(60000),
			wantErr: false,
		},
		{
			name:    "minutes - minutes",
			args:    args{val: "2minutes"},
			want:    float64(2 * 60000),
			wantErr: false,
		},
		{
			name:    "min - min",
			args:    args{val: "1min"},
			want:    float64(60000),
			wantErr: false,
		},
		{
			name:    "mins - mins",
			args:    args{val: "2min"},
			want:    float64(2 * 60000),
			wantErr: false,
		},

		/* second */
		{
			name:    "second - s",
			args:    args{val: "1s"},
			want:    float64(1000),
			wantErr: false,
		},
		{
			name:    "second - second",
			args:    args{val: "1second"},
			want:    float64(1000),
			wantErr: false,
		},
		{
			name:    "seconds - seconds",
			args:    args{val: "2seconds"},
			want:    float64(2 * 1000),
			wantErr: false,
		},
		{
			name:    "seconds - s",
			args:    args{val: "2s"},
			want:    float64(2 * 1000),
			wantErr: false,
		},

		/* ms */
		{
			name:    "millsecond - ms",
			args:    args{val: "1ms"},
			want:    float64(1),
			wantErr: false,
		},
		{
			name:    "millseconds - ms",
			args:    args{val: "2ms"},
			want:    float64(2),
			wantErr: false,
		},

		/* error */
		{
			name:    "month - month",
			args:    args{val: "1M"},
			want:    float64(0),
			wantErr: true,
		},
		{
			name:    "month - month",
			args:    args{val: "1month"},
			want:    float64(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isErrUnsupportedUnitError(t *testing.T) {
	err1 := errors.New("some error")
	if errors.Is(err1, ErrUnsupportedUnit) {
		t.Fail()
	}

	err2 := ErrUnsupportedUnit
	if !errors.Is(err2, ErrUnsupportedUnit) {
		t.Fail()
	}
}
