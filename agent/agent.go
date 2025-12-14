package agent

import "fmt"

type Executor func(Action) Observation

type Agent struct {
	Planner *Planner
	Context *Context
	Execute Executor
}

func NewAgent(planner *Planner, executor Executor) *Agent {
	return &Agent{
		Planner: planner,
		Context: NewContext(),
		Execute: executor,
	}
}

func (a *Agent) Run(goal string) {
	a.Context.AddUserMessage(goal)

	for i := 0; i < 10; i++ {
		action, err := a.Planner.Plan(a.Context)
		if err != nil {
			fmt.Printf("Planning error: %v\n", err)
			break
		}

		fmt.Printf("Thought: Acting on %s\n", action.ToolName)
		
		if action.ToolName == "done" {
			fmt.Println("Goal achieved!")
			break
		}

		obs := a.Execute(*action)
		a.Context.AddAssistantMessage(fmt.Sprintf(`{"tool_name": "%s", "args": %v}`, action.ToolName, action.Args))
		a.Context.AddObservation(obs)
		
		fmt.Printf("Observation: %s\n", obs.Output)
		if obs.Error != nil {
			fmt.Printf("Error: %v\n", obs.Error)
		}
	}
}
