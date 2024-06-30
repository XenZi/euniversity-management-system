import React, { useEffect, useState } from "react";
import { Dorm } from "../../models/dorm.model";
import { useDispatch, useSelector } from "react-redux";
import { useModalContext } from "../../context/modal.context";
import { Application } from "../../models/application.model";
import { axiosInstance } from "../../services/axios.service";
import { RootState } from "../../redux/store/store";
import { Admission } from "../../models/admission.model";
import {
  castFromApplicationStatusToActualString,
  castFromApplicationTypeNumberToActualString,
} from "../../utils/converter.utils";
import { closeModal, setModalOpen } from "../../redux/slices/modal.slice";
import DeleteDialog from "../dialogs/delete-dialog/delete-dialog.component";

const UserApplicationsTable = () => {
  const [dormList, setDormList] = useState<Dorm[]>([]);
  const [dormitoryAdmissions, setDormitoryAdmissions] = useState<Admission[]>(
    []
  );
  const [applications, setApplications] = useState<Application[]>([]);
  const dispatch = useDispatch();
  const user = useSelector((state: RootState) => state.user.user);
  const { setContent } = useModalContext();

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
      .get(`/dorm/applications/pin/${user?.personalIdentificationNumber}`)
      .then((data) => {
        console.log(data);
        setApplications(data.data.data);
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

  const openDialogForDelete = (applicationID: string) => {
    dispatch(closeModal());
    dispatch(setModalOpen());
    setContent(
      <DeleteDialog
        functionToProceedOnDelete={() => {
          deleteApplication(applicationID);
        }}
      />
    );
  };

  const deleteApplication = (applicationID: string) => {
    axiosInstance
      .delete(`/dorm/applications/${applicationID}`)
      .then((data) => {
        console.log(data);
      })
      .catch((err) => {
        console.log(err);
      });
    dispatch(closeModal());
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

  const sendRequestForUpdateInformationOfMyRequests = (app: Application) => {
    axiosInstance
      .put(`/dorm/applications/${app.id}`, {
        ...app,
      })
      .then((data) => {
        setApplications((prevApplications) => {
          const index = prevApplications.findIndex((app) => app.id === app.id);
          if (index === -1) return prevApplications;

          const newApplications = [...prevApplications];
          newApplications[index] = data.data.data;

          return newApplications;
        });
        console.log(data.data.data);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  return (
    <>
      <>
        <h2 className="text-3xl text-center mb-3 font-bold">My applications</h2>
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
                        {application.healthInsurance
                          ? "Insured"
                          : "Not insured"}
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
                        <button
                          className="border bg-auburn-500 border-auburn-500 font-semibold py-1 px-2 rounded focus:border-auburn-700 text-white"
                          onClick={(e) => {
                            e.preventDefault();
                            sendRequestForUpdateInformationOfMyRequests(
                              application
                            );
                          }}
                        >
                          Resend this application
                        </button>
                      </td>
                      <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                        <button
                          className="border bg-auburn-500 border-auburn-500 font-semibold py-1 px-2 rounded focus:border-auburn-700 text-white"
                          onClick={(e) => {
                            e.preventDefault();
                            openDialogForDelete(application.id as string);
                          }}
                        >
                          Delete this application
                        </button>
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
    </>
  );
};

export default UserApplicationsTable;
