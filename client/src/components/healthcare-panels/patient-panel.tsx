import { UserRecord } from "../../models/record.model"
import { useSelector } from "react-redux";
import { RootState } from "../../redux/store/store";
import { useEffect, useState } from "react";
import { axiosInstance } from "../../services/axios.service";
import { useDispatch } from "react-redux";
import { useModalContext } from "../../context/modal.context";
import { closeModal, setModalOpen } from "../../redux/slices/modal.slice";
import CreateReferralForm from "../forms/healthcare-forms/create-referral";
import { User } from "../../models/user.model";






const PatientPanel: React.FC<{
    userID? : string
}> = ({userID}) => {    
    const user = useSelector((state: RootState) => state.user.user);
    const [userRecord, setUserRecord] = useState<UserRecord>();
    const [doctors, setDoctors] = useState<User[]>([]);

    const [referrals, setReferrals] = useState<any[]>([]); // State for referrals

    const dispatch = useDispatch();
    const { setContent } = useModalContext();

    useEffect(() => {
        if ((user?.roles[0] == "Patient" && userID == user?.personalIdentificationNumber) || user?.roles[0]=="Doctor"){
          axiosInstance
          .get(`/healthcare/records/${userID}`)
          .then((data) => {
            setUserRecord(data.data.data);
            setReferrals(data.data.data.referrals); // Set referrals from fetched data

            console.log("Probao", userID)
            console.log("data:", data)
          })
          .catch((err) => {
            console.error(err)
          });
          axiosInstance
          .get("/auth/getUsers/Doctor")
          .then((data) => {
              setDoctors(data.data.data)
              console.log(doctors)
          })
          .catch((err) => {
              console.log(err)
          })
        }
    
      }, [userID])  

    const handleMakeReferral = (patientId: string | undefined) => {
        dispatch(closeModal())
        dispatch(setModalOpen());
        setContent(<CreateReferralForm patientId={patientId}/>)
    }

    return (
        <div className="flex flex-col items-center justify-center w-full">
        <div className="flex flex-row items-center justify-center w-full">
            <div className="bg-white h-full w-full">
                <div className="flex flex-row">
                    <div className="flex flex-col w-1/2">
                    <h3 className="text-3xl text-center">Record Details</h3>
                    <ul className="list-none pl-3">
                    <li className="mb-2">
                        Patient Name: <span className="font-bold">{user?.fullName}</span>
                    </li>
                    <li className="mb-2">
                        Patient PIN: <span className="font-bold">{userRecord?.patientID}</span>
                    </li>
                    <li className="mb-2">
                        Record ID: <span className="font-bold">{userRecord?.id}</span>
                    </li>
                    </ul>
                    </div>
                    <div className="flex flex-col w-1/2">
                    <h3 className="text-3xl text-center">Certificate Details</h3>
                    <ul className="list-none pl-3">
                    <li className="mb-2">
                       Title: <span className="font-bold">{userRecord?.certificate?.report?.title}</span>
                    </li>
                    <li className="mb-2">
                        Content: <span className="font-bold">{userRecord?.certificate?.report?.content}</span>
                    </li>
                    <li className="mb-2">
                        Date of Issue: <span className="font-bold">{userRecord?.certificate?.dateOfIssue}</span>
                    </li>
                    </ul>
                    </div>
                    
                </div>
            </div>
            
        </div>
        <br></br>
          {referrals && referrals.length > 0 && (
                <div className="overflow-x-auto mt-4">
                    <table className="min-w-full bg-white">
                        <thead>
                            <tr>
                                <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                                    Date of Issue
                                </th>
                                <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                                    Doctor Name
                                </th>
                                <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                                    Doctor ID
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            {referrals.map((referral, index) => (
                                <tr key={index}  className="cursor-pointer hover:bg-gray-100">
                                    <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                                        {referral.dateOfIssue}
                                    </td>
                                    <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                                    {doctors.find(doc => doc.personalIdentificationNumber === referral.doctorID)?.fullName}                                    </td>
                                    <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                                        {referral.doctorID}
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            )}
        {user?.roles[0] == "Doctor" ? (
            <div className="bg-white h-full w-full">
                <button
                          className="border bg-auburn-500 border-auburn-500 font-semibold py-1 px-2 rounded focus:border-auburn-700 text-white"
                          onClick={(e) => {
                            e.preventDefault();
                            handleMakeReferral(userRecord?.patientID);
                          }}
                        >
                          Make Referral
                        </button>
            </div>) : ("")}

        </div>
    );
};

export default PatientPanel;