import './App.css';
import Todo from './components/Todo';
import { TextField } from '@mui/material';
import {useState} from "react"

interface Todo { 
  value: string
  editMode: boolean
  completed: boolean
}

function App() {
  const [state,setState] = useState<{todos: string[],input: string}>({
    todos: [],
    input: ""
  });
  const onChangeInput = (event: React.ChangeEvent<HTMLInputElement>) : void=>{
    setState((prevState)=>{
      return {
        ...prevState,
        input: event.target.value
      }
    });
  }; 

  const onDeleteTodo = (index: number) =>{
    const todos = state.todos;
    todos.splice(index,1)
    setState((prevState)=>{
      return {
        ...prevState,
        todos: todos
      }
    })
  };
  
  const onSubmitInput = (event: React.KeyboardEvent<HTMLInputElement>)=>{
    if (event.key == "Enter"){
      const todos = state.todos;
      todos.push(state.input)
       
      setState((prevState)=>{
        return {
          ...prevState,
          todos: todos,
          input: ""
        }
       });
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>Todos</h1>
      </header>
      <div className="Todo-input">
        <TextField 
          value={state.input} 
          onChange={onChangeInput} 
          onKeyDown={onSubmitInput}
          sx={{backgroundColor: 'white',borderRadius: '10px',width: "100%"}} 
          label="Add Todo" 
          variant='filled'/>
      </div>
      <div className="Todo-list">
        {state.todos.map((todo,index)=>{
          return <Todo value={todo} onDeleteTodo={()=>onDeleteTodo(index)}/>
        })}
      </div>
    </div>
  );
}

export default App;
