import React from 'react';
import { useLocation } from 'react-router-dom';

const StudentProfilePage: React.FC = () => {
  const location = useLocation();
  const student = location.state?.student;

  if (!student) {
    return <div>No student data available</div>;
  }

  return (
    <div className="profile-page">
      <h2 className="text-center text-3xl font-semibold mb-3">Student Profile</h2>
      <p><strong>Full Name:</strong> {student.fullName}</p>
      <p><strong>Gender:</strong> {student.gender}</p>
      <p><strong>Identity Card Number:</strong> {student.identityCardNumber}</p>
      <p><strong>Citizenship:</strong> {student.citizenship}</p>
      <p><strong>Personal Identification Number:</strong> {student.personalIdentificationNumber}</p>
      <p><strong>Email:</strong> {student.email}</p>
      {/* Add more fields as necessary */}
      <h3>Residence</h3>
      <p><strong>Address:</strong> {student.residence.address}</p>
      <p><strong>Place of Residence:</strong> {student.residence.placeOfResidence}</p>
      <p><strong>Municipality of Residence:</strong> {student.residence.municipalityOfResidence}</p>
      <p><strong>Country of Residence:</strong> {student.residence.countryOfResidence}</p>
      <h3>Birth Data</h3>
      <p><strong>Date of Birth:</strong> {student.birthData.dateOfBirth}</p>
      <p><strong>Municipality of Birth:</strong> {student.birthData.municapilityOfBirth}</p>
      <p><strong>Country of Birth:</strong> {student.birthData.countryOfBirth}</p>
      <h3>University</h3>
      <p><strong>ID:</strong> {student.university.id}</p>
      <p><strong>Name:</strong> {student.university.name}</p>
      <p><strong>Address:</strong> {student.university.address}</p>
    </div>
  );
};

export default StudentProfilePage;
