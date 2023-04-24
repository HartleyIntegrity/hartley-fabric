import React from 'react';

function ContractTable({ contracts, onDelete }) {
  return (
    <table className="table table-striped">
      <thead>
        <tr>
          <th>ID</th>
          <th>Property ID</th>
          <th>Tenant</th>
          <th>Start Date</th>
          <th>End Date</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        {contracts.map((contract) => (
          <tr key={contract.id}>
            <td>{contract.id}</td>
            <td>{contract.property_id}</td>
            <td>{contract.tenant}</td>
            <td>{contract.start_date}</td>
            <td>{contract.end_date}</td>
            <td>
              <button
                className="btn btn-danger"
                onClick={() => onDelete(contract.id)}
              >
                Delete
              </button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}

export default ContractTable;
