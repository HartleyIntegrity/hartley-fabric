import React, { useEffect, useState } from 'react';
import ContractTable from '../components/ContractTable';
import ContractForm from '../components/ContractForm';
import { getContracts, createContract } from '../services/contracts';

function ContractsPage() {
  const [contracts, setContracts] = useState([]);

  useEffect(() => {
    getContracts().then((data) => setContracts(data));
  }, []);

  const handleCreate = (contract) => {
    createContract(contract).then((data) => setContracts([...contracts, data]));
  };

  const handleDelete = (id) => {
    setContracts(contracts.filter((contract) => contract.id !== id));
  };

  return (
    <div className="container">
      <h1 className="my-5">Contracts</h1>
      <div className="row">
        <div className="col-md-8">
          <ContractTable contracts={contracts} onDelete={handleDelete} />
        </div>
        <div className="col-md-4">
          <ContractForm onCreate={handleCreate} />
        </div>
      </div>
    </div>
  );
}

export default ContractsPage;
