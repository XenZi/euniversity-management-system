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
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import ConfirmExtendStatus from "../../components/forms/confirm-extend-status/confirm-extend-status.form";
import CreateScholarship from "../../components/forms/create-scholarship/create-scholarship.form";
import jsPDF from 'jspdf';
import ProfilePage from "../../components/forms/profile/profile-page.form";



const UniPage = () => {
  const user = useSelector((state: RootState) => state.user.user);
  
  const dispatch = useDispatch();

  const { setContent } = useModalContext();
  const openModal = () => {
    dispatch(setModalOpen());
  };

 
    const handleDownloadPDF = () => {
        // Create a new jsPDF instance
        const doc = new jsPDF();

        // Add text to the PDF document
        doc.text(`Student Confirmation for ${user?.fullName}`, 10, 10);

        // Save the PDF
        doc.save(`student-confirmation-${user?.fullName}.pdf`);
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
    ></PanelBox>,
    <PanelBox
      key="confirm-application"
      panelBoxTitle="Confirm Extend Status"
     onClick={() => {
      openModal();
      setContent(<ConfirmExtendStatus></ConfirmExtendStatus>)
     }}
    >
    </PanelBox>,
    <PanelBox key="confirm-scholarship-application" panelBoxTitle="Confirm Scholarship" onClick={() =>{
      openModal();
      setContent(<CreateScholarship></CreateScholarship>)
    }}>

    </PanelBox>

  

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
        setContent(<ProfilePage />);
      }}
    />,
    <PanelBox
      key="extend-status-application"
      panelBoxTitle="Extend Status"
      onClick={() => {
        axiosInstance
          .post("/university/extendStatusApplication",{
            citizen: user
          
          })
          .then((resp) => {
            console.log(resp)
            toast.success('Successfuly applied for extending student status!');
          })
          .catch((error) => {
            console.error(error);
            toast.error('Something went wrong!');
          });
          
      }}
      >

    </PanelBox>,
    <PanelBox key="scholarship-application" panelBoxTitle="Apply for scholarship" onClick={()=>{
      axiosInstance
      .post("/university/scholarshipApplication",{
        student: user
      })
      .then((resp) =>{
        console.log(resp)
        toast('Scholarship application successful!')
      })
      .catch((err) =>{
        console.error(err);
        toast.error('Something went wrong!')
      })
    }}>

    </PanelBox>,

        <PanelBox
        key="student-confirmation"
        panelBoxTitle="Download Student Confirmation"
        onClick={handleDownloadPDF}
        >
        
        </PanelBox>

    
    
  ];

  return (
    <div className="h-screen bg-papaya-500 w-full p-3">
      <Navigation />
      <div className="max-w-7xl mx-auto w-100 flex">
        {user?.roles[0] === "Admin" ? (
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
      <ToastContainer />
    </div>
  );
};

export default UniPage;
