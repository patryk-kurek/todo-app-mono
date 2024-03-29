import { Card,Box,Fab } from "@mui/material";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import CheckIcon from '@mui/icons-material/Check';

const Todo = (props:any)=> {
    return ( 
        <Card sx={{width: '25%',margin: 'auto',marginTop: '1%'}}>
            <Box sx={{marginLeft: "50%",marginTop:"2%"}}>
                <Fab size="small" sx={{position: 'absolute',marginLeft: '3%',marginTop: '0.25%'}}>
                    <CheckIcon/>
                </Fab>
                <Fab size="small" sx={{position: 'absolute',marginLeft: '6%',marginTop: '0.25%'}}>
                    <EditIcon/>
                </Fab>
                <Fab onClick={props.onDeleteTodo} size="small" sx={{position: 'absolute',marginLeft: '9%',marginTop: '0.25%'}}>
                    <DeleteIcon />
                </Fab> 
            </Box>
            <Box sx={{marginLeft: "20px"}}>
                <h2>{props.value}</h2>
            </Box>
        </Card>
    )
};

export default Todo;
