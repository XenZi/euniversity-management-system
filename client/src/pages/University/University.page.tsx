import React, { useEffect, useState } from "react";
import AdminComponent from "../../components/admin/admin.component";
import CreateUniForm from "../../components/forms/createUniForm/create-uni.form";
import Navigation from "../../components/navigation/navigation.component";
import PanelBox from "../../components/panel-box/panel-box.component";
import { useModalContext } from "../../context/modal.context";
import { setModalOpen } from "../../redux/slices/modal.slice";
import { RootState } from "../../redux/store/store";
import { useDispatch, useSelector } from "react-redux";
import { axiosInstance } from "../../services/axios.service";
import { useNavigate } from "react-router-dom";
import { University } from "../../models/university.model";
import CreateEntranceForm from "../../components/forms/create-entrance-exam/create-entrance-exam.form";
import CreateStudentForm from "../../components/forms/create-student-form/create-student.form";
import AllUniversities from "../../components/universities-table/all-universities.table";





const UniPage = () => {
  const user = useSelector((state: RootState) => state.user.user);
  
  const dispatch = useDispatch();

  const { setContent } = useModalContext();
  const openModal = () => {
    dispatch(setModalOpen());
  };
 
  const uniAdminComponents: React.JSX.Element[] = [
    <PanelBox
      key="create-uni"
      panelBoxTitle="Create New University"
      onClick={() => {
        openModal();
        setContent(<CreateUniForm />);
      }}
    />,  
    <PanelBox
      key="create-student"
      panelBoxTitle="Create New Student"
      onClick={() => {
        openModal();
        setContent(<CreateStudentForm></CreateStudentForm>)
      }}
    >  
    </PanelBox>,
    <PanelBox
      key="all-unies"
      panelBoxTitle="All Universities"
      onClick={() => {
        openModal();
        setContent(<AllUniversities></AllUniversities>)
      }}
    ></PanelBox>
  ];

  const citizenComponents: React.JSX.Element[] = [
    
    <PanelBox 
      key="create-entrance-exam"
      panelBoxDescription="Apply for entranece exam"
      onClick={() => {
        openModal();
        setContent(<CreateEntranceForm></CreateEntranceForm>)
      }}        
      ></PanelBox>
  ]

  const studentComponents: React.JSX.Element[] = [
    <PanelBox
      key="profile"
      panelBoxTitle="Profile"
      onClick={() => {
        openModal();
        setContent(<CreateUniForm />);
      }}
    />
  ];

  return (
    <div className="h-screen bg-papaya-500 w-full p-3">
      <Navigation />
      <div className="max-w-7xl mx-auto w-100 flex">
        {user?.roles[0] === "Citizen" ? (
          <AdminComponent>{uniAdminComponents}</AdminComponent>
        ) : (
          <div>No admin access</div>
        )}
      </div>
      <div className="max-w-7xl mx-auto w-100 flex">
        {user?.roles[0] === "Student" ? (
          <AdminComponent>{studentComponents}</AdminComponent>
        ) : (
          <div>No admin access</div>
        )}
      </div>
      <div className="max-w-7xl mx-auto w-100 flex">
        {user?.roles[0] === "Citizen" ? (
          <AdminComponent>{citizenComponents}</AdminComponent>
        ) : (
          <div>No admin access</div>
        )}
      </div>
    </div>
  );
};

export default UniPage;
