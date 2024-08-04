package main

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/theduql/duql/internal/logger"
	"github.com/theduql/duql/internal/validator"
	"go.uber.org/zap"
)

type model struct {
	choices  []string
	cursor   int
	selected int
	path     string
}

func initialModel() model {
	return model{
		choices:  []string{"Validate", "Generate SQL", "Quit"},
		selected: -1,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.cursor
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "What would you like to do?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

func main() {
	logger.InitLogger()
	log := logger.GetLogger()

	if len(os.Args) > 1 {
		// Non-interactive mode
		handleCommand(os.Args[1:])
	} else {
		// Interactive mode
		p := tea.NewProgram(initialModel())
		m, err := p.Run()
		if err != nil {
			log.Fatal("Error running program", zap.Error(err))
		}

		if m, ok := m.(model); ok && m.selected != -1 {
			path, err := promptForPath()
			if err != nil {
				log.Fatal("Error getting path", zap.Error(err))
			}

			switch m.choices[m.selected] {
			case "Validate":
				err = validator.Validate(path)
			case "Generate SQL":
				err = generateSQL(path)
			case "Quit":
				return
			}

			if err != nil {
				log.Error("Operation Failed! Ending.", zap.Error(err))
			} else {
				log.Info("Operation Successful!")
			}
		}
	}
}

func handleCommand(args []string) {
	log := logger.GetLogger()

	if len(args) < 2 {
		log.Error("Invalid command. To use try: duql [validate|generate] [file|directory]")
		os.Exit(1)
	}

	command := args[0]
	path := args[1]

	switch command {
	case "validate":
		err := validator.Validate(path)
		if err != nil {
			log.Error(fmt.Sprintf("Validation Failed: %s", path))
			os.Exit(1)
		}
		log.Info("Validation Successful!")
	case "generate":
		err := generateSQL(path)
		if err != nil {
			log.Error(fmt.Sprintf("SQL Generation Failed: %s", err))
			os.Exit(1)
		}
		log.Info("SQL Generation Successful!")
	default:
		log.Error(fmt.Sprintf("Unkonwn Command: %s", command))
		os.Exit(1)
	}
}

func generateSQL(path string) error {
	log := logger.GetLogger()

	// First, validate the input
	err := validator.Validate(path)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// TODO: Implement DUQL to PRQL conversion
	// TODO: Implement PRQL to SQL conversion using prqlc

	log.Info("SQL generation not yet implemented")
	return nil
}

func promptForPath() (string, error) {
	fmt.Print("Enter the path to the file or directory: ")
	var path string
	_, err := fmt.Scanln(&path)
	if err != nil {
		return "", err
	}
	return filepath.Abs(path)
}
