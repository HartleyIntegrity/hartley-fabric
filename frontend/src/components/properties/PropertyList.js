import React from "react";
import Property from "./Property";

const PropertyList = ({ properties }) => {
  return (
    <div>
      <h2>Properties</h2>
      <ul>
        {properties.map((property, index) => (
          <Property key={index} property={property} />
        ))}
      </ul>
    </div>
  );
};

export default PropertyList;
