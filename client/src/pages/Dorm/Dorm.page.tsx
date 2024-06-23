import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../../redux/store/store";
import Navigation from "../../components/navigation/navigation.component";
import AdminComponent from "../../components/admin/admin.component";
import PanelBox from "../../components/panel-box/panel-box.component";
import Modal from "../../components/modal/modal.component";
import { setModalOpen } from "../../redux/slices/modal.slice";
import CreateDormForm from "../../components/forms/create-dorm/create-dorm.form";
import { useState } from "react";
import CreateRoomForm from "../../components/forms/create-room/create-room.form";

const DormPage = () => {
  const user = useSelector((state: RootState) => state.user.user);
  const [componentToReadInModal, setComponentToReadInModal] =
    useState<string>("");
  const dispatch = useDispatch();
  const openModal = (componentName: string) => {
    setComponentToReadInModal(componentName);
    dispatch(setModalOpen());
  };
  const dormAdminComponents: React.JSX.Element[] = [
    <PanelBox
      panelBoxTitle="Create new dorm"
      onClick={() => {
        openModal("createDorm");
      }}
    ></PanelBox>,
    <PanelBox
      panelBoxDescription="Create new dorm room"
      onClick={() => {
        openModal("createRoom");
      }}
    ></PanelBox>,
  ];

  const chooseWhatComponentToRenderInModal = (): JSX.Element | null => {
    switch (componentToReadInModal) {
      case "createDorm":
        return <CreateDormForm />;
      case "createRoom":
        return <CreateRoomForm />;
      default:
        return null;
    }
  };
  return (
    <div className="h-screen bg-papaya-500 w-full p-3">
      <Navigation />
      <Modal>{chooseWhatComponentToRenderInModal()}</Modal>
      <div className="max-w-7xl mx-auto w-100 flex ">
        {user?.roles[0] == "Admin" ? (
          <AdminComponent children={dormAdminComponents} />
        ) : (
          "Student"
        )}
      </div>
    </div>
  );
};

export default DormPage;
