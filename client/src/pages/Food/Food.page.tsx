import { useDispatch, useSelector } from "react-redux";
import { useModalContext } from "../../context/modal.context";
import { setModalOpen } from "../../redux/slices/modal.slice";

import Navigation from "../../components/navigation/navigation.component";
import { RootState } from "../../redux/store/store";
import PanelBox from "../../components/panel-box/panel-box.component";
import CreateDormForm from "../../components/forms/create-dorm/create-dorm.form";

import DormTable from "../../components/dorms-table/dorms-table.component";
import RoomsTable from "../../components/rooms-table/rooms-table.component";
import Admissions from "../../components/forms/admissions/admissions.form";
import AdmissionsTable from "../../components/admissions-table/admissions-table.component";
import AdminDormitoryApplicationTable from "../../components/admin-dormitory-applications/admin-dormitory-applications.table";
import AdminComponent from "../../components/admin/admin.component";
import DormitoryApplication from "../../components/forms/dormitory-application/dormitory-application.form";
import CreateMessForm from "../../components/forms/create-mess/create-mess.form";
import MessTable from "../../components/forms/messes-table/mess-table.component";
import StudentTable from "../../components/forms/mess-users-table/students-table.component";
import CreateSupplierForm from "../../components/forms/create-supplier-form/create-supplier.component";

const FoodPage = () => {
  const user = useSelector((state: RootState) => state.user.user);
  const dispatch = useDispatch();
  const { setContent } = useModalContext();

  const openModal = () => {
    dispatch(setModalOpen());
  };

  const foodAdminComponents: React.JSX.Element[] = [
    <PanelBox
      panelBoxTitle="Create new mess"
      onClick={() => {
        openModal();
        setContent(<CreateMessForm />);
      }}
    ></PanelBox>,

    <PanelBox
      panelBoxDescription="View all messes"
      onClick={() => {
        openModal();
        setContent(<MessTable adminView={true} />);
      }}
    />,
    <PanelBox
      panelBoxDescription="View all mess users"
      onClick={() => {
        openModal();
        setContent(<StudentTable adminView={true} />);
      }}
    />,
    <PanelBox
      panelBoxDescription="Create new supplier admission"
      onClick={() => {
        openModal();
        setContent(<CreateSupplierForm />);
      }}
    />,
  ];

  const foodStudentComponents: React.JSX.Element[] = [
    <PanelBox
      panelBoxTitle="List all messes"
      onClick={() => {
        openModal();
        setContent(<MessTable />);
      }}
    />,
    <PanelBox
      panelBoxTitle="Buy credits"
      onClick={() => {
        openModal();
        setContent(<AdmissionsTable />);
      }}
    />,
    <PanelBox
      panelBoxDescription="Rate a mess"
      onClick={() => {
        openModal();
        setContent(<DormitoryApplication />);
      }}
    />,
    <PanelBox
      panelBoxDescription="Create new supplier admission"
      onClick={() => {
        openModal();
        setContent(<DormitoryApplication />);
      }}
    />,
    <PanelBox
      panelBoxDescription="My card"
      onClick={() => {
        openModal();
        setContent(<DormitoryApplication />);
      }}
    />,
  ];
  return (
    <div className="h-screen bg-papaya-500 w-full p-3">
      <Navigation />
      <div className="max-w-7xl mx-auto w-100 flex">
        {user?.roles[0] === "Admin" ? (
          <AdminComponent children={foodAdminComponents} />
        ) : (
          foodStudentComponents
        )}
      </div>
    </div>
  );
};

export default FoodPage;
