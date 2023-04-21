import React from "react";

const BlockList = ({ blocks }) => {
  return (
    <div>
      <h3>Blocks</h3>
      <table>
        <thead>
          <tr>
            <th>Index</th>
            <th>Timestamp</th>
            <th>Data</th>
            <th>Previous Hash</th>
            <th>Hash</th>
          </tr>
        </thead>
        <tbody>
          {blocks.map((block, index) => (
            <tr key={index}>
              <td>{block.Index}</td>
              <td>{block.Timestamp}</td>
              <td>{block.Data}</td>
              <td>{block.PrevHash}</td>
              <td>{block.Hash}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default BlockList;
