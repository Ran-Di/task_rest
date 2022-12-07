package crypting

import "testing"

func Test_Crypt(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"AAABBACACACAC", "3A2B4(AC)"},
		{"AAABBACACACACD", "3A2B4(AC)D"},
		{"AAABBACACACACDDDD", "3A2B4(AC)4D"},
		{"AAABBACACACACADDADDACCDD", "3A2B4(AC)2(A2D)A2C2D"},
		{"AAABABABACACCACC", "3A3(BA)C2(A2C)"},
		{"ABBFBBFBBFBBDBBFBBFBBFBBDE", "A2(3(2BF)2BD)E"},
	}
	for _, c := range cases {
		got := Crypt(c.in)
		if got != c.want {
			t.Errorf("Crypting: (%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func Test_Decrypt(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"A2(3(2BF)2BD)E", "ABBFBBFBBFBBDBBFBBFBBFBBDE"},
		{"A2(2B3(F2B)D)E", "ABBFBBFBBFBBDBBFBBFBBFBBDE"},
		{"3A2B4(AC)4D", "AAABBACACACACDDDD"},
		{"3A2B4(AC)2(A2D)A2C2D", "AAABBACACACACADDADDACCDD"},
		{"3A3(BA)C2(A2C)", "AAABABABACACCACC"},
		{"", ""},
	}
	for _, c := range cases {
		got := Decrypt(c.in)
		if got != c.want {
			t.Errorf("Derypting: (%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func Test_All(t *testing.T) {
	cases := []string{
		"AAABBACACACAC",
		"AAABBACACACACADDADDACCDD",
		"ABBFBBFBBFBBDBBFBBFBBFBBDE",
	}
	for _, c := range cases {
		got_crypt := Crypt(c)
		got_decrypt := Decrypt(got_crypt)
		if got_decrypt != c {
			t.Errorf("Crypting-Decrypting: (%q) == %q, intermediate line %q", c, got_decrypt, got_crypt)
		}
	}
}
