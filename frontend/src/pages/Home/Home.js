import React from 'react';
import { Box, Grid } from '@mui/material';
import { makeStyles } from '@mui/styles';
import { Link } from 'react-router-dom';
import { PropertyCardHorizontal } from '../../components/PropertyCard';
import properties from '../../data/properties';

const useStyles = makeStyles(() => ({
  title: {
    margin: '2rem',
  },
}));

function Home() {
  const classes = useStyles();

  return (
    <div>
      <Typography variant="h4" className={classes.title}>
        Properties
      </Typography>
      <Grid container spacing={3} justifyContent="center">
        {properties.map((property) => (
          <Grid item xs={12} sm={6} md={4} key={property.id}>
            <Link to={`/edit/${property.id}`}>
              <PropertyCard property={property} />
            </Link>
          </Grid>
        ))}
      </Grid>
    </div>
  );
}

export default Home;
