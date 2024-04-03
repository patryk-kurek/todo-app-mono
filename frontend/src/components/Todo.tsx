import { Card, Box, Fab, Grid } from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import EditIcon from "@mui/icons-material/Edit";
import CheckIcon from "@mui/icons-material/Check";
import CloseIcon from "@mui/icons-material/Close";
import { TodoInterface } from "../interfaces";
 
interface TodoProps {
    todo: TodoInterface;
    onCompleteTodo: ()=>void;
    onEditTodo: ()=>void;
    onDeleteTodo: ()=>void;
}

const Todo = (props: TodoProps) => {
  return (
    <Card
      sx={{
        width: "30%",
        margin: "auto",
        marginTop: "1%",
        verticalAlign: "middle",
      }}
    >
      <Grid container spacing={2}>
        <Grid item xs={7}>
          <Box
            sx={{ wordWrap: "break-word", paddingLeft: "20px", margin: "auto" }}
          >
            <h2
              style={
                props.todo.completed
                  ? { textDecoration: "line-through", color: "grey" }
                  : {}
              }
            >
              {props.todo.value}
            </h2>
          </Box>
        </Grid>
        <Grid item xs={5}>
          <Box
            sx={{
              margin: "auto",
              display: "flex",
              marginLeft: "20%",
              alignItems: "center",
              height: "100%",
            }}
          >
            <Fab onClick={props.onCompleteTodo} size="small">
              {props.todo.completed ? <CloseIcon /> : <CheckIcon />}
            </Fab>
            <Fab
              size="small"
              onClick={props.onEditTodo}
              sx={{ marginLeft: "10px" }}
            >
              <EditIcon />
            </Fab>
            <Fab
              onClick={props.onDeleteTodo}
              sx={{ marginLeft: "10px" }}
              size="small"
            >
              <DeleteIcon />
            </Fab>
          </Box>
        </Grid>
      </Grid>
    </Card>
  );
};

export default Todo;
