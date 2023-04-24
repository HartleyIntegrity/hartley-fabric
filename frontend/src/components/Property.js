import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { getProperty } from "../utils/api";
import { Typography, Card, CardMedia, CardContent } from "@mui/material";

function Property() {
  const [property, setProperty] = useState({});
  const { id } = useParams();

  useEffect(() => {
    getProperty(id).then((data) => setProperty(data));
  }, [id]);

  return (
    <div>
      <Typography variant="h2" component="h2" gutterBottom>
        {property.name}
      </Typography>
      <Card sx={{ maxWidth: 345 }}>
        <CardMedia component="img" height="140" image={property.image} alt={property.name} />
        <CardContent>
          <Typography variant="body2" color="text.secondary">
            {property.description}
          </Typography>
        </CardContent>
      </Card>
    </div>
  );
}

export default Property;
