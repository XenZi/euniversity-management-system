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

const DormPage = () => {
  const user = useSelector((state: RootState) => state.user.user);
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
        setContent(<DormTable />);
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
  ];

  return (
    <div className="h-screen bg-papaya-500 w-full p-3">
      <Navigation />
      <div className="max-w-7xl mx-auto w-100 flex ">
        {user?.roles[0] == "Admin" ? (
          <>
            <AdminComponent children={dormAdminComponents} />
          </>
        ) : (
          "Student"
        )}
      </div>
    </div>
  );
};

export default DormPage;
