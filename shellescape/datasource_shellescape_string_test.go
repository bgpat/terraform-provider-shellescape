package shellescape_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/bgpat/terraform-provider-shellescape/shellescape"
)

var chars = map[string]string{
	`\u0000`: "'\x00'", `\u0001`: "'\x01'", `\u0002`: "'\x02'", `\u0003`: "'\x03'",
	`\u0004`: "'\x04'", `\u0005`: "'\x05'", `\u0006`: "'\x06'", `\u0007`: "'\a'",
	`\u0008`: "'\b'", `\t`: "'\t'", `\n`: "'\n'", `\u000b`: "'\v'",
	`\u000c`: "'\f'", `\r`: "'\r'", `\u000e`: "'\x0e'", `\u000f`: "'\x0f'",
	`\u0010`: "'\x10'", `\u0011`: "'\x11'", `\u0012`: "'\x12'", `\u0013`: "'\x13'",
	`\u0014`: "'\x14'", `\u0015`: "'\x15'", `\u0016`: "'\x16'", `\u0017`: "'\x17'",
	`\u0018`: "'\x18'", `\u0019`: "'\x19'", `\u001a`: "'\x1a'", `\u001b`: "'\x1b'",
	`\u001c`: "'\x1c'", `\u001d`: "'\x1d'", `\u001e`: "'\x1e'", `\u001f`: "'\x1f'",
	" ": "' '", "!": "'!'", `\"`: "'\"'", "#": "'#'", "$": "'$'", "%": "%", "&": "'&'", "'": "''\"'\"''",
	"(": "'('", ")": "')'", "*": "'*'", "+": "+", ",": ",", "-": "-", ".": ".", "/": "/",
	"0": "0", "1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9",
	":": ":", ";": "';'", "<": "'<'", "=": "=", ">": "'>'", "?": "'?'", "@": "@",
	"A": "A", "B": "B", "C": "C", "D": "D", "E": "E", "F": "F", "G": "G", "H": "H", "I": "I",
	"J": "J", "K": "K", "L": "L", "M": "M", "N": "N", "O": "O", "P": "P", "Q": "Q", "R": "R",
	"S": "S", "T": "T", "U": "U", "V": "V", "W": "W", "X": "X", "Y": "Y", "Z": "Z",
	"[": "'['", "\\\\": "'\\'", "]": "']'", "^": "'^'", "_": "_",
	"`": "'`'", "{": "'{'", "|": "'|'", "}": "'}'", "~": "'~'",
}

func TestDataSourceShellescapeQuote(t *testing.T) {
	for k, v := range chars {
		t.Run(k, func(t *testing.T) {
			resource.ParallelTest(t, resource.TestCase{
				IsUnitTest: true,
				Providers:  map[string]*schema.Provider{"shellescape": shellescape.Provider()},
				Steps: []resource.TestStep{
					{
						Config: fmt.Sprintf(`data "shellescape_quote" "test" { string = "%s" }`, k),
						Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("data.shellescape_quote.test", "quoted", v)),
					},
				},
			})
		})
	}
}
