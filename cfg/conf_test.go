package cfg

import (
	"errors"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestConf_GetConf(t *testing.T) {
	type want struct {
		conf *Conf
		err  error
	}
	tests := []struct {
		name string
		file string
		want want
	}{
		{"basic", "nodes:\n  fish: 3c:97:0e:a3:22:40", want{&Conf{map[string]string{"fish": "3c:97:0e:a3:22:40"}}, nil}},
		{"many", "nodes:\n  fish: 3c:97:0e:a3:22:40\n  cow: 3c:97:0e:9e:cf:93\n  panda: 3c:97:0e:c3:93:76\n",
			want{&Conf{map[string]string{
				"fish":  "3c:97:0e:a3:22:40",
				"cow":   "3c:97:0e:9e:cf:93",
				"panda": "3c:97:0e:c3:93:76",
			}}, nil},
		},
		{"duplicates disallowed", "nodes:\n  fish: 3c:97:0e:a3:22:40\n  cow: 3c:97:0e:a3:22:40\n",
			want{nil, errors.New("multiple nodes may not share the same MAC Address")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeTestFile(tt.file)
			c := Conf{}
			got, err := c.GetConf()
			if !reflect.DeepEqual(got, tt.want.conf) {
				t.Errorf("GetConf() = %v, want %v", got, tt.want.conf)
			}
			if !reflect.DeepEqual(err, tt.want.err) {
				t.Errorf("GetConf() err = #{err}, want #{tt.want.err}")
			}
		})
	}
}

func writeTestFile(content string) {
	f, err := os.Create("conf.yaml")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}
}
