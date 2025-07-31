import { useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import { useModalContext } from "../../../context/modal.context";
import { axiosInstance } from "../../../services/axios.service";
import { closeModal, setModalOpen } from "../../../redux/slices/modal.slice";

import DeleteDialog from "../../dialogs/delete-dialog/delete-dialog.component";
import { Student } from "../../../models/student.model";

type MessRoom = {
  id: string;
  name: string;
  mess_room_users?: string[];
  [k: string]: any;
};

const StudentTable: React.FC<{ adminView?: boolean }> = ({ adminView }) => {
  const [students, setStudents] = useState<Student[]>([]);
  const [messRooms, setMessRooms] = useState<MessRoom[]>([]);
  const dispatch = useDispatch();
  const { setContent } = useModalContext();

  useEffect(() => {
    // fetch students
    axiosInstance
      .get("/food/allStudents")
      .then((data) => {
        setStudents(data.data.data);
      })
      .catch((err) => console.log(err));

    // fetch messrooms
    axiosInstance
      .get("/food/allMessRooms")
      .then((res) => {
        setMessRooms(res.data.data);
      })
      .catch((err) => console.log("failed to load mess rooms", err));
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
      .then(() => {
        setStudents((prev) =>
          prev.filter((student) => student.ID !== studentID)
        );
      })
      .catch((err) => {
        console.log(err);
      });
    dispatch(closeModal());
  };

  // helper to get mess names for a student based on their personal_id_number
  const getMessNamesForStudent = (student: Student): string => {
    const pid = student.personal_id_number;
    if (!pid) return "—";
    const names = messRooms
      .filter((m) => m.mess_room_users?.includes(pid))
      .map((m) => m.name);
    if (names.length === 0) return "—";
    return Array.from(new Set(names)).join(", ");
  };

  return (
    <div className="space-y-4">
      <div className="overflow-x-auto">
        <table className="min-w-full bg-white border-collapse">
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
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Mess(es)
              </th>
              {adminView && (
                <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                  Delete
                </th>
              )}
            </tr>
          </thead>
          <tbody>
            {students && students.length > 0 ? (
              students.map((student, index) => {
                const messNames = getMessNamesForStudent(student);
                return (
                  <tr
                    key={student.ID || index}
                    className="even:bg-gray-50 hover:bg-gray-100"
                  >
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
                    <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                      {messNames}
                    </td>
                    {adminView && (
                      <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                        <button
                          className="border bg-auburn-500 font-semibold py-1 px-2 rounded text-white text-xs"
                          onClick={(e) => {
                            e.preventDefault();
                            openDialogForDelete(student.ID);
                          }}
                        >
                          Delete Student
                        </button>
                      </td>
                    )}
                  </tr>
                );
              })
            ) : (
              <tr>
                <td
                  colSpan={adminView ? 6 : 5}
                  className="py-4 px-4 text-center text-sm text-gray-600"
                >
                  Students are not available...
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default StudentTable;
