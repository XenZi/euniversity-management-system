import { RootState } from "@reduxjs/toolkit/query";
import { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useModalContext } from "../../context/modal.context";
import { setModalOpen } from "../../redux/slices/modal.slice";
import { axiosInstance } from "../../services/axios.service";
import Navigation from "../../components/navigation/navigation.component";
import AdminComponent from "../../components/admin/admin.component";

const FoodPage = () => {
  const dispatch = useDispatch();
  const { setContent } = useModalContext();
  
  const openModal = () => {
    dispatch(setModalOpen());
  };

  const foodAdminComponents: React.JSX.Element[] = [
    // Add components for admin here
  ];

  const foodStudentComponents: React.JSX.Element[] = [
    // Add components for students here
  ];
  return (
    <div className="h-screen bg-papaya-500 w-full p-3">
      <Navigation />
      <div className="max-w-7xl mx-auto w-full flex">
        
      </div>
    </div>
  );
};

export default FoodPage;
