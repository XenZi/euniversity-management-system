import { useEffect, useState } from "react";

import { useDispatch } from "react-redux";
import { useModalContext } from "../../../context/modal.context";
import { axiosInstance } from "../../../services/axios.service";
import { closeModal, setModalOpen } from "../../../redux/slices/modal.slice";

import DeleteDialog from "../../dialogs/delete-dialog/delete-dialog.component";
import { Student } from "../../../models/student.model";

const StudentTable: React.FC<{ adminView?: boolean }> = ({ adminView }) => {
  const [students, setStudents] = useState<Student[]>([]);
  const dispatch = useDispatch();
  const { setContent } = useModalContext();

  useEffect(() => {
    axiosInstance
      .get("/food/allStudents")
      .then((data) => {
        console.log(data.data.data);
        setStudents(data.data.data);
      })
      .catch((err) => console.log(err));
  }, []);

  const openDialogForDelete = (studentID: string) => {
    dispatch(closeModal());
    dispatch(setModalOpen());

    setContent(
      <DeleteDialog
        functionToProceedOnDelete={() => {
          deleteStudent(studentID);
        }}
      />
    );
  };

  const deleteStudent = (studentID: string) => {
    axiosInstance
      .delete(`/food/deleteStudent/${studentID}`)
      .then((data) => {
        console.log(data);
        setStudents(students.filter((student) => student.ID !== studentID));
      })
      .catch((err) => {
        console.log(err);
      });
    dispatch(closeModal());
  };

  return (
    <>
      <div className="overflow-x-auto">
        <table className="min-w-full bg-white">
          <thead>
            <tr>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Name
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Birth Date
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Personal ID
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Status
              </th>
              {adminView && (
                <>
                  <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                    Edit
                  </th>
                  <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                    Delete
                  </th>
                </>
              )}
            </tr>
          </thead>
          <tbody>
            {students && students.length > 0 ? (
              students.map((student, index) => (
                <tr key={index}>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                    {student.full_name}
                  </td>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                    {student.birth_date}
                  </td>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                    {student.personal_id_number}
                  </td>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                    {student.status}
                  </td>
                  {adminView && (
                    <>
                      <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                        <button
                          className="border bg-auburn-500 border-auburn-500 font-semibold py-1 px-2 rounded focus:border-auburn-700 text-white"
                          onClick={(e) => {
                            e.preventDefault();
                            openDialogForDelete(student.ID);
                            console.log(student.ID);
                          }}
                        >
                          Delete Student
                        </button>
                      </td>
                    </>
                  )}
                </tr>
              ))
            ) : (
              <p className="my-2">Students are not available..</p>
            )}
          </tbody>
        </table>
      </div>
    </>
  );
};

export default StudentTable;
