import { useEffect, useState } from "react";
import { ExtendStatusApplication } from "../../../models/extend-status-application.model";
import { User } from "../../../models/user.model";
import { axiosInstance } from "../../../services/axios.service";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

const ConfirmExtendStatus = () => {
  const [loadedApplications, setLoadedApplications] = useState<
    ExtendStatusApplication[]
  >([]);
  const [selectedUser, setSelectedUser] = useState<User | null>(null);

  useEffect(() => {
    axiosInstance
      .get("/university/extendStatusApplication")
      .then((res) => {
        console.log("BLABLA", res);
        setLoadedApplications(res.data.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (selectedUser) {
      axiosInstance
        .put(
          `/university/student/status1/${selectedUser.personalIdentificationNumber}`
        )
        .then((res) => {
          console.log("Submission successful", res);
          toast.success("Status Extended");
        })
        .catch((err) => {
          console.log("Extention Error", err);
          toast.error(
            "Something went wrong! Student doesn't have a medical exam!"
          );
        });
    }
  };

  const handleUserChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedId = e.target.value;
    const exam = loadedApplications.find(
      (application) => application.id === selectedId
    );
    setSelectedUser(exam?.citizen || null);
  };

  return (
    <>
      <div>Extend Status</div>
      <form className="w-full" onSubmit={handleSubmit}>
        <label>Choose</label>
        <select
          id="select-student-extend"
          className="mb-3 p-3 border-2 border-battleship-500 w-full"
          onChange={handleUserChange}
        >
          <option value="" disabled selected>
            Select a student
          </option>
          {loadedApplications?.map((application) => (
            <option value={application.id} key={application.id}>
              {application.citizen.fullName} -{" "}
              {application.citizen.personalIdentificationNumber}
            </option>
          ))}
        </select>
        <button
          className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white w-full"
          type="submit"
        >
          Submit
        </button>
      </form>
    </>
  );
};
export default ConfirmExtendStatus;
