import { useSelector } from "react-redux";
import { RootState } from "../redux/store/store";
import Navigation from "../components/navigation/navigation.component";

const HealthcarePag = () => {
  const user = useSelector((state: RootState) => state.user.user);
  console.log(user);
  return (
    <div className="h-screen bg-papaya-500 w-full p-3">
      <Navigation />
      {user?.roles[0] == "Admin" ? "Admin" : "Student"}
    </div>
  );
};

export default HealthcarePag;
