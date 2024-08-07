// package main

// import (
// 	"errors"
// 	"fmt"
// 	"strings"
// 	"time"

// 	"github.com/charmbracelet/bubbles/filepicker"
// 	tea "github.com/charmbracelet/bubbletea"
// )

// //modif pour quitter et avoir le chemin du fichier quand on appui sur entré sur un fichier valide

// type model struct {
// 	filepicker   filepicker.Model
// 	selectedFile string
// 	quitting     bool
// 	err          error
// }

// type clearErrorMsg struct{}

// func clearErrorAfter(t time.Duration) tea.Cmd {
// 	return tea.Tick(t, func(_ time.Time) tea.Msg {
// 		return clearErrorMsg{}
// 	})
// }

// func (m model) Init() tea.Cmd {
// 	return m.filepicker.Init()
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "ctrl+c", "q":
// 			m.quitting = true
// 			return m, tea.Quit
// 		}
// 	case clearErrorMsg:
// 		m.err = nil
// 	}

// 	var cmd tea.Cmd
// 	m.filepicker, cmd = m.filepicker.Update(msg)

// 	// Did the user select a file?
// 	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
// 		// Get the path of the selected file.
// 		m.selectedFile = path
// 		m.quitting = true
// 		return m, tea.Quit
// 	}

// 	// Did the user select a disabled file?
// 	// This is only necessary to display an error to the user.
// 	if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
// 		// Let's clear the selectedFile and display an error.
// 		m.err = errors.New(path + " is not valid.")
// 		m.selectedFile = ""
// 		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
// 	}

// 	return m, cmd
// }

// func (m model) View() string {
// 	if m.quitting {
// 		return ""
// 	}
// 	var s strings.Builder
// 	s.WriteString("\n  ")
// 	if m.err != nil {
// 		s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.err.Error()))
// 	} else if m.selectedFile == "" {
// 		s.WriteString("Pick a file:")
// 	} else {
// 		s.WriteString("Selected file: " + m.filepicker.Styles.Selected.Render(m.selectedFile))
// 	}
// 	s.WriteString("\n\n" + m.filepicker.View() + "\n")
// 	return s.String()
// }

// func main() {
// 	fp := filepicker.New()
// 	fp.AllowedTypes = []string{".mod", ".sum", ".go", ".txt", ".md"}
// 	// pour definire le repertoire courant au repertoire user  remplacer par :
// 	// fp.CurrentDirectory,_ = os.UserHomeDir() 
// 	fp.CurrentDirectory = "./"


// 	m := model{
// 		filepicker: fp,
// 	}
// 	tm, _ := tea.NewProgram(&m).Run()
// 	mm := tm.(model)

// 	file:= mm.selectedFile

// 	if file != "" {
// 		fmt.Println("\n  You selected: " + m.filepicker.Styles.Selected.Render(file) + "\n")
// 	}


// 		// fmt.Println("\n  You selected: " + m.filepicker.Styles.Selected.Render(mm.selectedFile) + "\n")
	
	
// }
