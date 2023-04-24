import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { getTenancyAgreements, deleteTenancyAgreement } from "../utils/api";
import { Table, TableHead, TableBody, TableRow, TableCell, IconButton } from "@mui/material";
import DeleteIcon from '@mui/icons-material/Delete';

function TenancyAgreementList() {
  const [agreements, setAgreements] = useState([]);

  useEffect(() => {
    getTenancyAgreements().then((data) => setAgreements(data));
}, []);


  const handleDelete = (id) => {
    deleteTenancyAgreement(id).then(() => {
      setAgreements(agreements.filter((agreement) => agreement.id !== id));
    });
  };

  return (
    <div>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>ID</TableCell>
            <TableCell>Property</TableCell>
            <TableCell>Landlord</TableCell>
            <TableCell>Tenant</TableCell>
            <TableCell>Start Date</TableCell>
            <TableCell>End Date</TableCell>
            <TableCell>Actions</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {agreements.length > 0 ? (
            agreements.map((agreement) => (
              <TableRow key={agreement.ID}>
                <TableCell>{agreement.ID}</TableCell>
                <TableCell>
                  <Link to={`/property/${agreement.Property}`}>{agreement.Property}</Link>
                </TableCell>
                <TableCell>{agreement.Landlord}</TableCell>
                <TableCell>{agreement.Tenant}</TableCell>
                <TableCell>{agreement.StartDate}</TableCell>
                <TableCell>{agreement.EndDate}</TableCell>
                <TableCell>
                  <Link to={`/update/${agreement.ID}`}>Edit</Link>
                  <IconButton onClick={() => handleDelete(agreement.ID)} aria-label="delete">
                    <DeleteIcon />
                  </IconButton>
                </TableCell>
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={7} align="center">No tenancy agreements found.</TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}

export default TenancyAgreementList;
