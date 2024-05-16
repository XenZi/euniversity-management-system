import { useNavigate } from "react-router-dom";
import Navigation from "../components/navigation/navigation.component";

const HomePage = () => {
  const navigate = useNavigate();
  return (
    <div className="h-screen bg-papaya-500">
      <Navigation />
      <div className="max-w-7xl mx-auto w-100 h-full">
        <div className="flex h-full w-100 items-center justify-center">
          <button
            className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white"
            onClick={() => {
              navigate("/dorm");
            }}
          >
            Go to Dorm Page
          </button>
        </div>
      </div>
    </div>
  );
};

export default HomePage;
