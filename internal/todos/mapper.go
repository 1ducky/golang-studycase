package todos

func MapperArgsToTodos(args []string) (CreateRequest, error) {
	return CreateRequest{
		Task: args[TaskColumn],
	}, nil
}
