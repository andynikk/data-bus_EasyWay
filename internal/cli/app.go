package cli

import (
	"fmt"
	"sync"

	"databuseasyway/internal/cli/forms"
	"databuseasyway/internal/constants"
	"databuseasyway/internal/environment"

	tea "github.com/charmbracelet/bubbletea"
)

type Client struct {
	Config *environment.ConfigClient
	Models []tea.Model
	Tea    *tea.Program
	sync.RWMutex
}

func NewClient() Client {

	fmt.Print("loading...")
	client := Client{}
	client.Models = []tea.Model{}

	client.Models = []tea.Model{
		&forms.FormLogin{},
	}

	client.Config = environment.NewConfigClient()
	formLogin := client.Models[constants.FormLogin].(*forms.FormLogin)
	formLogin.Init()

	client.Tea = tea.NewProgram(formLogin)

	return client
}

func (c *Client) Run() error {
	_, err := c.Tea.Run()
	return err
}
