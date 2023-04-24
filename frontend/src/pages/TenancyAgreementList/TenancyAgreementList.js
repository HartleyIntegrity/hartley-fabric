import React, { useState, useEffect } from "react";
import { Grid, Typography } from "@mui/material";
import axios from "axios";
import TenancyAgreementCard from "../../components/TenancyAgreementCard/TenancyAgreementCard";

const TenancyAgreementList = () => {
  const [tenancyAgreements, setTenancyAgreements] = useState([]);

  useEffect(() => {
    axios
      .get("/api/tenancyAgreements")
      .then((response) => {
        setTenancyAgreements(response.data);
      })
      .catch((error) => {
        console.log(error);
      });
  }, []);

  return (
    <Grid container spacing={2}>
      <Grid item xs={12}>
        <Typography variant="h4">Tenancy Agreements</Typography>
      </Grid>
      {tenancyAgreements.map((tenancyAgreement) => (
        <Grid item xs={12} sm={6} md={4} key={tenancyAgreement.id}>
          <TenancyAgreementCard tenancyAgreement={tenancyAgreement} />
        </Grid>
      ))}
    </Grid>
  );
};

export default TenancyAgreementList;
