import React from 'react';
import {
  Container,
  Typography,
  Box,
  Button,
  TextField,
  makeStyles,
  Paper,
} from '@material-ui/core';
import { useHistory } from 'react-router-dom';

const useStyles = makeStyles((theme) => ({
  container: {
    paddingTop: theme.spacing(4),
  },
  paper: {
    padding: theme.spacing(4),
  },
  textField: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
}));

function CreateTenancyAgreement() {
  const classes = useStyles();
  const history = useHistory();

  const handleSubmit = (e) => {
    e.preventDefault();
    // Handle form submission here
    history.push('/tenancy-agreements');
  };

  return (
    <Container className={classes.container}>
      <Typography variant="h4">Create Tenancy Agreement</Typography>
      <Box mt={4}>
        <Paper className={classes.paper}>
          <form onSubmit={handleSubmit}>
            <TextField
              label="Tenant Name"
              variant="outlined"
              fullWidth
              className={classes.textField}
            />
            <TextField
              label="Property Address"
              variant="outlined"
              fullWidth
              className={classes.textField}
            />
            <TextField
              label="Start Date"
              type="date"
              InputLabelProps={{ shrink: true }}
              variant="outlined"
              fullWidth
              className={classes.textField}
            />
            <TextField
              label="End Date"
              type="date"
              InputLabelProps={{ shrink: true }}
              variant="outlined"
              fullWidth
              className={classes.textField}
            />
            <TextField
              label="Rent Amount"
              type="number"
              InputProps={{
                startAdornment: <span>&pound;</span>,
              }}
              variant="outlined"
              fullWidth
              className={classes.textField}
            />
            <Button variant="contained" color="primary" type="submit">
              Create
            </Button>
          </form>
        </Paper>
      </Box>
    </Container>
  );
}

export default CreateTenancyAgreement;
