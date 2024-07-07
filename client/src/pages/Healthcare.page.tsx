import { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../redux/store/store";
import { setModalOpen } from "../redux/slices/modal.slice";
import Navigation from "../components/navigation/navigation.component";
import { useModalContext } from "../context/modal.context";
import PanelBox from "../components/panel-box/panel-box.component";
import { UserRecord } from "../models/record.model";
import { axiosInstance } from "../services/axios.service";
import PatientPanel from "../components/healthcare-panels/patient-panel";
import DoctorPanel from "../components/healthcare-panels/doctor-panel";

const HealthcarePage = () => {
  const user = useSelector((state: RootState) => state.user.user);
  const [userRecord, setUserRecord] = useState<UserRecord>();
  const dispatch = useDispatch()
  const { setContent } = useModalContext();
  const openModal = () => {
    dispatch(setModalOpen());
  };



  useEffect(() => {
    if (user?.roles[0] == "Patient"){
      axiosInstance
      .get(`/healthcare/records/${user?.personalIdentificationNumber}`)
      .then((data) => {
        setUserRecord(data.data.data);
        console.log("Probao", user?.personalIdentificationNumber)
      })
      .catch((err) => {
        console.error(err)
      });
    }

  }, [user])

  useEffect(() => {
    // console.log("Record: ", userRecord)
  }, [userRecord])

  return (
    <div className="h-screen bg-papaya-500 w-full p-3">
      <Navigation />
      {user?.roles[0] == "Doctor" ?
       (<>
       <DoctorPanel/>
       </>) :
        (<>
          {userRecord  ? (
            <>
            <PatientPanel
            userID = {user?.personalIdentificationNumber}
            />
                       
            </>
            
            
          ) : (
            <div className="flex flex-col items-center justify-center w-full">
            <div className="bg-white w-full p-4 rounded-lg shadow-md mb-4">
                <h3 className="text-3xl text-center mb-4">You must chose a healtcare role</h3>

                </div>
                </div>
          )}
        </>)}
      
    </div>
  );
};

export default HealthcarePage;
