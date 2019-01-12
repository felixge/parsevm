package abnf

import "testing"

func TestABNF(t *testing.T) {
	tests := []struct {
		Name    string
		Grammar string
	}{
		{Name: "Simple", Grammar: `foo = "- "` + "\n"},
		// From https://github.com/felixge/taskpaper
		{
			Name: "TaskPaper Simple",
			Grammar: `
; The last item has no trailing CR.
document  = *(item CR) *1item
item      = (task / project / note)
task      = "- " content
project   = content ":"
note      = content
; The first tag needs no space in front of it, and the last tag doesn't need
; a space after it.
content   = *1(tag " ") *(" " tag " " / OCTET) *1(" " tag)
; A tag name or value can be empty which is weird, but you can try it in
; TaskPaper.
tag       = "@" tag_name *1("(" tag_value ")")
tag_name  = *OCTET
tag_value = *OCTET
`,
		},
		// https://en.wikipedia.org/wiki/Augmented_Backus%E2%80%93Naur_form#Example
		{
			Name: "WikiPedia postal address",
			Grammar: `
postal-address   = name-part street zip-part

name-part        = *(personal-part SP) last-name [SP suffix] CRLF
name-part        =/ personal-part CRLF

personal-part    = first-name / (initial ".")
first-name       = *ALPHA
initial          = ALPHA
last-name        = *ALPHA
suffix           = ("Jr." / "Sr." / 1*("I" / "V" / "X"))

street           = [apt SP] house-num SP street-name CRLF
apt              = 1*4DIGIT
house-num        = 1*8(DIGIT / ALPHA)
street-name      = 1*VCHAR

zip-part         = town-name "," SP state 1*2SP zip-code CRLF
town-name        = 1*(ALPHA / SP)
state            = 2ALPHA
zip-code         = 5DIGIT ["-" 4DIGIT]
`,
		},
		{
			Name: "ABNF in ABNF from RFC5234",
			Grammar: `
         rulelist       =  1*( rule / (*c-wsp c-nl) )

         rule           =  rulename defined-as elements c-nl
                                ; continues if next line starts
                                ;  with white space

         rulename       =  ALPHA *(ALPHA / DIGIT / "-")

         defined-as     =  *c-wsp ("=" / "=/") *c-wsp
                                ; basic rules definition and
                                ;  incremental alternatives

         elements       =  alternation *c-wsp

         c-wsp          =  WSP / (c-nl WSP)

         c-nl           =  comment / CRLF
                                ; comment or newline

         comment        =  ";" *(WSP / VCHAR) CRLF

         alternation    =  concatenation
                           *(*c-wsp "/" *c-wsp concatenation)

         concatenation  =  repetition *(1*c-wsp repetition)

         repetition     =  [repeat] element

         repeat         =  1*DIGIT / (*DIGIT "*" *DIGIT)

         element        =  rulename / group / option /
                           char-val / num-val / prose-val

         group          =  "(" *c-wsp alternation *c-wsp ")"

         option         =  "[" *c-wsp alternation *c-wsp "]"

         char-val       =  DQUOTE *(%x20-21 / %x23-7E) DQUOTE
                                ; quoted string of SP and VCHAR
                                ;  without DQUOTE

         num-val        =  "%" (bin-val / dec-val / hex-val)

         bin-val        =  "b" 1*BIT
                           [ 1*("." 1*BIT) / ("-" 1*BIT) ]
                                ; series of concatenated bit values
                                ;  or single ONEOF range

         dec-val        =  "d" 1*DIGIT
                           [ 1*("." 1*DIGIT) / ("-" 1*DIGIT) ]

         hex-val        =  "x" 1*HEXDIG
                           [ 1*("." 1*HEXDIG) / ("-" 1*HEXDIG) ]
`,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if !Validate(test.Grammar) {
				t.Fail()
			}
		})
	}
}
