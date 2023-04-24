import React from 'react';
import { Card, CardContent, CardMedia, Typography } from '@mui/material';

const PropertyCard = ({ property }) => {
  return (
    <Card sx={{ maxWidth: 345 }}>
      <CardMedia
        component="img"
        height="140"
        image={property.image}
        alt="Property Image"
      />
      <CardContent>
        <Typography gutterBottom variant="h5" component="div">
          {property.name}
        </Typography>
        <Typography variant="body2" color="text.secondary">
          {property.description}
        </Typography>
      </CardContent>
    </Card>
  );
};

export default PropertyCard;
