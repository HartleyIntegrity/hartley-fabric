import React, { useState } from "react";
import { useHistory } from "react-router-dom";
import {
  Box,
  Button,
  Container,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  TextField,
  Typography,
} from "@mui/material";
import { makeStyles } from "@mui/styles";
import api from "../../utils/api";

const useStyles = makeStyles({
  form: {
    display: "flex",
    flexDirection: "column",
    gap: 20,
    maxWidth: 600,
    margin: "auto",
  },
  submitButton: {
    marginTop: 20,
  },
});

const AddTenancyAgreement = () => {
  const classes = useStyles();
  const history = useHistory();
  const [formData, setFormData] = useState({
    Property: "",
    Tenant: "",
    StartDate: "",
    EndDate: "",
    Description: "",
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    api
      .post("/blockchain", { TenancyAgree: formData })
      .then(() => {
        history.push("/");
      })
      .catch((error) => {
        console.error(error);
      });
  };

  return (
    <Container maxWidth="md">
      <Typography variant="h4" gutterBottom>
        Add New Tenancy Agreement
      </Typography>
      <Box component="form" className={classes.form} onSubmit={handleSubmit}>
        <FormControl>
          <InputLabel id="property-label">Property</InputLabel>
          <Select
            labelId="property-label"
            id="property"
            name="Property"
            value={formData.Property}
            onChange={handleChange}
          >
            <MenuItem value="1">Property 1</MenuItem>
            <MenuItem value="2">Property 2</MenuItem>
            <MenuItem value="3">Property 3</MenuItem>
          </Select>
        </FormControl>
        <TextField
          label="Tenant Name"
          name="Tenant"
          value={formData.Tenant}
          onChange={handleChange}
        />
        <TextField
          label="Start Date"
          name="StartDate"
          type="date"
          value={formData.StartDate}
          onChange={handleChange}
          InputLabelProps={{
            shrink: true,
          }}
        />
        <TextField
          label="End Date"
          name="EndDate"
          type="date"
          value={formData.EndDate}
          onChange={handleChange}
          InputLabelProps={{
            shrink: true,
          }}
        />
        <TextField
          label="Description"
          name="Description"
          value={formData.Description}
          onChange={handleChange}
          multiline
          rows={4}
        />
        <Button
          variant="contained"
          color="primary"
          type="submit"
          className={classes.submitButton}
        >
          Submit
        </Button>
      </Box>
    </Container>
  );
};

export default AddTenancyAgreement;
