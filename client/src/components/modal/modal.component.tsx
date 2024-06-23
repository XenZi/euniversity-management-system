import React, { FC, useRef } from "react";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../../redux/store/store";
import { closeModal } from "../../redux/slices/modal.slice";
import useOutsideClick from "../../hooks/outside-click.hook";

interface ModalProps {
  children: React.ReactNode;
}

const Modal: FC<ModalProps> = ({ children }) => {
  const modal = useSelector((state: RootState) => state.modal);
  const dispatch = useDispatch();
  const modalRef = useRef<HTMLDivElement>(null);

  const callCloseModalFunction = () => {
    dispatch(closeModal());
  };

  useOutsideClick(modalRef, callCloseModalFunction, modal.isVisible);
  return (
    <div
      className={`fixed inset-0 bg-black bg-opacity-75 ${
        modal.isVisible ? "flex items-center justify-center" : "hidden"
      }`}
    >
      <div
        className="relative bg-white p-10 rounded-lg shadow-lg max-w-md w-full"
        ref={modalRef}
      >
        <button
          onClick={() => {
            callCloseModalFunction();
          }}
          className="absolute top-1 right-2"
        >
          X
        </button>
        {children}
      </div>
    </div>
  );
};

export default Modal;
