import { useRef } from "react";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../../redux/store/store";
import { closeModal } from "../../redux/slices/modal.slice";
import useOutsideClick from "../../hooks/outside-click.hook";
import { useModalContext } from "../../context/modal.context";

const Modal = () => {
  const modal = useSelector((state: RootState) => state.modal);
  const dispatch = useDispatch();
  const modalRef = useRef<HTMLDivElement>(null);
  const { content } = useModalContext();

  const callCloseModalFunction = () => {
    dispatch(closeModal());
  };

  console.log(modal.isVisible);
  useOutsideClick(modalRef, callCloseModalFunction, modal.isVisible);
  return (
    <div
      className={`fixed inset-0 bg-black bg-opacity-75 ${
        modal.isVisible ? "flex items-center justify-center" : "hidden"
      }`}
    >
      <div
        className="relative bg-white p-10 rounded-lg shadow-lg max-w-5xl w-full"
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
        {content}
      </div>
    </div>
  );
};

export default Modal;
