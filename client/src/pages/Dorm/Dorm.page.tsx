import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../../redux/store/store";
import Navigation from "../../components/navigation/navigation.component";
import AdminComponent from "../../components/admin/admin.component";
import PanelBox from "../../components/panel-box/panel-box.component";
import { setModalOpen } from "../../redux/slices/modal.slice";
import CreateDormForm from "../../components/forms/create-dorm/create-dorm.form";
import { useModalContext } from "../../context/modal.context";
import CreateRoomForm from "../../components/forms/create-room/create-room.form";
import DormTable from "../../components/dorms-table/dorms-table.component";
import RoomsTable from "../../components/rooms-table/rooms-table.component";
import Admissions from "../../components/forms/admissions/admissions.form";
import AdmissionsTable from "../../components/admissions-table/admissions-table.component";
import { useEffect, useState } from "react";
import { Room } from "../../models/room.model";
import { axiosInstance } from "../../services/axios.service";
import { Dorm } from "../../models/dorm.model";
import StudentDormPanel from "../../components/student-dorm-panel/student-dorm-panel";
import DormitoryApplication from "../../components/forms/dormitory-application/dormitory-application.form";

const DormPage = () => {
  const user = useSelector((state: RootState) => state.user.user);
  const [userRoom, setUserRoom] = useState<Room>();
  const [dormDetails, setDormDetails] = useState<Dorm>();
  const dispatch = useDispatch();
  const { setContent } = useModalContext();
  const openModal = () => {
    dispatch(setModalOpen());
  };
  const dormAdminComponents: React.JSX.Element[] = [
    <PanelBox
      panelBoxTitle="Create new dorm"
      onClick={() => {
        openModal();
        setContent(<CreateDormForm />);
      }}
    ></PanelBox>,
    <PanelBox
      panelBoxDescription="Create new dorm room"
      onClick={() => {
        openModal();
        setContent(<CreateRoomForm />);
      }}
    ></PanelBox>,
    <PanelBox
      panelBoxDescription="View all dorms"
      onClick={() => {
        openModal();
        setContent(<DormTable adminView={true} />);
      }}
    />,
    <PanelBox
      panelBoxDescription="View all rooms"
      onClick={() => {
        openModal();
        setContent(<RoomsTable />);
      }}
    />,
    <PanelBox
      panelBoxDescription="Create new admission"
      onClick={() => {
        openModal();
        setContent(<Admissions />);
      }}
    />,
    <PanelBox
      panelBoxDescription="List all admissions"
      onClick={() => {
        openModal();
        setContent(<AdmissionsTable adminView={true} />);
      }}
    />,
  ];
  const dormStudentComponents: React.JSX.Element[] = [
    <PanelBox
      panelBoxTitle="List all dorms"
      onClick={() => {
        openModal();
        setContent(<DormTable />);
      }}
    />,
    <PanelBox
      panelBoxTitle="List all active admissions"
      onClick={() => {
        openModal();
        setContent(<AdmissionsTable />);
      }}
    />,
    <PanelBox
      panelBoxDescription="Create application"
      onClick={() => {
        openModal();
        setContent(<DormitoryApplication />);
      }}
    />,
  ];

  useEffect(() => {
    if (user?.roles[0] === "Citizen") {
      axiosInstance
        .get(`/dorm/room/student/${user.personalIdentificationNumber}`)
        .then((data) => {
          setUserRoom(data.data.data);
        })
        .catch((err) => {
          console.error(err);
        });
    }
  }, [user]);

  useEffect(() => {
    if (!userRoom) return;
    axiosInstance
      .get(`/dorm/${userRoom?.dormID}`)
      .then((data) => {
        setDormDetails(data.data.data);
      })
      .catch((err) => {
        console.error(err);
      });
  }, [userRoom]);

  useEffect(() => {
    console.log("User Room:", userRoom);
    console.log("Dorm Details:", dormDetails);
  }, [userRoom, dormDetails]);

  return (
    <div className="h-screen bg-papaya-500 w-full p-3">
      <Navigation />
      <div className="max-w-7xl mx-auto w-100 flex">
        {user?.roles[0] === "Admin" ? (
          <AdminComponent children={dormAdminComponents} />
        ) : (
          <>
            {userRoom && dormDetails ? (
              <StudentDormPanel
                room={userRoom as Room}
                dorm={dormDetails as Dorm}
              />
            ) : (
              dormStudentComponents.map((el) => el)
            )}
          </>
        )}
      </div>
    </div>
  );
};

export default DormPage;

// student can see all the dorms
// student can see all the currently open admissions
// student can see all of his applications
