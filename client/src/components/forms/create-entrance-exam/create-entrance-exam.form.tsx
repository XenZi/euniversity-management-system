import { useState, useEffect } from "react";
import { axiosInstance } from "../../../services/axios.service";
import { University } from "../../../models/university.model";

import { useSelector } from "react-redux";
import { RootState } from "../../../redux/store/store";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const CreateEntranceForm = () => {
  const user = useSelector((state: RootState) => state.user.user);
  const [formVisible, setFormVisible] = useState(true);
  const [loadedUniversities, setLoadedUniversities] = useState<University[]>([]);
  const [selectedUniversity, setSelectedUniversity] = useState<University | null>(null);

  useEffect(() => {
    axiosInstance
      .get("/university/")
      .then((res) => {
        console.log(res);
        setLoadedUniversities(res.data.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (selectedUniversity) {
      axiosInstance
        .post("/university/entranceExam", {
          citizen: user,
          university: selectedUniversity,
        })
        .then((res) => {
          console.log("Submission successful", res);
          toast.success('Successfuly applied for entrance exam!');
          setFormVisible(false);

        })
        .catch((err) => {
          console.log("Submission error", err);
          toast.error('Something went wrong!');

        });
    }
  };

  const handleUniversityChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedId = e.target.value;
    const university = loadedUniversities.find((uni) => uni.id === selectedId);
    setSelectedUniversity(university || null);
    
  };
  if (!formVisible) {
    return <div>Form submitted successfully!</div>;
  }

  return (
    <>
      <div>Create Entrance Form</div>
      <form className="w-full" onSubmit={handleSubmit}>
        <label>Choose University</label>
        <select
          id="university-select"
          className="mb-3 p-3 border-2 border-battleship-500 w-full"
          value={selectedUniversity?.id || ""}
          onChange={handleUniversityChange}
        >
          <option value="" disabled>
            Select a university
          </option>
          {loadedUniversities?.map((uni) => (
            <option value={uni.id} key={uni.id}>
              {uni.name}
            </option>
            
          ))}
          
        </select>
        <button
          className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white w-full"
          type="submit"
        >
          Apply
        </button>
      </form>
    </>
  );
};

export default CreateEntranceForm;
