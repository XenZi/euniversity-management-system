import { createBrowserRouter, RouterProvider } from "react-router-dom";
import LoginPage from "./pages/Login.page";
import DormPage from "./pages/Dorm.page";
import { Provider, useDispatch } from "react-redux";
import store from "./redux/store/user.store";
import HomePage from "./pages/Home.page";
import { useEffect } from "react";
import PrivateRoute from "./components/routing/private-route.component";
import useLocalStorage from "./hooks/local-storage.hook";
import { setUser } from "./redux/slices/user.slice";
import { User } from "./models/user.model";
import HealthcarePage from "./pages/Healthcare.page";
import UniversityPage from "./pages/University.page";
import FoodPage from "./pages/Food.page";

function App() {
  const [userFromLocalStorage, setUserInLocalStorage] = useLocalStorage(
    "user",
    null
  );
  const router = createBrowserRouter([
    {
      path: "/",
      Component: LoginPage,
    },
    {
      path: "/dorm",
      element: <PrivateRoute Component={DormPage} />,
    },
    {
      path: "/home",
      element: <PrivateRoute Component={HomePage} />,
    },
    {
      path: "/healthcare",
      element: <PrivateRoute Component={HealthcarePage} />,
    },
    {
      path: "/university",
      element: <PrivateRoute Component={UniversityPage} />,
    },
    {
      path: "/food",
      element: <PrivateRoute Component={FoodPage} />,
    },
  ]);
  const dispatch = useDispatch();
  useEffect(() => {
    if (userFromLocalStorage) {
      dispatch(setUser({ ...(userFromLocalStorage as User) }));
    }
  }, []);
  return (
    <>
      <RouterProvider router={router} />
    </>
  );
}

export default App;
