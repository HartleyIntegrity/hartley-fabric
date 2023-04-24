import React, { useState, useEffect } from 'react';
import { getTransactions } from '../api';

const Properties = () => {
  const [transactions, setTransactions] = useState([]);

  useEffect(() => {
    const fetchTransactions = async () => {
      const transactions = await getTransactions();
      setTransactions(transactions);
    };

    fetchTransactions();
  }, []);

  return (
    <div>
      <h1>Properties</h1>
      {transactions.map((transaction, index) => (
        <div key={index}>
          <h2>{transaction.Details.PropertyID}</h2>
          <p>Price: {transaction.Details.Price}</p>
          <p>Rent Price: {transaction.Details.RentPrice}</p>
          <p>Maintenance Fee: {transaction.Details.MaintenanceFee}</p>
          <p>Start Date: {transaction.Details.StartDate}</p>
          <p>End Date: {transaction.Details.EndDate}</p>
        </div>
      ))}
    </div>
  );
};

export default Properties;
