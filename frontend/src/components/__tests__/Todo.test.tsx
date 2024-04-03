import {fireEvent, render,screen} from "@testing-library/react";
import Todo from "../Todo"; 
import { TodoInterface } from "../../interfaces";
const onCompleteTodoMock = jest.fn();

const onEditTodoMock = jest.fn();

const onDeleteTodoMock = jest.fn();
beforeEach(()=>{
    const todo : TodoInterface = {
        value: "test",
        completed: false
    }; 
    render(<Todo todo={todo} onCompleteTodo={onCompleteTodoMock} onEditTodo={onEditTodoMock} onDeleteTodo={onDeleteTodoMock}/>);  
});

test("component render correctly",()=>{
    const component = screen.getByText("test");
    expect(component).toBeInTheDocument();
})

test("delete button works",()=>{ 
    const buttons = screen.getAllByRole('button');
    fireEvent.click(buttons[2]);
    expect(onDeleteTodoMock).toHaveBeenCalledTimes(1);
});

test("complete/uncomplete button works",()=>{
    const buttons = screen.getAllByRole("button");
    fireEvent.click(buttons[0]);
    expect(onCompleteTodoMock).toHaveBeenCalledTimes(1);
    fireEvent.click(buttons[0]);
    expect(onCompleteTodoMock).toHaveBeenCalledTimes(2);
});
 

test("edit button works",()=>{
    const buttons = screen.getAllByRole("button");
    fireEvent.click(buttons[1]);
    expect(onEditTodoMock).toHaveBeenCalledTimes(1);
});
