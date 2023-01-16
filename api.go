// package reportstyle
package reportstyle

import (
	"html"
)

// const
const (
	// ascii string token
	Linefeed  = "\n"
	Space     = " "
	Empty     = ""
	Separator = " : "
	Ok        = "[OK]"
	Valid     = "[VALID]"
	Fail      = "[FAIL]"
	Alert     = "[ALERT]"
	// ansi color terminal
	AnsiRed     = "\033[2;91m"
	AnsiGreen   = "\033[2;92m"
	AnsiYellow  = "\033[2;93m"
	AnsiBlue    = "\033[2;94m"
	AnsiMagenta = "\033[2;95m"
	AnsiCyan    = "\033[2;96m"
	AnsiWhite   = "\033[2;97m"
	AnsiGrey    = "\033[2;90m"
	AnsiEnd     = "\033[0m"
	// html color
	HtmlRed   = "<a style=\"color:#FF0000\"><strong>"
	HtmlGreen = "<a style=\"color:#00FF00\"><strong>"
	HtmlEnd   = "</strong></a>"
)

// Style defines the output style, feel free to design your own within your app, see example defaults below
type Style struct {
	Raw      bool                // add raw datasets
	Start    string              // start object with
	End      string              // end object with
	Ok       string              // string Ok statement
	Valid    string              // strong valid statement
	Fail     string              // strong fail statement
	Alert    string              // strong alert statement
	L1       string              // table start & style 1
	L2       string              // table style 2
	L3       string              // table style 3 ( multiline style 2  )
	L4       string              // table style 4
	LE       string              // table end
	PS       string              // preformated start
	PE       string              // preformated end
	SaniFunc func(string) string // sanitizer function for preformated blocks
}

// StyleNone ...
// Returns a Default empty Style
func StyleNone() *Style { return &Style{} }

// StylePlain ...
// Returns a Default Plain Text Report Style [currently:alias]
func StylePlain() *Style { return StyleText() }

// StyleText ...
// Returns as example a Ascii Plain Text Report Style
func StyleText() *Style {
	return &Style{
		Ok:    Ok,
		Valid: Valid,
		Fail:  Fail,
		Alert: Alert,
		L2:    Separator,
		L3:    Linefeed,
		LE:    Linefeed,
	}
}

// StyleAnsi ...
// Returns as example a Ansi Color Terminal Report Style
func StyleAnsi() *Style {
	return &Style{
		Ok:    AnsiGreen + Ok + AnsiEnd,
		Valid: AnsiGreen + Valid + AnsiEnd,
		Fail:  AnsiRed + Fail + AnsiEnd,
		Alert: AnsiRed + Alert + AnsiEnd,
		L1:    AnsiYellow,
		L2:    AnsiBlue + Separator,
		L3:    AnsiBlue + Linefeed,
		L4:    AnsiBlue + Linefeed,
		LE:    AnsiEnd + Linefeed,
	}
}

// StyleMarkdown ...
// Returns as example a [gh] Markdown Report Style
func StyleMarkdown() *Style {
	return &Style{
		Ok:    Ok,
		Valid: Valid,
		Fail:  Fail,
		Alert: Alert,
		L2:    Separator,
		L3:    Linefeed,
		LE:    Linefeed,
	}
}

// StyleHTML ...
// Returns as example a HTML/CSS Report Style for usage with any google compatible syntax highlighter css stylesheet
func StyleHTML() *Style {
	return &Style{
		Start:    "\n<pre class=\"pr\" style=\"line-height:0.5;\">\n",
		End:      "\n</pre>\n",
		Ok:       HtmlGreen + Ok + HtmlEnd,
		Valid:    HtmlGreen + Valid + HtmlEnd,
		Fail:     HtmlRed + Fail + HtmlEnd,
		Alert:    HtmlRed + Alert + HtmlEnd,
		L1:       "\t<li><span class=\"kwd\">",
		L2:       "</span>" + Separator + "<span class=\"str\">",
		L3:       "</span><br>\n",
		L4:       "</span>" + Separator + "<span class=\"fun\">",
		LE:       "</span></li>\n",
		PS:       "\n<div class=\"pr\" style=\"line-height:1.1\">\n",
		PE:       "\n</div>\n",
		SaniFunc: html.EscapeString,
	}
}
