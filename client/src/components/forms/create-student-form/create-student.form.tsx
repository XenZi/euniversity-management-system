import React, { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../../redux/store/store';
import { UniversityAdmission } from '../../../models/university-admission.model';
import { axiosInstance } from '../../../services/axios.service';
import { User } from '../../../models/user.model';
import { University } from '../../../models/university.model';

const CreateStudentForm = () => {
  const [loadedEntranceExams, setLoadedEntranceExams] = useState<UniversityAdmission[]>([]);
  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  const [formVisible, setFormVisible] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [selectedUni, setSelectedUni] = useState<University| null>(null);

  

  useEffect(() => {
    axiosInstance
      .get("/university/entranceExam")
      .then((res) => {
        console.log(res);
        setLoadedEntranceExams(res.data.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (selectedUser) {
      axiosInstance
        .post("/university/student", {
          personalIdentificationNumber: selectedUser.personalIdentificationNumber,
          citizenship: selectedUser.citizenship,
          fullName: selectedUser.fullName,
          gender: selectedUser.gender,
          identityCardNumber: selectedUser.identityCardNumber,
          residence: selectedUser.residence,
          university: selectedUni
        })
        
        .then((res) => {
          console.log("Submission successful", res);
          setFormVisible(false);
          setError(null);
        })
        .catch((err) => {
          console.log("Submission error", err);
          setError(err.response?.data?.message || "An error occurred during submission. Please try again.");
        });
    }
  };

  const handleUserChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedId = e.target.value;
    const exam = loadedEntranceExams.find((exam) => exam.id === selectedId);
    setSelectedUser(exam?.citizen || null);
    setSelectedUni(exam?.university || null)
  };

  if (!formVisible) {
    return <div>Form submitted successfully!</div>;
  }

  return (
    <>
      <div>Create Student Form</div>
      {error && <div className="error-message">{error}</div>}
      <form className="w-full" onSubmit={handleSubmit}>
        <label>Choose</label>
        <select
          id="select-student"
          className="mb-3 p-3 border-2 border-battleship-500 w-full"
          onChange={handleUserChange}
        >
          <option value="" disabled>
            Select a student
          </option>
          {loadedEntranceExams?.map((exam) => (
            <option value={exam.id} key={exam.id}>
              {exam.citizen.fullName} - {exam.university.name}
            </option>
          ))}
        </select>
        <button
          className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white w-full"
          type="submit"
        >
          Create
        </button>
      </form>
    </>
  );
};

export default CreateStudentForm;
