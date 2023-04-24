import { useState, useEffect } from "react";
import { Typography } from "@mui/material";
import axios from "axios";

const HashBlock = () => {
  const [currentHash, setCurrentHash] = useState("");
  const [isHashGenerated, setIsHashGenerated] = useState(false);

  let intervalId;

  useEffect(() => {
    const generateHash = () => {
      const nextChar = String.fromCharCode(
        Math.floor(Math.random() * 94) + 32
      );
      setCurrentHash((prev) => prev + nextChar);
    };

    axios.get("/api/latest-hash")
      .then(response => {
        const hash = response.data.latestHash;
        const intervalId = setInterval(() => {
          if (currentHash === hash) {
            clearInterval(intervalId);
            setIsHashGenerated(true);
            axios.post('/api/tenancy-agreements', { hashBlock: currentHash })
              .then(response => {
                console.log('Updated tenancy agreement with hash block:', response.data);
              })
              .catch(error => {
                console.log('Error updating tenancy agreement:', error);
              });
          } else {
            generateHash();
          }
        }, 20);
      })
      .catch(error => {
        console.log('Error fetching latest hash:', error);
      });

    return () => clearInterval(intervalId);
  }, [currentHash]);

  return (
    <div
      style={{
        width: "100%",
        height: "50px",
        background: "#eee",
        borderRadius: "4px",
        position: "relative",
        overflow: "hidden",
        margin: "16px 0",
        opacity: isHashGenerated ? "1" : "0",
        transition: "opacity 0.5s ease-in-out",
      }}
    >
      <div
        style={{
          width: isHashGenerated ? "100%" : "0",
          height: "100%",
          background: "#5E9C5E",
          position: "absolute",
          top: "0",
          left: "0",
          transition: "all 2s ease-in-out",
        }}
      ></div>
      <Typography
        variant="h6"
        align="center"
        sx={{
          position: "absolute",
          top: "50%",
          left: "50%",
          transform: "translate(-50%, -50%)",
          width: "100%",
          zIndex: "10",
          color: "#333",
        }}
      >
        {currentHash}
      </Typography>
    </div>
  );
};

export default HashBlock;
