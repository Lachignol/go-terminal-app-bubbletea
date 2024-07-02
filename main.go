// package main

// // A simple example that shows how to retrieve a value from a Bubble Tea
// // program after the Bubble Tea has exited.

// import (
// 	"fmt"
// 	"os"
// 	"strings"

// 	tea "github.com/charmbracelet/bubbletea"
// )

// var choices = []string{"Taro", "Coffee", "Lychee"}

// type model struct {
// 	cursor int
// 	choice string
// }

// func (m model) Init() tea.Cmd {
// 	return nil
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "ctrl+c", "q", "esc":
// 			return m, tea.Quit

// 		case "enter":
// 			// Send the choice on the channel and exit.
// 			m.choice = choices[m.cursor]
// 			return m, tea.Quit

// 		case "down", "j":
// 			m.cursor++
// 			if m.cursor >= len(choices) {
// 				m.cursor = 0
// 			}

// 		case "up", "k":
// 			m.cursor--
// 			if m.cursor < 0 {
// 				m.cursor = len(choices) - 1
// 			}
// 		}
// 	}

// 	return m, nil
// }

// func (m model) View() string {
// 	s := strings.Builder{}
// 	s.WriteString("What kind of Bubble Tea would you like to order?\n\n")

// 	for i := 0; i < len(choices); i++ {
// 		if m.cursor == i {
// 			s.WriteString("> ")
// 		} else {
// 			s.WriteString("[ ] ")
// 		}
// 		s.WriteString(choices[i])
// 		s.WriteString("\n")
// 	}
// 	s.WriteString("\n(press q to quit)\n")

//		return s.String()
//	}
package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s? Sounds good to me.", m.choice))
	}
	if m.quitting {
		return quitTextStyle.Render("Not hungry? That’s cool.")
	}
	return "\n" + m.list.View()
}

func main() {
	items := []list.Item{
		item("Ramen"),
		item("Tomato Soup"),
		item("Hamburgers"),
		item("Cheeseburgers"),
		item("Currywurst"),
		item("Okonomiyaki"),
		item("Pasta"),
		item("Fillet Mignon"),
		item("Caviar"),
		item("Just Wine"),
	}

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "What do you want for dinner?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	// Run returns the model as a tea.Model.
	m := model{list: l}
	var object tea.Model
	var err error
	if object, err = tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	if object, ok := object.(model); ok && object.choice != "" {
		fmt.Printf("\n---\nYou chose %s!\n", object.choice)
	}

}

/*
	var message = flag.String(
		"message",
		"Hello, World!",
		"The message you'd like to print to the terminal",
	)

	var number = flag.Int(
		"number",
		1,
		"The number you'd like to add to your message",
	)

	flag.Parse()

	fmt.Println("This is the message you want to display: " + *message + " with number " + strconv.Itoa(*number))

*/
