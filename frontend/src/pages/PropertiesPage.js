import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import PropertyCard from '../components/PropertyCard';
import { getProperties } from '../services/properties';

function PropertiesPage() {
  const [properties, setProperties] = useState([]);

  useEffect(() => {
    getProperties().then((data) => setProperties(data));
  }, []);

  return (
    <div className="container">
      <h1 className="my-5">Properties</h1>
      <div className="row">
        {properties.map((property) => (
<div key={property.id} className="col-md-4 mb-5">
<PropertyCard property={property} />
</div>
))}
</div>
<Link to="/contracts" className="btn btn-primary">
View Contracts
</Link>
</div>
);
}

export default PropertiesPage;