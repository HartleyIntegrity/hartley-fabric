import React, { useState, useEffect } from "react";
import axios from "axios";
import BlockList from "../blocks/BlockList";

const AdminDashboard = () => {
  const [blocks, setBlocks] = useState([]);

  useEffect(() => {
    fetchBlocks();
  }, []);

  const fetchBlocks = async () => {
    const response = await axios.get("http://localhost:8000/blocks");
    setBlocks(response.data);
  };

  return (
    <div>
    <h2>Admin Dashboard</h2>
    <BlockList blocks={blocks} />
    </div>
    );
};
    
    export default AdminDashboard;
