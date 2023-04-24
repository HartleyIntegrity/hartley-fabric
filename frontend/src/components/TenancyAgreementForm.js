import React, { useState, useEffect } from "react";
import { useHistory, useParams } from "react-router-dom";
import { createTenancyAgreement, updateTenancyAgreement, getTenancyAgreementById } from "../utils/api";
import { TextField, Button, Stack } from "@mui/material";

function TenancyAgreementForm() {
  const [agreement, setAgreement] = useState({
    property: "",
    landlord: "",
    tenant: "",
    start_date: "",
    end_date: "",
  });
  const history = useHistory();
  const { id } = useParams();

  useEffect(() => {
    if (id) {
      getTenancyAgreementById(id).then((data) => setAgreement(data));
    }
  }, [id]);

  const handleChange = (event) => {
    const { name, value } = event.target;
    setAgreement({ ...agreement, [name]: value });
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    if (id) {
      updateTenancyAgreement(id, agreement).then(() => history.push("/"));
    } else {
      createTenancyAgreement(agreement).then(() => history.push("/"));
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <Stack spacing={2}>
        <TextField
          required
          fullWidth
          label="Property"
          name="property"
          value={agreement.property}
          onChange={handleChange}
        />
        <TextField
          required
          fullWidth
          label="Landlord"
          name="landlord"
          value={agreement.landlord}
          onChange={handleChange}
        />
        <TextField
          required
          fullWidth
          label="Tenant"
          name="tenant"
          value={agreement.tenant}
          onChange={handleChange}
        />
        <TextField
          required
          fullWidth
          type="date"
          label="Start Date"
          name="start_date"
          value={agreement.start_date}
          onChange={handleChange}
        />
        <TextField
          required
          fullWidth
          type="date"
          label="End Date"
          name="end_date"
          value={agreement.end_date}
          onChange={handleChange}
        />
        <Button type="submit" variant="contained">
          {id ? "Update Agreement" : "Create Agreement"}
        </Button>
      </Stack>
    </form>
  );
}

export default TenancyAgreementForm;
