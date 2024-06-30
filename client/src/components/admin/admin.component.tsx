import React, { FC } from "react";

interface AdminComponentInterface {
  children: React.JSX.Element[];
}
const AdminComponent: FC<AdminComponentInterface> = ({ children }) => {
  return (
    <>
      {children.map((el) => {
        return el;
      })}
    </>
  );
};

export default AdminComponent;
