import { createBrowserRouter, RouterProvider } from "react-router-dom";
import LoginPage from "./pages/Login.page";
import DormPage from "./pages/Dorm.page";
import { Provider } from "react-redux";
import store from "./redux/store/user.store";

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
  return (
    <>
      <Provider store={store}>
        <RouterProvider router={router} />;
      </Provider>
    </>
  );
}

export default App;
