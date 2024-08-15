import { useSelector } from "react-redux";
import { RootState } from "../../redux/store/store";
import React, { useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import { User } from "../../models/user.model";
import { axiosInstance } from "../../services/axios.service";
import { closeModal, setModalOpen } from "../../redux/slices/modal.slice";
import PatientPanel from "./patient-panel";
import { useModalContext } from "../../context/modal.context";





const DoctorPanel: React.FC<{}> = () => {
    const user = useSelector((state: RootState) => state.user.user);
    const [users, setUsers] = useState<User[]>();
    const dispatch = useDispatch();
    const { setContent } = useModalContext();



    useEffect(() => {
        axiosInstance
        .get("/auth/getUsers/Patient")
        .then((data) => {
            setUsers(data.data.data)
            console.log(users)
        }) .catch((error) => {
          console.error("Error fetching users:", error);
      });
  }, []);

    const handleUserClick = (userId:string) => {
      dispatch(closeModal());
      dispatch(setModalOpen());
      setContent(<PatientPanel userID={userId} />);
    };

    

    return (
        <div className="overflow-x-auto">
          <table className="min-w-full bg-white">
            <thead>
              <tr>
                <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                  Full Name
                </th>
                <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                  Personal Identification Number
                </th>
              </tr>
            </thead>
            <tbody>
          {users && users.length > 0 ? (
            users.map((user, index) => (
              <tr key={index} onClick={(e) => {
                e.preventDefault();
                handleUserClick(user.personalIdentificationNumber);
              }}>
                <td className="py-2 px-4 border-b border-gray-300 text-sm text-center cursor-pointer text-blue-500">
                  {user.fullName}
                </td>
                <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                  {user.personalIdentificationNumber}
                </td>
              </tr>
            ))
          ) : (
            <tr>
              <td
                colSpan={2}
                className="py-2 px-4 border-b border-gray-300 text-sm text-center"
              >
                No users available.
              </td>
            </tr>
          )}
        </tbody>
          </table>
        </div>
      );
    };

export default DoctorPanel;