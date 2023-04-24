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
          {agreements.map((agreement) => (
            <TableRow key={agreement.id}>
              <TableCell>{agreement.id}</TableCell>
              <TableCell>
                <Link to={`/property/${agreement.property}`}>{agreement.property}</Link>
              </TableCell>
              <TableCell>{agreement.landlord}</TableCell>
              <TableCell>{agreement.tenant}</TableCell>
              <TableCell>{agreement.start_date}</TableCell>
              <TableCell>{agreement.end_date}</TableCell>
              <TableCell>
                <Link to={`/update/${agreement.id}`}>Edit</Link>
                <IconButton onClick={() => handleDelete(agreement.id)} aria-label="delete">
                  <DeleteIcon />
                </IconButton>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}

export default TenancyAgreementList;
