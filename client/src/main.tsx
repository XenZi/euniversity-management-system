import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import { Provider } from "react-redux";
import store from "./redux/store/store.ts";
import { ModalProvider } from "./context/modal.context.tsx";
import Modal from "./components/modal/modal.component.tsx";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <Provider store={store}>
      <ModalProvider>
        <App />
        <Modal />
      </ModalProvider>
    </Provider>
  </React.StrictMode>
);
