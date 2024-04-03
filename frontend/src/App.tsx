import "./App.css";
import Todo from "./components/Todo";
import { Modal, TextField, Box, FormControl } from "@mui/material";
import { useState } from "react";
import { TodoInterface } from "./interfaces";


interface EditMode {
  todoIndex: number | null;
  value: string;
}

const ModalStyle = {
  position: "absolute" as "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: 400,
  bgcolor: "background.paper",
  border: "2px solid #000",
  borderRadius: "10px",
  boxShadow: 24,
  p: 4,
};

function App() {
  const [state, setState] = useState<{
    todos: TodoInterface[];
    editMode: EditMode;
    input: string;
  }>({
    todos: [],
    editMode: {
      todoIndex: null,
      value: "",
    },
    input: "",
  });
  const onChangeInput = (event: React.ChangeEvent<HTMLInputElement>): void => {
    setState((prevState) => {
      return {
        ...prevState,
        input: event.target.value,
      };
    });
  };

  const onDeleteTodo = (index: number) => {
    const todos = state.todos;
    todos.splice(index, 1);
    setState((prevState) => {
      return {
        ...prevState,
        todos: todos,
      };
    });
  };

  const onSubmitInput = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter" && state.input !== "") {
      const todos = state.todos;
      const todo: TodoInterface = {
        value: state.input,
        completed: false,
      };
      todos.push(todo);

      setState((prevState) => {
        return {
          ...prevState,
          todos: todos,
          input: "",
        };
      });
    }
  };

  const onCompleteTodo = (index: number) => {
    const todos = state.todos;
    todos[index].completed = !todos[index].completed;
    setState((prevState) => {
      return {
        ...prevState,
        todos: todos,
      };
    });
  };

  const onClickEditTodo = (index: number) => { 
    setState((prevState) => {
      return {
        ...prevState,
        editMode: {
          todoIndex: index,
          value: prevState.todos[index].value,
        },
      };
    });
  };

  const onChangeInputEditMode = (
    event: React.ChangeEvent<HTMLInputElement>,
  ) => {
    event.preventDefault();
    setState((prevState) => {
      return {
        ...prevState,
        editMode: {
          ...prevState.editMode,
          value: event.target.value,
        },
      };
    });
  };

  const onSubmitTodoAndExitEditMode = () => {
    const todos = state.todos;
    todos[state.editMode.todoIndex as number].value = state.editMode.value;
    setState((prevState) => {
      return {
        ...prevState,
        // todos: todos,
        editMode: {
          todoIndex: null,
          value: "",
        },
      };
    });
  };

  const onCloseModal = () => {
    setState((prevState) => {
      return {
        ...prevState,
        editMode: {
          todoIndex: null,
          value: "",
        },
      };
    });
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>Todos</h1>
      </header>
      <Modal
        open={state.editMode.todoIndex != null ? true : false}
        onClose={onCloseModal}
      >
        <Box sx={ModalStyle}>
          <form onSubmit={onSubmitTodoAndExitEditMode}>
            <TextField
              sx={{ width: "100%" }}
              label="Edit Todo"
              value={state.editMode.value}
              onChange={onChangeInputEditMode}
            />
          </form>
        </Box>
      </Modal>
      <div className="Todo-input">
        <TextField
          value={state.input}
          onChange={onChangeInput}
          onKeyDown={onSubmitInput}
          sx={{ backgroundColor: "white", borderRadius: "10px", width: "100%" }}
          label="Add Todo"
          variant="filled"
        />
      </div>
      <div className="Todo-list">
        {state.todos.map((todo, index) => {
          return (
            <Todo
              key={index}
              todo={todo}
              onEditTodo={() => onClickEditTodo(index)}
              onCompleteTodo={() => onCompleteTodo(index)}
              onDeleteTodo={() => onDeleteTodo(index)}
            />
          );
        })}
      </div>
    </div>
  );
}

export default App;
