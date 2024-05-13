import { createBrowserRouter, RouterProvider } from "react-router-dom";
import TestPage from "./pages/Test.page";
import NovaPage from "./pages/Nova.page";

function App() {
  // <>
  //   {/* <div className="bg-papaya-500 h-screen">
  //     <h3 className="text-gunmetal-500 text-5xl">Test</h3>
  //     <h3 className="text-bittersweet-500 text-5xl">Test</h3>
  //     <h3 className="text-auburn-500 text-5xl">Test</h3>
  //     <h3 className="text-battleship-500 text-5xl">Test</h3>
  //   </div> */}
  // </>
  const router = createBrowserRouter([
    {
      path: "/",
      Component: TestPage,
    },
    {
      path: "/test",
      Component: NovaPage,
    },
  ]);
  return <RouterProvider router={router} fallbackElement={<p>Loading</p>} />;
}

export default App;
