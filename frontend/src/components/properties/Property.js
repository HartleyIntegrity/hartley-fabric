import React from "react";

const Property = ({ property }) => {
  return (
    <li>
      <h3>{property.title}</h3>
      <p>Location: {property.location}</p>
      <p>Price: ${property.price}</p>
    </li>
  );
};

export default Property;
