package luhn

import (
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {

	tests := []struct {
		val  string
		want bool
	}{
		{"49927398716", true},
		{"49927398717", false},
		{"323392A7", false},
		{"1234567812345678", false},
		{"1234567812345670", true},
		{"79927398713", true},
		{"12344", true},
	}

	for _, tt := range tests {
		got := Validate(tt.val)
		if got != tt.want {
			t.Errorf("For %s, expected: %t, got %t", tt.val, tt.want, got)
		}
	}
}

func ExampleValidate() {
	fmt.Println(Validate("49927398716"))
	//Output: true
}

func TestGenerateNumber(t *testing.T) {
	tests := []struct {
		val  string
		want string
	}{
		{"123", "0"},
		{"3214", "4"},
		{"87438443285", "0"},
		{"9812244564412", "8"},
	}

	for _, tt := range tests {
		got, err := GenerateNumber(tt.val)
		if err != nil {
			t.Fatal(err)
		}
		if got != tt.want {
			t.Errorf("For %s, expected: %s, got %s", tt.val, tt.want, got)
		}
	}
}

func ExampleGenerateNumber() {
	num, _ := GenerateNumber("123")
	fmt.Println(num)
	//Output: 0
}

func TestGenerate(t *testing.T) {
	tests := []struct {
		val  string
		want string
	}{
		{"123", "1230"},
		{"3214", "32144"},
		{"87438443285", "874384432850"},
		{"9812244564412", "98122445644128"},
	}

	for _, tt := range tests {
		got, err := Generate(tt.val)
		if err != nil {
			t.Fatal(err)
		}
		if got != tt.want {
			t.Errorf("For %s, expected: %s, got %s", tt.val, tt.want, got)
		}
	}

}

func ExampleGenerate() {
	num, _ := Generate("123")
	fmt.Println(num)
	//Output: 1230
}
