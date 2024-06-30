import { useEffect, useState } from "react";
import { Dorm } from "../../models/dorm.model";
import { axiosInstance } from "../../services/axios.service";
import { useDispatch } from "react-redux";
import { useModalContext } from "../../context/modal.context";
import { Admission } from "../../models/admission.model";
import { closeModal, setModalOpen } from "../../redux/slices/modal.slice";
import Admissions from "../forms/admissions/admissions.form";
import DeleteDialog from "../dialogs/delete-dialog/delete-dialog.component";

const AdmissionsTable: React.FC<{
  adminView?: boolean;
}> = ({ adminView }) => {
  const [loadedAdmissions, setLoadedAdmissions] = useState<Admission[]>();
  const [dormList, setDormList] = useState<Dorm[]>([]);
  const dispatch = useDispatch();
  const { setContent } = useModalContext();

  const openDialogForEdit = (admission: Admission) => {
    dispatch(closeModal());
    dispatch(setModalOpen());
    setContent(<Admissions admission={admission} />);
  };

  const openDialogForDelete = (admissionID: string) => {
    dispatch(closeModal());
    dispatch(setModalOpen());
    setContent(
      <DeleteDialog
        functionToProceedOnDelete={() => {
          deleteAdmission(admissionID);
        }}
      />
    );
  };
  const deleteAdmission = (admissionID: string) => {
    axiosInstance
      .delete(`/dorm/admissions/${admissionID}`)
      .then((data) => {
        console.log(data);
      })
      .catch((err) => {
        console.log(err);
      });
    dispatch(closeModal());
  };

  useEffect(() => {
    axiosInstance
      .get(`/dorm/admissions/all`)
      .then((data) => {
        console.log(data.data.data);
        setLoadedAdmissions(data.data.data);
      })
      .catch((err) => {
        console.log(err);
      });

    axiosInstance
      .get("/dorm/all")
      .then((data) => {
        setDormList(data.data.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  const findDormInDormListByID = (id: string): Dorm | undefined => {
    const foundDorm = dormList.find((dorm) => dorm.id == id);
    return foundDorm;
  };

  return (
    <>
      <h2 className="text-3xl text-center mb-3 font-bold">Admissions</h2>
      <div className="overflow-x-auto">
        <table className="min-w-full bg-white">
          <thead>
            <tr>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Dorm
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Start Date
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                End Date
              </th>
            </tr>
          </thead>
          <tbody>
            {loadedAdmissions && loadedAdmissions.length > 0 ? (
              loadedAdmissions.map((admission, index) => {
                const foundDorm = findDormInDormListByID(admission.dormID);
                return (
                  <tr key={index}>
                    <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                      {foundDorm?.name} - {foundDorm?.location}
                    </td>
                    <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                      {admission.start}
                    </td>
                    <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                      {admission.end}
                    </td>
                    {adminView ? (
                      <>
                        <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                          <button
                            className="border bg-auburn-500 border-auburn-500 font-semibold py-1 px-2 rounded focus:border-auburn-700 text-white"
                            onClick={(e) => {
                              e.preventDefault();
                              openDialogForEdit(admission);
                            }}
                          >
                            Edit admission
                          </button>
                        </td>
                        <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                          <button
                            className="border bg-auburn-500 border-auburn-500 font-semibold py-1 px-2 rounded focus:border-auburn-700 text-white"
                            onClick={(e) => {
                              e.preventDefault();
                              openDialogForDelete(admission.id);
                            }}
                          >
                            Delete admission
                          </button>
                        </td>
                        <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                          <button
                            className="border bg-auburn-500 border-auburn-500 font-semibold py-1 px-2 rounded focus:border-auburn-700 text-white"
                            onClick={(e) => {
                              e.preventDefault();
                            }}
                          >
                            End this admission
                          </button>
                        </td>
                      </>
                    ) : (
                      ""
                    )}
                  </tr>
                );
              })
            ) : (
              <tr>
                <td
                  colSpan={3}
                  className="py-2 px-4 border-b border-gray-300 text-sm text-center"
                >
                  No admissions available
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </>
  );
};

export default AdmissionsTable;
