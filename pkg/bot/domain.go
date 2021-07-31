package bot

import (
	"errors"
	"fmt"

	pluginostv1alpha1 "github.com/bryanl/lilbot/gen/go/pluginhost/v1alpha1"
	"github.com/bwmarrin/discordgo"
)

// Command is a command from a plugin.
type Command struct {
	Name    string
	Actions map[string]Action
}

// CommandFromProto converts a register request into a command.
func CommandFromProto(in *pluginostv1alpha1.RegisterRequest) (Command, error) {
	if in == nil {
		return Command{}, errors.New("proto RegisterRequest is nil")
	}

	out := Command{
		Name: in.GetName(),
	}

	for actionName, inAction := range in.GetActions() {
		action, err := ActionFromProto(inAction)
		if err != nil {
			return Command{}, fmt.Errorf("convert action '%s' from proto: %w", actionName, err)
		}

		out.Actions[actionName] = action
	}

	return out, nil
}

// Action is an action for a Command.
type Action struct {
	Name        string
	Description string
	Options     []ActionOption
}

// ActionFromProto converts proto to Action.
func ActionFromProto(in *pluginostv1alpha1.Action) (Action, error) {
	if in == nil {
		return Action{}, errors.New("action is nil")
	}

	out := Action{
		Name:        in.GetName(),
		Description: in.GetDescription(),
	}

	for _, inOption := range in.GetOptions() {
		option, err := ActionOptionFromProto(inOption)
		if err != nil {
			return Action{}, fmt.Errorf("convert action option from proto: %w", err)
		}

		out.Options = append(out.Options, option)
	}

	return out, nil
}

// ToApplicationCommand converts an Action to a discord ApplicationCommand.
func (action *Action) ToApplicationCommand(commandPrefix string) *discordgo.ApplicationCommand {
	name := fmt.Sprintf("%s.%s", commandPrefix, action.Name)

	command := &discordgo.ApplicationCommand{
		Name:        name,
		Description: action.Description,
	}

	for _, option := range action.Options {
		command.Options = append(command.Options, option.ToApplicationCommandOption())
	}

	return command
}

// ActionOption is an option for an Action.
type ActionOption struct {
	Name        string
	Description string
	Type        discordgo.ApplicationCommandOptionType
	Required    bool
}

// ActionOptionFromProto converts proto to an ActionOption.
func ActionOptionFromProto(in *pluginostv1alpha1.ActionOption) (ActionOption, error) {
	if in == nil {
		return ActionOption{}, errors.New("action option is nil")
	}

	var t discordgo.ApplicationCommandOptionType

	switch in.GetType() {
	case pluginostv1alpha1.ActionOption_TYPE_UNSPECIFIED, pluginostv1alpha1.ActionOption_TYPE_STRING:
		t = discordgo.ApplicationCommandOptionString
	case pluginostv1alpha1.ActionOption_TYPE_INTEGER:
		t = discordgo.ApplicationCommandOptionInteger
	case pluginostv1alpha1.ActionOption_TYPE_BOOLEAN:
		t = discordgo.ApplicationCommandOptionBoolean
	default:
		return ActionOption{}, fmt.Errorf("unknown action option type '%s'", in.GetType())
	}

	out := ActionOption{
		Name:        in.GetName(),
		Description: in.GetDescription(),
		Type:        t,
		Required:    in.GetRequired(),
	}

	return out, nil
}

// ToApplicationCommandOption converts an ApplicationOption to discord ApplicationCommandOption.
func (actionOption *ActionOption) ToApplicationCommandOption() *discordgo.ApplicationCommandOption {
	out := &discordgo.ApplicationCommandOption{
		Type:        actionOption.Type,
		Name:        actionOption.Description,
		Description: actionOption.Description,
		Required:    actionOption.Required,
	}

	return out
}
