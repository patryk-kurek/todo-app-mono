import { render, screen, fireEvent } from "@testing-library/react";
import App from "./App";

const addTodo = ()=>{
  const input = screen.getByLabelText("Add Todo");
  fireEvent.change(input,{target: {value: "testtodo"}});
  fireEvent.keyDown(input, {key: 'Enter', code: 'Enter', charCode: 13});
};

beforeEach(()=>{
  render(<App/>);
})
describe("app renders everything",()=>{
  test("renders todos h1", () => {
    const h1Element = screen.getByText(/Todos/i);
    expect(h1Element).toBeInTheDocument();
  });
  
  test("renders add todo input",()=>{
    const input = screen.getByLabelText("Add Todo");
    expect(input).toBeInTheDocument();
  });
  
  test("render modal",async ()=>{
    addTodo();
    const todo = screen.getByText("testtodo");
    expect(todo).toBeInTheDocument(); 
    const buttons = screen.getAllByRole('button');
    fireEvent.click(buttons[1]);
    const modal = screen.getByLabelText("Edit Todo");
    const input = modal.parentElement?.getElementsByTagName('input')[0]; 
    expect(input).toHaveValue('testtodo');
  });
});

describe("CRUD functionalities",()=>{
  test("adding todo",()=>{
    const input = screen.getByLabelText("Add Todo");
    expect(input.textContent).toBe("");
    fireEvent.change(input,{target: {value: "testtodo"}});
    fireEvent.keyDown(input, {key: 'Enter', code: 'Enter', charCode: 13});
    const todo = screen.getByText("testtodo");
    expect(todo).toBeInTheDocument(); 
  });
  
  test("deleting todo",()=>{
    addTodo();
    const todo = screen.getByText("testtodo");
    expect(todo).toBeInTheDocument(); 
    const buttons = screen.getAllByRole('button');
    fireEvent.click(buttons[2]); 
    const updatedTodo = screen.queryByText("testtodo"); 
    expect(updatedTodo).toBeNull();
  }); 
  
  test("complete todo",()=>{
    addTodo();
    const todo = screen.getByText("testtodo"); 
    expect(todo).toBeInTheDocument();
    const buttons = screen.getAllByRole('button');
    fireEvent.click(buttons[0]);
    expect(todo).toHaveStyle(`text-decoration: line-through`);
    expect(todo).toHaveStyle(`color: grey`);
  });
  
  test("editing todo", ()=>{
    addTodo();
    const todo = screen.getByText("testtodo");
    expect(todo).toBeInTheDocument(); 
    const buttons = screen.getAllByRole('button');
    fireEvent.click(buttons[1]);
    const inputEditMode = screen.getByLabelText("Edit Todo"); 
    expect(inputEditMode.getAttribute('value')).toBe("testtodo");
    fireEvent.change(inputEditMode,{target: {value: "testtodo2"}});
    fireEvent.submit(inputEditMode,{key: 'Enter',code: 'Enter', charCode: 13});  
    const updatedTodo = screen.getByText("testtodo2");
    expect(updatedTodo) .toBeInTheDocument();
  });
});

