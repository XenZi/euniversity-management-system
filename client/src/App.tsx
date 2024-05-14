import { createBrowserRouter, RouterProvider } from "react-router-dom";
import LoginPage from "./pages/Login.page";
import DormPage from "./pages/Dorm.page";

function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      Component: LoginPage,
    },
    {
      path: "/dorm",
      Component: DormPage,
    },
  ]);
  return <RouterProvider router={router} />;
}

export default App;
