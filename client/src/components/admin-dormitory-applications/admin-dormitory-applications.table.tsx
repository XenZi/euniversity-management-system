import { ChangeEvent, useEffect, useState } from "react";
import { Admission } from "../../models/admission.model";
import { Application } from "../../models/application.model";
import { Dorm } from "../../models/dorm.model";
import { axiosInstance } from "../../services/axios.service";
import {
  castFromApplicationStatusToActualString,
  castFromApplicationTypeNumberToActualString,
} from "../../utils/converter.utils";

const AdminDormitoryApplicationTable = () => {
  const [dormList, setDormList] = useState<Dorm[]>([]);
  const [dormitoryAdmissions, setDormitoryAdmissions] = useState<Admission[]>(
    []
  );
  const [applications, setApplications] = useState<Application[]>([]);
  const [selectedAdmission, setSelectedAdmission] = useState<Admission>();

  useEffect(() => {
    axiosInstance
      .get("/dorm/all")
      .then((data) => {
        setDormList(data.data.data);
      })
      .catch((err) => {
        console.log(err);
      });

    axiosInstance
      .get(`/dorm/admissions/all`)
      .then((data) => {
        setDormitoryAdmissions(data.data.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  useEffect(() => {
    if (selectedAdmission != null) {
      axiosInstance
        .get(`/dorm/applications/admission/${selectedAdmission?.id}`)
        .then((data) => {
          console.log(data.data.data);
          setApplications(data.data.data);
        })
        .catch((err) => {
          console.log(err);
        });
    }
  }, [selectedAdmission]);

  const handleSelectChange = (e: ChangeEvent<HTMLSelectElement>) => {
    const selectedId = e.target.value;
    const admission = dormitoryAdmissions.find((d) => d.id === selectedId);
    setSelectedAdmission(admission);
  };

  const findDormInDormListByID = (id: string): Dorm | undefined => {
    const foundDorm = dormList.find((dorm) => dorm.id == id);
    return foundDorm;
  };

  const findAdmissionInAdmissionsList = (id: string): Admission | undefined => {
    const foundAdmission = dormitoryAdmissions.find(
      (admission) => admission.id == id
    );
    return foundAdmission;
  };

  return (
    <>
      <h3 className="text-center font-bold text-3xl mb-2">
        All applications for one admission
      </h3>
      <form className="flex flex-col">
        <label htmlFor="options">Select a dorm:</label>
        <select
          id="options"
          className="mb-3 p-3 border-2 border-battleship-500 w-full"
          onChange={handleSelectChange}
        >
          <option value="" disabled selected>
            Select an option
          </option>
          {dormitoryAdmissions.map((admission, index) => {
            const dorm = findDormInDormListByID(admission.dormID) as Dorm;
            return (
              <option key={index} value={admission.id}>
                {dorm.name}, {dorm.location} - {admission.start} -{" "}
                {admission.end}
              </option>
            );
          })}
        </select>
      </form>
      <div className="overflow-x-auto">
        <table className="min-w-full bg-white">
          <thead>
            <tr>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                For dorm
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Start Date
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                End Date
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Health insurance
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Verified student
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Application Status
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Application Type
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Student
              </th>
            </tr>
          </thead>
          <tbody>
            {applications && applications.length > 0 ? (
              applications.map((application, index) => {
                const foundAdmission = findAdmissionInAdmissionsList(
                  application?.dormitoryAdmissionsID as string
                ) as Admission;
                const foundDorm = findDormInDormListByID(
                  foundAdmission?.dormID as string
                );
                return (
                  <tr key={index}>
                    <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                      {foundDorm?.name} - {foundDorm?.location}
                    </td>
                    <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                      {foundAdmission.start}
                    </td>
                    <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                      {foundAdmission.end}
                    </td>
                    <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                      {application.healthInsurance ? "Insured" : "Not insured"}
                    </td>
                    <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                      {application.verifiedStudent
                        ? "Verified"
                        : "Not verified"}
                    </td>
                    <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                      {castFromApplicationStatusToActualString(
                        application.applicationStatus as number
                      )}
                    </td>
                    <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                      {castFromApplicationTypeNumberToActualString(
                        application.applicationType as number
                      )}
                    </td>
                    <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                      {application.student?.fullName}
                    </td>
                  </tr>
                );
              })
            ) : (
              <tr>
                <td
                  colSpan={3}
                  className="py-2 px-4 border-b border-gray-300 text-sm text-center"
                >
                  No applications available
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </>
  );
};

export default AdminDormitoryApplicationTable;
