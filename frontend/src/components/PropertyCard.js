import React from 'react';

function PropertyCard({ property }) {
  return (
    <div className="card h-100">
      <img src={property.image} className="card-img-top" alt={property.address} />
      <div className="card-body">
        <h5 className="card-title">{property.address}</h5>
        <p className="card-text">${property.rent} / month</p>
      </div>
    </div>
  );
}

export default PropertyCard;