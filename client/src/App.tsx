import { createBrowserRouter, RouterProvider } from "react-router-dom";
import LoginPage from "./pages/Login.page";
import DormPage from "./pages/Dorm.page";
import { Provider } from "react-redux";
import store from "./redux/store/user.store";
import HomePage from "./pages/Home.page";
import { useEffect } from "react";
import PrivateRoute from "./components/routing/private-route.component";

function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      Component: LoginPage,
    },
    {
      path: "/dorm",
      element: <PrivateRoute Component={DormPage} requiredRole="Citizen" />,
    },
    {
      path: "/home",
      element: <HomePage />,
    },
  ]);

  useEffect(() => {
    console.log("Example of reload gr");
  }, []);
  return (
    <>
      <Provider store={store}>
        <RouterProvider router={router} />
      </Provider>
    </>
  );
}

export default App;
