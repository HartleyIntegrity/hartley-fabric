import React, { useState, useEffect } from "react";
import { useHistory } from "react-router-dom";
import { TextField, Button } from "@mui/material";
import { getDecodedToken } from "../../utils/api";
import axios from "axios";

function EditTenancyAgreement(props) {
  const history = useHistory();
  const [formData, setFormData] = useState({});

  useEffect(() => {
    axios
      .get(`/api/tenancy-agreements/${props.match.params.id}`)
      .then((response) => {
        setFormData(response.data);
      })
      .catch((error) => {
        console.log(error);
      });
  }, [props.match.params.id]);

  const handleFormSubmit = (e) => {
    e.preventDefault();
    axios
      .put(`/api/tenancy-agreements/${formData.id}`, formData, {
        headers: {
          Authorization: `Bearer ${getDecodedToken().token}`,
        },
      })
      .then(() => {
        history.push("/tenancy-agreements");
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  return (
    <div>
      <form onSubmit={handleFormSubmit}>
        <TextField
          label="Start Date"
          name="startDate"
          type="date"
          defaultValue={formData.startDate}
          onChange={handleInputChange}
          InputLabelProps={{
            shrink: true,
          }}
        />
        <br />
        <TextField
          label="End Date"
          name="endDate"
          type="date"
          defaultValue={formData.endDate}
          onChange={handleInputChange}
          InputLabelProps={{
            shrink: true,
          }}
        />
        <br />
        <Button variant="contained" color="primary" type="submit">
          Submit
        </Button>
      </form>
    </div>
  );
}

export default EditTenancyAgreement;
