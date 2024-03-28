import './App.css';
import Todo from './components/Todo';
import { TextField } from '@mui/material';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h1>Todos</h1>
      </header>
      <div className="Todo-input">
        <TextField sx={{backgroundColor: 'white',borderRadius: '10px',width: "100%"}} label="Add Todo" variant='filled'/>
      </div>
      <div className="Todo-list">
        <Todo/>
        <Todo/>
      </div>
    </div>
  );
}

export default App;
