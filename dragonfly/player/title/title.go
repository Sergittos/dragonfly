package title

import (
	"fmt"
	"strings"
	"time"
)

// Title represents a title that may be sent to the player. The title will show up as large text in the middle
// of the screen, with optional subtitle and action text.
type Title struct {
	text, subtitle, actionText                string
	fadeInDuration, fadeOutDuration, duration time.Duration
}

// New returns a new title using the text passed. The text is formatted according to the formatting rules of
// fmt.Sprintln, but with no newline at the end.
// The title has default durations set, which will generally suffice.
func New(text ...interface{}) *Title {
	return &Title{
		text:            format(text),
		fadeInDuration:  time.Second / 20,
		fadeOutDuration: time.Second / 20,
		duration:        time.Second * 2,
	}
}

// Text returns the text of the title, as passed to New when created.
func (title *Title) Text() string {
	return title.text
}

// SetText changes the text of the title after it is created. The text passed is formatted according to the
// formatting rules of fmt.Sprintln, but without a newline.
func (title *Title) SetText(text ...interface{}) *Title {
	title.text = format(text)
	return title
}

// SetSubTitle sets the subtitle of the title. The text passed will be formatted according to the formatting
// rules of fmt.Sprintln, but without the newline.
// The subtitle is shown under the title in a somewhat smaller font.
func (title *Title) SetSubtitle(text ...interface{}) *Title {
	title.subtitle = format(text)
	return title
}

// Subtitle returns the subtitle of the title, as passed to SetSubtitle. Subtitle returns an empty string if
// no subtitle was previously set.
func (title *Title) Subtitle() string {
	return title.subtitle
}

// SetActionText sets the action text of the title. This text is roughly the same as sending a tip/popup, but
// will synchronise with the title.
// SetActionText will format the text passed using the formatting rules of fmt.Sprintln, but without newline.
func (title *Title) SetActionText(text ...interface{}) *Title {
	title.actionText = format(text)
	return title
}

// ActionText returns the action text added to the title. This text is roughly the same as sending a tip, but
// will synchronise with the title. By default, the action text is empty.
func (title *Title) ActionText() string {
	return title.actionText
}

// Duration returns the duration that the title will be visible for, without fading in or out. By default,
// this is two seconds.
func (title *Title) Duration() time.Duration {
	return title.duration
}

// SetDuration sets the duration that the title will be visible for without fading in or fading out.
func (title *Title) SetDuration(d time.Duration) *Title {
	title.duration = d
	return title
}

// SetFadeInDuration sets the duration that the title takes to fade in on the screen. The duration will be
// rounded to ticks.
func (title *Title) SetFadeInDuration(d time.Duration) *Title {
	title.fadeInDuration = d
	return title
}

// FadeInDuration returns the duration that the fade-in of the title takes. By default, this is a quarter of
// a second.
func (title *Title) FadeInDuration() time.Duration {
	return title.fadeInDuration
}

// SetFadeOutDuration sets the duration that the title takes to fade out of the screen. The duration will be
// rounded to ticks.
func (title *Title) SetFadeOutDuration(d time.Duration) *Title {
	title.fadeOutDuration = d
	return title
}

// FadeOutDuration returns the duration that the fade-out of the title takes.By default, this is a quarter of
// a second.
func (title *Title) FadeOutDuration() time.Duration {
	return title.fadeOutDuration
}

// format is a utility function to format a list of values to have spaces between them, but no newline at the
// end.
func format(a []interface{}) string {
	return strings.TrimSuffix(strings.TrimSuffix(fmt.Sprintln(a...), "\n"), "\n")
}