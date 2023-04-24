import React, { useState } from 'react';

function ContractForm({ onCreate }) {
  const [propertyId, setPropertyId] = useState('');
  const [tenant, setTenant] = useState('');
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');

  const handleSubmit = (event) => {
    event.preventDefault();
    const contract = {
      property_id: propertyId,
      tenant: tenant,
      start_date: startDate,
      end_date: endDate,
    };
    onCreate(contract);
    setPropertyId('');
    setTenant('');
    setStartDate('');
    setEndDate('');
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className="mb-3">
        <label htmlFor="propertyId" className="form-label">
          Property ID
        </label>
        <input
          type="text"
          className="form-control"
          id="propertyId"
          value={propertyId}
          onChange={(event) => setPropertyId(event.target.value)}
        />
      </div>
      <div className="mb-3">
        <label htmlFor="tenant" className="form-label">
          Tenant
        </label>
        <input
          type="text"
          className="form-control"
          id="tenant"
          value={tenant}
          onChange={(event) => setTenant(event.target.value)}
        />
      </div>
      <div className="mb-3">
        <label htmlFor="startDate" className="form-label">
          Start Date
        </label>
        <input type="date"
      className="form-control"
      id="startDate"
      value={startDate}
      onChange={(event) => setStartDate(event.target.value)}
    />
  </div>
  <div className="mb-3">
    <label htmlFor="endDate" className="form-label">
      End Date
    </label>
    <input
      type="date"
      className="form-control"
      id="endDate"
      value={endDate}
      onChange={(event) => setEndDate(event.target.value)}
    />
  </div>
  <button type="submit" className="btn btn-primary">
    Create
  </button>
</form>
);
}

export default ContractForm;
