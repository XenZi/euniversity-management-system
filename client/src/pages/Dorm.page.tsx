import React from "react";
import { useSelector } from "react-redux";
import { RootState } from "../redux/store/user.store";
import Navigation from "../components/navigation/navigation.component";

const DormPage = () => {
  const user = useSelector((state: RootState) => state.user.user);
  console.log(user);
  return (
    <div className="h-screen bg-papaya-500 w-full p-3">
      <Navigation />
      {user?.roles[0] == "Admin" ? "Admin" : "Student"}
    </div>
  );
};

export default DormPage;
