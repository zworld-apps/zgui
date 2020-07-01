package zgui

import (
	"fmt"

	rl "github.com/xzebra/raylib-go/raylib"
)

type textFieldCallback func(tf *TextFieldComponent)

type TextFieldOptions struct {
	Box  *BoxOptions
	Text *TextOptions

	IsPassword bool

	SubmitCallback textFieldCallback
}

const (
	selectedBarTime = 0.5
)

type TextFieldComponent struct {
	*baseComponent
	Box   *LineBoxComponent
	Label *LabelComponent

	// Real text inserted (Label contains the "display text").
	Text string
	// showBar shows the blinking focus indicator
	showBar bool
	// timer is used for the blinking
	timer float32

	opt *TextFieldOptions
}

func NewTextFieldComponent(options *TextFieldOptions) *TextFieldComponent {
	tf := &TextFieldComponent{
		baseComponent: newBaseComponent(),
		Box:           NewLineBoxComponent(options.Box),
		Label:         NewLabelComponent("", options.Text),
		opt:           options,
	}

	tf.Box.Add(tf.Label, DefaultConstraints())
	tf.baseComponent.Add(tf.Box, DefaultConstraints())

	return tf
}

func (tf *TextFieldComponent) SetState(val GuiState) {
	if val != StatePressed {
		tf.showBar = false
	} else {
		tf.timer = 0
		rl.ShowKeyboard(true)
	}
	tf.baseComponent.SetState(val)
}

func (tf *TextFieldComponent) handleKeyPressed(key int32) {
	// Update the real text variable.
	tf.Text += string(key)
	// According to text representation, update the display text.
	if tf.opt.IsPassword {
		tf.Label.Text += "*"
	} else {
		tf.Label.Text += string(key)
	}
}

func (tf *TextFieldComponent) removeChar() {
	textLen := len(tf.Text)
	if textLen == 0 {
		// avoid removing from empty string
		return
	}
	// Remove last character
	tf.Text = tf.Text[:textLen-1]
	tf.Label.Text = tf.Label.Text[:textLen-1]
}

func (tf *TextFieldComponent) Update(dt float32) {
	hover := tf.MouseInBounds(rl.GetMouseX(), rl.GetMouseY())
	tapped := rl.IsMouseButtonPressed(rl.MouseLeftButton) ||
		rl.IsGestureDetected(rl.GestureTap)
	touched := (hover && tapped) || tf.TouchInBounds()

	if hover {
		if touched {
			if tf.State != StatePressed {
				tf.SetState(StatePressed)
				tf.timer = 0
				rl.ShowKeyboard(true)
			}
		} else if tf.State == StateNormal {
			tf.SetState(StateFocused)
			tf.showBar = false
		}
	} else if tf.State == StateFocused || tapped {
		tf.SetState(StateNormal)
		tf.showBar = false
		// rl.ShowKeyboard(false)
	}

	if tf.IsSelected() {
		tf.timer += dt
		if tf.timer >= selectedBarTime {
			tf.timer = 0
			tf.showBar = !tf.showBar
		}

		// text field keyboard handling
		latestKey := rl.GetKeyPressed()
		if latestKey != 0 {
			// only non special keys are listed as KeyPressed,
			// the rest of them can be checked with rl.IsKeyPressed
			tf.handleKeyPressed(latestKey)
		} else if rl.IsKeyPressed(rl.KeyBackspace) {
			tf.removeChar()
		} else if rl.IsKeyPressed(rl.KeyEnter) {
			if tf.opt.SubmitCallback != nil && len(tf.Text) > 0 {
				tf.opt.SubmitCallback(tf)
				tf.Text = ""
				tf.Label.Text = ""
			}
		}
	}
}

func (tf *TextFieldComponent) Draw() {
	tf.baseComponent.Draw()
}

func (b *TextFieldComponent) String() string {
	return fmt.Sprintf("TextFieldComponent%+v", b.opt)
}
