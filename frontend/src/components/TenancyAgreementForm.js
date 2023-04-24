import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import {
  Button,
  Container,
  FormControl,
  FormHelperText,
  InputLabel,
  MenuItem,
  Select,
  TextField,
  Typography,
} from "@mui/material";
import { createTenancyAgreement } from "../utils/api";
import HashBlock from "./HashBlock";

const TenancyAgreementForm = () => {
  const [property, setProperty] = useState("");
  const [landlord, setLandlord] = useState("");
  const [tenant, setTenant] = useState("");
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");
  const [error, setError] = useState("");
  const [hash, setHash] = useState("");
  const history = useNavigate();

  const [latestHash, setLatestHash] = useState("");

  const fetchLatestHash = async () => {
    try {
      const response = await fetch("/api/latest-hash");
      const data = await response.json();
      setLatestHash(data.latestHash);
    } catch (error) {
      console.log("Error fetching latest hash:", error);
    }
  };
  

useEffect(() => {
  fetchLatestHash();
}, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const data = {
        property,
        landlord,
        tenant,
        start_date: startDate,
        end_date: endDate,
      };
      const response = await createTenancyAgreement(data);
      console.log("Transaction Hash:", response.transaction.txid); // Log the transaction hash
      setHash(response.transaction.txid);
      setLatestHash(response.hash);
    } catch (error) {
      console.log(error);
    }
  };
  
  

  return (
    <Container maxWidth="sm">
      <Typography variant="h4" align="center" sx={{ mt: 2, mb: 4 }}>
        Create Tenancy Agreement
      </Typography>
      <form onSubmit={handleSubmit}>
        <FormControl fullWidth sx={{ mb: 2 }}>
          <InputLabel id="property-label">Property</InputLabel>
          <Select
            labelId="property-label"
            id="property"
            value={property}
            label="Property"
            onChange={(e) => setProperty(e.target.value)}
          >
            <MenuItem value={"123 Main St"}>123 Main St</MenuItem>
            <MenuItem value={"456 Elm St"}>456 Elm St</MenuItem>
            <MenuItem value={"789 Oak St"}>789 Oak St</MenuItem>
          </Select>
        </FormControl>
        <TextField
          fullWidth
          id="landlord"
          label="Landlord"
          value={landlord}
          onChange={(e) => setLandlord(e.target.value)}
          sx={{ mb: 2 }}
        />
        <TextField
          fullWidth
          id="tenant"
          label="Tenant"
          value={tenant}
          onChange={(e) => setTenant(e.target.value)}
          sx={{ mb: 2 }}
        />
        <TextField
          fullWidth
          id="start-date"
          label="Start Date"
          type="date"
          value={startDate}
          onChange={(e) => setStartDate(e.target.value)}
          sx={{ mb: 2 }}
          InputLabelProps={{
            shrink: true,
          }}
        />
        <TextField
          fullWidth
          id="end-date"
          label="End Date"
          type="date"
          value={endDate}
          onChange={(e) => setEndDate(e.target.value)}
          sx={{ mb: 2 }}
          InputLabelProps={{
            shrink: true,
          }}
        />
        {error && (
          <FormHelperText error sx={{ mb: 2 }}>
            {error}
          </FormHelperText>
        )}
        <HashBlock hash={hash}/>
        <Button variant="contained" fullWidth type="submit">
          Create
        </Button>
      </form>
    </Container>
  );
};

export default TenancyAgreementForm;
